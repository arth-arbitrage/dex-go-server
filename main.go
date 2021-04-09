/*
 * Arth-Arbitrage
 *
 * Arth Arbitrage API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	restapi "github.com/arth-arbitrage/dex-go-server/go"
	"github.com/arth-arbitrage/dex-go-server/dexes"
	"github.com/rs/cors"
)

func main() {
	config := dexes.NewConfig() 
	log.Printf("Server started %v", config.InfuraId)

	//DefaultApiService := restapi.NewDefaultApiService()
	DefaultApiService, err := dexes.NewDefaultApiService(config)
	if(err != nil) {
		log.Fatal(err)
	}
	DefaultApiController := restapi.NewDefaultApiController(DefaultApiService)

	router := restapi.NewRouter(DefaultApiController)

	handler := cors.Default().Handler(router)
	//log.Fatal(http.ListenAndServe(":4000", router))
	log.Fatal(http.ListenAndServe(":4000", handler))
}
