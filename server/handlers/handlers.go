package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/db"
	models "server/models"

	"github.com/gorilla/mux"
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

// =================== HOMEWORK ====================

// - updateVault:id put { name, password, desc }
func UpdateVault(res http.ResponseWriter, req *http.Request) {
	
	params := mux.Vars(req) // Get URL parameters
	id := params["id"]

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
	insertVault, err := DB.Prepare("UPDATE vault SET name = ?, password = ?, desc = ? WHERE id = ?"); 
	if err != nil {
		http.Error(res, "Falied to prepare vault update query", http.StatusInternalServerError)
		return
	}
	
	defer insertVault.Close()

	// // exec query

	result, err := insertVault.Exec(vault.Name, vault.Password, vault.Desc, id)
	if err != nil {
		http.Error(res, "Falied to Exec update query..!", http.StatusInternalServerError)
		return
	}	

	rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        res.WriteHeader(http.StatusNotFound)
        json.NewEncoder(res).Encode(map[string]string{"error": "Item not found"})
        return
    }

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200);
    json.NewEncoder(res).Encode(map[string]string{"message": "Item updated successfully"})
}

// - deleteVault:id delete
func DeleteVaultById(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) // Get URL parameters
    id := params["id"]

	DB := db.OpenConn()
	if DB == nil {
		http.Error(res, "Database Init fail..!", http.StatusInternalServerError)
		return
	}

	// write insert query
	query, err := DB.Prepare("DELETE FROM vault WHERE id = ?"); 
	if err != nil {
		http.Error(res, "Falied to prepare Delete Vault query", http.StatusInternalServerError)
		return
	}
	
	defer query.Close()

	// // exec query

	result, err := query.Exec(id)
	if err != nil {
		http.Error(res, "Falied to Delete vault by this id : "+ id, http.StatusInternalServerError)
		return
	}	

	rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        res.WriteHeader(http.StatusNotFound)
        json.NewEncoder(res).Encode(map[string]string{"error": "Item not found"})
        return
    }

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200);
    json.NewEncoder(res).Encode(map[string]string{"message": "Item deleted successfully"})
}

// - createCredential: post - { vid, credential: { name:"xyz", cid:"ejgweufifgh" } }
// - getCredentialByVaultId:vid - get

// - updateCredentialById:id put - { "credential": { name:"xyz", cid:"ejgweufifgh" } }
// - deleteCredentialById:id delete 