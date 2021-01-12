package models

// Device entity
type Device struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	Labels         []string               `json:"labels"`
	Description    string                 `json:"description"`
	Location       string                 `json:"location"`
	Created        int64                  `json:"created"`
	Modified       int64                  `json:"modified"`
	Origin         int64                  `json:"origin"`
	AdminState     string                 `json:"adminState"`
	OperatingState string                 `json:"operatingState"`
	AutoEvents     []*AutoEvent           `json:"autoEvents"`
	LastConnected  int64                  `json:"lastConnected"`
	LastReported   int64                  `json:"lastReported"`
	Service        *Service               `json:"service"`
	Profile        *DeviceProfile         `json:"profile"`
	Protocols      map[string]interface{} `json:"protocols" faker:"-"`
}

// Addressable entity
type Addressable struct {
	Created  int64  `json:"created"`
	Modified int64  `json:"modified"`
	Origin   int64  `json:"origin"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Method   string `json:"method"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Path     string `json:"path"`
	BaseURL  string `json:"baseURL"`
	URL      string `json:"url"`
}

// Service entity
type Service struct {
	Created        int64        `json:"created"`
	Modified       int64        `json:"modified"`
	Origin         int64        `json:"origin"`
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	OperatingState string       `json:"operatingState"`
	Addressable    *Addressable `json:"addressable"`
	AdminState     string       `json:"adminState"`
}

// Value entity
type Value struct {
	Type          string `json:"type"`
	ReadWrite     string `json:"readWrite"`
	Minimum       string `json:"minimum"`
	Maximum       string `json:"maximum"`
	DefaultValue  string `json:"defaultValue"`
	Mask          uint64 `json:"mask"`
	Shift         uint64 `json:"shift"`
	Scale         int64  `json:"scale"`
	Offset        int64  `json:"offset"`
	Base          int64  `json:"base"`
	Assertion     string `json:"assertion"`
	FloatEncoding string `json:"floatEncoding"`
	MediaType     string `json:"mediaType"`
}

// Unit entity
type Unit struct {
	Type         string `json:"type"`
	ReadWrite    string `json:"readWrite"`
	DefaultValue string `json:"defaultValue"`
}

// Propertie entity
type Propertie struct {
	Value *Value `json:"value"`
	Units *Unit  `json:"units"`
}

// DeviceResource entity
type DeviceResource struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Tag         string     `json:"tag"`
	Properties  *Propertie `json:"properties"`
}

// CoreCommandGetResponse entity
type CoreCommandGetResponse struct {
	Code           string   `json:"code"`
	ExpectedValues []string `json:"expectedValues"`
	Description    string   `json:"description"`
}

// CoreCommandGet entity
type CoreCommandGet struct {
	Path      string                    `json:"path"`
	Responses []*CoreCommandGetResponse `json:"responses"`
}

// CoreCommandPutResponse entity
type CoreCommandPutResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// CoreCommandPut entity
type CoreCommandPut struct {
	Path           string                    `json:"path"`
	Responses      []*CoreCommandPutResponse `json:"responses"`
	ParameterNames []string                  `json:"parameterNames"`
}

// CoreCommand entity
type CoreCommand struct {
	Name string          `json:"name"`
	Get  *CoreCommandGet `json:"get"`
	Put  *CoreCommandPut `json:"put"`
}

// DeviceProfile entity
type DeviceProfile struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Labels          []string          `json:"labels"`
	Description     string            `json:"description"`
	Created         int64             `json:"created"`
	Modified        int64             `json:"modified"`
	Manufacturer    string            `json:"manufacturer"`
	Model           string            `json:"model"`
	DeviceResources []*DeviceResource `json:"deviceResources"`
	DeviceCommands  []*DeviceCommand  `json:"deviceCommands"`
	CoreCommands    []*CoreCommand    `json:"coreCommands"`
}

// AutoEvent entity
type AutoEvent struct {
	Frequency string `json:"frequency"`
	Onchange  bool   `json:"onchange"`
	Resource  string `json:"resource"`
}

// DeviceCommand entity
type DeviceCommand struct {
	Name string              `json:"name"`
	Get  []*DeviceCommandGet `json:"get"`
	Set  []*DeviceCommandSet `json:"set"`
}

// DeviceCommandGet entity
type DeviceCommandGet struct {
	Operation      string            `json:"operation"`
	Object         string            `json:"object"`
	DeviceResource string            `json:"deviceResource"`
	Mappings       map[string]string `json:"mappings"`
}

// DeviceCommandSet entity
type DeviceCommandSet struct {
	Operation      string            `json:"operation"`
	Object         string            `json:"object"`
	DeviceResource string            `json:"deviceResource"`
	Parameter      string            `json:"parameter"`
	Mappings       map[string]string `json:"mappings"`
}
