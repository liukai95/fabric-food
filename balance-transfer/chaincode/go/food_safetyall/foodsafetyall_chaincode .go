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

//ID
type Value struct {
	Valueid string `json:"valueid"` //保存id，根据ID查询数据
}

//用户
type User struct {
	Userid   string `json:"userid"`   //用户id,主键
	Name     string `json:"name"`     //姓名
	Password string `json:"password"` //密码
}

//质检员
type Person struct {
	Personid  string `json:"personid"`  //质检员id,主键
	Name      string `json:"name"`      //姓名
	Sex       string `json:"sex"`       //性别
	Workplace string `json:"workplace"` //工作单位
	Job       string `json:"job"`       //岗位
	Password  string `json:"password"`  //密码
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

	if function == "deleteOne" { //删除一个记录
		return t.deleteOne(stub, args)
	} else if function == "readOneById" { //根据id读取一个记录
		return t.readOneById(stub, args)
	} else if function == "readOneByIdOther" { //根据id调用另一个链码读取
		return t.readOneByIdOther(stub, args)
	} else if function == "readAll" { //根据范围查询一个类型的所有
		return t.readAll(stub, args)
	} else if function == "readAllOther" { //根据范围调用另一个链码查询一个类型的所有
		return t.readAllOther(stub, args)
	} else if function == "initLedger" { //初始化所有类型的数据
		return t.initLedger(stub)
	} else if function == "initPerson" { //新增一个质检员
		return t.initPerson(stub, args)
	} else if function == "initUser" { //新增一个用户
		return t.initUser(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *SimpleChaincode) initLedger(stub shim.ChaincodeStubInterface) pb.Response {

	//初始化质检员
	persons := []Person{
		Person{Personid: "FOOD00264", Name: "夏良军", Sex: "男", Workplace: "米业生产", Job: "监仓", Password: "123"},
		Person{Personid: "FOOD006133", Name: "徐炎浩", Sex: "男", Workplace: "福天下合作社", Job: "基地管理员", Password: "123"},
		Person{Personid: "FOOD006222", Name: "张道海", Sex: "男", Workplace: "福天下合作社", Job: "稻虾基地管理员", Password: "123"},
		Person{Personid: "FOOD00014", Name: "齐永新", Sex: "男", Workplace: "福天下合作社", Job: "技术员", Password: "123"},
		Person{Personid: "FOOD006232", Name: "杜宝元", Sex: "男", Workplace: "福天下合作社", Job: "稻虾基地管理员", Password: "123"},
		Person{Personid: "FOOD006281", Name: "张从仿", Sex: "男", Workplace: "福天下合作社", Job: "基地管理员", Password: "123"},
		Person{Personid: "FOOD00630", Name: "谢申华", Sex: "男", Workplace: "福天下合作社", Job: "工程部经理", Password: "123"},
		Person{Personid: "FOOD00656", Name: "谢星", Sex: "男", Workplace: "米业生产", Job: "技术员", Password: "123"},
		Person{Personid: "FOOD00702", Name: "谢均涛", Sex: "男", Workplace: "米业生产", Job: "原料保管", Password: "123"},
		Person{Personid: "FOOD00803", Name: "谢松才", Sex: "男", Workplace: "福天下合作社", Job: "基地管理员", Password: "123"},
		Person{Personid: "FOOD00001", Name: "aa", Sex: "女", Workplace: "米业生产", Job: "原料保管", Password: "123"},
		Person{Personid: "FOOD00002", Name: "bb", Sex: "女", Workplace: "福天下合作社", Job: "基地管理员", Password: "123"},
	}
	i := 0
	for i < len(persons) {
		fmt.Println("i is ", i)
		personAsBytes, _ := json.Marshal(persons[i])
		//把Personid作为键
		//stub.PutState(persons[i].Personid, personAsBytes)
		newvalueAsBytes, _ := json.Marshal(&Value{"PERSON" + strconv.Itoa(i+1)})
		stub.PutState("PERSON"+strconv.Itoa(i+1), personAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(persons[i].Personid, newvalueAsBytes)
		fmt.Println("Added", persons[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ := json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YPERSON", valueAsBytes)

	//初始化用户
	users := []User{
		User{Userid: "123456", Name: "Tom", Password: "123456"},
		User{Userid: "123457", Name: "Jim", Password: "123456"},
		User{Userid: "123458", Name: "张三", Password: "123456"},
	}
	i = 0
	for i < len(users) {
		fmt.Println("i is ", i)
		userAsBytes, _ := json.Marshal(users[i])
		newvalueAsBytes, _ := json.Marshal(&Value{"USER" + strconv.Itoa(i+1)})
		stub.PutState("USER"+strconv.Itoa(i+1), userAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(users[i].Userid, newvalueAsBytes)
		fmt.Println("Added", users[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YUSER", valueAsBytes)

	return shim.Success(nil)
}

// ==================================================
// 读取
// ==================================================
func (t *SimpleChaincode) readOneById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	id = args[0]
	//由ID得到新键
	valAsbytes, err := stub.GetState(id)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"ValueID does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}
	valueToTransfer := Value{}
	err = json.Unmarshal(valAsbytes, &valueToTransfer) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("id" + valueToTransfer.Valueid)

	valAsbytes2, err := stub.GetState(valueToTransfer.Valueid) //由id得到新键
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + valueToTransfer.Valueid + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes2 == nil {
		jsonResp = "{\"Error\":\"ValueID does not exist: " + valueToTransfer.Valueid + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsbytes2)
}

// ==================================================
// 删除
// ==================================================
func (t *SimpleChaincode) deleteOne(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	//var personJSON Person
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	id := args[0]
	//由ID得到新键
	valAsbytes, err := stub.GetState(id)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"ValueID does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}
	valueToTransfer := Value{}
	err = json.Unmarshal(valAsbytes, &valueToTransfer) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	valAsbytes2, err := stub.GetState(valueToTransfer.Valueid)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + valueToTransfer.Valueid + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes2 == nil {
		jsonResp = "{\"Error\":\"ValueID does not exist: " + valueToTransfer.Valueid + "\"}"
		return shim.Error(jsonResp)
	}
	//删除新键和旧键
	err = stub.DelState(valueToTransfer.Valueid)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	err = stub.DelState(id)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) readAll(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	startKey := args[0]
	endKey := args[1]

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

// ============================================================
// 新增用户
// ============================================================
func (t *SimpleChaincode) initUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0           1       2
	// '123456', 'Tom', '123'
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init user")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	// ==== 检查用户是否存在 ====
	userAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get user: " + err.Error())
	} else if userAsBytes != nil {
		fmt.Println("The user already exists: " + args[0])
		return shim.Error("The user already exists: " + args[0])
	}

	// ==== 创建实体转成JSON ====
	user := &User{args[0], args[1], args[2]}
	userJSONasBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YUSER")
	if err != nil {
		return shim.Error("Failed to get YUSER:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YUSER does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("USER"+strconv.Itoa(intValue), userJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"USER" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YUSER", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增质检员
// ============================================================
func (t *SimpleChaincode) initPerson(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0           1       2     3                4        5
	// 'FOOD00014', '齐永新', '男', '福天下合作社', '技术员', "aa"
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init person")
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
		return shim.Error("6si argument must be a non-empty string")
	}

	// ==== 检查质检员是否存在 ====
	personAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get person: " + err.Error())
	} else if personAsBytes != nil {
		fmt.Println("The person already exists: " + args[0])
		return shim.Error("The person already exists: " + args[0])
	}

	// ==== 创建实体转成JSON ====
	//person := &Person{personid, name, sex, workplace, job, password}
	person := &Person{args[0], args[1], args[2], args[3], args[4], args[5]}
	personJSONasBytes, err := json.Marshal(person)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YPERSON")
	if err != nil {
		return shim.Error("Failed to get YPERSON:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YPERSON does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("PERSON"+strconv.Itoa(intValue), personJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"PERSON" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YPERSON", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 调用其他的链码读取信息
// ============================================================
func (t *SimpleChaincode) readAllOther(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	stringss := args[1:]
	bargs := make([][]byte, len(stringss))
	for i, arg := range stringss {
		bargs[i] = []byte(arg)
	}
	return stub.InvokeChaincode(args[0], bargs, "")
}

// ==================================================
// 根据id调用另一个链码读取信息
// ==================================================
func (t *SimpleChaincode) readOneByIdOther(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	stringss := args[1:]
	bargs := make([][]byte, len(stringss))
	for i, arg := range stringss {
		bargs[i] = []byte(arg)
	}
	return stub.InvokeChaincode(args[0], bargs, "")
}

