package main

import (
	"encoding/json"
	"log"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func modifyIdentity(fargs CCFuncArgs) pb.Response {
	log.Printf("starting modifyIdentity\n")

	//get identity to be modified
	var qparm = &IdentityParams{AID: fargs.req.AID}
	qpbytes, err := json.Marshal(*qparm)
	if err != nil {
		log.Printf("[addIdentity] Could not marshal campaign info object: %+v\n", err)
		return shim.Error(err.Error())
	}

	var qresp = &QRsp{}
	qr := getIdentity(CCFuncArgs{req: Message{Params: string(qpbytes)}})
	err = json.Unmarshal([]byte(qr.Payload), qresp)
	if err != nil {
		return shim.Error("[getIdentity] Error unable to unmarshall msg: " + err.Error())
	}

	var id = &Identity{}
	err = json.Unmarshal([]byte(qresp.Elem[0].Value), id)
	if err != nil {
		return shim.Error("[getIdentity] Error unable to unmarshall msg: " + err.Error())
	}

	applyIdentityModsFromParam(qparm, id)

	apbytes, err := json.Marshal(*id)
	if err != nil {
		log.Printf("[addIdentity] Could not marshal campaign info object: %+v\n", err)
		return shim.Error(err.Error())
	}

	err = fargs.stub.PutState(id.AID, apbytes)
	if err != nil {
		log.Printf("[addIdentity] Error storing data in the ledger %+v\n", err)
		return shim.Error(err.Error())
	}

	return shim.Success(nil) //change nil to appropriate response
}
