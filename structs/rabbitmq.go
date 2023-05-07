package structs

type MessagePayload struct {
	Id         int64       `json:"Id"`
	Command    string      `json:"Command"`
	Time       string      `json:"Time"`
	ModuleId   string      `json:"ModuleId"`
	Properties interface{} `json:"Properties"`
	Signature  string      `json:"Signature"`
	Data       interface{} `json:"Data"`
}
