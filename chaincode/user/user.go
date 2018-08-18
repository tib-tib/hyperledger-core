package main

import (
    "fmt"
    "encoding/json"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)

type UserChaincode struct {
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
}

func (t *UserChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
  fmt.Println("User Chaincode Init")
  return shim.Success(nil)
}

func (t *UserChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("User Chaincode Invoke")
	function, args := stub.GetFunctionAndParameters()

	fmt.Println(function)

	switch function {
	case "create":
		return t.create(stub, args)
	case "delete":
		return t.delete(stub, args)
	case "get":
		return t.get(stub, args)
	default:
		return peer.Response{Status: 500, Message: "Invalid invoke function name. Expecting \"create\", \"delete\", \"get\"", Payload: []byte{}}
	}
}

func (t *UserChaincode) create(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("Create user")

	if len(args) != 2 {
		return peer.Response{Status: 500, Message: "Invalid number of arguments. Expecting 2.", Payload: []byte{}}
	}

	userInStore, err := stub.GetState(args[0])
	if err != nil {
		fmt.Println("Error: Failed to create user. Can check if exists")
		return peer.Response{Status: 500, Message: "Failed to create user. Can check if exists", Payload: []byte{}}
	}
	if len(userInStore) != 0 {
		fmt.Println("Error: User exists")
		return peer.Response{Status: 423, Message: "User exists", Payload: []byte{}}
	}
	var user = User{Id: args[0], Name: args[1]}

	userAsBytes, _ := json.Marshal(user)
	fmt.Println(userAsBytes)

	err = stub.PutState(args[0], userAsBytes)
	if err != nil {
		return shim.Error("Failed to create user")
	}

	fmt.Println("User created successfully")

	return shim.Success(nil)
}

func (t *UserChaincode) get(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("Get user")

	if len(args) != 1 {
		fmt.Println("Error: Incorrect number of arguments. Expecting 1")
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userAsBytes, err := stub.GetState(args[0])
	if err != nil {
		fmt.Println("Error: Incorrect number of arguments. Expecting 1")
		return shim.Error("Failed to get user")
	}

	if len(userAsBytes) == 0 {
		return peer.Response{Status: 404, Message: "User doesn't exist", Payload: []byte{}}
	}

	var user User
	json.Unmarshal(userAsBytes, &user)
	fmt.Println(user)

	userString := "{\"id\":\"" + user.Id + "\",\"name\":\"" + user.Name + "\"}"
	fmt.Println(userString)
	b := []byte(userString)
	fmt.Println(b)
	return shim.Success(userAsBytes)
}

func (t *UserChaincode) delete(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("Delete user")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("Failed to delete user")
	}

	return shim.Success(nil)

}

func main() {
	err := shim.Start(new(UserChaincode))
	if err != nil {
		fmt.Printf("Error starting UserChaincode: %s", err)
	}
}

