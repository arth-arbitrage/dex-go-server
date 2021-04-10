import json
import sys
import os
from subprocess import call

package_folder = "./go-contracts/"
abi_folder = "./contract-abi/"

abi_files = [
    {"input": "AddressProvider.abi", "output": "AddressProvider.go",  "package_name": "curvefi", "type_name": "AddressProvider"},
    {"input": "Registry.abi", "output": "Registry.go",  "package_name": "curvefi", "type_name": "Registry"},
    {"input": "PoolInfo.abi", "output": "PoolInfo.go",  "package_name": "curvefi", "type_name": "PoolInfo"},
    {"input": "Swaps.abi", "output": "Swaps.go",  "package_name": "curvefi", "type_name": "Swaps"},

    {"input": "IUniswapV2Factory.abi", "output": "IUniswapV2Factory.go",  "package_name": "uniswapv2", "type_name": "IUniswapV2Factory"},
    {"input": "IUniswapV2Router02.abi", "output": "IUniswapV2Router02.go", "package_name": "uniswapv2", "type_name": "IUniswapV2Router02"},
    {"input": "IUniswapV2Pair.abi", "output": "IUniswapV2Pair.go",  "package_name": "uniswapv2", "type_name": "IUniswapV2Pair"},
    {"input": "IUniswapv2Swap.abi", "output": "IUniswapv2Swap.go",  "package_name": "uniswapv2", "type_name": "IUniswapv2Swap"},

    {"input": "ILendingPoolAddressesProvider.abi", "output": "ILendingPoolAddressesProvider.go",  "package_name": "aave", "type_name": "ILendingPoolAddressesProvider"},
    {"input": "ILendingPool.abi", "output": "ILendingPool.go", "package_name": "aave", "type_name": "ILendingPool"},   
    {"input": "ArthArbV1MultiSwap.abi", "output": "ArthArbV1MultiSwap.go", "package_name": "arbitrage", "type_name": "ArthArbV1MultiSwap"},
    {"input": "AaveArbMultiSwapV1.abi", "output": "AaveArbMultiSwapV1.go", "package_name": "arbitrage", "type_name": "AaveArbMultiSwapV1"},    
]

def generate_go(input, output, package_name, type_name):
    call(["abigen", "--abi", input, "--pkg", package_name,  "--type", type_name, "--out", output])

for abi_file in abi_files:
    input_file = abi_folder + abi_file["input"]
    out_file = package_folder + abi_file["package_name"] + "/" + abi_file["output"]
    os.makedirs(os.path.dirname(out_file), exist_ok=True)
    generate_go(input_file, out_file, abi_file["package_name"], abi_file["type_name"])
