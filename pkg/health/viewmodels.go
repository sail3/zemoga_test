package health

// Here we have all the models that are used as inputs or outputs from the transports.

// HealthService represents a service in the response from handler Health.
type HealthService struct {
	Name  string `json:"name"`
	Alive bool   `json:"alive"`
}

// HealthResponse is the response of the handler Health.
type HealthResponse struct {
	Services []HealthService `json:"services"`
}
