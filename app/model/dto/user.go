package dto

type GetUserCond struct {
	Account string `uri:"account"`
}

type UserInfoResp struct {
	Id          uint   `json:"id"`
	Account     string `json:"account"`
	Nickname    string `json:"nickname"`
	Birthdate   string `json:"birthdate"`
	Gender      int    `json:"gender"`
	CountryCode string `json:"country_code"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type UserRegCond struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserLoginCond struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	Token string `json:"token"`
}

type UpdateUserIdCond struct {
	Id string `uri:"id"`
}

type UpdateUserInfoCond struct {
	Password    *string `json:"password"`
	Nickname    *string `json:"nickname"`
	Birthdate   *string `json:"birthdate"`
	Gender      *int    `json:"gender"`
	CountryCode *string `json:"country_code"`
	Address     *string `json:"address"`
	PhoneNumber *string `json:"phone_number"`
}
