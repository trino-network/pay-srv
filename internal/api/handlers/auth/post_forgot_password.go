package auth

import (
	"database/sql"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/trino-network/pay-srv/internal/api"
	"github.com/trino-network/pay-srv/internal/models"
	"github.com/trino-network/pay-srv/internal/types"
	"github.com/trino-network/pay-srv/internal/util"
	"github.com/trino-network/pay-srv/internal/util/db"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func PostForgotPasswordRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Auth.POST("/forgot-password", postForgotPasswordHandler(s))
}

func postForgotPasswordHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var body types.PostForgotPasswordPayload
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		log := util.LogFromContext(ctx).With().Str("username", body.Username.String()).Logger()

		user, err := models.Users(models.UserWhere.Username.EQ(null.StringFrom(body.Username.String()))).One(ctx, s.DB)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Debug().Str("username", body.Username.String()).Err(err).Msg("User not found")
				return c.NoContent(http.StatusNoContent)
			}

			log.Debug().Str("username", body.Username.String()).Err(err).Msg("Failed to load user")
			return err
		}

		if !user.IsActive {
			log.Debug().Msg("User is deactivated, rejecting password reset")
			return c.NoContent(http.StatusNoContent)
		}

		if !user.Password.Valid {
			log.Debug().Msg("User is missing password, forbidding password reset")
			return c.NoContent(http.StatusNoContent)
		}

		if err := db.WithTransaction(ctx, s.DB, func(tx boil.ContextExecutor) error {
			passwordResetToken, err := user.PasswordResetTokens(
				models.PasswordResetTokenWhere.CreatedAt.GT(time.Now().Add(time.Minute*1)),
			).One(ctx, tx)
			if err != nil {
				if err == sql.ErrNoRows {
					log.Debug().Err(err).Msg("No valid password reset token exists, creating new one")

					passwordResetToken = &models.PasswordResetToken{
						UserID:     user.ID,
						ValidUntil: time.Now().Add(s.Config.Auth.PasswordResetTokenValidity),
					}

					if err := passwordResetToken.Insert(ctx, tx, boil.Infer()); err != nil {
						log.Debug().Err(err).Msg("Failed to insert password reset token")
						return err
					}
				} else {
					log.Debug().Err(err).Msg("Failed to check for existing valid password reset token")
					return err
				}
			}

			u, err := url.Parse(s.Config.Frontend.BaseURL)
			if err != nil {
				log.Error().Err(err).Msg("Failed to parse frontend base URL")
				return err
			}

			u.Path = path.Join(u.Path, s.Config.Frontend.PasswordResetEndpoint)

			q := u.Query()
			q.Set("token", passwordResetToken.Token)
			u.RawQuery = q.Encode()

			if err := s.Mailer.SendPasswordReset(ctx, user.Username.String, u.String()); err != nil {
				log.Debug().Err(err).Msg("Failed to send password reset email")
				return err
			}

			return nil
		}); err != nil {
			log.Debug().Err(err).Msg("Failed to initiate password reset")
			return err
		}

		log.Debug().Msg("Successfully initiated password reset")

		return c.NoContent(http.StatusNoContent)
	}
}
