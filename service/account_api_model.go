package service

import (
	"encoding/json"
	"encoding/xml"
)

// CreateAccountInput is a type auto-generated from OpenAPI spec
// It is the input of CreateAccount operation
// stride:generate create-account-input
type CreateAccountInput struct {
}

// CreateAccountInternalServerErrorOutput is a type auto-generated from OpenAPI spec
// It is the output of CreateAccount operation with code: 500
// stride:generate create-account-internal-server-error-output
type CreateAccountInternalServerErrorOutput struct {
	// stride:generate body
	Body *Error `body:"~"`
}

// MarshalJSON marshals into valid JSON
// stride:generate create-account-internal-server-error-output:marshal-j-s-o-n
func (x CreateAccountInternalServerErrorOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Body)
}

// MarshalXML marshals into valid XML
// stride:generate create-account-internal-server-error-output:marshal-x-m-l
func (x CreateAccountInternalServerErrorOutput) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(x.Body, start)
}

// Status returns an http status code
// stride:generate create-account-internal-server-error-output:status
func (x *CreateAccountInternalServerErrorOutput) Status() int {
	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end
	return 500
}

// CreateAccountCreatedOutput is a type auto-generated from OpenAPI spec
// It is the output of CreateAccount operation with code: 201
// stride:generate create-account-created-output
type CreateAccountCreatedOutput struct {
}

// Status returns an http status code
// stride:generate create-account-created-output:status
func (x *CreateAccountCreatedOutput) Status() int {
	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end
	return 201
}

// CreateAccountOutput is a type auto-generated from OpenAPI spec
// It is the alias to the default output of CreateAccount operation
// stride:generate create-account-output
type CreateAccountOutput CreateAccountCreatedOutput

// DeleteAccountByIDInput is a type auto-generated from OpenAPI spec
// It is the input of DeleteAccountByID operation
// stride:generate delete-account-by-id-input
type DeleteAccountByIDInput struct {
	// stride:generate path
	Path *DeleteAccountByIDInputPath `path:"~"`
}

// DeleteAccountByIDInputPath is a type auto-generated from OpenAPI spec
// It is the path of DeleteAccountByIDInput
// stride:generate delete-account-by-id-input-path
type DeleteAccountByIDInputPath struct {
	// stride:generate account-id
	AccountID string `path:"account_id,simple" validate:"required,gt=0"`
}

// DeleteAccountByIDInternalServerErrorOutput is a type auto-generated from OpenAPI spec
// It is the output of DeleteAccountByID operation with code: 500
// stride:generate delete-account-by-id-internal-server-error-output
type DeleteAccountByIDInternalServerErrorOutput struct {
	// stride:generate body
	Body *Error `body:"~"`
}

// MarshalJSON marshals into valid JSON
// stride:generate delete-account-by-id-internal-server-error-output:marshal-j-s-o-n
func (x DeleteAccountByIDInternalServerErrorOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Body)
}

// MarshalXML marshals into valid XML
// stride:generate delete-account-by-id-internal-server-error-output:marshal-x-m-l
func (x DeleteAccountByIDInternalServerErrorOutput) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(x.Body, start)
}

// Status returns an http status code
// stride:generate delete-account-by-id-internal-server-error-output:status
func (x *DeleteAccountByIDInternalServerErrorOutput) Status() int {
	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end
	return 500
}

// DeleteAccountByIDOKOutput is a type auto-generated from OpenAPI spec
// It is the output of DeleteAccountByID operation with code: 200
// stride:generate delete-account-by-id-ok-output
type DeleteAccountByIDOKOutput struct {
}

// Status returns an http status code
// stride:generate delete-account-by-id-ok-output:status
func (x *DeleteAccountByIDOKOutput) Status() int {
	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end
	return 200
}

// DeleteAccountByIDOutput is a type auto-generated from OpenAPI spec
// It is the alias to the default output of DeleteAccountByID operation
// stride:generate delete-account-by-id-output
type DeleteAccountByIDOutput DeleteAccountByIDOKOutput

// GetAccountByIDInput is a type auto-generated from OpenAPI spec
// It is the input of GetAccountByID operation
// stride:generate get-account-by-id-input
type GetAccountByIDInput struct {
	// stride:generate path
	Path *GetAccountByIDInputPath `path:"~"`
}

// GetAccountByIDInputPath is a type auto-generated from OpenAPI spec
// It is the path of GetAccountByIDInput
// stride:generate get-account-by-id-input-path
type GetAccountByIDInputPath struct {
	// stride:generate account-id
	AccountID string `path:"account_id,simple" validate:"required,gt=0"`
}

// GetAccountByIDNotFoundOutput is a type auto-generated from OpenAPI spec
// It is the output of GetAccountByID operation with code: 404
// stride:generate get-account-by-id-not-found-output
type GetAccountByIDNotFoundOutput struct {
	// stride:generate body
	Body *Error `body:"~"`
}

// MarshalJSON marshals into valid JSON
// stride:generate get-account-by-id-not-found-output:marshal-j-s-o-n
func (x GetAccountByIDNotFoundOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Body)
}

// MarshalXML marshals into valid XML
// stride:generate get-account-by-id-not-found-output:marshal-x-m-l
func (x GetAccountByIDNotFoundOutput) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(x.Body, start)
}

// Status returns an http status code
// stride:generate get-account-by-id-not-found-output:status
func (x *GetAccountByIDNotFoundOutput) Status() int {
	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end
	return 404
}

// GetAccountByIDInternalServerErrorOutput is a type auto-generated from OpenAPI spec
// It is the output of GetAccountByID operation with code: 500
// stride:generate get-account-by-id-internal-server-error-output
type GetAccountByIDInternalServerErrorOutput struct {
	// stride:generate body
	Body *Error `body:"~"`
}

// MarshalJSON marshals into valid JSON
// stride:generate get-account-by-id-internal-server-error-output:marshal-j-s-o-n
func (x GetAccountByIDInternalServerErrorOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Body)
}

// MarshalXML marshals into valid XML
// stride:generate get-account-by-id-internal-server-error-output:marshal-x-m-l
func (x GetAccountByIDInternalServerErrorOutput) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(x.Body, start)
}

// Status returns an http status code
// stride:generate get-account-by-id-internal-server-error-output:status
func (x *GetAccountByIDInternalServerErrorOutput) Status() int {
	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end
	return 500
}

// GetAllAccountsInput is a type auto-generated from OpenAPI spec
// It is the input of GetAllAccounts operation
// stride:generate get-all-accounts-input
type GetAllAccountsInput struct {
	// stride:generate query
	Query *GetAllAccountsInputQuery `query:"~"`
}

// GetAllAccountsInputQuery is a type auto-generated from OpenAPI spec
// It is the query of GetAllAccountsInput
// stride:generate get-all-accounts-input-query
type GetAllAccountsInputQuery struct {
	// stride:generate limit
	Limit int32 `query:"limit,form,explode" validate:"-"`
}

// GetAllAccountsOKOutput is a type auto-generated from OpenAPI spec
// It is the output of GetAllAccounts operation with code: 200
// stride:generate get-all-accounts-ok-output
type GetAllAccountsOKOutput struct {
	// stride:generate body
	Body *Account `body:"~"`
}

// MarshalJSON marshals into valid JSON
// stride:generate get-all-accounts-ok-output:marshal-j-s-o-n
func (x GetAllAccountsOKOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Body)
}

// MarshalXML marshals into valid XML
// stride:generate get-all-accounts-ok-output:marshal-x-m-l
func (x GetAllAccountsOKOutput) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(x.Body, start)
}

// Status returns an http status code
// stride:generate get-all-accounts-ok-output:status
func (x *GetAllAccountsOKOutput) Status() int {
	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end
	return 200
}

// GetAllAccountsOutput is a type auto-generated from OpenAPI spec
// It is the alias to the default output of GetAllAccounts operation
// stride:generate get-all-accounts-output
type GetAllAccountsOutput GetAllAccountsOKOutput

// GetAllAccountsInternalServerErrorOutput is a type auto-generated from OpenAPI spec
// It is the output of GetAllAccounts operation with code: 500
// stride:generate get-all-accounts-internal-server-error-output
type GetAllAccountsInternalServerErrorOutput struct {
	// stride:generate body
	Body *Error `body:"~"`
}

// MarshalJSON marshals into valid JSON
// stride:generate get-all-accounts-internal-server-error-output:marshal-j-s-o-n
func (x GetAllAccountsInternalServerErrorOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Body)
}

// MarshalXML marshals into valid XML
// stride:generate get-all-accounts-internal-server-error-output:marshal-x-m-l
func (x GetAllAccountsInternalServerErrorOutput) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(x.Body, start)
}

// Status returns an http status code
// stride:generate get-all-accounts-internal-server-error-output:status
func (x *GetAllAccountsInternalServerErrorOutput) Status() int {
	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end
	return 500
}
