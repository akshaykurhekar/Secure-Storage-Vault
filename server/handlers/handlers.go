package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/db"
	models "server/models"
	"strconv"
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
func CreateCredential(res http.ResponseWriter, req *http.Request) {
	// define model
	var credentials models.Credentials
	if err := json.NewDecoder(req.Body).Decode(&credentials); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	DB := db.OpenConn()
	if DB == nil {
		http.Error(res, "Database Init fail..!", http.StatusInternalServerError)
		return
	}

	// Marshal CredentialDetails to a JSON string to store in the database
	credentialJSON, err := json.Marshal(credentials.Credential)
	if err != nil {
		http.Error(res, "Failed to encode credentials", http.StatusInternalServerError)
		return
	}
   fmt.Println("Credential DATA:", credentials, credentialJSON)

	// write insert query
	insertVault, err := DB.Prepare("INSERT INTO credentials (vid, credential) VALUES (?,?)"); 
	if err != nil {
		http.Error(res, "Falied to prepare insert query", http.StatusInternalServerError)
		return
	}
	
	defer insertVault.Close()

	// // exec query

	result, err := insertVault.Exec(credentials.VId, string(credentialJSON))
	if err != nil {
		http.Error(res, "Falied to Exec insert query..!", http.StatusInternalServerError)
		return
	}	

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(res, "Falied to get credentials Id..!", http.StatusInternalServerError)
		return
	}	
	credentials.Id = id

	// send response
	fmt.Println("Created New credentials entry:", credentials)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200);
	json.NewEncoder(res).Encode(credentials)
}

// - getCredentialByVaultId:vid - get
func GetCredentialByVaultId(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) // Get URL parameters
    vid := params["vid"]
	// define model
	var credentials []models.Credentials

	DB := db.OpenConn()
	if DB == nil {
		http.Error(res, "Database Init fail..!", http.StatusInternalServerError)
		return
	}

	// Check if uid is a valid integer
    intVid, err := strconv.ParseInt(vid, 10, 64)
    if err != nil {
        http.Error(res, "Invalid user ID format", http.StatusBadRequest)
        return
    }

	// selete query
	rows, err := DB.Query("SELECT id, vid, credential FROM credentials WHERE vid = ?", intVid);
	if err != nil {
		http.Error(res, "Failed to query table...!"+ err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	//scan rows
	for rows.Next() {
		var credentialRecord models.Credentials
		var credentialJSON string
		if err := rows.Scan( &credentialRecord.Id, &credentialRecord.VId, &credentialJSON); err != nil {
			http.Error(res, "Failed to scan credential record", http.StatusInternalServerError);
			return
		}

		// Unmarshal JSON if credential is a JSON object
		if err := json.Unmarshal([]byte(credentialJSON), &credentialRecord.Credential); err != nil {
			fmt.Println("Failed to unmarshal credential JSON:", err)
			http.Error(res, "Failed to unmarshal credential JSON", http.StatusInternalServerError)
			return
		}

		credentials = append(credentials, credentialRecord)
		fmt.Println("Credential record",credentialRecord);
	}

	if err := rows.Err(); err != nil {
		http.Error(res, "Error in row iteration", http.StatusInternalServerError);
		return
	}

		//send response
	// res.WriteHeader(http.StatusFound)
	json.NewEncoder(res).Encode(credentials)
}

// - updateCredentialById:id put - { "credential": { name:"xyz", cid:"ejgweufifgh" } }
func UpdateCredential(res http.ResponseWriter, req *http.Request) {
	
	params := mux.Vars(req) // Get URL parameters
	id := params["id"]

	var credentials models.Credentials
	if err := json.NewDecoder(req.Body).Decode(&credentials); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	DB := db.OpenConn()
	if DB == nil {
		http.Error(res, "Database Init fail..!", http.StatusInternalServerError)
		return
	}
	
	// Marshal CredentialDetails to a JSON string to store in the database
	credentialJSON, err := json.Marshal(credentials.Credential)
	if err != nil {
		http.Error(res, "Failed to encode credentials", http.StatusInternalServerError)
		return
	}
   fmt.Println("Credential DATA:", credentials, credentialJSON)

	// write insert query
	insertVault, err := DB.Prepare("UPDATE credentials SET vid = ?, credential = ? WHERE id = ?"); 
	if err != nil {
		http.Error(res, "Falied to prepare credentials update query", http.StatusInternalServerError)
		return
	}
	
	defer insertVault.Close()

	// // exec query

	result, err := insertVault.Exec( credentials.VId, string(credentialJSON), id)
	if err != nil {
		http.Error(res, "Falied to Exec credentials update query..!", http.StatusInternalServerError)
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
    json.NewEncoder(res).Encode(map[string]string{"message": "credentials Item updated successfully"})
}

// - deleteCredential:id delete
func DeleteCredential(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) // Get URL parameters
    id := params["id"]

	DB := db.OpenConn()
	if DB == nil {
		http.Error(res, "Database Init fail..!", http.StatusInternalServerError)
		return
	}

	// write insert query
	query, err := DB.Prepare("DELETE FROM credentials WHERE id = ?"); 
	if err != nil {
		http.Error(res, "Falied to prepare Delete credentials query", http.StatusInternalServerError)
		return
	}
	
	defer query.Close()

	result, err := query.Exec(id)
	if err != nil {
		http.Error(res, "Falied to Delete credentials by this id : "+ id, http.StatusInternalServerError)
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
    json.NewEncoder(res).Encode(map[string]string{"message": "credentials Item deleted successfully"})
}