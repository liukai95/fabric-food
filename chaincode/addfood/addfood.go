package main

import (
	"fmt"
	"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/examples/chaincode/go/addfood/util"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// PassthruChaincode passes thru invoke and query to another chaincode where
//     called ChaincodeID = function
//     called chaincode's function = args[0]
//     called chaincode's args = args[1:]
type PassthruChaincode struct {
}

//Init func will return error if function has string "error" anywhere
func (p *PassthruChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	function, _ := stub.GetFunctionAndParameters()
	if strings.Index(function, "error") >= 0 {
		return shim.Error(function)
	}
	return shim.Success([]byte(function))
}

//helper
func (p *PassthruChaincode) iq(stub shim.ChaincodeStubInterface, function string, args []string) pb.Response {
	if function == "" {
		return shim.Error("Chaincode ID not provided")
	}
	chaincodeID := function

	return stub.InvokeChaincode(chaincodeID, util.ToChaincodeArgs(args...), "")
}

// Invoke passes through the invoke call
func (p *PassthruChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	return p.iq(stub, function, args)
}

func main() {
	err := shim.Start(new(PassthruChaincode))
	if err != nil {
		fmt.Printf("Error starting Passthru chaincode: %s", err)
	}
}
