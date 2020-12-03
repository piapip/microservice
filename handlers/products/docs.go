// Package classification Product API
//
// Documentation for Product API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
// swagger:meta
package handlers

import "github.com/piapip/microservice/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Error returns when there's no matching result
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Error returns when it fails the validation test
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of errors
	// in: body
	Body ValidationError
}

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All current products in the system
	// in: body
	Body []data.Product
}

// Data structure representing a single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Newly created in the system
	// in: body
	Body data.Product
}

// swagger:response noContentResponse
type productsNoContent struct {
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to pass into Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}