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

//药品
type Drug struct {
	Drugid   string `json:"drugid"`   //药品id
	Name     string `json:"name"`     //名称
	Dosage   string `json:"dosage"`   //剂型
	Standard string `json:"standard"` //生产标准
	Effect   string `json:"effect"`   //作用
	Personid string `json:"personid"` //质检员ID
}

//种子
type Seed struct {
	Seedid   string `json:"seedid"`   //种子ID
	Variety  string `json:"variety"`  //品种
	Type     string `json:"type"`     //类型
	Personid string `json:"personid"` //质检员ID
}

//秧苗
type Seedling struct {
	Seedlingid   string `json:"seedlingid"`   //秧苗id
	Nurseryplace string `json:"nurseryplace"` //育秧地点
	Startdate    string `json:"startdate"`    //育秧起始日期
	Personid     string `json:"personid"`     //质检员ID
}

//种子浸药
type Seedsoakdrug struct {
	Seedid        string `json:"seedid"`        //种子id
	Drugid        string `json:"drugid"`        //药品id
	Concentration string `json:"concentration"` //浓度（ml/L）
	Startdate     string `json:"startdate"`     //起始日期
	Enddata       string `json:"enddata"`       //结束日期
	Personid      string `json:"personid"`      //质检员ID
}

//秧苗喷药
type Seedlingspraydrug struct {
	Seedlingid string `json:"seedlingid"` //秧苗id
	Drugid     string `json:"drugid"`     //药品id
	Dosage     string `json:"dosage"`     //用量(ml/kg)
	Data       string `json:"data"`       //日期
	Personid   string `json:"personid"`   //质检员ID
}

//种植
type Plant struct {
	Plantid   string `json:"plantid"`   //种植id
	Place     string `json:"place"`     //种植地点
	Startdate string `json:"startdate"` //种植起始日期
	Personid  string `json:"personid"`  //质检员ID
}

//种植用药
type Plantusedrug struct {
	Plantid  string `json:"plantid"`  //种植id
	Drugid   string `json:"drugid"`   //药品id
	Dosage   string `json:"dosage"`   //浓度（ml/L）
	Effect   string `json:"effect"`   //作用
	Date     string `json:"date"`     //日期
	Personid string `json:"personid"` //质检员ID
}

//种子秧苗
type Seed2seedling struct {
	Seedid     string `json:"seedid"`     //种子id
	Seedlingid string `json:"seedlingid"` //秧苗id
	Personid   string `json:"personid"`   //质检员ID
}

//秧苗种植
type Seedling2plant struct {
	Seedlingid string `json:"seedlingid"` //秧苗id
	Plantid    string `json:"plantid"`    //种植id
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
	} else if function == "initDrug" { //新增一个药品
		return t.initDrug(stub, args)
	} else if function == "initSeed" { //新增一个种子
		return t.initSeed(stub, args)
	} else if function == "initSeedling" { //新增一个秧苗
		return t.initSeedling(stub, args)
	} else if function == "initSeedsoakdrug" { //新增一个种子浸药
		return t.initSeedsoakdrug(stub, args)
	} else if function == "initSeedlingspraydrug" { //新增一个种子喷药
		return t.initSeedlingspraydrug(stub, args)
	} else if function == "initPlant" { //新增一个种植
		return t.initPlant(stub, args)
	} else if function == "initPlantusedrug" { //新增一个种植用药
		return t.initPlantusedrug(stub, args)
	} else if function == "initSeed2seedling" { //新增一个种子秧苗
		return t.initSeed2seedling(stub, args)
	} else if function == "initSeedling2plant" { //新增一个秧苗种植
		return t.initSeedling2plant(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *SimpleChaincode) initLedger(stub shim.ChaincodeStubInterface) pb.Response {

	//初始化药品
	drugs := []Drug{
		Drug{Drugid: "CCJHB0317012501", Name: "除虫剂", Dosage: "液态", Standard: "国标", Effect: "除虫", Personid: "FOOD"},
		Drug{Drugid: "FHCBJ0317021001", Name: "氟环唑", Dosage: "悬浮剂", Standard: "Q/SSCC101-2011", Effect: "杀菌", Personid: "FOOD"},
		Drug{Drugid: "KJJBJ0317020101", Name: "咯菌腈", Dosage: "悬浮剂", Standard: "Q/320583GQB", Effect: "杀菌", Personid: "FOOD"},
		Drug{Drugid: "MXABJ0317021801", Name: "咪鲜胺", Dosage: "水乳剂", Standard: "GB22625-2008", Effect: "杀菌", Personid: "FOOD"},
		Drug{Drugid: "XDJBJ0317012001", Name: "消毒剂", Dosage: "液态", Standard: "国标", Effect: "消毒", Personid: "FOOD"},
	}
	i := 0
	for i < len(drugs) {
		fmt.Println("i is ", i)
		drugAsBytes, _ := json.Marshal(drugs[i])
		//stub.PutState("Person"+strconv.Itoa(i+1), personAsBytes)
		valuedrugAsBytes, _ := json.Marshal(&Value{"DRUG" + strconv.Itoa(i+1)})
		stub.PutState("DRUG"+strconv.Itoa(i+1), drugAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(drugs[i].Drugid, valuedrugAsBytes)
		fmt.Println("Added", drugs[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ := json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YDRUG", valueAsBytes)

	//初始化种子
	seeds := []Seed{
		Seed{Seedid: "JZ361JX0117020101", Variety: "江早361", Type: "早稻", Personid: "FOOD123"},
		Seed{Seedid: "TLY83HN0117020101", Variety: "潭两优83", Type: "早稻", Personid: "FOOD123"},
		Seed{Seedid: "ZJZ17ZJJ0217020101", Variety: "中嘉早17", Type: "早稻", Personid: "FOOD123"},
		Seed{Seedid: "ZLY171ZJ0217020101", Variety: "株两优171", Type: "早稻", Personid: "FOOD123"},
	}
	i = 0
	for i < len(seeds) {
		fmt.Println("i is ", i)
		seedAsBytes, _ := json.Marshal(seeds[i])
		valueseedAsBytes, _ := json.Marshal(&Value{"SEED" + strconv.Itoa(i+1)})
		stub.PutState("SEED"+strconv.Itoa(i+1), seedAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(seeds[i].Seedid, valueseedAsBytes)
		fmt.Println("Added", seeds[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YSEED", valueAsBytes)

	//初始化秧苗
	seedlings := []Seedling{
		Seedling{Seedlingid: "JZ361JX0117020101XDJBJ031701200117031401DP217031601", Nurseryplace: "2号大棚", Startdate: "2017-03-16", Personid: "FOOD123"},
		Seedling{Seedlingid: "TLY83HN0117020101XDJBJ031701200117031701HY17031901", Nurseryplace: "红阳育秧工厂", Startdate: "2017-03-19", Personid: "FOOD123"},
		Seedling{Seedlingid: "ZJZ17ZJJL0217020101XDJBJ031701200117031501DP117031701", Nurseryplace: "1号大棚", Startdate: "2017-03-17", Personid: "FOOD123"},
		Seedling{Seedlingid: "ZLY171ZJ0217020101XDJBJ031701200117031601HY17031801", Nurseryplace: "红阳育秧工厂", Startdate: "2017-03-18", Personid: "FOOD123"},
	}
	i = 0
	for i < len(seedlings) {
		fmt.Println("i is ", i)
		seedlingAsBytes, _ := json.Marshal(seedlings[i])
		valueseedlingAsBytes, _ := json.Marshal(&Value{"SEEDLING" + strconv.Itoa(i+1)})
		stub.PutState("SEEDLING"+strconv.Itoa(i+1), seedlingAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(seedlings[i].Seedlingid, valueseedlingAsBytes)
		fmt.Println("Added", seedlings[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YSEEDLING", valueAsBytes)

	//初始化种子浸药
	seedsoakdrugs := []Seedsoakdrug{
		Seedsoakdrug{Seedid: "JZ361JX0117020101", Drugid: "XDJBJ0317012001", Concentration: "29.5", Startdate: "2017-03-14", Enddata: "2017-03-15", Personid: "FOOD123"},
		Seedsoakdrug{Seedid: "TLY83HN0117020101", Drugid: "XDJBJ0317012001", Concentration: "32", Startdate: "2017-03-17", Enddata: "2017-03-18", Personid: "FOOD123"},
		Seedsoakdrug{Seedid: "ZJZ17ZJJ0217020101", Drugid: "XDJBJ0317012001", Concentration: "30", Startdate: "2017-03-15", Enddata: "2017-03-16", Personid: "FOOD123"},
		Seedsoakdrug{Seedid: "ZLY171ZJ0217020101", Drugid: "XDJBJ0317012001", Concentration: "33", Startdate: "2017-03-16", Enddata: "2017-03-17", Personid: "FOOD123"},
	}
	i = 0
	for i < len(seedsoakdrugs) {
		fmt.Println("i is ", i)
		seedsoakdrugAsBytes, _ := json.Marshal(seedsoakdrugs[i])
		valueseedsoakdrugAsBytes, _ := json.Marshal(&Value{"SEEDSOAKDRUG" + strconv.Itoa(i+1)})
		stub.PutState("SEEDSOAKDRUG"+strconv.Itoa(i+1), seedsoakdrugAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(seedsoakdrugs[i].Seedid+","+seedsoakdrugs[i].Drugid, valueseedsoakdrugAsBytes)

		stub.PutState("XSEEDSOAKDRUG"+seedsoakdrugs[i].Seedid, valueseedsoakdrugAsBytes)
		stub.PutState("XSEEDSOAKDRUG"+seedsoakdrugs[i].Drugid, valueseedsoakdrugAsBytes)
		fmt.Println("Added", seedsoakdrugs[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YSEEDSOAKDRUG", valueAsBytes)

	//初始化秧苗喷药
	seedlingspraydrugs := []Seedlingspraydrug{
		Seedlingspraydrug{Seedlingid: "JZ361JX0117020101XDJBJ031701200117031401DP217031601", Drugid: "KJJBJ0317020101", Dosage: "20.5", Data: "2017-03-16", Personid: "FOOD123"},
		Seedlingspraydrug{Seedlingid: "TLY83HN0117020101XDJBJ031701200117031701HY17031901", Drugid: "KJJBJ0317020101", Dosage: "21", Data: "2017-03-19", Personid: "FOOD123"},
		Seedlingspraydrug{Seedlingid: "ZJZ17ZJJL0217020101XDJBJ031701200117031501DP117031701", Drugid: "KJJBJ0317020101", Dosage: "22", Data: "2017-03-17", Personid: "FOOD123"},
		Seedlingspraydrug{Seedlingid: "ZLY171ZJ0217020101XDJBJ031701200117031601HY17031801", Drugid: "KJJBJ0317020101", Dosage: "20.5", Data: "2017-03-18", Personid: "FOOD123"},
	}
	i = 0
	for i < len(seedlingspraydrugs) {
		fmt.Println("i is ", i)
		seedlingspraydrugAsBytes, _ := json.Marshal(seedlingspraydrugs[i])
		valueseedlingspraydrugAsBytes, _ := json.Marshal(&Value{"SEEDLINGSPRAYDRUG" + strconv.Itoa(i+1)})
		stub.PutState("SEEDLINGSPRAYDRUG"+strconv.Itoa(i+1), seedlingspraydrugAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(seedlingspraydrugs[i].Seedlingid+","+seedlingspraydrugs[i].Drugid, valueseedlingspraydrugAsBytes)
		stub.PutState("XSEEDLINGSPRAYDRUG"+seedlingspraydrugs[i].Seedlingid, valueseedlingspraydrugAsBytes)
		stub.PutState("XSEEDLINGSPRAYDRUG"+seedlingspraydrugs[i].Drugid, valueseedlingspraydrugAsBytes)
		fmt.Println("Added", seedlingspraydrugs[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YSEEDLINGSPRAYDRUG", valueAsBytes)

	//初始化种植
	plants := []Plant{
		Plant{Plantid: "NT1ZLY171ZJ0217020101XDJBJ031701200117031601HY1703180117040401", Place: "1号农田", Startdate: "2017-04-04", Personid: "FOOD123"},
		Plant{Plantid: "NT2TLY83HN0117020101XDJBJ031701200117031701HY1703190117040501", Place: "2号农田", Startdate: "2017-04-05", Personid: "FOOD123"},
		Plant{Plantid: "NT3ZJZ17ZJJL0217020101XDJBJ031701200117031501DP11703170117040301", Place: "3号农田", Startdate: "2017-04-03", Personid: "FOOD123"},
		Plant{Plantid: "NT4JZ361JX0117020101XDJBJ031701200117031401DP21703160117040201", Place: "4号农田", Startdate: "2017-04-05", Personid: "FOOD123"},
	}
	i = 0
	for i < len(plants) {
		fmt.Println("i is ", i)
		plantAsBytes, _ := json.Marshal(plants[i])
		valueplantAsBytes, _ := json.Marshal(&Value{"PLANT" + strconv.Itoa(i+1)})
		stub.PutState("PLANT"+strconv.Itoa(i+1), plantAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(plants[i].Plantid, valueplantAsBytes)
		fmt.Println("Added", plants[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YPLANT", valueAsBytes)

	//初始化种植用药
	plantusedrugs := []Plantusedrug{
		Plantusedrug{Plantid: "NT1ZLY171ZJ0217020101XDJBJ031701200117031601HY1703180117040401", Drugid: "CCJHB0317012501", Dosage: "9", Effect: "防治稻瘟、叶枯、纹枯", Date: "2017-06-18", Personid: "FOOD123"},
		Plantusedrug{Plantid: "NT2TLY83HN0117020101XDJBJ031701200117031701HY1703190117040501", Drugid: "CCJHB0317012501", Dosage: "8.7", Effect: "防治稻瘟、叶枯、纹枯", Date: "2017-06-18", Personid: "FOOD123"},
		Plantusedrug{Plantid: "NT3ZJZ17ZJJL0217020101XDJBJ031701200117031501DP11703170117040301", Drugid: "CCJHB0317012501", Dosage: "9", Effect: "防治稻瘟、叶枯、纹枯", Date: "2017-06-18", Personid: "FOOD123"},
		Plantusedrug{Plantid: "NT4JZ361JX0117020101XDJBJ031701200117031401DP21703160117040201", Drugid: "CCJHB0317012501", Dosage: "8.5", Effect: "防治稻瘟、叶枯、纹枯", Date: "2017-06-18", Personid: "FOOD123"},
	}
	i = 0
	for i < len(plantusedrugs) {
		fmt.Println("i is ", i)
		plantusedrugAsBytes, _ := json.Marshal(plantusedrugs[i])
		valueplantusedrugAsBytes, _ := json.Marshal(&Value{"PLANTUSEDRUG" + strconv.Itoa(i+1)})
		stub.PutState("PLANTUSEDRUG"+strconv.Itoa(i+1), plantusedrugAsBytes)
		//放入主键ID，可以根据ID查询
		stub.PutState(plantusedrugs[i].Plantid+","+plantusedrugs[i].Drugid, valueplantusedrugAsBytes)
		stub.PutState("XPLANTUSEDRUG"+plantusedrugs[i].Plantid, valueplantusedrugAsBytes)
		stub.PutState("XPLANTUSEDRUG"+plantusedrugs[i].Drugid, valueplantusedrugAsBytes)
		fmt.Println("Added", plantusedrugs[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YPLANTUSEDRUG", valueAsBytes)

	//初始化种子秧苗
	seed2seedlings := []Seed2seedling{
		Seed2seedling{Seedid: "JZ361JX0117020101", Seedlingid: "JZ361JX0117020101XDJBJ031701200117031401DP217031601", Personid: "FOOD123"},
		Seed2seedling{Seedid: "TLY83HN0117020101", Seedlingid: "TLY83HN0117020101XDJBJ031701200117031701HY17031901", Personid: "FOOD123"},
		Seed2seedling{Seedid: "ZJZ17ZJJ0217020101", Seedlingid: "ZJZ17ZJJL0217020101XDJBJ031701200117031501DP117031701", Personid: "FOOD123"},
		Seed2seedling{Seedid: "ZLY171ZJ0217020101", Seedlingid: "ZLY171ZJ0217020101XDJBJ031701200117031601HY17031801", Personid: "FOOD123"},
	}
	i = 0
	for i < len(seed2seedlings) {
		fmt.Println("i is ", i)
		seed2seedlingAsBytes, _ := json.Marshal(seed2seedlings[i])
		valueseed2seedlingAsBytes, _ := json.Marshal(&Value{"SEEDSEEDLING" + strconv.Itoa(i+1)})
		stub.PutState("SEEDSEEDLING"+strconv.Itoa(i+1), seed2seedlingAsBytes)
		//放入主键ID，可以根据ID查询
		//stub.PutState(seed2seedlings[i].Seedid+","+seed2seedlings[i].Seedlingid, valueseed2seedlingAsBytes)
		stub.PutState("XSEEDSEEDLING"+seed2seedlings[i].Seedid, valueseed2seedlingAsBytes)
		stub.PutState("XSEEDSEEDLING"+seed2seedlings[i].Seedlingid, valueseed2seedlingAsBytes)
		fmt.Println("Added", seed2seedlings[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YSEEDSEEDLING", valueAsBytes)

	//初始化秧苗种植
	seedling2plants := []Seedling2plant{
		Seedling2plant{Seedlingid: "ZLY171ZJ0217020101XDJBJ031701200117031601HY17031801", Plantid: "NT1ZLY171ZJ0217020101XDJBJ031701200117031601HY1703180117040401", Personid: "FOOD123"},
		Seedling2plant{Seedlingid: "TLY83HN0117020101XDJBJ031701200117031701HY17031901", Plantid: "NT2TLY83HN0117020101XDJBJ031701200117031701HY1703190117040501", Personid: "FOOD123"},
		Seedling2plant{Seedlingid: "ZJZ17ZJJL0217020101XDJBJ031701200117031501DP117031701", Plantid: "NT3ZJZ17ZJJL0217020101XDJBJ031701200117031501DP11703170117040301", Personid: "FOOD123"},
		Seedling2plant{Seedlingid: "JZ361JX0117020101XDJBJ031701200117031401DP217031601", Plantid: "NT4JZ361JX0117020101XDJBJ031701200117031401DP21703160117040201", Personid: "FOOD123"},
	}
	i = 0
	for i < len(seedling2plants) {
		fmt.Println("i is ", i)
		seedling2plantAsBytes, _ := json.Marshal(seedling2plants[i])
		valueseedling2plantAsBytes, _ := json.Marshal(&Value{"SEEDLINGPLANT" + strconv.Itoa(i+1)})
		stub.PutState("SEEDLINGPLANT"+strconv.Itoa(i+1), seedling2plantAsBytes)
		//放入主键ID，可以根据ID查询
		//stub.PutState(seedling2plants[i].Seedlingid+","+seedling2plants[i].Plantid, valueseedling2plantAsBytes)
		stub.PutState("XSEEDLINGPLANT"+seedling2plants[i].Seedlingid, valueseedling2plantAsBytes)
		stub.PutState("XSEEDLINGPLANT"+seedling2plants[i].Plantid, valueseedling2plantAsBytes)
		fmt.Println("Added", seedling2plants[i])
		i = i + 1
	}
	//存放最大标号的下一个，便于插入新数据
	valueAsBytes, _ = json.Marshal(&Value{strconv.Itoa(i)})
	stub.PutState("YSEEDLINGPLANT", valueAsBytes)

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
// 新增药品
// ============================================================
func (t *SimpleChaincode) initDrug(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2      3     4                5
	//  'XDJBJ0317012001', '消毒剂', '液态', '国标', '消毒', 'a'
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
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

	// ==== 检查主键是否存在 ====
	drugAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get drug : " + err.Error())
	} else if drugAsBytes != nil {
		fmt.Println("The drug  already exists: " + args[0])
		return shim.Error("The drug  already exists: " + args[0])
	}

	valueAsBytes, err := stub.GetState("YDRUG")
	if err != nil {
		return shim.Error("Failed to get YDRUG:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YDRUG does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)
	// ==== 创建实体转成JSON ====
	drug := &Drug{args[0], args[1], args[2], args[3], args[4], args[5]}
	drugJSONasBytes, err := json.Marshal(drug)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === 保存 ===
	err = stub.PutState("DRUG"+strconv.Itoa(intValue), drugJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	newvalueAsBytes, _ := json.Marshal(&Value{"DRUG" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YDRUG", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增种子
// ============================================================
func (t *SimpleChaincode) initSeed(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2      3
	// 'JZ361JX0117020101', '江早361', '早稻', 'a'
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
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
	// ==== 检查主键是否存在 ====
	seedAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get seed : " + err.Error())
	} else if seedAsBytes != nil {
		fmt.Println("The seed  already exists: " + args[0])
		return shim.Error("The seed  already exists: " + args[0])
	}

	// ==== 创建实体转成JSON ====
	seed := &Seed{args[0], args[1], args[2], args[3]}
	seedJSONasBytes, err := json.Marshal(seed)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YSEED")
	if err != nil {
		return shim.Error("Failed to get YSEED:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YSEED does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("SEED"+strconv.Itoa(intValue), seedJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"SEED" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YSEED", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0], newvalueAsBytes)

	return shim.Success(nil)
}

// ============================================================
// 新增秧苗
// ============================================================
func (t *SimpleChaincode) initSeedling(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	// 0  1          2      3
	// 'JZ361JX0117020101XDJBJ031701200117031401DP217031601', '2号大棚', '2017-03-16',  'FOOD001'
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
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
	// ==== 检查主键是否存在 ====
	seedlingAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get seedling : " + err.Error())
	} else if seedlingAsBytes != nil {
		fmt.Println("The seedling  already exists: " + args[0])
		return shim.Error("The seedling  already exists: " + args[0])
	}

	// ==== 创建实体转成JSON ====
	seedling := &Seedling{args[0], args[1], args[2], args[3]}
	seedlingJSONasBytes, err := json.Marshal(seedling)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YSEEDLING")
	if err != nil {
		return shim.Error("Failed to get YSEEDLING:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YSEEDLING does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("SEEDLING"+strconv.Itoa(intValue), seedlingJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"SEEDLING" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YSEEDLING", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0], newvalueAsBytes)

	return shim.Success(nil)
}

// ============================================================
// 新增种子侵药
// ============================================================
func (t *SimpleChaincode) initSeedsoakdrug(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2      3     4        5
	// 'JZ361JX0117020101', 'XDJBJ0317012001', '29.5', '2017-03-14', '2017-03-15', 'FOOD'
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

	// ==== 创建实体转成JSON ====
	seedsoakdrug := &Seedsoakdrug{args[0], args[1], args[2], args[3], args[4], args[5]}
	seedsoakdrugJSONasBytes, err := json.Marshal(seedsoakdrug)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YSEEDSOAKDRUG")
	if err != nil {
		return shim.Error("Failed to get YSEEDSOAKDRUG:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YSEEDSOAKDRUG does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("SEEDSOAKDRUG"+strconv.Itoa(intValue), seedsoakdrugJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"SEEDSOAKDRUG" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YSEEDSOAKDRUG", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0]+","+args[1], newvalueAsBytes)
	stub.PutState("XSEEDSOAKDRUG"+args[0], newvalueAsBytes)
	stub.PutState("XSEEDSOAKDRUG"+args[1], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增秧苗喷药
// ============================================================
func (t *SimpleChaincode) initSeedlingspraydrug(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2      3     4
	// 'JZ361JX0117020101XDJBJ031701200117031401DP217031601', 'KJJBJ0317020101', '20.5', '2017-03-16', 'FOOD'
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
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

	// ==== 创建实体转成JSON ====
	seedlingspraydrug := &Seedlingspraydrug{args[0], args[1], args[2], args[3], args[4]}
	seedlingspraydrugJSONasBytes, err := json.Marshal(seedlingspraydrug)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YSEEDLINGSPRAYDRUG")
	if err != nil {
		return shim.Error("Failed to get YSEEDLINGSPRAYDRUG:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YSEEDLINGSPRAYDRUG does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("SEEDLINGSPRAYDRUG"+strconv.Itoa(intValue), seedlingspraydrugJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"SEEDLINGSPRAYDRUG" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YSEEDLINGSPRAYDRUG", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0]+","+args[1], newvalueAsBytes)
	stub.PutState("XSEEDLINGSPRAYDRUG"+args[0], newvalueAsBytes)
	stub.PutState("XSEEDLINGSPRAYDRUG"+args[1], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增种植
// ============================================================
func (t *SimpleChaincode) initPlant(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2      3
	//  '('NT1ZLY171ZJ0217020101XDJBJ031701200117031601HY1703180117040401', '1号农田', '2017-04-04', 'FOOD'
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
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

	// ==== 检查主键是否存在 ====
	plantAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get plant : " + err.Error())
	} else if plantAsBytes != nil {
		fmt.Println("The plant  already exists: " + args[0])
		return shim.Error("The plant  already exists: " + args[0])
	}

	// ==== 创建实体转成JSON ====
	plant := &Plant{args[0], args[1], args[2], args[3]}
	plantJSONasBytes, err := json.Marshal(plant)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YPLANT")
	if err != nil {
		return shim.Error("Failed to get YPLANT:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YPLANT does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("PLANT"+strconv.Itoa(intValue), plantJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"PLANT" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YPLANT", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增种植用药
// ============================================================
func (t *SimpleChaincode) initPlantusedrug(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2      3     4        5
	// ''NT1ZLY171ZJ0217020101XDJBJ031701200117031601HY1703180117040401', 'CCJHB0317012501', '9', '防治稻瘟、叶枯、纹枯', '2017-06-18','FOOD'
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

	// ==== 创建实体转成JSON ====
	plantusedrug := &Plantusedrug{args[0], args[1], args[2], args[3], args[4], args[5]}
	plantusedrugJSONasBytes, err := json.Marshal(plantusedrug)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YPLANTUSEDRUG")
	if err != nil {
		return shim.Error("Failed to get YPLANTUSEDRUG:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YPLANTUSEDRUG does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("PLANTUSEDRUG"+strconv.Itoa(intValue), plantusedrugJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"PLANTUSEDRUG" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YPLANTUSEDRUG", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	stub.PutState(args[0]+","+args[1], newvalueAsBytes)
	stub.PutState("XPLANTUSEDRUG"+args[0], newvalueAsBytes)
	stub.PutState("XPLANTUSEDRUG"+args[1], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增种子秧苗
// ============================================================
func (t *SimpleChaincode) initSeed2seedling(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	// 0  1          2
	// 'ZLYMAS0417082001', 'JC1CMGUO01170828ZS01','FOOD'
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
	seed2seedling := &Seed2seedling{args[0], args[1], args[2]}
	seed2seedlingJSONasBytes, err := json.Marshal(seed2seedling)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YSEEDSEEDLING")
	if err != nil {
		return shim.Error("Failed to get YSEEDSEEDLING:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YSEEDSEEDLING does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("SEEDSEEDLING"+strconv.Itoa(intValue), seed2seedlingJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"SEEDSEEDLING" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YSEEDSEEDLING", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	//stub.PutState(args[0]+","+args[1], newvalueAsBytes)
	stub.PutState("XSEEDSEEDLING"+args[0], newvalueAsBytes)
	stub.PutState("XSEEDSEEDLING"+args[1], newvalueAsBytes)
	return shim.Success(nil)
}

// ============================================================
// 新增秧苗种植
// ============================================================
func (t *SimpleChaincode) initSeedling2plant(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// 0  1          2
	//'ZLYMAS0417082001', 'JC1CMGUO01170828ZS01','FOOD'
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
	seedling2plant := &Seedling2plant{args[0], args[1], args[2]}
	seedling2plantJSONasBytes, err := json.Marshal(seedling2plant)
	if err != nil {
		return shim.Error(err.Error())
	}

	valueAsBytes, err := stub.GetState("YSEEDLINGPLANT")
	if err != nil {
		return shim.Error("Failed to get YSEEDLINGPLANT:" + err.Error())
	} else if valueAsBytes == nil {
		return shim.Error("YSEEDLINGPLANT does not exist")
	}

	valueToTransfer := Value{}
	err = json.Unmarshal(valueAsBytes, &valueToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	//字符串转数字
	intValue, _ := strconv.Atoi(valueToTransfer.Valueid)

	// === 保存 ===
	stub.PutState("SEEDLINGPLANT"+strconv.Itoa(intValue), seedling2plantJSONasBytes)
	newvalueAsBytes, _ := json.Marshal(&Value{"SEEDLINGPLANT" + strconv.Itoa(intValue)})
	//保存新的最大值
	valueToTransfer.Valueid = strconv.Itoa(intValue + 1)
	valueJSONasBytes, _ := json.Marshal(valueToTransfer)
	stub.PutState("YSEEDLINGPLANT", valueJSONasBytes)
	//放入主键ID，可以根据ID查询
	//stub.PutState(args[0]+","+args[1], newvalueAsBytes)
	stub.PutState("XSEEDLINGPLANT"+args[0], newvalueAsBytes)
	stub.PutState("XSEEDLINGPLANT"+args[1], newvalueAsBytes)
	return shim.Success(nil)
}

