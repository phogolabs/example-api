package service

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/phogolabs/restify"
)

// AccountAPI is a type auto-generated from OpenAPI spec
// stride:generate account-api
type AccountAPI struct {
}

// Mount mounts the controller to the router
// stride:generate account-api:mount
func (x *AccountAPI) Mount(r chi.Router) {
	r.Post("/accounts", x.CreateAccount)
	r.Delete("/accounts/{account_id}", x.DeleteAccountByID)
	r.Get("/accounts/{account_id}", x.GetAccountByID)
	r.Get("/accounts", x.GetAllAccounts)

	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end
}

// CreateAccount handles endpoint POST /accounts
// Create an account
// stride:generate account-api:create-account
func (x *AccountAPI) CreateAccount(w http.ResponseWriter, r *http.Request) {
	reactor := restify.NewReactor(w, r)

	var (
		input  = &CreateAccountInput{}
		output = &CreateAccountOutput{}
	)

	if err := reactor.Bind(input); err != nil {
		reactor.Render(err)
		return
	}

	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end

	if err := reactor.Render(output); err != nil {
		reactor.Render(err)
	}
}

// DeleteAccountByID handles endpoint DELETE /accounts/{account_id}
// Delete an account by id
// stride:generate account-api:delete-account-by-id
func (x *AccountAPI) DeleteAccountByID(w http.ResponseWriter, r *http.Request) {
	reactor := restify.NewReactor(w, r)

	var (
		input  = &DeleteAccountByIDInput{}
		output = &DeleteAccountByIDOutput{}
	)

	if err := reactor.Bind(input); err != nil {
		reactor.Render(err)
		return
	}

	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end

	if err := reactor.Render(output); err != nil {
		reactor.Render(err)
	}
}

// GetAccountByID handles endpoint GET /accounts/{account_id}
// Fetch an account by id
// stride:generate account-api:get-account-by-id
func (x *AccountAPI) GetAccountByID(w http.ResponseWriter, r *http.Request) {
	reactor := restify.NewReactor(w, r)

	var (
		input  = &GetAccountByIDInput{}
		output = &GetAccountByIDOutput{}
	)

	if err := reactor.Bind(input); err != nil {
		reactor.Render(err)
		return
	}

	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end

	if err := reactor.Render(output); err != nil {
		reactor.Render(err)
	}
}

// GetAllAccounts handles endpoint GET /accounts
// List all accounts
// stride:generate account-api:get-all-accounts
func (x *AccountAPI) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	reactor := restify.NewReactor(w, r)

	var (
		input  = &GetAllAccountsInput{}
		output = &GetAllAccountsOutput{}
	)

	if err := reactor.Bind(input); err != nil {
		reactor.Render(err)
		return
	}

	// stride:define body:start
	// NOTE: not implemented
	// stride:define body:end

	if err := reactor.Render(output); err != nil {
		reactor.Render(err)
	}
}
