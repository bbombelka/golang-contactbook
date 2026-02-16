package types

type FullUser struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	Address     Address `json:"address"`
	PhoneNumber string  `json:"phoneNumber"`
}

type ShortUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Users []FullUser
