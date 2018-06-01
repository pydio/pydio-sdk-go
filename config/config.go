package config

// SdkConfig contains necessary data for connection to Pydio 8 API
type SdkConfig struct {	
	// Full Http(s) URL to the server	
	Host     string `json:"host"`
	BasePath string `json:"basepath"`

	Schemes  []string `json:"schemas"`
	SkipVerify bool   `json:"skipVerify"`

	// Pydio User Authentication
	User     string `json:"user"`
	Password string `json:"password"`
}

var (
	DefaultConfig *SdkConfig
)
