package main

import (
	"encoding/json"
	"log"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//Init Chaincode function
func (t *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	log.Printf("Nothing to Initialize!")
	return shim.Success(nil)
}

//Invoke Chaincode function
func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	//initialize function map
	t.FMap = map[string]ccfunc{
		"addIdentity":    addIdentity,
		"getIdentity":    getIdentity,
		"modifyIdentity": modifyIdentity,
	}

	function, args := stub.GetFunctionAndParameters()
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}
	log.Printf("In Invoke with function %s", function)

	//extract message from args
	err := json.Unmarshal([]byte(args[0]), &t.Msg)
	if err != nil {
		return shim.Error("[Invoke] unable to unmarshall args[0]: " + err.Error())
	}

	if t.FMap[function] != nil {
		fargs := CCFuncArgs{function: function, req: t.Msg, stub: stub}
		return t.FMap[function](fargs)
	}

	log.Printf("BEACON IDENTITY Received unknown invoke function name - %s" + function)
	return shim.Error("BEACON IDENTITY Received unknown invoke function name - '" + function + "'")
}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		log.Printf("Error starting BEACON IDENTITY chaincode: %s", err)
	} else {
		log.Printf("BEACON IDENTITY Chaincode successfully started")
	}
}
