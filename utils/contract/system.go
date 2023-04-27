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
		Code:    common.Hex2Bytes(`608060405234801561001057600080fd5b50600436106100415760003560e01c80638da5cb5b14610046578063f2fde38b14610064578063f6d7d88a14610080575b600080fd5b61004e61009c565b60405161005b9190610329565b60405180910390f35b61007e60048036038101906100799190610375565b6100c0565b005b61009a600480360381019061009591906103d8565b610212565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461014e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161014590610475565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029790610475565b60405180910390fd5b808273ffffffffffffffffffffffffffffffffffffffff167f635e06dffef29c1697e96c78c75aff20ea871ce1e4fb5a444b5e1fa692afe79f60405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610313826102e8565b9050919050565b61032381610308565b82525050565b600060208201905061033e600083018461031a565b92915050565b600080fd5b61035281610308565b811461035d57600080fd5b50565b60008135905061036f81610349565b92915050565b60006020828403121561038b5761038a610344565b5b600061039984828501610360565b91505092915050565b6000819050919050565b6103b5816103a2565b81146103c057600080fd5b50565b6000813590506103d2816103ac565b92915050565b600080604083850312156103ef576103ee610344565b5b60006103fd85828601610360565b925050602061040e858286016103c3565b9150509250929050565b600082825260208201905092915050565b7f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000600082015250565b600061045f601783610418565b915061046a82610429565b602082019050919050565b6000602082019050818103600083015261048e81610452565b905091905056fea26469706673582212208bd1e409c064c5f446c93290b5054d0fd779feb5a1c8230fbcd603fb90cd276f64736f6c63430008110033`),
		Storage: map[common.Hash]common.Hash{
			common.BigToHash(big.NewInt(0)): common.HexToHash("0x000000000000000000000000771d63a1d58Eb53c874Fb87475cC9eb1Cb1F5d2d"),
		},
	}
	NativeTokenBurnerContract = Config{
		Address: common.HexToAddress("0000000000000000000000000000000000001001"),
		Code:    common.Hex2Bytes(`608060405260043610601c5760003560e01c80639fc845ef146021575b600080fd5b60276029565b005b600073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015606f573d6000803e3d6000fd5b50343373ffffffffffffffffffffffffffffffffffffffff167ff6152bc59dfebb6482c2f5f3a17b81bc3fc8a8eeafc964768b17f77c1b12409c60405160405180910390a356fea26469706673582212209fa87adbf1e8c77764713c4108049842571910e962882ffb09051af6ec778c9964736f6c63430008110033`),
	}

	ZeroFeeContract = Config{
		Address: common.HexToAddress("0000000000000000000000000000000000001002"),
		Code:    common.Hex2Bytes("608060405234801561001057600080fd5b506004361061002b5760003560e01c806350be45bb14610030575b600080fd5b61004a600480360381019061004591906101da565b61004c565b005b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546001146100cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100c490610264565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506002819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101a78261017c565b9050919050565b6101b78161019c565b81146101c257600080fd5b50565b6000813590506101d4816101ae565b92915050565b6000602082840312156101f0576101ef610177565b5b60006101fe848285016101c5565b91505092915050565b600082825260208201905092915050565b7f63616c6c206e6f742061646d696e000000000000000000000000000000000000600082015250565b600061024e600e83610207565b915061025982610218565b602082019050919050565b6000602082019050818103600083015261027d81610241565b905091905056fea26469706673582212206ecb4f40716f2d7d92d1a1877bcc16a606cd75b991cd2b76bfc358f6e772bf7064736f6c63430008110033"),
		Storage: map[common.Hash]common.Hash{},
	}

	zeroFeeAddFuncID     = common.Hex2Bytes("50be45bb") // add_zero(address)
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
	zeroGasLists = make(map[common.Address]struct{})
)

func HandleSystemContract(state StateDB, caller common.Address, contract common.Address, input []byte) {
	queryZeroGasFee(state)
	switch {
	case contract == NativeTokenAdderContract.Address && bytes.HasPrefix(input, mintNativeTokenFunID) && len(input) == 68:
		user := common.BytesToAddress(input[4:36])
		amount := big.NewInt(0).SetBytes(input[36:])
		state.AddBalance(user, amount)
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

func queryZeroGasFee(state StateDB) {
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

	var index int64
	for index = 0; index < alen; index++ {
		key := big.NewInt(0).Add(start, big.NewInt(index))
		ah := state.GetState(ZeroFeeContract.Address, common.BigToHash(key))
		zeroGasLists[common.BigToAddress(ah.Big())] = struct{}{}
	}
}

func ListZeroGas() []common.Address {
	rs := make([]common.Address, 0, len(zeroGasLists))
	for key := range zeroGasLists {
		rs = append(rs, key)
	}
	return rs
}
