package ethadaptor

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/palletone/adaptor"
)

func TestSignTransaction(t *testing.T) {
	keyHex := "8e87ebb3b00565aaf3675e1f7d16ed68b300c7302267934f3831105b48e8a3e7"
	key := Hex2Bytes(keyHex)

	var input adaptor.SignTransactionInput
	input.PrivateKey = key
	//input.Transaction = Hex2Bytes("f9024981848203e883200b20946817cfb2c442693d850332c3b755b2342ec4afb280b902248c2e032100000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000aaa919a7c465be9b053673c567d73be8603179630000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000000407d7116a8706ae08baa7f4909e26728fa7a5f0365aaa919a7c465be9b053673c567d73be8603179636c7110482920e0af149a82189251f292a84148a85b7cd70d00000000000000000000000000000000000000000000000000000000000000417197961c5ae032ed6f33650f1f3a3ba111e8548a3dad14b3afa1cb6bc8f4601a6cb2b21aedcd575784e923942f3130f3290d56522ab2b28afca478e489426a4601000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041ae94b0e599ef0508ba7bec41db5b46d5a065b30d3d5c4b0a4c85ea2d4899d6607e80e3314ee0741049963d30fb3aceaa5506e13835a41ef54a8f44a04ef0f1e40100000000000000000000000000000000000000000000000000000000000000808080")
	input.Transaction = Hex2Bytes("ee8201e88502540be40082520894fa7fb320b6801336d5c044d52f70e282d48264a088016345785d8a000080808080")
	result, err := SignTransaction(&input)
	if err != nil {
		fmt.Println("failed ", err.Error())
	} else {
		fmt.Printf("%x\n", result.Signature)
		fmt.Printf("%x\n", result.Extra)
	}
}

func TestBindETHTxAndSignature(t *testing.T) {
	sigHex := "cbb1b8ba8d4460159338a06ef077706b45eb13f3112ae90fe907a7e8e9c5c7ea1a7bc41ec24256fe62edb4724d0553bffa7d1b8f9a153b78360218f2bfcb5b7a00"
	sig := Hex2Bytes(sigHex)

	var input adaptor.BindTxAndSignatureInput
	input.Transaction = Hex2Bytes("ee8201df8502540be40082520894aaa919a7c465be9b053673c567d73be860317963880de0b6b3a764000080808080")
	input.Signs = append(input.Signs, sig)

	result, err := BindETHTxAndSignature(&input)
	if err != nil {
		fmt.Println("failed ", err.Error())
	} else {
		testTxHex := "f86e8201df8502540be40082520894aaa919a7c465be9b053673c567d73be860317963880de0b6b3a7640000801ba0cbb1b8ba8d4460159338a06ef077706b45eb13f3112ae90fe907a7e8e9c5c7eaa01a7bc41ec24256fe62edb4724d0553bffa7d1b8f9a153b78360218f2bfcb5b7a"
		testTx := Hex2Bytes(testTxHex)
		fmt.Printf("%x\n", testTx)
		if bytes.Equal(testTx, result.SignedTx) {
			fmt.Println("same")
		} else {
			fmt.Println("different")
		}
		fmt.Printf("%x\n", result.SignedTx)
	}
}
func TestBindTxAndSignature(t *testing.T) {
	sigHex := "d9cf5c24e7e8fed515770010560dff2d449562086a0676039937fa87108c241b593002e8114d2c29c8c68ed92f906c98203c7bca895baa5993a7434c5d051fcc00"
	sig := Hex2Bytes(sigHex)

	var input adaptor.BindTxAndSignatureInput
	input.Transaction = Hex2Bytes("6d0000000000000000000000005b8c8b8aa705bf555f0b8e556bf0d58956ecd6e9000000000000000000000000aaa919a7c465be9b053673c567d73be8603179630000000000000000000000000000000000000000000000000de0b6b3a76400002b9d23bffc64aaba7607445760434037a18e95f9501cf2bd49eedfb0115e5bea")
	input.Signs = append(input.Signs, sig)
	input.Extra = []byte("withdraw(address,uint,string,bytes,bytes,bytes)")
	//input.Extra = []byte("transfer(address,uint256)")

	result, err := BindTxAndSignature(&input)
	if err != nil {
		fmt.Println("failed ", err.Error())
	} else {
		testTxHex := "a78a8274000000000000000000000000aaa919a7c465be9b053673c567d73be8603179630000000000000000000000000000000000000000000000000de0b6b3a76400002b9d23bffc64aaba7607445760434037a18e95f9501cf2bd49eedfb0115e5bead9cf5c24e7e8fed515770010560dff2d449562086a0676039937fa87108c241b593002e8114d2c29c8c68ed92f906c98203c7bca895baa5993a7434c5d051fcc00"
		testTx := Hex2Bytes(testTxHex)
		fmt.Printf("%x\n", testTx)
		if bytes.Equal(testTx, result.SignedTx) {
			fmt.Println("same")
		} else {
			fmt.Println("different")
		}
		fmt.Printf("%x\n", result.SignedTx)
	}
}
func TestCalcTxHash(t *testing.T) {
	var txs [][]byte
	hashHex := []string{"072984d536835f130f9e120c06026ae0638b27e3f49f478d79731cebb64e2b51", "a0826794e0381b52c49eb4e8a13d906db797165856dbdc3506bee1043117ca13"}

	txHex := "ee8201df8502540be40082520894aaa919a7c465be9b053673c567d73be860317963880de0b6b3a764000080808080"
	tx := Hex2Bytes(txHex)
	txs = append(txs, tx)

	txSignedHex := "f9028981848203e883200b20946817cfb2c442693d850332c3b755b2342ec4afb280b902248c2e032100000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000aaa919a7c465be9b053673c567d73be8603179630000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000000407d7116a8706ae08baa7f4909e26728fa7a5f0365aaa919a7c465be9b053673c567d73be8603179636c7110482920e0af149a82189251f292a84148a85b7cd70d00000000000000000000000000000000000000000000000000000000000000417197961c5ae032ed6f33650f1f3a3ba111e8548a3dad14b3afa1cb6bc8f4601a6cb2b21aedcd575784e923942f3130f3290d56522ab2b28afca478e489426a4601000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041ae94b0e599ef0508ba7bec41db5b46d5a065b30d3d5c4b0a4c85ea2d4899d6607e80e3314ee0741049963d30fb3aceaa5506e13835a41ef54a8f44a04ef0f1e401000000000000000000000000000000000000000000000000000000000000001ca054e4cd625aaac2c7b9eb57108438c1749affe1417192775a4a549ff198bc2cfba0236c19b5646916cdf276e1d925771d70ee50f4e261e415deeef8a03e0f7e5a13"
	txSigned := Hex2Bytes(txSignedHex)
	txs = append(txs, txSigned)

	var input adaptor.CalcTxHashInput

	for i := 0; i < len(txs); i++ {
		input.Transaction = txs[i]

		result, err := CalcTxHash(&input)
		if err != nil {
			fmt.Println("failed ", err.Error())
		} else {
			testHash := Hex2Bytes(hashHex[i])
			fmt.Printf("test result %d ", i)
			if bytes.Equal(testHash, result.Hash) {
				fmt.Println("same")
			} else {
				fmt.Println("different")
			}
			fmt.Printf("%x\n", result.Hash)
		}
	}

}

func TestSendTransaction(t *testing.T) {
	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",
	}
	var input adaptor.SendTransactionInput
	//txSignedHex := "f90dd18201e18502540be40083200b208080b90d7c6080604052633b9aca0060035534801561001857600080fd5b5060008054600160a060020a03191673a54880da9a63cdd2ddacf25af68daf31a1bcc0c9179055610d2e8061004e6000396000f3006080604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100d4578063095ea7b31461015e5780630ac74e721461019657806318160ddd146101c757806323b872dd146101ee578063313ce567146102185780634e11092f1461024357806370a08231146102645780638c5cecaa14610285578063927f526f146102a657806395d89b41146102c7578063a9059cbb146102dc578063dd62ed3e14610300578063e1a0cbd314610327575b600080fd5b3480156100e057600080fd5b506100e9610348565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561012357818101518382015260200161010b565b50505050905090810190601f1680156101505780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561016a57600080fd5b50610182600160a060020a036004351660243561037f565b604080519115158252519081900360200190f35b3480156101a257600080fd5b506101ab610388565b60408051600160a060020a039092168252519081900360200190f35b3480156101d357600080fd5b506101dc610397565b60408051918252519081900360200190f35b3480156101fa57600080fd5b50610182600160a060020a0360043581169060243516604435610435565b34801561022457600080fd5b5061022d61043e565b6040805160ff9092168252519081900360200190f35b34801561024f57600080fd5b506101ab600160a060020a0360043516610443565b34801561027057600080fd5b506101dc600160a060020a036004351661045e565b34801561029157600080fd5b506101ab600160a060020a03600435166104b3565b3480156102b257600080fd5b506100e9600160a060020a03600435166104ce565b3480156102d357600080fd5b506100e96104f3565b3480156102e857600080fd5b50610182600160a060020a036004351660243561052a565b34801561030c57600080fd5b506101dc600160a060020a036004358116906024351661037f565b34801561033357600080fd5b506100e9600160a060020a03600435166105df565b60408051808201909152600b81527f50544e204d617070696e67000000000000000000000000000000000000000000602082015281565b60005b92915050565b600054600160a060020a031681565b60008060009054906101000a9004600160a060020a0316600160a060020a03166318160ddd6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561040457600080fd5b505af1158015610418573d6000803e3d6000fd5b505050506040513d602081101561042e57600080fd5b5051905090565b60009392505050565b600081565b600260205260009081526040902054600160a060020a031681565b600160a060020a0381811660009081526001602052604081205490911615156104aa57600160a060020a038281166000908152600260205260409020541615156104aa575060016104ae565b5060005b919050565b600160205260009081526040902054600160a060020a031681565b606081816104eb6104e66104e184610606565b6107de565b610977565b949350505050565b60408051808201909152600681527f50544e4d61700000000000000000000000000000000000000000000000000000602082015281565b33600090815260016020526040812054600160a060020a031615156105d7573360008181526001602090815260408083208054600160a060020a03891673ffffffffffffffffffffffffffffffffffffffff199182168117909255818552600284529382902080549094168517909355805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001610382565b506000610382565b600160a060020a0380821660009081526001602052604090205460609161038291166104ce565b6040805160198082528183019092526060916c01000000000000000000000000840291839160009182918291906020820161032080388339505081519195506000918691508290811061065557fe5b906020010190600160f860020a031916908160001a905350600092505b60148360ff1610156106cb578460ff84166014811061068d57fe5b1a60f860020a02848460010160ff168151811015156106a857fe5b906020010190600160f860020a031916908160001a905350600190920191610672565b6040805160008082526bffffffffffffffffffffffff1988166001830152915160029283926015808201936020939092839003909101908290865af1158015610718573d6000803e3d6000fd5b5050506040513d602081101561072d57600080fd5b505160408051918252516020828101929091908190038201816000865af115801561075c573d6000803e3d6000fd5b5050506040513d602081101561077157600080fd5b50519150600090505b60048160ff1610156107d0578160ff82166020811061079557fe5b1a60f860020a02848260150160ff168151811015156107b057fe5b906020010190600160f860020a031916908160001a90535060010161077a565b8395505b5050505050919050565b6060806000806000808651600014156108075760408051602081019091526000815295506107d4565b6040805160288082526105208201909252906020820161050080388339019050509450600085600081518110151561083b57fe5b60ff90921660209283029091019091015260019350600092505b86518360ff16101561095257868360ff1681518110151561087257fe5b90602001015160f860020a900460f860020a0260f860020a900460ff169150600090505b8360ff168160ff16101561090757848160ff168151811015156108b557fe5b9060200190602002015160ff166101000282019150603a828115156108d657fe5b06858260ff168151811015156108e857fe5b60ff909216602092830290910190910152603a82049150600101610896565b600082111561094757603a8206858560ff1681518110151561092557fe5b60ff909216602092830290910190910152600190930192603a82049150610907565b826001019250610855565b61096c6109676109628787610ac4565b610b59565b610bef565b979650505050505050565b606080606060008085516002016040519080825280601f01601f1916602001820160405280156109b1578160200160208202803883390190505b508051909450849350600192507f500000000000000000000000000000000000000000000000000000000000000090849060009081106109ed57fe5b906020010190600160f860020a031916908160001a905350825160018301927f3100000000000000000000000000000000000000000000000000000000000000918591908110610a3957fe5b906020010190600160f860020a031916908160001a905350600090505b85518160ff161015610aba57858160ff16815181101515610a7357fe5b90602001015160f860020a900460f860020a028383806001019450815181101515610a9a57fe5b906020010190600160f860020a031916908160001a905350600101610a56565b5091949350505050565b60608060008360ff16604051908082528060200260200182016040528015610af6578160200160208202803883390190505b509150600090505b8360ff168160ff161015610b5157848160ff16815181101515610b1d57fe5b90602001906020020151828260ff16815181101515610b3857fe5b60ff909216602092830290910190910152600101610afe565b509392505050565b60608060008351604051908082528060200260200182016040528015610b89578160200160208202803883390190505b509150600090505b83518160ff161015610be8578351849060ff8316810360001901908110610bb457fe5b90602001906020020151828260ff16815181101515610bcf57fe5b60ff909216602092830290910190910152600101610b91565b5092915050565b606080600083516040519080825280601f01601f191660200182016040528015610c23578160200160208202803883390190505b509150600090505b83518160ff161015610be857606060405190810160405280603a81526020017f31323334353637383941424344454647484a4b4c4d4e5051525354555657585981526020017f5a6162636465666768696a6b6d6e6f707172737475767778797a000000000000815250848260ff16815181101515610ca557fe5b9060200190602002015160ff16815181101515610cbe57fe5b90602001015160f860020a900460f860020a02828260ff16815181101515610ce257fe5b906020010190600160f860020a031916908160001a905350600101610c2b5600a165627a7a72305820e3859334108c189e50db001d7a3605fb2d4ca63b2eae29be13fd188932d31a7900291ca05e84ee18b819e591059b0ed6dede39cb79918982023df214c8b7cc20487cadeba04319cb487496e30c7226b079cdf72a8e780533dfd19762f249bd6ae75df342e5"
	txSignedHex := "f86e8201e88502540be40082520894fa7fb320b6801336d5c044d52f70e282d48264a088016345785d8a0000801ba0a2f7ee8133336f23d68ef8600d05eaded682991c9292d94562d940a75bb53d74a0542b462d7eb4ba5db603f85b276a93a2fa610dfb7f1b1d944f010f7aa24f3c90"
	input.Transaction = Hex2Bytes(txSignedHex)
	result, err := SendTransaction(&input, &rpcParams)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%x\n", result.TxID)
	}
}

func TestCreateETHTx(t *testing.T) {
	rpcParams := RPCParams{
		Rawurl: "https://ropsten.infura.io/", //"\\\\.\\pipe\\geth.ipc",
	}
	var input adaptor.CreateTransferTokenTxInput
	input.FromAddress = "0x7D7116A8706Ae08bAA7F4909e26728fa7A5f0365"
	//input.FromAddress = "0x5b8c8B8Aa705bF555F0B8E556Bf0d58956eCD6e9"
	//input.ToAddress = "0xaAA919a7c465be9b053673C567D73Be860317963"
	input.ToAddress = "0xfa7fb320b6801336d5c044d52f70e282d48264a0"
	input.Amount = adaptor.NewAmountAssetString("100000000000000000", "ETH")
	input.Fee = adaptor.NewAmountAssetString("10000000000", "ETH") //10g wei,

	result, err := CreateETHTx(&input, &rpcParams)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("unsigned tx: %x\n", result.Transaction)
	}
}

func TestCreateTx(t *testing.T) {
	var input adaptor.CreateTransferTokenTxInput
	//input.FromAddress = "0x7D7116A8706Ae08bAA7F4909e26728fa7A5f0365"
	input.FromAddress = "0x5b8c8B8Aa705bF555F0B8E556Bf0d58956eCD6e9"
	input.ToAddress = "0xaAA919a7c465be9b053673C567D73Be860317963"
	input.Amount = adaptor.NewAmountAssetString("1000000000000000000", "ETH") //1eth
	input.Fee = adaptor.NewAmountAssetString("10000000000", "ETH")            //10g wei,
	input.Extra = Hex2Bytes("2b9d23bffc64aaba7607445760434037a18e95f9501cf2bd49eedfb0115e5bea")

	result, err := CreateTx(&input)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("unsigned tx: %x\n", result.Transaction)
	}
}
func TestSignMessage(t *testing.T) {
	keyHex := "8e87ebb3b00565aaf3675e1f7d16ed68b300c7302267934f3831105b48e8a3e7"
	key := Hex2Bytes(keyHex)

	var input adaptor.SignMessageInput
	input.PrivateKey = key
	input.Message = Hex2Bytes("")

	result, err := SignMessage(&input)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Signature: %x\n", result.Signature)
	}
}
func TestVerifySignature(t *testing.T) {
	var input adaptor.VerifySignatureInput
	input.Message = Hex2Bytes("6d0000000000000000000000005b8c8b8aa705bf555f0b8e556bf0d58956ecd6e9000000000000000000000000aaa919a7c465be9b053673c567d73be8603179630000000000000000000000000000000000000000000000000de0b6b3a76400002b9d23bffc64aaba7607445760434037a18e95f9501cf2bd49eedfb0115e5bea")
	input.Signature = Hex2Bytes("d9cf5c24e7e8fed515770010560dff2d449562086a0676039937fa87108c241b593002e8114d2c29c8c68ed92f906c98203c7bca895baa5993a7434c5d051fcc00")
	input.PublicKey = Hex2Bytes("021c183161f5d96f59d6078d0123021876b5a0982b131ffa021b4437f49b93588a")
	result, err := VerifySignature(&input)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Signature Pass: %v\n", result.Pass)
	}
}
