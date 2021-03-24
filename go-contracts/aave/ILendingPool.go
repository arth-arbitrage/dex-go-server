// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ILendingPoolABI is the input ABI used to generate the binding from.
const ILendingPoolABI = "[{\"inputs\":[],\"name\":\"addressesProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_referralCode\",\"type\":\"uint16\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveConfigurationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationDiscount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"fixedBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsFixed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsVariable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fixedBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageFixedBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalLiquidityETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentUnderlyingBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_purchaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_receiveAToken\",\"type\":\"bool\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"rebalanceFixedBorrowRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"swapBorrowRateMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ILendingPool is an auto generated Go binding around an Ethereum contract.
type ILendingPool struct {
	ILendingPoolCaller     // Read-only binding to the contract
	ILendingPoolTransactor // Write-only binding to the contract
	ILendingPoolFilterer   // Log filterer for contract events
}

// ILendingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILendingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILendingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILendingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILendingPoolSession struct {
	Contract     *ILendingPool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILendingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILendingPoolCallerSession struct {
	Contract *ILendingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ILendingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILendingPoolTransactorSession struct {
	Contract     *ILendingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ILendingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILendingPoolRaw struct {
	Contract *ILendingPool // Generic contract binding to access the raw methods on
}

// ILendingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILendingPoolCallerRaw struct {
	Contract *ILendingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// ILendingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILendingPoolTransactorRaw struct {
	Contract *ILendingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILendingPool creates a new instance of ILendingPool, bound to a specific deployed contract.
func NewILendingPool(address common.Address, backend bind.ContractBackend) (*ILendingPool, error) {
	contract, err := bindILendingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILendingPool{ILendingPoolCaller: ILendingPoolCaller{contract: contract}, ILendingPoolTransactor: ILendingPoolTransactor{contract: contract}, ILendingPoolFilterer: ILendingPoolFilterer{contract: contract}}, nil
}

// NewILendingPoolCaller creates a new read-only instance of ILendingPool, bound to a specific deployed contract.
func NewILendingPoolCaller(address common.Address, caller bind.ContractCaller) (*ILendingPoolCaller, error) {
	contract, err := bindILendingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolCaller{contract: contract}, nil
}

// NewILendingPoolTransactor creates a new write-only instance of ILendingPool, bound to a specific deployed contract.
func NewILendingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ILendingPoolTransactor, error) {
	contract, err := bindILendingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolTransactor{contract: contract}, nil
}

// NewILendingPoolFilterer creates a new log filterer instance of ILendingPool, bound to a specific deployed contract.
func NewILendingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ILendingPoolFilterer, error) {
	contract, err := bindILendingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolFilterer{contract: contract}, nil
}

// bindILendingPool binds a generic wrapper to an already deployed contract.
func bindILendingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILendingPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILendingPool *ILendingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILendingPool.Contract.ILendingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILendingPool *ILendingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILendingPool.Contract.ILendingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILendingPool *ILendingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILendingPool.Contract.ILendingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILendingPool *ILendingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILendingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILendingPool *ILendingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILendingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILendingPool *ILendingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILendingPool.Contract.contract.Transact(opts, method, params...)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_ILendingPool *ILendingPoolCaller) AddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "addressesProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_ILendingPool *ILendingPoolSession) AddressesProvider() (common.Address, error) {
	return _ILendingPool.Contract.AddressesProvider(&_ILendingPool.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_ILendingPool *ILendingPoolCallerSession) AddressesProvider() (common.Address, error) {
	return _ILendingPool.Contract.AddressesProvider(&_ILendingPool.CallOpts)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationDiscount, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool fixedBorrowRateEnabled, bool isActive)
func (_ILendingPool *ILendingPoolCaller) GetReserveConfigurationData(opts *bind.CallOpts, _reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationDiscount         *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	FixedBorrowRateEnabled      bool
	IsActive                    bool
}, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getReserveConfigurationData", _reserve)

	outstruct := new(struct {
		Ltv                         *big.Int
		LiquidationThreshold        *big.Int
		LiquidationDiscount         *big.Int
		InterestRateStrategyAddress common.Address
		UsageAsCollateralEnabled    bool
		BorrowingEnabled            bool
		FixedBorrowRateEnabled      bool
		IsActive                    bool
	})

	outstruct.Ltv = out[0].(*big.Int)
	outstruct.LiquidationThreshold = out[1].(*big.Int)
	outstruct.LiquidationDiscount = out[2].(*big.Int)
	outstruct.InterestRateStrategyAddress = out[3].(common.Address)
	outstruct.UsageAsCollateralEnabled = out[4].(bool)
	outstruct.BorrowingEnabled = out[5].(bool)
	outstruct.FixedBorrowRateEnabled = out[6].(bool)
	outstruct.IsActive = out[7].(bool)

	return *outstruct, err

}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationDiscount, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool fixedBorrowRateEnabled, bool isActive)
func (_ILendingPool *ILendingPoolSession) GetReserveConfigurationData(_reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationDiscount         *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	FixedBorrowRateEnabled      bool
	IsActive                    bool
}, error) {
	return _ILendingPool.Contract.GetReserveConfigurationData(&_ILendingPool.CallOpts, _reserve)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationDiscount, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool fixedBorrowRateEnabled, bool isActive)
func (_ILendingPool *ILendingPoolCallerSession) GetReserveConfigurationData(_reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationDiscount         *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	FixedBorrowRateEnabled      bool
	IsActive                    bool
}, error) {
	return _ILendingPool.Contract.GetReserveConfigurationData(&_ILendingPool.CallOpts, _reserve)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsFixed, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 fixedBorrowRate, uint256 averageFixedBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_ILendingPool *ILendingPoolCaller) GetReserveData(opts *bind.CallOpts, _reserve common.Address) (struct {
	TotalLiquidity         *big.Int
	AvailableLiquidity     *big.Int
	TotalBorrowsFixed      *big.Int
	TotalBorrowsVariable   *big.Int
	LiquidityRate          *big.Int
	VariableBorrowRate     *big.Int
	FixedBorrowRate        *big.Int
	AverageFixedBorrowRate *big.Int
	UtilizationRate        *big.Int
	LiquidityIndex         *big.Int
	VariableBorrowIndex    *big.Int
	ATokenAddress          common.Address
	LastUpdateTimestamp    *big.Int
}, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getReserveData", _reserve)

	outstruct := new(struct {
		TotalLiquidity         *big.Int
		AvailableLiquidity     *big.Int
		TotalBorrowsFixed      *big.Int
		TotalBorrowsVariable   *big.Int
		LiquidityRate          *big.Int
		VariableBorrowRate     *big.Int
		FixedBorrowRate        *big.Int
		AverageFixedBorrowRate *big.Int
		UtilizationRate        *big.Int
		LiquidityIndex         *big.Int
		VariableBorrowIndex    *big.Int
		ATokenAddress          common.Address
		LastUpdateTimestamp    *big.Int
	})

	outstruct.TotalLiquidity = out[0].(*big.Int)
	outstruct.AvailableLiquidity = out[1].(*big.Int)
	outstruct.TotalBorrowsFixed = out[2].(*big.Int)
	outstruct.TotalBorrowsVariable = out[3].(*big.Int)
	outstruct.LiquidityRate = out[4].(*big.Int)
	outstruct.VariableBorrowRate = out[5].(*big.Int)
	outstruct.FixedBorrowRate = out[6].(*big.Int)
	outstruct.AverageFixedBorrowRate = out[7].(*big.Int)
	outstruct.UtilizationRate = out[8].(*big.Int)
	outstruct.LiquidityIndex = out[9].(*big.Int)
	outstruct.VariableBorrowIndex = out[10].(*big.Int)
	outstruct.ATokenAddress = out[11].(common.Address)
	outstruct.LastUpdateTimestamp = out[12].(*big.Int)

	return *outstruct, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsFixed, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 fixedBorrowRate, uint256 averageFixedBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_ILendingPool *ILendingPoolSession) GetReserveData(_reserve common.Address) (struct {
	TotalLiquidity         *big.Int
	AvailableLiquidity     *big.Int
	TotalBorrowsFixed      *big.Int
	TotalBorrowsVariable   *big.Int
	LiquidityRate          *big.Int
	VariableBorrowRate     *big.Int
	FixedBorrowRate        *big.Int
	AverageFixedBorrowRate *big.Int
	UtilizationRate        *big.Int
	LiquidityIndex         *big.Int
	VariableBorrowIndex    *big.Int
	ATokenAddress          common.Address
	LastUpdateTimestamp    *big.Int
}, error) {
	return _ILendingPool.Contract.GetReserveData(&_ILendingPool.CallOpts, _reserve)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsFixed, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 fixedBorrowRate, uint256 averageFixedBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_ILendingPool *ILendingPoolCallerSession) GetReserveData(_reserve common.Address) (struct {
	TotalLiquidity         *big.Int
	AvailableLiquidity     *big.Int
	TotalBorrowsFixed      *big.Int
	TotalBorrowsVariable   *big.Int
	LiquidityRate          *big.Int
	VariableBorrowRate     *big.Int
	FixedBorrowRate        *big.Int
	AverageFixedBorrowRate *big.Int
	UtilizationRate        *big.Int
	LiquidityIndex         *big.Int
	VariableBorrowIndex    *big.Int
	ATokenAddress          common.Address
	LastUpdateTimestamp    *big.Int
}, error) {
	return _ILendingPool.Contract.GetReserveData(&_ILendingPool.CallOpts, _reserve)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns()
func (_ILendingPool *ILendingPoolCaller) GetReserves(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getReserves")

	if err != nil {
		return err
	}

	return err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns()
func (_ILendingPool *ILendingPoolSession) GetReserves() error {
	return _ILendingPool.Contract.GetReserves(&_ILendingPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns()
func (_ILendingPool *ILendingPoolCallerSession) GetReserves() error {
	return _ILendingPool.Contract.GetReserves(&_ILendingPool.CallOpts)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_ILendingPool *ILendingPoolCaller) GetUserAccountData(opts *bind.CallOpts, _user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getUserAccountData", _user)

	outstruct := new(struct {
		TotalLiquidityETH           *big.Int
		TotalCollateralETH          *big.Int
		TotalBorrowsETH             *big.Int
		AvailableBorrowsETH         *big.Int
		CurrentLiquidationThreshold *big.Int
		Ltv                         *big.Int
		HealthFactor                *big.Int
	})

	outstruct.TotalLiquidityETH = out[0].(*big.Int)
	outstruct.TotalCollateralETH = out[1].(*big.Int)
	outstruct.TotalBorrowsETH = out[2].(*big.Int)
	outstruct.AvailableBorrowsETH = out[3].(*big.Int)
	outstruct.CurrentLiquidationThreshold = out[4].(*big.Int)
	outstruct.Ltv = out[5].(*big.Int)
	outstruct.HealthFactor = out[6].(*big.Int)

	return *outstruct, err

}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_ILendingPool *ILendingPoolSession) GetUserAccountData(_user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _ILendingPool.Contract.GetUserAccountData(&_ILendingPool.CallOpts, _user)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_ILendingPool *ILendingPoolCallerSession) GetUserAccountData(_user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _ILendingPool.Contract.GetUserAccountData(&_ILendingPool.CallOpts, _user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentUnderlyingBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_ILendingPool *ILendingPoolCaller) GetUserReserveData(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentUnderlyingBalance *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getUserReserveData", _reserve, _user)

	outstruct := new(struct {
		CurrentATokenBalance     *big.Int
		CurrentUnderlyingBalance *big.Int
		CurrentBorrowBalance     *big.Int
		PrincipalBorrowBalance   *big.Int
		BorrowRateMode           *big.Int
		BorrowRate               *big.Int
		LiquidityRate            *big.Int
		OriginationFee           *big.Int
		VariableBorrowIndex      *big.Int
		LastUpdateTimestamp      *big.Int
		UsageAsCollateralEnabled bool
	})

	outstruct.CurrentATokenBalance = out[0].(*big.Int)
	outstruct.CurrentUnderlyingBalance = out[1].(*big.Int)
	outstruct.CurrentBorrowBalance = out[2].(*big.Int)
	outstruct.PrincipalBorrowBalance = out[3].(*big.Int)
	outstruct.BorrowRateMode = out[4].(*big.Int)
	outstruct.BorrowRate = out[5].(*big.Int)
	outstruct.LiquidityRate = out[6].(*big.Int)
	outstruct.OriginationFee = out[7].(*big.Int)
	outstruct.VariableBorrowIndex = out[8].(*big.Int)
	outstruct.LastUpdateTimestamp = out[9].(*big.Int)
	outstruct.UsageAsCollateralEnabled = out[10].(bool)

	return *outstruct, err

}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentUnderlyingBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_ILendingPool *ILendingPoolSession) GetUserReserveData(_reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentUnderlyingBalance *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _ILendingPool.Contract.GetUserReserveData(&_ILendingPool.CallOpts, _reserve, _user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentUnderlyingBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_ILendingPool *ILendingPoolCallerSession) GetUserReserveData(_reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentUnderlyingBalance *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _ILendingPool.Contract.GetUserReserveData(&_ILendingPool.CallOpts, _reserve, _user)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_ILendingPool *ILendingPoolTransactor) Borrow(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "borrow", _reserve, _amount, _interestRateMode, _referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_ILendingPool *ILendingPoolSession) Borrow(_reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.Contract.Borrow(&_ILendingPool.TransactOpts, _reserve, _amount, _interestRateMode, _referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_ILendingPool *ILendingPoolTransactorSession) Borrow(_reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.Contract.Borrow(&_ILendingPool.TransactOpts, _reserve, _amount, _interestRateMode, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_ILendingPool *ILendingPoolTransactor) Deposit(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "deposit", _reserve, _amount, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_ILendingPool *ILendingPoolSession) Deposit(_reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.Contract.Deposit(&_ILendingPool.TransactOpts, _reserve, _amount, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_ILendingPool *ILendingPoolTransactorSession) Deposit(_reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.Contract.Deposit(&_ILendingPool.TransactOpts, _reserve, _amount, _referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_ILendingPool *ILendingPoolTransactor) FlashLoan(opts *bind.TransactOpts, _receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "flashLoan", _receiver, _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_ILendingPool *ILendingPoolSession) FlashLoan(_receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _ILendingPool.Contract.FlashLoan(&_ILendingPool.TransactOpts, _receiver, _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_ILendingPool *ILendingPoolTransactorSession) FlashLoan(_receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _ILendingPool.Contract.FlashLoan(&_ILendingPool.TransactOpts, _receiver, _reserve, _amount, _params)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_ILendingPool *ILendingPoolTransactor) LiquidationCall(opts *bind.TransactOpts, _collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "liquidationCall", _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_ILendingPool *ILendingPoolSession) LiquidationCall(_collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _ILendingPool.Contract.LiquidationCall(&_ILendingPool.TransactOpts, _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_ILendingPool *ILendingPoolTransactorSession) LiquidationCall(_collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _ILendingPool.Contract.LiquidationCall(&_ILendingPool.TransactOpts, _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// RebalanceFixedBorrowRate is a paid mutator transaction binding the contract method 0x66f8cd93.
//
// Solidity: function rebalanceFixedBorrowRate(address _reserve, address _user) returns()
func (_ILendingPool *ILendingPoolTransactor) RebalanceFixedBorrowRate(opts *bind.TransactOpts, _reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "rebalanceFixedBorrowRate", _reserve, _user)
}

// RebalanceFixedBorrowRate is a paid mutator transaction binding the contract method 0x66f8cd93.
//
// Solidity: function rebalanceFixedBorrowRate(address _reserve, address _user) returns()
func (_ILendingPool *ILendingPoolSession) RebalanceFixedBorrowRate(_reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.RebalanceFixedBorrowRate(&_ILendingPool.TransactOpts, _reserve, _user)
}

// RebalanceFixedBorrowRate is a paid mutator transaction binding the contract method 0x66f8cd93.
//
// Solidity: function rebalanceFixedBorrowRate(address _reserve, address _user) returns()
func (_ILendingPool *ILendingPoolTransactorSession) RebalanceFixedBorrowRate(_reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.RebalanceFixedBorrowRate(&_ILendingPool.TransactOpts, _reserve, _user)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x935642cf.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount) returns()
func (_ILendingPool *ILendingPoolTransactor) RedeemUnderlying(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "redeemUnderlying", _reserve, _user, _amount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x935642cf.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount) returns()
func (_ILendingPool *ILendingPoolSession) RedeemUnderlying(_reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ILendingPool.Contract.RedeemUnderlying(&_ILendingPool.TransactOpts, _reserve, _user, _amount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x935642cf.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount) returns()
func (_ILendingPool *ILendingPoolTransactorSession) RedeemUnderlying(_reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ILendingPool.Contract.RedeemUnderlying(&_ILendingPool.TransactOpts, _reserve, _user, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_ILendingPool *ILendingPoolTransactor) Repay(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "repay", _reserve, _amount, _onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_ILendingPool *ILendingPoolSession) Repay(_reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.Repay(&_ILendingPool.TransactOpts, _reserve, _amount, _onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_ILendingPool *ILendingPoolTransactorSession) Repay(_reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.Repay(&_ILendingPool.TransactOpts, _reserve, _amount, _onBehalfOf)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_ILendingPool *ILendingPoolTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, _reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "setUserUseReserveAsCollateral", _reserve, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_ILendingPool *ILendingPoolSession) SetUserUseReserveAsCollateral(_reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _ILendingPool.Contract.SetUserUseReserveAsCollateral(&_ILendingPool.TransactOpts, _reserve, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_ILendingPool *ILendingPoolTransactorSession) SetUserUseReserveAsCollateral(_reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _ILendingPool.Contract.SetUserUseReserveAsCollateral(&_ILendingPool.TransactOpts, _reserve, _useAsCollateral)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_ILendingPool *ILendingPoolTransactor) SwapBorrowRateMode(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "swapBorrowRateMode", _reserve)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_ILendingPool *ILendingPoolSession) SwapBorrowRateMode(_reserve common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.SwapBorrowRateMode(&_ILendingPool.TransactOpts, _reserve)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_ILendingPool *ILendingPoolTransactorSession) SwapBorrowRateMode(_reserve common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.SwapBorrowRateMode(&_ILendingPool.TransactOpts, _reserve)
}
