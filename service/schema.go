package service

import "time"

// Account is a type auto-generated from OpenAPI spec
// stride:generate account
type Account struct {
	// stride:generate id
	ID int64 `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty" field:"id,omitempty" validate:"-"`
	// stride:generate date-of-birth
	DateOfBirth time.Time `json:"date_of_birth,omitempty" xml:"date_of_birth,omitempty" form:"date_of_birth,omitempty" field:"date_of_birth,omitempty" validate:"-"`
	// stride:generate email
	Email string `json:"email" xml:"email" form:"email" field:"email" validate:"required,gt=0"`
	// stride:generate first-name
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty" form:"first_name,omitempty" field:"first_name,omitempty" validate:"gte=0"`
	// stride:generate last-name
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty" form:"last_name,omitempty" field:"last_name,omitempty" validate:"gte=0"`
}

// Accounts is a type auto-generated from OpenAPI spec
// stride:generate accounts
type Accounts []*Account

// Error is a type auto-generated from OpenAPI spec
// stride:generate error
type Error struct {
	// stride:generate code
	Code int32 `json:"code" xml:"code" form:"code" field:"code" validate:"required"`
	// stride:generate message
	Message string `json:"message" xml:"message" form:"message" field:"message" validate:"required,gt=0"`
}
