package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// 初始化链码
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	if function == "test" { //删除一个记录
		return t.Test(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *SimpleChaincode) Test(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) <= 1 {
		return shim.Error("Incorrect number of arguments")
	}
	stringss := args[1:]
	bargs := make([][]byte, len(stringss))
	for i, arg := range stringss {
		bargs[i] = []byte(arg)
	}
	return stub.InvokeChaincode(args[0], bargs, "")
}
