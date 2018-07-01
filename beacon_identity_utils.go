package main

import (
	"encoding/json"
	"fmt"
)

func createQueryString(params *IdentityParams) (qstring string, err error) {
	//ex: {"selector":{"CharityID":"marble","Status":1}
	var selector = IdentityParamSelector{Selector: *params}
	serialized, err := json.Marshal(selector)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	qstring = string(serialized)
	return qstring, nil
}

func applyIdentityModsFromParam(src *IdentityParams, dest *Identity) {
	if "" != src.AID {
		dest.AID = src.AID
	}
	if "" != src.UserType {
		dest.UserType = src.UserType
	}
	if "" != src.UserName {
		dest.UserName = src.UserName
	}
	if "" != src.Password {
		dest.Password = src.Password
	}
	if "" != src.Name {
		dest.Name = src.Name
	}
	if "" != src.WalletAddr {
		dest.WalletAddr = src.WalletAddr
	}
	if "" != src.Email {
		dest.Email = src.Email
	}
	if "" != src.MobileNo {
		dest.MobileNo = src.MobileNo
	}
	if "" != src.Description {
		dest.Description = src.Description
	}
	if "" != src.ApprvlStatus {
		dest.ApprvlStatus = src.ApprvlStatus
	}
	if "" != src.Surname {
		dest.Surname = src.Surname
	}
	if "" != src.MMN {
		dest.MMN = src.MMN
	}
}
