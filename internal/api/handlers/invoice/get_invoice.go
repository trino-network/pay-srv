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

func GetRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Invoice.GET("/:id", getShow(s))
}

func getShow(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		invoiceId := fmt.Sprintf("%v", c.Param("id"))
		invoice, err := models.Invoices(models.InvoiceWhere.ID.EQ(invoiceId)).One(ctx, s.DB)
		if err != nil || invoice == nil {
			return httperrors.ErrNotFoundInvoiceNotFound
		}

		rsp := types.GetInvoiceResponse{
			Status:   &invoice.Status,
			Metadata: &invoice.Metadata,
			TxnHash:  &invoice.TXNHash.String,
		}

		return util.ValidateAndReturn(c, http.StatusOK, &rsp)
	}
}
