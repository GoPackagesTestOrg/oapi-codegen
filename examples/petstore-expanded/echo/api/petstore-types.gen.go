// Package PetstoreExpanded provides primitives to interact the openapi HTTP API.
//
// Code generated by sios.tech/indigo/oapi-codegen DO NOT EDIT.
package api

// Error defines model for Error.
type Error struct {

	// Error code
	Code int32 `json:"code" validate:"required"`

	// Error message
	Message string `json:"message" validate:"required"`
}

// NewPet defines model for NewPet.
type NewPet struct {

	// Name of the pet
	Name string `json:"name" validate:"required"`

	// Type of the pet
	Tag *string `json:"tag,omitempty" validate:"min=2,max=100,regex=^[A-Za-z]+"`
}

// Pet defines model for Pet.
type Pet struct {
	// Embedded struct due to allOf(#/components/schemas/NewPet)
	NewPet
	// Embedded fields due to inline allOf schema

	// Unique id of the pet
	Id int64 `json:"id" validate:"required,min=1,max=100"`
}

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {

	// tags to filter by
	Tags *[]string `json:"tags,omitempty"`

	// maximum number of results to return
	Limit *int32 `json:"limit,omitempty"`
}

// AddPetJSONBody defines parameters for AddPet.
type AddPetJSONBody NewPet

// AddPetRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody AddPetJSONBody