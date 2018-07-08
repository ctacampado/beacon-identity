package main

import (
	"encoding/json"
	"log"
	"time"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/satori/go.uuid"
)

func addIdentity(fargs CCFuncArgs) pb.Response {
	log.Printf("starting addIdentity\n")

	u := uuid.Must(uuid.NewV4())
	var ustr = u.String()

	id := Identity{AID: ustr, ApprvlStatus: "NEW", DateCreated: string(time.Now().Format("2006-Jan-02"))}

	err := json.Unmarshal([]byte(fargs.msg.Params), &id)
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
	fargs.msg.Data = string(bytes)
	log.Printf("fargs: %+v\n", fargs.msg)
	rspbytes, err := json.Marshal(fargs.msg)
	if err != nil {
		log.Printf("[addCampaign] Could not marshal fargs object: %+v\n", err)
		return shim.Error(err.Error())
	}
	fargs.stub.SetEvent("newidentity", rspbytes)
	log.Printf("rspbytes: %+v\n", rspbytes)
	log.Println("- end addIdentity")
	return shim.Success(rspbytes) //change nil to appropriate response
}
