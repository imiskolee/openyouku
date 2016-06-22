package openyouku

type ResponseError struct {
	Code     int    `json:"code"`
	Provider string `json:"provider"`
	Desc     string `json:"desc"`
}

type Response struct {
	Error ResponseError `json:"e"`
	Cost  float64       `json:"cost"`
	Data  interface{}   `json:"data"`
}
