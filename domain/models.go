package domain

type RequestGenerate struct {
	Number string `json:"number"`
}

// RequestWithCode used in req and response both
type RequestWithCode struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
}

type ResponseVerified struct {
	VerifiedAt int `json:"verifiedAt"`
}
