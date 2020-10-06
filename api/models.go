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

type Service struct {
	IP                    string               `json:"ip" validate:"ipv4" swagger:"ignoreParam"`
	Port                  string               `json:"port" validate:"required" swagger:"ignoreParam"`
	IsUp                  bool                 `json:"isUp,omitempty" swagger:"ignoreParam"`
	BalanceType           string               `json:"balanceType" validate:"required" example:"rr"`
	RoutingType           string               `json:"routingType" validate:"required" example:"masquerading,tunneling"`
	Protocol              string               `json:"protocol" validate:"required" example:"tcp,udp"`
	AlivedAppServersForUp int                  `json:"alivedAppServersForUp" validate:"required,gt=0,lte=100"`
	HCType                string               `json:"hcType" validate:"required" example:"tcp"`
	HCRepeat              time.Duration        `json:"hcRepeat" validate:"required" example:"3000000000"`
	HCTimeout             time.Duration        `json:"hcTimeout" validate:"required" example:"1000000000"`
	HCNearFieldsMode      bool                 `json:"hcNearFieldsMode,omitempty"`
	HCUserDefinedData     map[string]string    `json:"hcUserDefinedData,omitempty"`
	HCRetriesForUP        int                  `json:"hcRetriesForUP" validate:"required,gt=0" example:"3"`
	HCRetriesForDown      int                  `json:"hcRetriesForDown" validate:"required,gt=0" example:"10"`
	ApplicationServers    []*ApplicationServer `json:"applicationServers" validate:"required,dive,required"`
}

type ApplicationServer struct {
	IP                  string `json:"ip" validate:"ipv4" example:"1.1.1.1"`
	Port                string `json:"port" validate:"required" example:"1111"`
	IsUp                bool   `json:"isUp,omitempty" swagger:"ignoreParam"`
	HCAddress           string `json:"hcAddress" validate:"required" example:"http://1.1.1.1:1234"`
	ExampleBashCommands string `json:"exampleBashCommands,omitempty" swagger:"ignoreParam"`
}

// GetAllServicesResponse ...
type GetAllServicesResponse struct {
	AllServices []Service `json:"services,omitempty"`
}
