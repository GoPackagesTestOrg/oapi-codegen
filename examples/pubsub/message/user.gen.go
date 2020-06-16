// Package message provides primitives to interact the openapi HTTP API.
//
// Code generated by sios.tech/indigo/oapi-codegen DO NOT EDIT.
package message

// User defines model for User.
type User struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"min=1,max=13,required"`
}
