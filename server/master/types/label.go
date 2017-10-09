package types

type Label struct {
	System map[string]interface{} `json:"sys"`
	User   map[string]interface{} `json:"user"`
	Other  map[string]interface{} `json:"other"`
}
