/*
 * Arth-Arbitrage
 *
 * Arth Arbitrage API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package restapi

type Swap struct {

	Token0 string `json:"token0"`

	Token1 string `json:"token1"`

	Amount float32 `json:"amount"`

	Exchange int64 `json:"exchange"`
}
