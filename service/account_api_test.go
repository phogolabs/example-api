package service_test

import (
	"github.com/go-chi/chi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/phogolabs/example-api/service"
)

var _ = Describe("AccountAPI", func() {
	var (
		router chi.Router
		// TODO: Uncomment if you are going to test your code
		// request  *http.Request
		// recorder *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		router = chi.NewRouter()

		controller := &service.CustomerAPI{}
		controller.Mount(router)

		Expect(router.Routes()).NotTo(BeEmpty())

		// TODO: Uncomment if you are going to test your code
		// recorder = httptest.NewRecorder()
	})

	Describe("POST /accounts", func() {
		// TODO: Implement the test cases for CreateAccount operation
	})

	Describe("DELETE /accounts/{account_id}", func() {
		// TODO: Implement the test cases for DeleteAccountByID operation
	})

	Describe("GET /accounts/{account_id}", func() {
		// TODO: Implement the test cases for GetAccountByID operation
	})

	Describe("GET /accounts", func() {
		// TODO: Implement the test cases for GetAllAccounts operation
	})
})
