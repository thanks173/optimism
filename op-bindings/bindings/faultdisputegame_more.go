// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const FaultDisputeGameStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"gameStart\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_userDefinedValueType(Timestamp)1012\"},{\"astId\":1001,\"contract\":\"contracts/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"status\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_enum(GameStatus)1006\"},{\"astId\":1002,\"contract\":\"contracts/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"bondManager\",\"offset\":9,\"slot\":\"0\",\"type\":\"t_contract(IBondManager)1005\"},{\"astId\":1003,\"contract\":\"contracts/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"claimData\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_struct(ClaimData)1007_storage)dyn_storage\"},{\"astId\":1004,\"contract\":\"contracts/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"claims\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_userDefinedValueType(ClaimHash)1009,t_bool)\"}],\"types\":{\"t_array(t_struct(ClaimData)1007_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct IFaultDisputeGame.ClaimData[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(ClaimData)1007_storage\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(IBondManager)1005\":{\"encoding\":\"inplace\",\"label\":\"contract IBondManager\",\"numberOfBytes\":\"20\"},\"t_enum(GameStatus)1006\":{\"encoding\":\"inplace\",\"label\":\"enum GameStatus\",\"numberOfBytes\":\"1\"},\"t_mapping(t_userDefinedValueType(ClaimHash)1009,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(ClaimHash =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(ClaimHash)1009\",\"value\":\"t_bool\"},\"t_struct(ClaimData)1007_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IFaultDisputeGame.ClaimData\",\"numberOfBytes\":\"96\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_userDefinedValueType(Claim)1008\":{\"encoding\":\"inplace\",\"label\":\"Claim\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(ClaimHash)1009\":{\"encoding\":\"inplace\",\"label\":\"ClaimHash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Clock)1010\":{\"encoding\":\"inplace\",\"label\":\"Clock\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Position)1011\":{\"encoding\":\"inplace\",\"label\":\"Position\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Timestamp)1012\":{\"encoding\":\"inplace\",\"label\":\"Timestamp\",\"numberOfBytes\":\"8\"}}}"

var FaultDisputeGameStorageLayout = new(solc.StorageLayout)

var FaultDisputeGameDeployedBin = "0x60806040526004361061016a5760003560e01c80638129fc1c116100cb578063bcef3b551161007f578063cf09e0d011610059578063cf09e0d01461049c578063e4c290c4146104bb578063fa24f743146104db57600080fd5b8063bcef3b55146103e8578063c55cd0c714610425578063c6f0308c1461043857600080fd5b80638b85902b116100b05780638b85902b146103585780639293129814610398578063bbdc02db146103cc57600080fd5b80638129fc1c1461032e5780638980e0cc1461034357600080fd5b8063363cc4271161012257806354fd4d501161010757806354fd4d50146102e4578063609d333414610306578063632247ea1461031b57600080fd5b8063363cc427146102515780634778efe8146102b057600080fd5b80632810e1d6116101535780632810e1d6146101ed5780633218b99d1461020257806335fef5671461023c57600080fd5b8063200d2ed21461016f578063266198f9146101ab575b600080fd5b34801561017b57600080fd5b506000546101959068010000000000000000900460ff1681565b6040516101a29190611af6565b60405180910390f35b3480156101b757600080fd5b506101df7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016101a2565b3480156101f957600080fd5b506101956104ff565b34801561020e57600080fd5b506000546102239067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016101a2565b61024f61024a366004611b37565b6108f4565b005b34801561025d57600080fd5b5060005461028b906901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a2565b3480156102bc57600080fd5b506101df7f000000000000000000000000000000000000000000000000000000000000000081565b3480156102f057600080fd5b506102f9610904565b6040516101a29190611bd3565b34801561031257600080fd5b506102f96109a7565b61024f610329366004611c02565b6109b9565b34801561033a57600080fd5b5061024f610f73565b34801561034f57600080fd5b506001546101df565b34801561036457600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003602001356101df565b3480156103a457600080fd5b5061028b7f000000000000000000000000000000000000000000000000000000000000000081565b3480156103d857600080fd5b50604051600081526020016101a2565b3480156103f457600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003356101df565b61024f610433366004611b37565b6110b4565b34801561044457600080fd5b50610458610453366004611c37565b6110c0565b6040805163ffffffff90961686529315156020860152928401919091526fffffffffffffffffffffffffffffffff908116606084015216608082015260a0016101a2565b3480156104a857600080fd5b5060005467ffffffffffffffff16610223565b3480156104c757600080fd5b5061024f6104d6366004611c99565b611131565b3480156104e757600080fd5b506104f06116aa565b6040516101a293929190611d2d565b60008060005468010000000000000000900460ff16600281111561052557610525611ac7565b1461055c576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001805460009161056c91611d87565b90506fffffffffffffffffffffffffffffffff815b67ffffffffffffffff811015610656576000600182815481106105a6576105a6611d9e565b6000918252602090912060039091020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9093019290915060ff64010000000090910416156105f75750610581565b600281015460009061063b906fffffffffffffffffffffffffffffffff167f00000000000000000000000000000000000000000000000000000000000000006116e8565b90508381101561064f578093508260010194505b5050610581565b5060006001838154811061066c5761066c611d9e565b600091825260208220600390910201805490925063ffffffff908116919082146106d657600182815481106106a3576106a3611d9e565b906000526020600020906003020160020160109054906101000a90046fffffffffffffffffffffffffffffffff16610702565b600283015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff165b905062049d4061071c67ffffffffffffffff831642611d87565b610738836fffffffffffffffffffffffffffffffff1660401c90565b67ffffffffffffffff1661074c9190611dcd565b11610783576040517ff2440b5300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600283810154610825906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b61082f9190611e14565b67ffffffffffffffff1615801561085657506fffffffffffffffffffffffffffffffff8414155b156108645760029550610869565b600195505b600080548791907fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000008360028111156108ae576108ae611ac7565b0217905560028111156108c3576108c3611ac7565b6040517f5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da6090600090a2505050505090565b610900828260006109b9565b5050565b606061092f7f000000000000000000000000000000000000000000000000000000000000000061179d565b6109587f000000000000000000000000000000000000000000000000000000000000000061179d565b6109817f000000000000000000000000000000000000000000000000000000000000000061179d565b60405160200161099393929190611e3b565b604051602081830303815290604052905090565b60606109b46020806118da565b905090565b6000805468010000000000000000900460ff1660028111156109dd576109dd611ac7565b14610a14576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82158015610a20575080155b15610a57576040517fa42637bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060018481548110610a6c57610a6c611d9e565b60009182526020918290206040805160a0810182526003909302909101805463ffffffff8116845260ff64010000000090910416151593830193909352600180840154918301919091526002909201546fffffffffffffffffffffffffffffffff80821660608401527001000000000000000000000000000000009091041660808201528154909250819086908110610b0757610b07611d9e565b6000918252602082206003909102018054921515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff909316929092179091556060820151610b71906fffffffffffffffffffffffffffffffff1684151760011b90565b90507f0000000000000000000000000000000000000000000000000000000000000000610c30826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff161115610c72576040517f56f57b2b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160009063ffffffff90811614610cd2576001836000015163ffffffff1681548110610ca157610ca1611d9e565b906000526020600020906003020160020160109054906101000a90046fffffffffffffffffffffffffffffffff1690505b608083015160009067ffffffffffffffff1667ffffffffffffffff1642610d0b846fffffffffffffffffffffffffffffffff1660401c90565b67ffffffffffffffff16610d1f9190611dcd565b610d299190611d87565b905062049d4067ffffffffffffffff82161115610d72576040517f3381d11400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000604082901b421790506000610d93888660009182526020526040902090565b60008181526002602052604090205490915060ff1615610ddf576040517f80497e3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260026020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155815160a08101835263ffffffff808f1682529381018581528184018e81526fffffffffffffffffffffffffffffffff808d16606085019081528a82166080860190815286548088018855968a52945160039096027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf68101805495511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000909616979099169690961793909317909655517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf78401555190518416700100000000000000000000000000000000029316929092177fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf8909201919091555133918a918c917f9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be91a4505050505050505050565b600080547fffffffffffffffffffffffffffffffffffffffffffffff000000000000000000164267ffffffffffffffff161781556040805160a08101825263ffffffff81526020810192909252600191908101610ff87ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c90033590565b815260016020820152604001426fffffffffffffffffffffffffffffffff908116909152825460018181018555600094855260209485902084516003909302018054958501511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000090961663ffffffff909316929092179490941781556040830151938101939093556060820151608090920151811670010000000000000000000000000000000002911617600290910155565b610900828260016109b9565b600181815481106110d057600080fd5b600091825260209091206003909102018054600182015460029092015463ffffffff8216935064010000000090910460ff1691906fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041685565b6000805468010000000000000000900460ff16600281111561115557611155611ac7565b1461118c576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600187815481106111a1576111a1611d9e565b6000918252602082206003919091020160028101549092506fffffffffffffffffffffffffffffffff16908715821760011b90506112007f00000000000000000000000000000000000000000000000000000000000000006001611dcd565b61129c826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff16146112dd576040517f5f53dd9800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806112fb836fffffffffffffffffffffffffffffffff16611971565b67ffffffffffffffff1660000361135b577f0000000000000000000000000000000000000000000000000000000000000000915060018b8154811061134257611342611d9e565b906000526020600020906003020160010154905061157b565b6000808b156113dd5760018e8154811061137757611377611d9e565b906000526020600020906003020160020160009054906101000a90046fffffffffffffffffffffffffffffffff16915060018e815481106113ba576113ba611d9e565b90600052602060002090600302016001015493508590508660010154925061146c565b600287015460018089015481549096506fffffffffffffffffffffffffffffffff9092169350908f90811061141457611414611d9e565b906000526020600020906003020160020160009054906101000a90046fffffffffffffffffffffffffffffffff16905060018e8154811061145757611457611d9e565b90600052602060002090600302016001015492505b60016114aa6fffffffffffffffffffffffffffffffff83167f0000000000000000000000000000000000000000000000000000000000000000611a17565b6114b49190611eb1565b6fffffffffffffffffffffffffffffffff1661150b7f0000000000000000000000000000000000000000000000000000000000000000846fffffffffffffffffffffffffffffffff16611a1790919063ffffffff16565b6fffffffffffffffffffffffffffffffff161415806115415750838b8b604051611536929190611ee2565b604051809103902014155b15611578576040517f696550ff00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505b6040517ff8e0cb96000000000000000000000000000000000000000000000000000000008152819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063f8e0cb96906115f5908d908d908d908d90600401611f3b565b6020604051808303816000875af1158015611614573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116389190611f6d565b0361166f576040517ffb4e40dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505082547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff1664010000000017909255505050505050505050565b6000367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033560606116e16109a7565b9050909192565b600080611775847e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1690508083036001841b600180831b0386831b17039250505092915050565b6060816000036117e057505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561180a57806117f481611f86565b91506118039050600a83611fbe565b91506117e4565b60008167ffffffffffffffff81111561182557611825611fd2565b6040519080825280601f01601f19166020018201604052801561184f576020820181803683370190505b5090505b84156118d257611864600183611d87565b9150611871600a86612001565b61187c906030611dcd565b60f81b81838151811061189157611891611d9e565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506118cb600a86611fbe565b9450611853565b949350505050565b6060600061191184367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003611dcd565b90508267ffffffffffffffff1667ffffffffffffffff81111561193657611936611fd2565b6040519080825280601f01601f191660200182016040528015611960576020820181803683370190505b509150828160208401375092915050565b6000806119fe837e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b600167ffffffffffffffff919091161b90920392915050565b600080611aa4847e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169050808303600180821b0385821b179250505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160038310611b31577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60008060408385031215611b4a57600080fd5b50508035926020909101359150565b60005b83811015611b74578181015183820152602001611b5c565b83811115611b83576000848401525b50505050565b60008151808452611ba1816020860160208601611b59565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611be66020830184611b89565b9392505050565b80358015158114611bfd57600080fd5b919050565b600080600060608486031215611c1757600080fd5b8335925060208401359150611c2e60408501611bed565b90509250925092565b600060208284031215611c4957600080fd5b5035919050565b60008083601f840112611c6257600080fd5b50813567ffffffffffffffff811115611c7a57600080fd5b602083019150836020828501011115611c9257600080fd5b9250929050565b600080600080600080600060a0888a031215611cb457600080fd5b8735965060208801359550611ccb60408901611bed565b9450606088013567ffffffffffffffff80821115611ce857600080fd5b611cf48b838c01611c50565b909650945060808a0135915080821115611d0d57600080fd5b50611d1a8a828b01611c50565b989b979a50959850939692959293505050565b60ff84168152826020820152606060408201526000611d4f6060830184611b89565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611d9957611d99611d58565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008219821115611de057611de0611d58565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680611e2f57611e2f611de5565b92169190910692915050565b60008451611e4d818460208901611b59565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611e89816001850160208a01611b59565b60019201918201528351611ea4816002840160208801611b59565b0160020195945050505050565b60006fffffffffffffffffffffffffffffffff83811690831681811015611eda57611eda611d58565b039392505050565b8183823760009101908152919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000611f4f604083018688611ef2565b8281036020840152611f62818587611ef2565b979650505050505050565b600060208284031215611f7f57600080fd5b5051919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611fb757611fb7611d58565b5060010190565b600082611fcd57611fcd611de5565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008261201057612010611de5565b50069056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(FaultDisputeGameStorageLayoutJSON), FaultDisputeGameStorageLayout); err != nil {
		panic(err)
	}

	layouts["FaultDisputeGame"] = FaultDisputeGameStorageLayout
	deployedBytecodes["FaultDisputeGame"] = FaultDisputeGameDeployedBin
}
