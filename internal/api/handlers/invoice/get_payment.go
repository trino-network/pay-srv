package invoice

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/trino-network/pay-srv/internal/api"
	"github.com/trino-network/pay-srv/internal/api/httperrors"
	"github.com/trino-network/pay-srv/internal/models"
	"github.com/trino-network/pay-srv/internal/types"
	"github.com/trino-network/pay-srv/internal/util"
	"net/http"
)

func GetPaymentRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Invoice.GET("/payment/:id", getPayment(s))
}

func getPayment(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		invoiceId := fmt.Sprintf("%v", c.Param("id"))
		invoice, err := models.Invoices(models.InvoiceWhere.ID.EQ(invoiceId)).One(ctx, s.DB)
		if err != nil || invoice == nil {
			return httperrors.ErrNotFoundInvoiceNotFound
		}

		rsp := types.GetInvoicePaymentInfoResponse{
			Title:     &invoice.Title,
			Amount:    &invoice.Amount.Int64,
			JumpURL:   &invoice.JumpURL,
			Network:   &invoice.Network,
			Recipient: &invoice.Recipient,
		}

		return util.ValidateAndReturn(c, http.StatusOK, &rsp)
	}
}
