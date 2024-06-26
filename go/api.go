/*
 * Arth-Arbitrage
 *
 * Arth Arbitrage API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package restapi

import (
	"context"
	"net/http"
)



// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request, 
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface { 
	GetExchange(http.ResponseWriter, *http.Request)
	GetLender(http.ResponseWriter, *http.Request)
	GetLenderPools(http.ResponseWriter, *http.Request)
	GetSwapPools(http.ResponseWriter, *http.Request)
	ListExchanges(http.ResponseWriter, *http.Request)
	ListLenders(http.ResponseWriter, *http.Request)
	Multiswap(http.ResponseWriter, *http.Request)
	Swap(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type DefaultApiServicer interface { 
	GetExchange(context.Context, int64) (ImplResponse, error)
	GetLender(context.Context, int64) (ImplResponse, error)
	GetLenderPools(context.Context, int64) (ImplResponse, error)
	GetSwapPools(context.Context, int64) (ImplResponse, error)
	ListExchanges(context.Context) (ImplResponse, error)
	ListLenders(context.Context) (ImplResponse, error)
	Multiswap(context.Context, MultiSwapExec) (ImplResponse, error)
	Swap(context.Context, SwapExec) (ImplResponse, error)
}
