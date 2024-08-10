package models

type Vault struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Desc string `json:"desc"`
}

type Credential struct {
	Name string `json:"name"`
	Cid string `json:"cid"`
}

type Credentials struct {
	Id int64 `json:"id"`
	VId int64 `json:"vid"`
	Credential Credential `json:"credential"`
}