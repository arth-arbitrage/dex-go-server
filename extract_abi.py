import json
import sys
import os

contract_folder = "/opt/arth-arbitrage/arth-contract/build/contracts/"
package_folder = "./go-contracts/"
abi_folder = "./contract-abi/"

json_files = [
    {"input": "AddressProvider.json", "output": "AddressProvider.abi"},
    {"input": "Registry.json", "output": "Registry.abi"},
    {"input": "PoolInfo.json", "output": "PoolInfo.abi"},
    {"input": "Swaps.json", "output": "Swaps.abi"},

    {"input": "IUniswapV2Factory.json", "output": "IUniswapV2Factory.abi"},
    {"input": "IUniswapV2Router02.json", "output": "IUniswapV2Router02.abi"},
    {"input": "IUniswapV2Pair.json", "output": "IUniswapV2Pair.abi"},
    {"input": "IUniswapv2Swap.json", "output": "IUniswapv2Swap.abi"},

    {"input": "ILendingPoolAddressesProvider.json", "output": "ILendingPoolAddressesProvider.abi"},
    {"input": "ILendingPool.json", "output": "ILendingPool.abi"},
    {"input": "ArthArbV1MultiSwap.json", "output": "ArthArbV1MultiSwap.abi"},
    {"input": "AaveArbMultiSwapV1.json", "output": "AaveArbMultiSwapV1.abi"},
]

for json_file in json_files:
    abi = None
    input_file = contract_folder + json_file["input"]
    with open(input_file) as fi:
        info_json = json.load(fi)
        abi = info_json["abi"]
    out_file = abi_folder + json_file["output"]
    os.makedirs(os.path.dirname(out_file), exist_ok=True)
    with open(out_file, "w") as fo:
        print(abi)
        json.dump(abi, fo)



