package mini

type ResponseBase struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseMiniProgram struct {
	ResponseBase

	Msg     string `json:"msg,omitempty"`
	ErrCode int    `json:"errcode,omitempty"`
	ErrMSG  string `json:"errmsg,omitempty"`

	ResultCode string `json:"resultcode,omitempty"`
	ResultMSG  string `json:"resultmsg,omitempty"`
}

type ResponseCode2Session struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`

	ResponseMiniProgram
}
