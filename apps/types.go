package apps

type appsResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Service string `json:"service"`
	Name    string `json:"name"`
	Error   bool   `json:"error,omitempty"`
	Url     string `json:"url,omitempty"`
}
