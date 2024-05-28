// Code generated by goctl. DO NOT EDIT.
package types

type Captcha struct {
	Server string `json:"server"`
	Token  string `json:"token"`
}

type CheckLoginReq struct {
	Token string `json:"token"`
}

type CheckLoginResp struct {
	IsValid bool `json:"isValid"`
}

type LoginReq struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Captcha  *Captcha `json:"captcha,optional"`
	Ip       string   `json:"ip,optional"`
}

type LoginResp struct {
	Username      string `json:"username"`
	Token         string `json:"token"`
	MemberLevel   string `json:"memberLevel"`
	RealName      string `json:"realName"`
	Country       string `json:"country"`
	Avatar        string `json:"avatar"`
	PromotionCode string `json:"promotionCode"`
	Id            int64  `json:"id"`
	LoginCount    int    `json:"loginCount"`
	SuperPartner  string `json:"superPartner"`
	MemberRate    int    `json:"memberRate"`
}

type RegisterReq struct {
	Username     string   `json:"username,optional"`
	Password     string   `json:"password,optional"`
	Captcha      *Captcha `json:"captcha,optional"`
	Phone        string   `json:"phone,optional"`
	Promotion    string   `json:"promotion,optional"`
	Code         string   `json:"code,optional"`
	Country      string   `json:"country,optional"`
	SuperPartner string   `json:"superPartner,optional"`
	Ip           string   `json:"ip,optional"`
}

type RegisterResp struct {
}

type SendCodeReq struct {
	Phone   string `json:"phone"`
	Country string `json:"country"`
}

type SendCodeResp struct {
}
