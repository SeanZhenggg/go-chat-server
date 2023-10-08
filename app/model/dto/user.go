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
	Country     string `json:"country"`
	Address     string `json:"address"`
	RegionCode  string `json:"region_code"`
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
	Country     *string `json:"country"`
	Address     *string `json:"address"`
	RegionCode  *string `json:"region_code"`
	PhoneNumber *string `json:"phone_number"`
}
