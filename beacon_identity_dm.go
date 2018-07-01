package main

import (
	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//--------------------------------------------------------------------------
//Start adding Chaincode-related Structures here

//CCFuncArgs common cc func args
type CCFuncArgs struct {
	function string
	req      Message
	stub     shim.ChaincodeStubInterface
}

type ccfunc func(args CCFuncArgs) pb.Response

//Chaincode cc structure
type Chaincode struct {
	FMap map[string]ccfunc //ccfunc map
	Msg  Message           //data
}

//Message Charity Org Chain Code Message Structure
type Message struct {
	CID    string `json:"CID"`    //ClientID --for websocket push (event-based messaging readyness)
	AID    string `json:"AID"`    //ActorID (Donor ID/Charity Org ID/Auditor ID/etc.)
	Type   string `json:"type"`   //Chaincode Function
	Params string `json:"params"` //Function Parameters
}

//End of Chaincode-related Structures
//--------------------------------------------------------------------------
//Start adding Query Parameter (Parm) Structures here

//QRes Structure for Query Response from ledger
type QRes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//QRsp Structure for Query Response to chaincode invoke
type QRsp struct {
	Elem []QRes `json:"elem"`
}

//IdentityParams Structure for Query Parameters
type IdentityParams struct {
	AID          string `json:"AID,omitempty"`
	UserType     string `json:"UserType,omitempty"`
	UserName     string `json:"UserName,omitempty"`
	Password     string `json:"Password,omitempty"`
	Name         string `json:"Name,omitempty"`
	WalletAddr   string `json:"WalletAddr,omitempty"`
	Email        string `json:"Email,omitempty"`
	MobileNo     string `json:"MobileNo,omitempty"`
	Description  string `json:"Description,omitempty"`
	ApprvlStatus string `json:"ApprvlStatus, omitempty"`
	Surname      string `json:"Surname, omitempty"`
	MMN          string `json:"MMN, omitempty"`
}

//IdentityParamSelector Structure for Query Selector
type IdentityParamSelector struct {
	Selector IdentityParams `json:"selector"`
}

//End of Query Paramter Structures
//--------------------------------------------------------------------------
//Start adding Data Models here

/****Approval Status
	NEW - newly added identity
	APPRVD - approved identity
	RJCTD  - rejected identity
	RVKD   - revoked identity
****/

//Identity data model
type Identity struct {
	AID          string `json:"AID"`
	UserType     string `json:"UserType"`
	UserName     string `json:"UserName"`
	Password     string `json:"Password"`
	Name         string `json:"Name"`
	WalletAddr   string `json:"WalletAddr,omitempty"`
	Email        string `json:"Email"`
	MobileNo     string `json:"MobileNo,omitempty"`
	ApprvlStatus string `json:"ApprvlStatus"`
	Description  string `json:"Description,omitempty"`
	Surname      string `json:"Surname, omitempty"`
	MMN          string `json:"MMN, omitempty"`
}

//End of Data Models
//--------------------------------------------------------------------------
