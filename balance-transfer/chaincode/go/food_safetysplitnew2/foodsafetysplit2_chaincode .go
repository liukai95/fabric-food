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

//入库管理
type Input struct {
	Inputid     string `json:"inputid"`     //入库id
	Harvestdate string `json:"harvestdate"` //收割日期
	Quantity    string `json:"quantity"`    //入库量（kg）
	Inputdate   string `json:"inputdate"`   //入库日期
	Personid    string `json:"personid"`    //质检员ID
}

//仓库管理
type Warehouse struct {
	Warehouseid string `json:"warehouseid"` //仓库id
	Place       string `json:"place"`       //地点
	Capacity    string `json:"capacity"`    //容量（t）
	Standard    string `json:"standard"`    //生产标准
	Personid    string `json:"personid"`    //质检员ID
}

//出库管理
type Warehouse2feed struct {
	Warehouseid string `json:"warehouseid"` //仓库id
	Feedid      string `json:"feedid"`      //进料id
	Personid    string `json:"personid"`    //质检员ID
}

//种植入库批次转换
type Plant2input struct {
	Plantid  string `json:"plantid"`  //种植id
	Inputid  string `json:"inputid"`  //入库id
	Personid string `json:"personid"` //质检员ID
}

//入库仓库批次转换
type Input2warehouse struct {
	Inputid     string `json:"inputid"`     //入库id
	Warehouseid string `json:"warehouseid"` //仓库id
	Personid    string `json:"personid"`    //质检员ID
}

//进料管理
type Feed struct {
	Feedid       string `json:"feedid"`       //进料id
	Weight       string `json:"weight"`       //稻谷重量（kg）
	Watercontent string `json:"watercontent"` //大米水分含量（%）
	Brokenrice   string `json:"brokenrice"`   //碎米率（%）
	Qingmilv     string `json:"qingmilv"`     //青米率（%）
	Date         string `json:"date"`         //日期
	Personid     string `json:"personid"`     //质检员ID
}

//进料产品批次转换
type Feed2product struct {
	Feedid    string `json:"feedid"`    //进料id
	Productid string `json:"productid"` //产品id
	Personid  string `json:"personid"`  //质检员ID
}

//原料管理
type Material struct {
	Materialid string `json:"materialid"` //原料id
	Kind       string `json:"kind"`       //种类
	Weight     string `json:"weight"`     //重量（t）
	Source     string `json:"source"`     //来源
	Date       string `json:"date"`       //日期
	Personid   string `json:"personid"`   //质检员ID
}

//产品
type Product struct {
	Productid     string `json:"productid"`     //产品id
	Name          string `json:"name"`          //名称
	Specification string `json:"specification"` //口味
	Flavor        string `json:"flavor"`        //规格
	Date          string `json:"date"`          //日期
	Personid      string `json:"personid"`      //质检员ID
}

//原料产品批次转换
type Material2product struct {
	Materialid string `json:"materialid"` //原料id
	Productid  string `json:"productid"`  //产品id
	Personid   string `json:"personid"`   //质检员ID
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
	} else if function == "initInput" { //新增一个入库
		return t.initInput(stub, args)
	} else if function == "initWarehouse" { //新增一个仓库
		return t.initWarehouse(stub, args)
	} else if function == "initWarehouse2feed" { //新增一个出库
		return t.initWarehouse2feed(stub, args)
	} else if function == "initPlant2input" { //新增一个种植入库批次转换
		return t.initPlant2input(stub, args)
	} else if function == "initInput2warehouse" { //新增一个入库仓库批次转换
		return t.initInput2warehouse(stub, args)
	} else if function == "initFeed2product" { //新增一个进料产品批次转换
		return t.initFeed2product(stub, args)
	} else if function == "initFeed" { //新增进料管理
		return t.initFeed(stub, args)
	} else if function == "initMaterial" { //新增一个原料
		return t.initMaterial(stub, args)
	} else if function == "initProduct" { //新增一个产品
		return t.initProduct(stub, args)
	} else if function == "initMaterial2product" { //新增一个原料产品批次转换
		return t.initMaterial2product(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *SimpleChaincode) initLedger(stub shim.ChaincodeStubInterface) pb.Response {

	//初始化入库
	inputs := []Input{
		Input{Inputid: "NT12017072243990601", Harvestdate: "2017-07-22", Quantity: "439906", Inputdate: "2017-07-26", Personid: "FOOD123"},
		Input{Inputid: "NT22017072342178101", Harvestdate: "2017-07-23", Quantity: "421781", Inputdate: "2017-07-27", Personid: "FOOD123"},
		Input{Inputid: "NT32017072141156501", Harvestdate: "2017-07-21", Quantity: "411565", Inputdate: "2017-07-25", Personid: "FOOD123"},
		Input{Inputid: "NT4201707202627101", Harvestdate: "2017-07-20", Quantity: "26271", Inputdate: "2017-07-24", Personid: "FOOD123"},
	}
	i := 0
	for i < len(inputs) {
		fmt.Println("i is ", i)
		inputAsBytes, _ := json.Marshal(inputs[i])
		valueinputAsBytes, _ := json.Marshal(&Value{"INPUT" + strconv.Itoa(i+1)})
		stub.PutState("INPUT"+strconv.Itoa(i+1), inputAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(inputs[i].Inputid, valueinputAsBytes)
		fmt.Println("Added", inputs[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ := json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YINPUT", valueAsBytes)

	//初始化仓库
	warehouses := []Warehouse{
		Warehouse{Warehouseid: "CK5", Place: "湖北监利新沟银欣米业厂", Capacity: "9676", Standard: "标准", Personid: "FOOD123"},
		Warehouse{Warehouseid: "CK6", Place: "湖北监利新沟银欣米业厂", Capacity: "9370", Standard: "标准", Personid: "FOOD123"},
		Warehouse{Warehouseid: "CK7", Place: "湖北监利新沟银欣米业厂", Capacity: "8869", Standard: "标准", Personid: "FOOD123"},
		Warehouse{Warehouseid: "CK8", Place: "湖北监利新沟银欣米业厂", Capacity: "9641", Standard: "标准", Personid: "FOOD123"},
	}

	i = 0
	for i < len(warehouses) {
		fmt.Println("i is ", i)
		warehouseAsBytes, _ := json.Marshal(warehouses[i])
		valuewarehouseAsBytes, _ := json.Marshal(&Value{"WAREHOUSE" + strconv.Itoa(i+1)})
		//stub.PutState("Input"+strconv.Itoa(i+1), inputAsBytes)
		stub.PutState("WAREHOUSE"+strconv.Itoa(i+1), warehouseAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(warehouses[i].Warehouseid, valuewarehouseAsBytes)
		fmt.Println("Added", warehouses[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YWAREHOUSE", valueAsBytes)

	//初始化出库
	warehouse2feeds := []Warehouse2feed{
		Warehouse2feed{Warehouseid: "CK5", Feedid: "CK5DMJG0117082801", Personid: "FOOD123"},
		Warehouse2feed{Warehouseid: "CK6", Feedid: "CK6DMJG0117082801", Personid: "FOOD123"},
		Warehouse2feed{Warehouseid: "CK7", Feedid: "CK7DMJG0117083001", Personid: "FOOD123"},
		Warehouse2feed{Warehouseid: "CK8", Feedid: "CK8DMJG0117083001", Personid: "FOOD123"},
	}
	i = 0
	for i < len(warehouse2feeds) {
		fmt.Println("i is ", i)
		warehouse2feedAsBytes, _ := json.Marshal(warehouse2feeds[i])
		valuewarehouse2feedAsBytes, _ := json.Marshal(&Value{"WAREHOUSEFEED" + strconv.Itoa(i+1)})
		stub.PutState("WAREHOUSEFEED"+strconv.Itoa(i+1), warehouse2feedAsBytes)
		//放入主键ID，可以根据ID查询
		//stub.PutState(warehouse2feeds[i].Warehouseid+","+warehouse2feeds[i].Feedid, valuewarehouse2feedAsBytes)
		stub.PutState("XWAREHOUSEFEED"+warehouse2feeds[i].Warehouseid, valuewarehouse2feedAsBytes)
		stub.PutState("XWAREHOUSEFEED"+warehouse2feeds[i].Feedid, valuewarehouse2feedAsBytes)
		fmt.Println("Added", warehouse2feeds[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YWAREHOUSEFEED", valueAsBytes)

	//初始化种植入库批次转换
	plant2inputs := []Plant2input{
		Plant2input{Plantid: "NT1ZLY171ZJ0217020101XDJBJ031701200117031601HY1703180117040401", Inputid: "NT12017072243990601", Personid: "FOOD123"},
		Plant2input{Plantid: "NT2TLY83HN0117020101XDJBJ031701200117031701HY1703190117040501", Inputid: "NT22017072342178101", Personid: "FOOD123"},
		Plant2input{Plantid: "NT3ZJZ17ZJJL0217020101XDJBJ031701200117031501DP11703170117040301", Inputid: "NT32017072141156501", Personid: "FOOD123"},
		Plant2input{Plantid: "NT4JZ361JX0117020101XDJBJ031701200117031401DP21703160117040201", Inputid: "NT4201707202627101", Personid: "FOOD123"},
	}
	i = 0
	for i < len(plant2inputs) {
		fmt.Println("i is ", i)
		plant2inputAsBytes, _ := json.Marshal(plant2inputs[i])
		valueplant2inputAsBytes, _ := json.Marshal(&Value{"PLANTINPUT" + strconv.Itoa(i+1)})
		stub.PutState("PLANTINPUT"+strconv.Itoa(i+1), plant2inputAsBytes)
		//放入主键ID，可以根据ID查询
		//stub.PutState(plant2inputs[i].Plantid+","+plant2inputs[i].Inputid, valueplant2inputAsBytes)
		stub.PutState("XPLANTINPUT"+plant2inputs[i].Plantid, valueplant2inputAsBytes)
		stub.PutState("XPLANTINPUT"+plant2inputs[i].Inputid, valueplant2inputAsBytes)
		fmt.Println("Added", plant2inputs[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YPLANTINPUT", valueAsBytes)

	//初始化入库仓库批次转换
	input2warehouses := []Input2warehouse{
		Input2warehouse{Inputid: "NT12017072243990601", Warehouseid: "CK6", Personid: "FOOD123"},
		Input2warehouse{Inputid: "NT22017072342178101", Warehouseid: "CK7", Personid: "FOOD123"},
		Input2warehouse{Inputid: "NT22017072342178101", Warehouseid: "CK8", Personid: "FOOD123"},
		Input2warehouse{Inputid: "NT32017072141156501", Warehouseid: "CK6", Personid: "FOOD123"},
		Input2warehouse{Inputid: "NT4201707202627101", Warehouseid: "CK5", Personid: "FOOD123"},
	}
	i = 0
	for i < len(input2warehouses) {
		fmt.Println("i is ", i)
		input2warehouseAsBytes, _ := json.Marshal(input2warehouses[i])
		valueinput2warehouseAsBytes, _ := json.Marshal(&Value{"INPUTWAREHOUSE" + strconv.Itoa(i+1)})
		stub.PutState("INPUTWAREHOUSE"+strconv.Itoa(i+1), input2warehouseAsBytes)
		//放入主键ID，可以根据ID查询
		//stub.PutState(input2warehouses[i].Inputid+","+input2warehouses[i].Warehouseid, valueinput2warehouseAsBytes)

		//相同的入库ID可能对应不同的仓库，构造复合键
		indexinput := "inputid~warehouseid"
		inputidwarehouseidIndexKey, err := stub.CreateCompositeKey(indexinput, []string{input2warehouses[i].Inputid, input2warehouses[i].Warehouseid})
		if err != nil {
			return shim.Error(err.Error())
		}
		//  传个空字符为值，方便删除
		valueAsBytes := []byte{0x00}
		stub.PutState(inputidwarehouseidIndexKey, valueAsBytes)
		//相同的仓库id可能对应不同的入库，构造复合键
		indexwarehouse := "warehouseid~inputid"
		warehouseidinputidIndexKey, err := stub.CreateCompositeKey(indexwarehouse, []string{input2warehouses[i].Warehouseid, input2warehouses[i].Inputid})
		if err != nil {
			return shim.Error(err.Error())
		}
		stub.PutState(warehouseidinputidIndexKey, valueAsBytes)

		stub.PutState("XINPUTWAREHOUSE"+input2warehouses[i].Inputid, valueinput2warehouseAsBytes)
		stub.PutState("XINPUTWAREHOUSE"+input2warehouses[i].Warehouseid, valueinput2warehouseAsBytes)

		fmt.Println("Added", input2warehouses[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YINPUTWAREHOUSE", valueAsBytes)

	//初始化进料产品批次转换
	feed2products := []Feed2product{
		Feed2product{Feedid: "CK5DMJG0117082801", Productid: "JC1CMGUO01170828ZS01", Personid: "FOOD123"},
		Feed2product{Feedid: "CK6DMJG0117082801", Productid: "JC2CMGUO01170828HG01", Personid: "FOOD123"},
		Feed2product{Feedid: "CK7DMJG0117083001", Productid: "JC3CMS01170830ZM01", Personid: "FOOD123"},
		Feed2product{Feedid: "CK8DMJG0117083001", Productid: "JC4CMGEN01170830HZ01", Personid: "FOOD123"},
	}
	i = 0
	for i < len(feed2products) {
		fmt.Println("i is ", i)
		feed2productAsBytes, _ := json.Marshal(feed2products[i])
		valuefeed2productAsBytes, _ := json.Marshal(&Value{"FEEDPRODUCT" + strconv.Itoa(i+1)})
		stub.PutState("FEEDPRODUCT"+strconv.Itoa(i+1), feed2productAsBytes)
		//放入主键ID，可以根据ID查询
		//stub.PutState(feed2products[i].Feedid+","+feed2products[i].Productid, valuefeed2productAsBytes)
		stub.PutState("XFEEDPRODUCT"+feed2products[i].Feedid, valuefeed2productAsBytes)
		stub.PutState("XFEEDPRODUCT"+feed2products[i].Productid, valuefeed2productAsBytes)

		fmt.Println("Added", feed2products[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YFEEDPRODUCT", valueAsBytes)

	//初始化进料管理
	feeds := []Feed{
		Feed{Feedid: "CK5DMJG0117082801", Weight: "169440", Watercontent: "14.5", Brokenrice: "3.5", Qingmilv: "2.5", Date: "2017-08-28", Personid: "FOOD123"},
		Feed{Feedid: "CK6DMJG0117082801", Weight: "97780", Watercontent: "14.1", Brokenrice: "2", Qingmilv: "3", Date: "2017-08-28", Personid: "FOOD123"},
		Feed{Feedid: "CK7DMJG0117083001", Weight: "203330", Watercontent: "14.1", Brokenrice: "2", Qingmilv: "3", Date: "2017-08-30", Personid: "FOOD123"},
		Feed{Feedid: "CK8DMJG0117083001", Weight: "92060", Watercontent: "14.5", Brokenrice: "3", Qingmilv: "2", Date: "2017-08-30", Personid: "FOOD123"},
	}
	i = 0
	for i < len(feeds) {
		fmt.Println("i is ", i)
		feedAsBytes, _ := json.Marshal(feeds[i])
		valuefeedAsBytes, _ := json.Marshal(&Value{"FEED" + strconv.Itoa(i+1)})
		stub.PutState("FEED"+strconv.Itoa(i+1), feedAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(feeds[i].Feedid, valuefeedAsBytes)
		fmt.Println("Added", feeds[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YFEED", valueAsBytes)

	//初始化原料
	materials := []Material{
		Material{Materialid: "BSTMS0517082201", Kind: "白砂糖", Weight: "10", Source: "四川眉山王五糖业", Date: "2017-08-22", Personid: "FOOD123"},
		Material{Materialid: "XMFSJZ0617082501", Kind: "小麦粉", Weight: "10", Source: "河北石家庄晋务粉业", Date: "2017-08-25", Personid: "FOOD123"},
		Material{Materialid: "ZLYMAS0417082001", Kind: "棕榈油", Weight: "10", Source: "安徽马鞍山金龙鱼有限公司", Date: "2017-08-20", Personid: "FOOD123"},
	}
	i = 0
	for i < len(materials) {
		fmt.Println("i is ", i)
		materialAsBytes, _ := json.Marshal(materials[i])
		valuematerialAsBytes, _ := json.Marshal(&Value{"MATERIAL" + strconv.Itoa(i+1)})
		stub.PutState("MATERIAL"+strconv.Itoa(i+1), materialAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(materials[i].Materialid, valuematerialAsBytes)
		fmt.Println("Added", materials[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YMATERIAL", valueAsBytes)

	//初始化产品
	products := []Product{
		Product{Productid: "JC1CMGUO01170828ZS01", Name: "糙米果", Specification: "芝士味", Flavor: "2kg*1", Date: "2017-08-28", Personid: "FOOD123"},
		Product{Productid: "JC2CMGUO01170828HG01", Name: "糙米果", Specification: "黄瓜味", Flavor: "2kg*1", Date: "2017-08-28", Personid: "FOOD123"},
		Product{Productid: "JC3CMS01170830ZM01", Name: "糙米酥", Specification: "芝麻味", Flavor: "32g*12*6", Date: "2017-08-30", Personid: "FOOD123"},
		Product{Productid: "JC4CMGEN01170830HZ01", Name: "糙米羹", Specification: "红枣味", Flavor: "360g*12", Date: "2017-08-30", Personid: "FOOD123"},
	}
	i = 0
	for i < len(products) {
		fmt.Println("i is ", i)
		productAsBytes, _ := json.Marshal(products[i])
		valueproductAsBytes, _ := json.Marshal(&Value{"PRODUCT" + strconv.Itoa(i+1)})
		stub.PutState("PRODUCT"+strconv.Itoa(i+1), productAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(products[i].Productid, valueproductAsBytes)
		fmt.Println("Added", products[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YPRODUCT", valueAsBytes)

	//初始化原料产品批次转换
	material2products := []Material2product{
		Material2product{Materialid: "ZLYMAS0417082001", Productid: "JC1CMGUO01170828ZS01", Personid: "FOOD123"},
		Material2product{Materialid: "XMFSJZ0617082501", Productid: "JC2CMGUO01170828HG01", Personid: "FOOD123"},
		Material2product{Materialid: "BSTMS0517082201", Productid: "JC3CMS01170830ZM01", Personid: "FOOD123"},
		Material2product{Materialid: "BSTMS0517082201", Productid: "JC4CMGEN01170830HZ01", Personid: "FOOD123"},
	}
	i = 0
	for i < len(material2products) {
		fmt.Println("i is ", i)
		material2productAsBytes, _ := json.Marshal(material2products[i])
		valuematerial2productAsBytes, _ := json.Marshal(&Value{"MATERIALPRODUCT" + strconv.Itoa(i+1)})
		stub.PutState("MATERIALPRODUCT"+strconv.Itoa(i+1), material2productAsBytes)
		//放入主键ID，可以根据ID查询
		//stub.PutState(material2products[i].Materialid+","+material2products[i].Productid, valuematerial2productAsBytes)
		stub.PutState("XMATERIALPRODUCT"+material2products[i].Materialid, valuematerial2productAsBytes)
		stub.PutState("XMATERIALPRODUCT"+material2products[i].Productid, valuematerial2productAsBytes)
		fmt.Println("Added", material2products[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YMATERIALPRODUCT", valueAsBytes)

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

// ==================================================
// 根据Inputid读取一个入库仓库批次转换
// ==================================================
func (t *SimpleChaincode) readI2WByInputid(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var inputid string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	inputid = args[0]
	inputidResultsIterator, err := stub.GetStateByPartialCompositeKey("inputid~warehouseid", []string{inputid})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer inputidResultsIterator.Close()
	var i int
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	for i = 0; inputidResultsIterator.HasNext(); i++ {
		responseRange, err := inputidResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		//得到复合键中的inputid和warehouseid
		objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		returnedinputid := compositeKeyParts[0]
		returnedwarehouseid := compositeKeyParts[1]
		fmt.Printf("index:%s inputid:%s warehouseid:%s\n", objectType, returnedinputid, returnedwarehouseid)

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(returnedinputid)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":{\"inputid\":")
		buffer.WriteString("\"")
		buffer.WriteString(returnedinputid)
		buffer.WriteString("\"")
		buffer.WriteString(", \"warehouseid\":")
		buffer.WriteString("\"")
		buffer.WriteString(returnedwarehouseid)
		buffer.WriteString("\"}}")

		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")
	fmt.Printf("- readInput2warehouse:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// ==================================================
// 根据Warehouseid读取一个入库仓库批次转换
// ==================================================
func (t *SimpleChaincode) readI2WByWarehouseid(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var warehouseid string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	warehouseid = args[0]
	warehouseidResultsIterator, err := stub.GetStateByPartialCompositeKey("warehouseid~inputid", []string{warehouseid})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer warehouseidResultsIterator.Close()
	var i int
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	for i = 0; warehouseidResultsIterator.HasNext(); i++ {
		responseRange, err := warehouseidResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		//得到复合键中的inputid和warehouseid
		objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		returnedwarehouseid := compositeKeyParts[0]
		returnedinputid := compositeKeyParts[1]
		fmt.Printf("index:%s inputid:%s warehouseid:%s\n", objectType, returnedinputid, returnedwarehouseid)

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(returnedinputid)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":{\"inputid\":")
		buffer.WriteString("\"")
		buffer.WriteString(returnedinputid)
		buffer.WriteString("\"")
		buffer.WriteString(", \"warehouseid\":")
		buffer.WriteString("\"")
		buffer.WriteString(returnedwarehouseid)
		buffer.WriteString("\"}}")

		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")
	fmt.Printf("- readInput2warehouse:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
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

// ============================================================
// 新增入库
// ============================================================
func (t *SimpleChaincode) initInput(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0           1       2     3                4
	//  'NT12017072243990601', '2017-07-22', 439906, '2017-07-26', 'FOOD"
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init input")
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
	// ==== 检查主键是否存在 ====
	inputAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get input : " + err.Error())
	} else if inputAsBytes != nil {
		fmt.Println("The input  already exists: " + args[0])
		return shim.Error("The input  already exists: " + args[0])
	}

	// ==== 创建实体转成JSON ====
	input := &Input{args[0], args[1], args[2], args[3], args[4]}
	inputJSONasBytes, err := json.Marshal(input)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YINPUT")
	if err != nil {
		return shim.Error("Failed to get YINPUT:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YINPUT does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("INPUT"+strconv.Itoa(intValue), inputJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"INPUT" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YINPUT", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增仓库
// ============================================================
func (t *SimpleChaincode) initWarehouse(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2      3     4
	//  'CK5', '湖北监利新沟银欣米业厂', '9676', '标准','fw'
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init warehouse")
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
	// ==== 检查主键是否存在 ====
	warehouseAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get warehouse : " + err.Error())
	} else if warehouseAsBytes != nil {
		fmt.Println("The warehouse  already exists: " + args[0])
		return shim.Error("The warehouse  already exists: " + args[0])
	}

	// ==== 创建实体转成JSON ====
	warehouse := &Warehouse{args[0], args[1], args[2], args[3], args[4]}
	warehouseJSONasBytes, err := json.Marshal(warehouse)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YWAREHOUSE")
	if err != nil {
		return shim.Error("Failed to get YWAREHOUSE:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YWAREHOUSE does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("WAREHOUSE"+strconv.Itoa(intValue), warehouseJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"WAREHOUSE" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YWAREHOUSE", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增出库
// ============================================================
func (t *SimpleChaincode) initWarehouse2feed(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2
	// 'CK5', 'CK5DMJG0117082801', 'a'
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting3")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init warehouse2feed")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	// ==== 创建实体转成JSON ====
	warehouse2feed := &Warehouse2feed{args[0], args[1], args[2]}
	warehouse2feedJSONasBytes, err := json.Marshal(warehouse2feed)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YWAREHOUSEFEED")
	if err != nil {
		return shim.Error("Failed to get YWAREHOUSEFEED:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YWAREHOUSEFEED does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("WAREHOUSEFEED"+strconv.Itoa(intValue), warehouse2feedJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"WAREHOUSEFEED" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YWAREHOUSEFEED", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	//stub.PutState(args[0]+","+args[1], newvalueAsBytes)
	stub.PutState("XWAREHOUSEFEED"+args[0], newvalueAsBytes)
	stub.PutState("XWAREHOUSEFEED"+args[1], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增种植入库批次转换
// ============================================================
func (t *SimpleChaincode) initPlant2input(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	// 0  1          2
	// 'NT1ZLY171ZJ0217020101XDJBJ031701200117031601HY1703180117040401', 'NT12017072243990601',  'FOOD001'
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init warehouse2feed")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	// ==== 创建实体转成JSON ====
	plant2input := &Plant2input{args[0], args[1], args[2]}
	plant2inputJSONasBytes, err := json.Marshal(plant2input)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YPLANTINPUT")
	if err != nil {
		return shim.Error("Failed to get YPLANTINPUT:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YPLANTINPUT does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("PLANTINPUT"+strconv.Itoa(intValue), plant2inputJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"PLANTINPUT" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YPLANTINPUT", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	//stub.PutState(args[0]+","+args[1], newvalueAsBytes)
	stub.PutState("XPLANTINPUT"+args[0], newvalueAsBytes)
	stub.PutState("XPLANTINPUT"+args[1], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增入库仓库批次转换
// ============================================================
func (t *SimpleChaincode) initInput2warehouse(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2
	//'NT12017072243990601', 'CK6', 'FOOD'
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init warehouse2feed")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	// ==== 创建实体转成JSON ====
	input2warehouse := &Input2warehouse{args[0], args[1], args[2]}
	input2warehouseJSONasBytes, err := json.Marshal(input2warehouse)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YINPUTWAREHOUSE")
	if err != nil {
		return shim.Error("Failed to get YINPUTWAREHOUSE:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YINPUTWAREHOUSE does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("INPUTWAREHOUSE"+strconv.Itoa(intValue), input2warehouseJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"INPUTWAREHOUSE" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YINPUTWAREHOUSE", valueJSONasBytes)

	//构造复合键
	//相同的入库ID可能对应不同的仓库
	indexinput := "inputid~warehouseid"
	inputidwarehouseidIndexKey, err := stub.CreateCompositeKey(indexinput, []string{args[0], args[1]})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  传个空字符为值，方便删除
	valueAsBytes2 := []byte{0x00}
	stub.PutState(inputidwarehouseidIndexKey, valueAsBytes2)
	//相同的仓库id可能对应不同的入库，构造复合键
	indexwarehouse := "warehouseid~inputid"
	warehouseidinputidIndexKey, err := stub.CreateCompositeKey(indexwarehouse, []string{args[1], args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(warehouseidinputidIndexKey, valueAsBytes2)

	//放入主键ID，可以根据ID查询
	//stub.PutState(args[0]+","+args[1], newvalueAsBytes)
	stub.PutState("XINPUTWAREHOUSE"+args[0], newvalueAsBytes)
	stub.PutState("XINPUTWAREHOUSE"+args[1], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增进料产品批次转换
// ============================================================
func (t *SimpleChaincode) initFeed2product(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2
	// ''CK5DMJG0117082801', 'JC1CMGUO01170828ZS01', 'FOOD'
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init warehouse2feed")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	// ==== 创建实体转成JSON ====
	feed2product := &Feed2product{args[0], args[1], args[2]}
	feed2productJSONasBytes, err := json.Marshal(feed2product)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YFEEDPRODUCT")
	if err != nil {
		return shim.Error("Failed to get YFEEDPRODUCT:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YFEEDPRODUCT does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("FEEDPRODUCT"+strconv.Itoa(intValue), feed2productJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"FEEDPRODUCT" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YFEEDPRODUCT", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	//stub.PutState(args[0]+","+args[1], newvalueAsBytes)
	stub.PutState("XFEEDPRODUCT"+args[0], newvalueAsBytes)
	stub.PutState("XFEEDPRODUCT"+args[1], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增进料管理
// ============================================================
func (t *SimpleChaincode) initFeed(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2      3     4        5     6
	// 'CK5DMJG0117082801', '169440', '14.5', '3.5', '2.5', '2017-08-28','FOOD'
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init warehouse2feed")
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
		return shim.Error("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th argument must be a non-empty string")
	}
	// ==== 检查主键是否存在 ====
	feedAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get feed : " + err.Error())
	} else if feedAsBytes != nil {
		fmt.Println("The feed  already exists: " + args[0])
		return shim.Error("The feed  already exists: " + args[0])
	}

	// ==== 创建实体转成JSON ====
	feed := &Feed{args[0], args[1], args[2], args[3], args[4], args[5], args[6]}
	feedJSONasBytes, err := json.Marshal(feed)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YFEED")
	if err != nil {
		return shim.Error("Failed to get YFEED:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YFEED does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("FEED"+strconv.Itoa(intValue), feedJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"FEED" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YFEED", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增原料
// ============================================================
func (t *SimpleChaincode) initMaterial(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2      3     4        5
	// 'BSTMS0517082201', '白砂糖', '10', '四川眉山王五糖业', '2017-08-22', 'FOOD'
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init seed")
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
		return shim.Error("6th argument must be a non-empty string")
	}
	// ==== 检查主键是否存在 ====
	materialAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get material : " + err.Error())
	} else if materialAsBytes != nil {
		fmt.Println("The material  already exists: " + args[0])
		return shim.Error("The material  already exists: " + args[0])
	}

	// ==== 创建实体转成JSON ====
	material := &Material{args[0], args[1], args[2], args[3], args[4], args[5]}
	materialJSONasBytes, err := json.Marshal(material)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YMATERIAL")
	if err != nil {
		return shim.Error("Failed to get YMATERIAL:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YMATERIAL does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("MATERIAL"+strconv.Itoa(intValue), materialJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"MATERIAL" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YMATERIAL", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增产品
// ============================================================
func (t *SimpleChaincode) initProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2      3     4         5
	// 'JC1CMGUO01170828ZS01', '糙米果', '芝士味', '2kg*1', '2017-08-28', 'FOOD'
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init seed")
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
		return shim.Error("6th argument must be a non-empty string")
	}
	// ==== 检查主键是否存在 ====
	productAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get product : " + err.Error())
	} else if productAsBytes != nil {
		fmt.Println("The product  already exists: " + args[0])
		return shim.Error("The product  already exists: " + args[0])
	}

	// ==== 创建实体转成JSON ====
	product := &Product{args[0], args[1], args[2], args[3], args[4], args[5]}
	productJSONasBytes, err := json.Marshal(product)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YPRODUCT")
	if err != nil {
		return shim.Error("Failed to get YPRODUCT:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YPRODUCT does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("PRODUCT"+strconv.Itoa(intValue), productJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"PRODUCT" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YPRODUCT", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增原料产品批次转换
// ============================================================
func (t *SimpleChaincode) initMaterial2product(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2
	// ''ZLYMAS0417082001', 'JC1CMGUO01170828ZS01','FOOD'
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init seed")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	// ==== 创建实体转成JSON ====
	material2product := &Material2product{args[0], args[1], args[2]}
	material2productJSONasBytes, err := json.Marshal(material2product)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YMATERIALPRODUCT")
	if err != nil {
		return shim.Error("Failed to get YMATERIALPRODUCT:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YMATERIALPRODUCT does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("MATERIALPRODUCT"+strconv.Itoa(intValue), material2productJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"MATERIALPRODUCT" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YMATERIALPRODUCT", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	//stub.PutState(args[0]+","+args[1], newvalueAsBytes)
	stub.PutState("XMATERIALPRODUCT"+args[0], newvalueAsBytes)
	stub.PutState("XMATERIALPRODUCT"+args[1], newvalueAsBytes)
	return shim.Success(nil)
}
