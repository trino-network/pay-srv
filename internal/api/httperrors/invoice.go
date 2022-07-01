package httperrors

import (
	"net/http"
)

var (
	ErrNotFoundInvoiceNotFound = NewHTTPError(http.StatusNotFound, "INVOICE_NOT_FOUND", "Provided invoice id was not found")
	ErrInvalidRecipient        = NewHTTPError(http.StatusBadRequest, "INVALID_RECIPIENT", "invalid recipient")
)
