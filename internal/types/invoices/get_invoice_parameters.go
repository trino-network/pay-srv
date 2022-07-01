// Code generated by go-swagger; DO NOT EDIT.

package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetInvoiceParams creates a new GetInvoiceParams object
// no default values defined in spec.
func NewGetInvoiceParams() GetInvoiceParams {

	return GetInvoiceParams{}
}

// GetInvoiceParams contains all the bound params for the get invoice operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetInvoice
type GetInvoiceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*id of invoice
	  Required: true
	  In: path
	*/
	ID string `param:"id"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetInvoiceParams() beforehand.
func (o *GetInvoiceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInvoiceParams) Validate(formats strfmt.Registry) error {
	var res []error

	// id
	// Required: true
	// Parameter is provided by construction from the route

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetInvoiceParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}
