package views

type Response struct {
	Code int `json:"code"`
	Body interface{} `json:"body"`
}

type Users struct {
	Userid string `json:"userid"`
	Name string `json:"name"`
}
