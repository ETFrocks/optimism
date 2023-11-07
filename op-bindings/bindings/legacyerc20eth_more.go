// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const LegacyERC20ETHStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/legacy/LegacyERC20ETH.sol:LegacyERC20ETH\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/legacy/LegacyERC20ETH.sol:LegacyERC20ETH\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1002,\"contract\":\"src/legacy/LegacyERC20ETH.sol:LegacyERC20ETH\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/legacy/LegacyERC20ETH.sol:LegacyERC20ETH\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":1004,\"contract\":\"src/legacy/LegacyERC20ETH.sol:LegacyERC20ETH\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var LegacyERC20ETHStorageLayout = new(solc.StorageLayout)

var LegacyERC20ETHDeployedBin = "0x608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e14610387578063e78cea921461033b578063ee9a31a2146103cd57600080fd5b8063ae1f6aaf1461033b578063c01e1bd614610361578063d6c0b2c41461036157600080fd5b80639dc29fac116100bd5780639dc29fac14610302578063a457c2d714610315578063a9059cbb1461032857600080fd5b806370a08231146102d257806395d89b41146102fa57600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461026e57806340c10f191461028157806354fd4d501461029657600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a366004610905565b6103f4565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f86104e5565b60405161019b919061094e565b61018f6102133660046109ea565b610577565b6002545b60405190815260200161019b565b61018f610238366004610a14565b610607565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161019b565b61018f61027c3660046109ea565b610692565b61029461028f3660046109ea565b61071d565b005b6101f86040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b61021c6102e0366004610a50565b73ffffffffffffffffffffffffffffffffffffffff163190565b6101f861077f565b6102946103103660046109ea565b61078e565b61018f6103233660046109ea565b6107f0565b61018f6103363660046109ea565b61087b565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c610395366004610a6b565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000085168314806104ad57507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b806104dc57507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104f490610a9e565b80601f016020809104026020016040519081016040528092919081815260200182805461052090610a9e565b801561056d5780601f106105425761010080835404028352916020019161056d565b820191906000526020600020905b81548152906001019060200180831161055057829003601f168201915b5050505050905090565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f4c656761637945524332304554483a20617070726f766520697320646973616260448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526000906084015b60405180910390fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f4c656761637945524332304554483a207472616e7366657246726f6d2069732060448201527f64697361626c656400000000000000000000000000000000000000000000000060648201526000906084016105fe565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4c656761637945524332304554483a20696e637265617365416c6c6f77616e6360448201527f652069732064697361626c65640000000000000000000000000000000000000060648201526000906084016105fe565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4c656761637945524332304554483a206d696e742069732064697361626c656460448201526064016105fe565b6060600480546104f490610a9e565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4c656761637945524332304554483a206275726e2069732064697361626c656460448201526064016105fe565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4c656761637945524332304554483a206465637265617365416c6c6f77616e6360448201527f652069732064697361626c65640000000000000000000000000000000000000060648201526000906084016105fe565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4c656761637945524332304554483a207472616e73666572206973206469736160448201527f626c65640000000000000000000000000000000000000000000000000000000060648201526000906084016105fe565b60006020828403121561091757600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461094757600080fd5b9392505050565b600060208083528351808285015260005b8181101561097b5785810183015185820160400152820161095f565b8181111561098d576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146109e557600080fd5b919050565b600080604083850312156109fd57600080fd5b610a06836109c1565b946020939093013593505050565b600080600060608486031215610a2957600080fd5b610a32846109c1565b9250610a40602085016109c1565b9150604084013590509250925092565b600060208284031215610a6257600080fd5b610947826109c1565b60008060408385031215610a7e57600080fd5b610a87836109c1565b9150610a95602084016109c1565b90509250929050565b600181811c90821680610ab257607f821691505b602082108103610aeb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b5091905056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(LegacyERC20ETHStorageLayoutJSON), LegacyERC20ETHStorageLayout); err != nil {
		panic(err)
	}

	layouts["LegacyERC20ETH"] = LegacyERC20ETHStorageLayout
	deployedBytecodes["LegacyERC20ETH"] = LegacyERC20ETHDeployedBin
}
