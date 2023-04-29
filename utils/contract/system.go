package contract

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Config struct {
	Address common.Address
	Code    []byte
	Storage map[common.Hash]common.Hash
}

var (
	NativeTokenAdderContract = Config{
		Address: common.HexToAddress("0000000000000000000000000000000000001000"),
		Code:    common.Hex2Bytes(`608060405234801561001057600080fd5b50600436106100575760003560e01c806313a070b71461005c57806340be5df0146100785780638da5cb5b14610094578063f2fde38b146100b2578063f6d7d88a146100ce575b600080fd5b6100766004803603810190610071919061051e565b6100ea565b005b610092600480360381019061008d919061051e565b6101b3565b005b61009c61027b565b6040516100a9919061055a565b60405180910390f35b6100cc60048036038101906100c7919061051e565b61029f565b005b6100e860048036038101906100e391906105ab565b6103f1565b005b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460011461016c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016390610648565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600114610235576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022c90610648565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905550565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461032d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610324906106b4565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600114610473576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046a90610648565b60405180910390fd5b808273ffffffffffffffffffffffffffffffffffffffff167f635e06dffef29c1697e96c78c75aff20ea871ce1e4fb5a444b5e1fa692afe79f60405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104eb826104c0565b9050919050565b6104fb816104e0565b811461050657600080fd5b50565b600081359050610518816104f2565b92915050565b600060208284031215610534576105336104bb565b5b600061054284828501610509565b91505092915050565b610554816104e0565b82525050565b600060208201905061056f600083018461054b565b92915050565b6000819050919050565b61058881610575565b811461059357600080fd5b50565b6000813590506105a58161057f565b92915050565b600080604083850312156105c2576105c16104bb565b5b60006105d085828601610509565b92505060206105e185828601610596565b9150509250929050565b600082825260208201905092915050565b7f63616c6c206e6f742061646d696e000000000000000000000000000000000000600082015250565b6000610632600e836105eb565b915061063d826105fc565b602082019050919050565b6000602082019050818103600083015261066181610625565b9050919050565b7f63616c6c206e6f74206f776e6572000000000000000000000000000000000000600082015250565b600061069e600e836105eb565b91506106a982610668565b602082019050919050565b600060208201905081810360008301526106cd81610691565b905091905056fea26469706673582212207cf8c90c444cd226160eae730c6256004c16fc8a3c707a11eb8777870ca5a35164736f6c63430008110033`),
		Storage: map[common.Hash]common.Hash{
			common.BigToHash(big.NewInt(0)): common.HexToHash("0x771d63a1d58Eb53c874Fb87475cC9eb1Cb1F5d2d"),
		},
	}
	NativeTokenBurnerContract = Config{
		Address: common.HexToAddress("0000000000000000000000000000000000001001"),
		Code:    common.Hex2Bytes(`608060405260043610601c5760003560e01c80639fc845ef146021575b600080fd5b60276029565b005b600073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015606f573d6000803e3d6000fd5b50343373ffffffffffffffffffffffffffffffffffffffff167ff6152bc59dfebb6482c2f5f3a17b81bc3fc8a8eeafc964768b17f77c1b12409c60405160405180910390a356fea26469706673582212209fa87adbf1e8c77764713c4108049842571910e962882ffb09051af6ec778c9964736f6c63430008110033`),
	}

	ZeroFeeContract = Config{
		Address: common.HexToAddress("0000000000000000000000000000000000001002"),
		Code:    common.Hex2Bytes("608060405234801561001057600080fd5b50600436106100575760003560e01c80631324921b1461005c57806313a070b71461007857806340be5df01461009457806350be45bb146100b05780638da5cb5b146100cc575b600080fd5b610076600480360381019061007191906106c6565b6100ea565b005b610092600480360381019061008d91906106c6565b610366565b005b6100ae60048036038101906100a991906106c6565b61043d565b005b6100ca60048036038101906100c591906106c6565b610512565b005b6100d461063d565b6040516100e19190610702565b60405180910390f35b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460011461016b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101629061077a565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009055600060016002805490506101c291906107d3565b905060005b600280549050811015610260578273ffffffffffffffffffffffffffffffffffffffff16600282815481106101ff576101fe610807565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361024d57809150610260565b808061025890610836565b9150506101c7565b506002805490508110156103625760026001600280549050610282919061087e565b8154811061029357610292610807565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600282815481106102d2576102d1610807565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600280548061032c5761032b6108b2565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590555b5050565b3373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146103f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ed9061092d565b60405180910390fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b3373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146104cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c49061092d565b60405180910390fd5b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905550565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600114610593576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058a9061077a565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506002819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061069382610668565b9050919050565b6106a381610688565b81146106ae57600080fd5b50565b6000813590506106c08161069a565b92915050565b6000602082840312156106dc576106db610663565b5b60006106ea848285016106b1565b91505092915050565b6106fc81610688565b82525050565b600060208201905061071760008301846106f3565b92915050565b600082825260208201905092915050565b7f6e6f742061646d696e0000000000000000000000000000000000000000000000600082015250565b600061076460098361071d565b915061076f8261072e565b602082019050919050565b6000602082019050818103600083015261079381610757565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006107de8261079a565b91506107e98361079a565b9250828201905080821115610801576108006107a4565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006108418261079a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610873576108726107a4565b5b600182019050919050565b60006108898261079a565b91506108948361079a565b92508282039050818111156108ac576108ab6107a4565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f6e6f74206f776e65720000000000000000000000000000000000000000000000600082015250565b600061091760098361071d565b9150610922826108e1565b602082019050919050565b600060208201905081810360008301526109468161090a565b905091905056fea26469706673582212208fe33b11404123a95a2b1c9690fd5006220536b8e8a2716e83dd7a4f5abf8ce664736f6c63430008110033"),
		Storage: map[common.Hash]common.Hash{
			common.BigToHash(big.NewInt(3)): common.HexToHash("0x771d63a1d58Eb53c874Fb87475cC9eb1Cb1F5d2d"),
		},
	}

	zeroFeeAddFuncID     = common.Hex2Bytes("50be45bb") // add_zero(address)
	zeroFeeRemoveFuncID  = common.Hex2Bytes("1324921b") // remove_zero(address)
	mintNativeTokenFunID = common.Hex2Bytes("f6d7d88a") // mintNativeToken(address,uint256)
)

func Constructor(statedb StateDB) {
	// NativeTokenAdder
	statedb.SetCode(NativeTokenAdderContract.Address, NativeTokenAdderContract.Code)
	for key, value := range NativeTokenAdderContract.Storage {
		statedb.SetState(NativeTokenAdderContract.Address, key, value)
	}

	// NativeTokenBurner
	statedb.SetCode(NativeTokenBurnerContract.Address, NativeTokenBurnerContract.Code)
	for key, value := range NativeTokenBurnerContract.Storage {
		statedb.SetState(NativeTokenBurnerContract.Address, key, value)
	}

	// add zero fee
	statedb.SetCode(ZeroFeeContract.Address, ZeroFeeContract.Code)
	for key, value := range ZeroFeeContract.Storage {
		statedb.SetState(ZeroFeeContract.Address, key, value)
	}
}

var (
	zeroGasLists []common.Address
)

func HandleSystemContract(state StateDB, caller common.Address, contract common.Address, input []byte) {
	switch {
	case contract == NativeTokenAdderContract.Address && bytes.HasPrefix(input, mintNativeTokenFunID) && len(input) == 68:
		user := common.BytesToAddress(input[4:36])
		amount := big.NewInt(0).SetBytes(input[36:])
		state.AddBalance(user, amount)
	case contract == ZeroFeeContract.Address && bytes.HasPrefix(input, zeroFeeAddFuncID) && len(input) == 36:
		QueryZeroGasFee(state)
	case contract == ZeroFeeContract.Address && bytes.HasPrefix(input, zeroFeeRemoveFuncID) && len(input) == 36:
		QueryZeroGasFee(state)
	}
}

func ExistZeroGasFee(state StateDB, addr common.Address) bool {
	ks := crypto.NewKeccakState()
	key := append([]byte{}, addr.Hash().Bytes()...)
	key = append(key, common.Big1Hash.Bytes()...)
	ks.Reset()
	ks.Write(key[:])
	skey := ks.Sum(nil)
	return state.GetState(ZeroFeeContract.Address, common.BytesToHash(skey)) == common.Big1Hash
}

func QueryZeroGasFee(state StateDB) {
	alen := state.GetState(ZeroFeeContract.Address, common.Big2Hash).Big().Int64()
	if alen <= 0 {
		return
	}

	ks := crypto.NewKeccakState()
	// read array length
	ks.Reset()
	ks.Write(common.Big2Hash.Bytes())
	skey := ks.Sum(nil)
	start := common.BytesToHash(skey).Big()

	zeroGasLists = make([]common.Address, alen, alen)

	var index int64
	for index = 0; index < alen; index++ {
		key := big.NewInt(0).Add(start, big.NewInt(index))
		ah := state.GetState(ZeroFeeContract.Address, common.BigToHash(key))
		zeroGasLists[index] = common.BigToAddress(ah.Big())
	}
}

func ListZeroGas() []common.Address {
	return zeroGasLists
}
