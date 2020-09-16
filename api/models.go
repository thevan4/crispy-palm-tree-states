package api

import "time"

// GetServicesRequest ...
type GetServicesRequest struct {
	ID string `json:"id"`
}

// TokenRequest ...
type TokenRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
	ID       string `json:"id"`
}

// TokenResponseOkay ...
type TokenResponseOkay struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ID           string `json:"id"`
}

// TokenResponseError ...
type TokenResponseError struct {
	Error string `json:"error"`
	ID    string `json:"id"`
}

// AdvancedHealthcheckParameters ...
type AdvancedHealthcheckParameters struct {
	NearFieldsMode  bool                   `json:"nearFieldsMode"`
	UserDefinedData map[string]interface{} `json:"userDefinedData"`
}

// ServerHealthcheck ...
type ServerHealthcheck struct {
	TypeOfCheck                   string                          `json:"typeOfCheck"`
	HealthcheckAddress            string                          `json:"healthcheckAddress"`
	AdvancedHealthcheckParameters []AdvancedHealthcheckParameters `json:"advancedHealthcheckParameters"`
}

// ServerApplicationWithStates ...
type ServerApplicationWithStates struct {
	ServerIP                    string            `json:"ip"`
	ServerPort                  string            `json:"port"`
	ServerHealthcheck           ServerHealthcheck `json:"serverHealthcheck"`
	IsUp                        bool              `json:"serverIsUp"`
	ServerСonfigurationCommands string            `json:"bashCommands"`
}

// ServiceHealthcheck ...
type ServiceHealthcheck struct {
	Type                 string        `json:"type"`
	Timeout              time.Duration `json:"timeout"`
	RepeatHealthcheck    time.Duration `json:"repeatHealthcheck"`
	PercentOfAlivedForUp int           `json:"percentOfAlivedForUp"`
}

// Service ...
type Service struct {
	ID                       string                        `json:"id"`
	ApplicationServers       []ServerApplicationWithStates `json:"applicationServers"`
	ServiceIP                string                        `json:"serviceIP"`
	ServicePort              string                        `json:"servicePort"`
	Healthcheck              ServiceHealthcheck            `json:"healthcheck"`
	JobCompletedSuccessfully bool                          `json:"jobCompletedSuccessfully"`
	ExtraInfo                string                        `json:"extraInfo"`
	BalanceType              string                        `json:"balanceType"`
	RoutingType              string                        `json:"routingType"`
	IsUp                     bool                          `json:"serviceIsUp"`
	Protocol                 string                        `json:"protocol"`
}

// GetAllServicesResponse ...
type GetAllServicesResponse struct {
	ID                       string    `json:"id"`
	JobCompletedSuccessfully bool      `json:"jobCompletedSuccessfully"`
	AllServices              []Service `json:"allServices,omitempty"`
	ExtraInfo                string    `json:"extraInfo,omitempty"`
}
