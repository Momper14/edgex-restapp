package models

// ValueDescriptor entity
type ValueDescriptor struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Labels       []string `json:"labels"`
	Description  string   `json:"description"`
	Created      int64    `json:"created"`
	Modified     int64    `json:"modified"`
	Origin       int64    `json:"origin"`
	Type         string   `json:"type"`
	UomLabel     string   `json:"uomLabel"`
	Min          string   `json:"min"`
	Max          string   `json:"max"`
	DefaultValue string   `json:"defaultValue"`
	Formatting   string   `json:"formatting"`
}
