package response

type ResponseApi struct {
	ResponseCode string      `json:"responseCode"`
	ResponseDesc string      `json:"responseDesc"`
	Data         interface{} `json:"data"`
}
