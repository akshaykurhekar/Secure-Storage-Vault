package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/db"
	models "server/models"
)


func TestFunc(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Testing..Handler..")
	res.WriteHeader(200);
	json.NewEncoder(res).Encode("Hi... we are in test handler funcation...")
}

// - getAllVaults
func GetAllVaults(res http.ResponseWriter, req *http.Request) {
	
	// define model
	var vaults []models.Vault

	DB := db.OpenConn()
	if DB == nil {
		http.Error(res, "Database Init fail..!", http.StatusInternalServerError)
		return
	}

	// selete query
	Rows, err := DB.Query("SELECT id, name, password, desc FROM vault ");
	if err != nil {
		http.Error(res, "Failed to query table...!"+ err.Error(), http.StatusInternalServerError)
		return
	}
	defer Rows.Close()
	
	// scan that row
	for Rows.Next() {
		var vault models.Vault
		if err := Rows.Scan(&vault.Id, &vault.Name, &vault.Password, &vault.Desc); err != nil {
			http.Error(res, "Failed to scan rows...!", http.StatusInternalServerError)
			return
		}

		vaults = append(vaults, vault)		
	} 
	// send response

	res.WriteHeader(200);
	json.NewEncoder(res).Encode(vaults)
}

// - createVault: post { name, password, desc }
func CreateVault(res http.ResponseWriter, req *http.Request) {
	// define model
	var vault models.Vault
	if err := json.NewDecoder(req.Body).Decode(&vault); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	DB := db.OpenConn()
	if DB == nil {
		http.Error(res, "Database Init fail..!", http.StatusInternalServerError)
		return
	}

	// write insert query
	insertVault, err := DB.Prepare("INSERT INTO vault (name, password, desc) VALUES (?,?,?)"); 
	if err != nil {
		http.Error(res, "Falied to prepare insert query", http.StatusInternalServerError)
		return
	}
	
	defer insertVault.Close()

	// // exec query

	result, err := insertVault.Exec(vault.Name, vault.Password, vault.Desc)
	if err != nil {
		http.Error(res, "Falied to Exec insert query..!", http.StatusInternalServerError)
		return
	}	

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(res, "Falied to get vault Id..!", http.StatusInternalServerError)
		return
	}	
	vault.Id = id

	// send response
	fmt.Println("Created New vault entry:", vault)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200);
	json.NewEncoder(res).Encode(vault)
}

// - updateVault:id put { name, password, desc }
// - deleteVault:id delete

// - createCredential: post - { vid, credential: { name:"xyz", cid:"ejgweufifgh" } }
// - getAllCredentials

// - getCredentialByVaultId:vid - get
// - updateCredentialById:id put - { "credential": { name:"xyz", cid:"ejgweufifgh" } }
// - deleteCredentialById:id delete 