package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func TestFunc(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Testing..Handler..")
	res.WriteHeader(200);
	json.NewEncoder(res).Encode("Hi... we are in test handler funcation...")
}

// - getAllVaults
func GetAllVaults(res http.ResponseWriter, req *http.Request) {

	res.WriteHeader(200);
	json.NewEncoder(res).Encode("Hi... we are in getAllVault funcation...")
}

// - getVaultById:id - get

// - createVault: post { name, password, desc }

// - updateVault:id put { name, password, desc }

// - deleteVault:id delete

// - getAllCredentials
// - getCredentialByVaultId:vid - get
// - createCredential: post - { vid, credential: { name:"xyz", cid:"ejgweufifgh" } }
// - updateCredentialById:id put - { "credential": { name:"xyz", cid:"ejgweufifgh" } }
// - deleteCredentialById:id delete 