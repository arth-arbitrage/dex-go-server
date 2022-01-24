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
	Config *Config

	Logger *logrus.Entry
	Transactor *bind.TransactOpts
	client *ethclient.Client
	opts   *bind.TransactOpts
 }

 // NewDefaultApiService creates a default api service
 func NewDefaultApiService(config *Config) (restapi.DefaultApiServicer, error) {

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
	for _, lend := range lendList{
		err = lend.Init(apiService)
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
	if(int(id) >= len(lendList)) {
		return restapi.Response(http.StatusNotImplemented, nil), errors.New("Lender unavailble")
	}
	name := lendList[id].Name(s)
	lender := &restapi.Lender{
		Id: id, 
		Name:name,
	}
	return restapi.Response(http.StatusOK, lender), nil
 }
 
 // GetLenderPools - Pool List
 func (s *DefaultApiService) GetLenderPools(ctx context.Context, id int64) (restapi.ImplResponse, error) {
	if(int(id) >= len(lendList)) {
		return restapi.Response(http.StatusNotImplemented, nil), errors.New("Lender unavailble")
	}
	return lendList[id].GetLenderPools(s)
 }
 
 // GetSwapPools - Pool List
 func (s *DefaultApiService) GetSwapPools(ctx context.Context, id int64) (restapi.ImplResponse, error) {
	if(int(id) >= len(dexList)) {
		return restapi.Response(http.StatusNotImplemented, nil), errors.New("Exchange unavailble")
	}
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
	lenders := make([]restapi.Lender, 0)
	for i, lender := range lendList{
		name := lender.Name(s)
		lender := &restapi.Lender{
			Id: int64(i), 
			Name:name,
		}
		lenders = append(lenders, *lender)
	}
	return restapi.Response(http.StatusOK, lenders), nil
 }
 
 // Multiswap - process swap
 func (s *DefaultApiService) Multiswap(ctx context.Context, multiSwapExec restapi.MultiSwapExec) (restapi.ImplResponse, error){
	return Multiswap(s, multiSwapExec)
}

// Swap - process swap
func (s *DefaultApiService) Swap(ctx context.Context, swapExec restapi.SwapExec) (restapi.ImplResponse, error)  {
	return restapi.Response(http.StatusNotImplemented, nil), errors.New("Swap method not implemented")
}

 