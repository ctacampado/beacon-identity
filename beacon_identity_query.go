package main

import (
	"encoding/json"
	"log"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func getIdentity(fargs CCFuncArgs) pb.Response {
	log.Println("[getIdentity]starting getIdentity")

	var qparams = &IdentityParams{}
	err := json.Unmarshal([]byte(fargs.req.Params), qparams)
	if err != nil {
		return shim.Error("[getIdentity] Error unable to unmarshall msg: " + err.Error())
	}
	log.Printf("[getIdentity] creating query string")
	qstring, err := createQueryString(qparams)
	if err != nil {
		return shim.Error("[getIdentity] Error unable to create query string: " + err.Error())
	}
	log.Printf("[getIdentity] query using querystring: %+v\n", qparams)
	resultsIterator, err := fargs.stub.GetQueryResult(qstring)
	log.Printf("- getQueryResultForQueryString resultsIterator:\n%+v\n", resultsIterator)
	defer resultsIterator.Close()
	if err != nil {
		return shim.Error("[getIdentity] Error unable to GetQueryResult: " + err.Error())
	}

	var qresp = QRsp{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("[getIdentity] Error unable to get next item in iterator: " + err.Error())
		}

		q := QRes{Key: queryResponse.Key, Value: string(queryResponse.Value)}
		qresp.Elem = append(qresp.Elem, q)
	}

	log.Printf("- getQueryResultForQueryString queryResult:\n%+v\n", qresp)
	log.Printf("- getQueryResultForQueryString querystring:\n%s\n", qstring)
	log.Printf("- getQueryResultForQueryString qparams:\n%+v\n", qparams)

	qr, err := json.Marshal(qresp)
	if err != nil {
		return shim.Error("[getCOCampaigns] Error unable to Marshall qresp: " + err.Error())
	}

	log.Println("- end getIdentity")
	return shim.Success(qr)
}
