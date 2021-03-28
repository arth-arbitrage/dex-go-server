/*
 * Arth-Arbitrage
 *
 * Arth Arbitrage API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package dexes

 import (
	 "fmt"
	 "context"
	 "net/http"
	 "errors"
	 "github.com/sirupsen/logrus"
	 "github.com/arth-arbitrage/dex-go-server/go"	 
	 "github.com/ethereum/go-ethereum/accounts/abi/bind"	
	 "github.com/ethereum/go-ethereum/ethclient"
	 stackdriver "github.com/TV4/logrus-stackdriver-formatter"	 	 	 
 )

 // DexApiService is derived from auto generated DefaultApiService  that implents the logic for the DefaultApiServicer
 // This service should implement the business logic for every endpoint for the DefaultApi API. 
 // Include any external packages or services that will be required by this service.
 type DefaultApiService struct {
	context.Context
	Config Config

	Logger *logrus.Entry
	Transactor *bind.TransactOpts
	client *ethclient.Client
	opts   *bind.TransactOpts
 }

 // NewDefaultApiService creates a default api service
 func NewDefaultApiService(config Config) (restapi.DefaultApiServicer, error) {

	logger := logrus.New()
	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %s", config.LogLevel)
	}
	logger.SetLevel(logLevel)
	logger.SetFormatter(stackdriver.NewFormatter(
		stackdriver.WithService("dexapi"),
	))
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + config.InfuraId)
	if err != nil {
		return nil, err
	}
	apiService := &DefaultApiService{
		Config: config,
		Logger: logrus.NewEntry(logger),
		client: client,
	}
	for _, dex := range dexList{
		err = dex.Init(apiService)
		if err != nil {
			return nil, err
		}
	}
	return apiService, nil
 }
 

 // GetExchange - Exchange details
 func (s *DefaultApiService) GetExchange(ctx context.Context, id int64) (restapi.ImplResponse, error) {

	if(int(id) >= len(dexList)) {
		return restapi.Response(http.StatusNotImplemented, nil), errors.New("Exchange unavailble")
	}
	name := dexList[id].Name(s)
	exchange := &restapi.Exchange{
		Id: id, 
		Name:name,
	}
	return restapi.Response(http.StatusOK, exchange), nil
 }
 
 // GetLender - Get lender
 func (s *DefaultApiService) GetLender(ctx context.Context, id int64) (restapi.ImplResponse, error) {
	 // TODO - update GetLender with the required logic for this service method.
	 // Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
 
	 //TODO: Uncomment the next line to return response Response(200, Lender{}) or use other options such as http.Ok ...
	 //return Response(200, Lender{}), nil
 
	 //TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	 //return Response(0, Error{}), nil
 
	 return restapi.Response(http.StatusNotImplemented, nil), errors.New("GetLender method not implemented")
 }
 
 // GetLenderPools - Pool List
 func (s *DefaultApiService) GetLenderPools(ctx context.Context, id int64) (restapi.ImplResponse, error) {
	 // TODO - update GetLenderPools with the required logic for this service method.
	 // Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
 
	 //TODO: Uncomment the next line to return response Response(200, []LenderPool{}) or use other options such as http.Ok ...
	 //return Response(200, []LenderPool{}), nil
 
	 //TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	 //return Response(0, Error{}), nil
 
	 return restapi.Response(http.StatusNotImplemented, nil), errors.New("GetLenderPools method not implemented")
 }
 
 // GetSwapPools - Pool List
 func (s *DefaultApiService) GetSwapPools(ctx context.Context, id int64) (restapi.ImplResponse, error) {
	
	 return dexList[id].GetSwapPools(s)
 }
 
 // ListExchanges - Get exchange list
 func (s *DefaultApiService) ListExchanges(ctx context.Context) (restapi.ImplResponse, error) {
	exchanges := make([]restapi.Exchange, 0)
	for i, dex := range dexList{
		name := dex.Name(s)
		exchange := &restapi.Exchange{
			Id: int64(i), 
			Name:name,
		}
		exchanges = append(exchanges, *exchange)
	}
	return restapi.Response(http.StatusOK, exchanges), nil
 }
 
 // ListLenders - Lending service list
 func (s *DefaultApiService) ListLenders(ctx context.Context) (restapi.ImplResponse, error) {
	 // TODO - update ListLenders with the required logic for this service method.
	 // Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
 
	 //TODO: Uncomment the next line to return response Response(200, []Lender{}) or use other options such as http.Ok ...
	 //return Response(200, []Lender{}), nil
 
	 //TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	 //return Response(0, Error{}), nil
 
	 return restapi.Response(http.StatusNotImplemented, nil), errors.New("ListLenders method not implemented")
 }
 
 