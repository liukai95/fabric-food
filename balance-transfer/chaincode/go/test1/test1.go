package main

import (
   "github.com/hyperledger/fabric/core/chaincode/shim"
   pb "github.com/hyperledger/fabric/protos/peer"
   "fmt"
   "encoding/pem"
   "crypto/x509"
   "bytes"
)

type SimpleChaincode struct {
}

func main() {
   err := shim.Start(new(SimpleChaincode))
   if err != nil {
      fmt.Printf("Error starting Simple chaincode: %s", err)
   }
}
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
   return shim.Success(nil)
}


func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
   function, args := stub.GetFunctionAndParameters()
   fmt.Println("invoke is running " + function)
   if function == "cert" {//自定义函数名称
      return t.testCertificate(stub, args)//定义调用的函数
   }
   return shim.Error("Received unknown function invocation")
}
func (t *SimpleChaincode) testCertificate(stub shim.ChaincodeStubInterface, args []string) pb.Response{
   creatorByte,_:= stub.GetCreator()
   certStart := bytes.IndexAny(creatorByte, "-----")// Devin:I don't know why sometimes -----BEGIN is invalid, so I use -----
   if certStart == -1 {
      fmt.Errorf("No certificate found")
   }
   certText := creatorByte[certStart:]
   bl, _ := pem.Decode(certText)
   if bl == nil {
      fmt.Errorf("Could not decode the PEM structure")
   }
   fmt.Println(string(certText))
   cert, err := x509.ParseCertificate(bl.Bytes)
   if err != nil {
      fmt.Errorf("ParseCertificate failed")
   }
   fmt.Println(cert)
   uname:=cert.Subject.CommonName
   fmt.Println("Name:"+uname)
   return shim.Success([]byte("Called testCertificate "+uname))
}
