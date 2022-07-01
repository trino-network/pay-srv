// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostCreateInvoicePayload post create invoice payload
//
// swagger:model postCreateInvoicePayload
type PostCreateInvoicePayload struct {

	// amount
	// Example: 234535435345
	// Required: true
	Amount *int64 `json:"amount"`

	// jump url
	// Example: http://xxx/order/xxx
	// Required: true
	JumpURL *string `json:"jump_url"`

	// 用户自定义元数据，base64
	// Example: e29c1B698F98BdFe6Ca4012dEE6FB350D73E40AE
	// Required: true
	// Max Length: 512
	// Min Length: 0
	Metadata *string `json:"metadata"`

	// 支付所用的区块链网络
	// Example: ethereum
	// Required: true
	// Enum: [ethereum bsc polygon solana near goerli]
	Network *string `json:"network"`

	// notify url
	// Example: http://xxx/order/xxx
	// Required: true
	NotifyURL *string `json:"notify_url"`

	// recipient
	// Example: 0xe29c1B698F98BdFe6Ca4012dEE6FB350D73E40AE
	// Required: true
	// Max Length: 256
	// Min Length: 1
	Recipient *string `json:"recipient"`

	// 标题，在支付页面显示
	// Example: 支付订单NO：23435435646
	// Required: true
	// Max Length: 256
	// Min Length: 1
	Title *string `json:"title"`
}

// Validate validates this post create invoice payload
func (m *PostCreateInvoicePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJumpURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifyURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCreateInvoicePayload) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *PostCreateInvoicePayload) validateJumpURL(formats strfmt.Registry) error {

	if err := validate.Required("jump_url", "body", m.JumpURL); err != nil {
		return err
	}

	return nil
}

func (m *PostCreateInvoicePayload) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if err := validate.MinLength("metadata", "body", *m.Metadata, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("metadata", "body", *m.Metadata, 512); err != nil {
		return err
	}

	return nil
}

var postCreateInvoicePayloadTypeNetworkPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ethereum","bsc","polygon","solana","near","goerli"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postCreateInvoicePayloadTypeNetworkPropEnum = append(postCreateInvoicePayloadTypeNetworkPropEnum, v)
	}
}

const (

	// PostCreateInvoicePayloadNetworkEthereum captures enum value "ethereum"
	PostCreateInvoicePayloadNetworkEthereum string = "ethereum"

	// PostCreateInvoicePayloadNetworkBsc captures enum value "bsc"
	PostCreateInvoicePayloadNetworkBsc string = "bsc"

	// PostCreateInvoicePayloadNetworkPolygon captures enum value "polygon"
	PostCreateInvoicePayloadNetworkPolygon string = "polygon"

	// PostCreateInvoicePayloadNetworkSolana captures enum value "solana"
	PostCreateInvoicePayloadNetworkSolana string = "solana"

	// PostCreateInvoicePayloadNetworkNear captures enum value "near"
	PostCreateInvoicePayloadNetworkNear string = "near"

	// PostCreateInvoicePayloadNetworkGoerli captures enum value "goerli"
	PostCreateInvoicePayloadNetworkGoerli string = "goerli"
)

// prop value enum
func (m *PostCreateInvoicePayload) validateNetworkEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postCreateInvoicePayloadTypeNetworkPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostCreateInvoicePayload) validateNetwork(formats strfmt.Registry) error {

	if err := validate.Required("network", "body", m.Network); err != nil {
		return err
	}

	// value enum
	if err := m.validateNetworkEnum("network", "body", *m.Network); err != nil {
		return err
	}

	return nil
}

func (m *PostCreateInvoicePayload) validateNotifyURL(formats strfmt.Registry) error {

	if err := validate.Required("notify_url", "body", m.NotifyURL); err != nil {
		return err
	}

	return nil
}

func (m *PostCreateInvoicePayload) validateRecipient(formats strfmt.Registry) error {

	if err := validate.Required("recipient", "body", m.Recipient); err != nil {
		return err
	}

	if err := validate.MinLength("recipient", "body", *m.Recipient, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("recipient", "body", *m.Recipient, 256); err != nil {
		return err
	}

	return nil
}

func (m *PostCreateInvoicePayload) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post create invoice payload based on context it is used
func (m *PostCreateInvoicePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostCreateInvoicePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCreateInvoicePayload) UnmarshalBinary(b []byte) error {
	var res PostCreateInvoicePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
