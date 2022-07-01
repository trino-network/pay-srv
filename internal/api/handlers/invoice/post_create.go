package invoice

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/labstack/echo/v4"
	"github.com/trino-network/pay-srv/internal/api"
	"github.com/trino-network/pay-srv/internal/api/auth"
	"github.com/trino-network/pay-srv/internal/api/httperrors"
	"github.com/trino-network/pay-srv/internal/models"
	"github.com/trino-network/pay-srv/internal/types"
	"github.com/trino-network/pay-srv/internal/util"
	"github.com/trino-network/pay-srv/internal/util/db"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http"
	"regexp"
	"time"
)

func PostCreateRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Invoice.POST("/create", postCreateHandler(s))
}

func postCreateHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log := util.LogFromContext(ctx)
		user := auth.UserFromContext(ctx)

		var body types.PostCreateInvoicePayload
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}
		if *body.Network == types.PostCreateInvoicePayloadNetworkGoerli {
			re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
			if !re.MatchString(*body.Recipient) {
				return httperrors.ErrInvalidRecipient
			}
		}

		if err := db.WithTransaction(ctx, s.DB, func(tx boil.ContextExecutor) error {
			invoice := &models.Invoice{
				UserID:    user.ID,
				Title:     *body.Title,
				Network:   *body.Network,
				Recipient: *body.Recipient,
				Amount:    null.Int64From(*body.Amount),
				JumpURL:   *body.JumpURL,
				NotifyURL: *body.NotifyURL,
				Status:    types.GetInvoiceResponseStatusUnpaid,
				Metadata:  *body.Metadata,
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			}
			if err := invoice.Insert(ctx, tx, boil.Infer()); err != nil {
				log.Debug().Err(err).Msg("Failed to insert invoice")
				return err
			}

			log.Debug().Msg("Successfully created invoice")
			jumpURL := fmt.Sprintf("%s/%s", s.Config.Frontend.BaseURL, invoice.ID)
			response := &types.PostCreateInvoiceResponse{
				InvoiceID: conv.UUID4(strfmt.UUID4(invoice.ID)),
				JumpURL:   &jumpURL,
			}

			return util.ValidateAndReturn(c, http.StatusOK, response)
		}); err != nil {
			log.Debug().Err(err).Msg("Failed to insert invoice")
			return err
		}
		return nil
	}
}
