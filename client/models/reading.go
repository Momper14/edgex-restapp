package models

// Reading entity
type Reading struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Created   int64  `json:"created"`
	Modified  int64  `json:"modified"`
	Origin    int64  `json:"origin"`
	Value     string `json:"value"`
	ValueType string `json:"valueType"`
	Pushed    int64  `json:"pushed"`
	Device    string `json:"device"`
}
