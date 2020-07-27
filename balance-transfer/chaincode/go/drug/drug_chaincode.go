package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	//"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type Drug struct {
	Drugid   string `json:"drugid"`   //药品id
	Name     string `json:"name"`     //名称
	Dosage   string `json:"dosage"`   //剂型
	Standard string `json:"standard"` //生产标准
	Effect   string `json:"effect"`   //作用
	Personid string `json:"personid"` //用户ID
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

	if function == "initDrug" { //新增药品
		return t.initDrug(stub, args)
	} else if function == "delete" { //删除一个药品
		return t.delete(stub, args)
	} else if function == "readDrugById" { //根据药品id读取一个药品
		return t.readDrugById(stub, args)
	} else if function == "readAllDrug" { //查询所有药品
		return t.readAllDrug(stub)
	} else if function == "initLedgerDrug" { //初始化
		return t.initLedgerDrug(stub)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *SimpleChaincode) initLedgerDrug(stub shim.ChaincodeStubInterface) pb.Response {
	drugs := []Drug{
		Drug{Drugid: "CCJHB0317012501", Name: "除虫剂", Dosage: "液态", Standard: "国标", Effect: "除虫", Personid: "FW"},
		Drug{Drugid: "FHCBJ0317021001", Name: "氟环唑", Dosage: "液态", Standard: "国标", Effect: "除虫", Personid: "FW"},
		Drug{Drugid: "KJJBJ0317020101", Name: "咯菌腈", Dosage: "液态", Standard: "国标", Effect: "除虫", Personid: "FW"},
		Drug{Drugid: "MXABJ0317021801", Name: "咪鲜胺", Dosage: "液态", Standard: "国标", Effect: "除虫", Personid: "FW"},
	}

	i := 0
	for i < len(drugs) {
		fmt.Println("i is ", i)
		drugAsBytes, _ := json.Marshal(drugs[i])
		//stub.PutState("Person"+strconv.Itoa(i), personAsBytes)
		stub.PutState("DRUG"+strconv.Itoa(i), drugAsBytes)
		fmt.Println("Added", drugs[i])
		i = i + 1
	}

	return shim.Success(nil)
}

// ============================================================
// initPerson - create a new person, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initDrug(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2      3     4                5        6
	//  主键DRUG+数字  'XDJBJ0317012001', '消毒剂', '液态', '国标', '消毒', 'a'
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init drug")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5fi argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("5fi argument must be a non-empty string")
	}
	//if len(args[6]) <= 0 {//personid
	//	return shim.Error("6si argument must be a non-empty string")
	//}

	// ==== 检查主键是否存在s ====
	drugAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get pdrug : " + err.Error())
	} else if drugAsBytes != nil {
		fmt.Println("The drug  already exists: " + args[0])
		return shim.Error("The drug  already exists: " + args[0])
	}

	// ==== 创建实体转成JSON ====
	drug := &Drug{args[1], args[2], args[3], args[4], args[5], args[6]}
	drugJSONasBytes, err := json.Marshal(drug)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === 保存 ===
	err = stub.PutState(args[0], drugJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// ===============================================
// readDrugById - read a drug from chaincode state
// ===============================================
func (t *SimpleChaincode) readDrugById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	id = args[0]
	valAsbytes, err := stub.GetState(id) //get the drug from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Marble does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ==================================================
// delete - remove a person key/value pair from state
// ==================================================
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var drugJSON Drug
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	id := args[0]

	// to maintain the color~name index, we need to read the person first and get its color
	valAsbytes, err := stub.GetState(id) //get the person from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Marble does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &drugJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + id + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(id) //remove the marble from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) readAllDrug(stub shim.ChaincodeStubInterface) pb.Response {

	startKey := "DRUG0001"
	endKey := "DRUG9999"

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- readAllPerson:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
