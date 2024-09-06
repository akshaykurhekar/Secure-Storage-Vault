# Backend using Go lang

## Database table

    vault - { id - p.k, name, password, desc }
    credentials - { id - p.k, vid - f.k, credential:text } => credential= ex:{"name":"xyz", cid:"sbdkjfhbs"}

## Api List Or endpoints

    - getAllVaults
        - getVaultById:id - get
        - createVault: post { name, password, desc }
        - updateVault:id put { name, password, desc }
        - deleteVault:id delete
    - getAllCredentials
        - getCredentialByVaultId:vid - get
        - createCredential: post - { vid, credential: { name:"xyz", cid:"ejgweufifgh" } }
        - updateCredentialById:id put - { "credential": { name:"xyz", cid:"ejgweufifgh" } }
        - deleteCredentialById:id delete 

## Go commands

$ go mod init project_name

$ go run main.go

## pacakages used

- github.com/gorilla/mux@v1.8.1 

- database/sql
- log
- modernc.org/sqlite@v1.30.1

## SQL query 

## create vault
"INSERT INTO vault (name, password, desc) VALUES (?,?,?)"

## get vault
"SELECT id, name, password, desc FROM vault"

## create credential
"INSERT INTO credentials (uid, credential) VALUES (?,?)"

## get credential by vaultId
"SELECT id, uid, credential FROM credentials where uid = ?",intUid
