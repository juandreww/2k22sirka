package views

type Response struct {
	Code int `json:"code"`
	Body interface{} `json:"body"`
}

type Kelapa struct {
	Type2 string `json:"Type2"`
	Quantity float64 `json:"Quantity"`
}

type Users struct {
	Userid string `json:"userid"`
	Name string `json:"name"`
}
