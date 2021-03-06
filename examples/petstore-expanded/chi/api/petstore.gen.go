// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/indigonote/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

// Error defines model for Error.
type Error struct {

	// Error code
	Code int32 `json:"code"`

	// Error message
	Message string `json:"message"`
}

// NewPet defines model for NewPet.
type NewPet struct {

	// Name of the pet
	Name *string `json:"name" validate:"omitempty,alphanum,max=1048576"`
	Size int     `json:"size" validate:"min=0,max=20"`

	// Type of the pet
	Tag *string `json:"tag,omitempty" validate:"omitempty,min=2,max=32,regex=^[A-Za-z]+"`
}

// Pet defines model for Pet.
type Pet struct {
	// Embedded struct due to allOf(#/components/schemas/NewPet)
	NewPet
	// Embedded fields due to inline allOf schema

	// Unique id of the pet
	Id int64 `json:"id" validate:"min=1,max=100"`
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

type ServerInterface interface {
	// Returns all pets (GET /pets)
	FindPets(w http.ResponseWriter, r *http.Request)
	// Creates a new pet (POST /pets)
	AddPet(w http.ResponseWriter, r *http.Request)
	// Deletes a pet by ID (DELETE /pets/{id})
	DeletePet(w http.ResponseWriter, r *http.Request)
	// Returns a pet by ID (GET /pets/{id})
	FindPetById(w http.ResponseWriter, r *http.Request)
}

// ParamsForFindPets operation parameters from context
func ParamsForFindPets(ctx context.Context) *FindPetsParams {
	return ctx.Value("FindPetsParams").(*FindPetsParams)
}

// FindPets operation middleware
func FindPetsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// Parameter object where we will unmarshal all parameters from the context
		var params FindPetsParams

		// ------------- Optional query parameter "tags" -------------
		if paramValue := r.URL.Query().Get("tags"); paramValue != "" {

		}

		err = runtime.BindQueryParameter("form", true, false, "tags", r.URL.Query(), &params.Tags)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter tags: %s", err), http.StatusBadRequest)
			return
		}

		// ------------- Optional query parameter "limit" -------------
		if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

		}

		err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter limit: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "FindPetsParams", &params)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AddPet operation middleware
func AddPetCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DeletePet operation middleware
func DeletePetCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// ------------- Path parameter "id" -------------
		var id int64

		err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "id", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// FindPetById operation middleware
func FindPetByIdCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// ------------- Path parameter "id" -------------
		var id int64

		err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "id", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerFromMux(si, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	r.Group(func(r chi.Router) {
		r.Use(FindPetsCtx)
		r.Get("/pets", si.FindPets)
	})
	r.Group(func(r chi.Router) {
		r.Use(AddPetCtx)
		r.Post("/pets", si.AddPet)
	})
	r.Group(func(r chi.Router) {
		r.Use(DeletePetCtx)
		r.Delete("/pets/{id}", si.DeletePet)
	})
	r.Group(func(r chi.Router) {
		r.Use(FindPetByIdCtx)
		r.Get("/pets/{id}", si.FindPetById)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
