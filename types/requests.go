package types

type AddUserRequestBody struct {
	Name        *string  `json:"name"`
	Username    *string  `json:"username"`
	Address     *Address `json:"address"`
	PhoneNumber *string  `json:"phoneNumber"`
}

type UpdateUserRequestBody struct {
	Name        *string  `json:"name"`
	Username    *string  `json:"username"`
	Address     *Address `json:"address"`
	PhoneNumber *string  `json:"phoneNumber"`
}
