/*
 * Arth-Arbitrage
 *
 * Arth Arbitrage API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package restapi

type SwapPool struct {

	Id int64 `json:"id"`

	Name string `json:"name"`

	Token0 string `json:"token0"`

	Token1 string `json:"token1"`

	Decimals0 int64 `json:"decimals0"`

	Decimals1 int64 `json:"decimals1"`

	SwapWrap string `json:"swapWrap"`

	Pool string `json:"pool"`

	Token0Address string `json:"token0Address"`

	Token1Address string `json:"token1Address"`

	Reserve0 float32 `json:"reserve0"`

	Reserve1 float32 `json:"reserve1"`

	Price0 float32 `json:"price0"`

	Price1 float32 `json:"price1"`

	Fees float32 `json:"fees"`
}
