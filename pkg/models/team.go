package models

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Auth Auth   `json:"auth"`
}

type Auth struct {
	Owner Owner `json:"owner"`
}

type Owner struct {
	Groups []string `json:"groups"`
	Users  []string `json:"users"`
}
