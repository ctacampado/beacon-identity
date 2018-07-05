package main

import (
	"encoding/json"
	"log"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/satori/go.uuid"
)

func addIdentity(fargs CCFuncArgs) pb.Response {
	log.Printf("starting addIdentity\n")

	u := uuid.Must(uuid.NewV4())
	var ustr = u.String()

	id := Identity{AID: ustr, ApprvlStatus: "NEW"}

	err := json.Unmarshal([]byte(fargs.req.Params), &id)
	if err != nil {
		return shim.Error("[addIdentity] Error unable to unmarshall msg: " + err.Error())
	}

	log.Printf("[addIdentity ] identity info: %+v\n", id)

	bytes, err := json.Marshal(id)
	if err != nil {
		log.Printf("[addIdentity] Could not marshal campaign info object: %+v\n", err)
		return shim.Error(err.Error())
	}

	err = fargs.stub.PutState(id.AID, bytes)
	if err != nil {
		log.Printf("[addIdentity] Error storing data in the ledger %+v\n", err)
		return shim.Error(err.Error())
	}

	log.Println("- end addIdentity")
	return shim.Success(bytes) //change nil to appropriate response
}
