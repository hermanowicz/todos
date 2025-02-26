package defs

type DefaultResponse struct {
	Status       int64  `json:"status"`
	ResponseBody string `json:"response_body"`
	YourIP       string `json:"your_ip"`
}
