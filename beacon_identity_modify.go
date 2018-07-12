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
	qparm := IdentityParams{}
	err := json.Unmarshal([]byte(fargs.msg.Params), &qparm)
	if err != nil {
		return shim.Error("[getIdentity] Error unable to unmarshall Params: " + err.Error())
	}

	qp := IdentityParams{AID: qparm.AID}
	qpbytes, err := json.Marshal(qp)
	if err != nil {
		log.Printf("[addIdentity] Could not marshal campaign info object: %+v\n", err)
		return shim.Error(err.Error())
	}

	var qresp = QRsp{}
	qr := getIdentity(CCFuncArgs{function: "modifyIdentity", stub: fargs.stub, msg: Message{Params: string(qpbytes)}})
	err = json.Unmarshal([]byte(qr.Payload), &qresp)
	if err != nil {
		return shim.Error("[getIdentity] Error unable to unmarshall Payload: " + err.Error())
	}

	id := Identity{}
	err = json.Unmarshal([]byte(qresp.Elem[0].Value), &id)
	if err != nil {
		return shim.Error("[getIdentity] Error unable to unmarshall Elem[0].Value: " + err.Error())
	}

	applyIdentityModsFromParam(&qparm, &id)

	apbytes, err := json.Marshal(id)
	if err != nil {
		log.Printf("[addIdentity] Could not marshal campaign info object: %+v\n", err)
		return shim.Error(err.Error())
	}

	err = fargs.stub.PutState(id.AID, apbytes)
	if err != nil {
		log.Printf("[addIdentity] Error storing data in the ledger %+v\n", err)
		return shim.Error(err.Error())
	}

	fargs.msg.Data = string(apbytes)
	rspbytes, err := json.Marshal(fargs.msg)
	if err != nil {
		log.Printf("[addIdentity] Error msg JSON marshal%+v\n", err)
		return shim.Error(err.Error())
	}
	log.Printf("- end modifyIdentity")
	fargs.stub.SetEvent("modidentity", rspbytes)
	log.Printf("rspbytes: %+v\n", rspbytes)
	return shim.Success(rspbytes) //change nil to appropriate response
}
