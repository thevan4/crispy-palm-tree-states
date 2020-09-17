package run

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	logger "github.com/thevan4/logrus-wrapper"
)

// Default values
const (
	defaultConfigFilePath      = "./diag.yaml"
	defaultURL                 = "http://127.0.0.1:7000/"
	defaultLogin               = ""
	defaultPasword             = ""
	defaultAutoMerge           = false
	defaultIPAndPortSearchMode = "nope"
)

var selectedColumns = []string{}
var defaultColumnsHeaders = []string{"VIP", "SRV STATE", "REAL", "HEALTH", "PROTO", "ROUTING", "TYPE", "HC TYPE", "HC ADDR", "HC TIMERS"}
var tmpColumnsHeadersForReplace = []string{"VIP", "SRV-STATE", "REAL", "HEALTH", "PROTO", "ROUTING", "TYPE", "HC-TYPE", "HC-ADDR", "HC-TIMERS"}

// Config names
const (
	configFilePathName      = "diag-config-file-path"
	urlName                 = "nlb-url"
	nlbLoginName            = "nlb-login"
	nlbPasswordName         = "nlb-password"
	autoMergeName           = "auto-merge"
	ipAndPortSearchModeName = "ip-and-port-search"
	selectedColumnsName     = "selected-columns"
)

var (
	viperConfig *viper.Viper
	logging     *logrus.Logger
)

func init() {
	var err error
	viperConfig = viper.New()
	// work with env
	viperConfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viperConfig.AutomaticEnv()

	// work with flags
	pflag.StringP(configFilePathName, "f", defaultConfigFilePath, "Path to config file. Example: './diag.yaml'")

	pflag.StringP(urlName, "u", defaultURL, "NLB tier 1 url")
	pflag.StringP(nlbLoginName, "l", defaultLogin, "Login")
	pflag.StringP(nlbPasswordName, "p", defaultPasword, "Password")
	pflag.BoolP(autoMergeName, "a", defaultAutoMerge, "Auto merge cells")

	pflag.StringP(ipAndPortSearchModeName, "o", defaultIPAndPortSearchMode, "Mode for IP + port search. Example: 1.1.1.1:1111")
	pflag.StringSliceP(selectedColumnsName, "c", []string{}, "Selected columns. Example: ./diag -l user -p password -c VIP,SRV-STATE,REAL")
	pflag.Parse()
	if err := viperConfig.BindPFlags(pflag.CommandLine); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// FIXME: viper spaces

	// work with config file
	viperConfig.SetConfigFile(viperConfig.GetString(configFilePathName))
	if err := viperConfig.ReadInConfig(); err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	// init logs
	newLogger := &logger.Logger{
		Output:           []string{"stdout"},
		Level:            "trace",
		Formatter:        "text",
		LogEventLocation: false,
	}
	logging, err = logger.NewLogrusLogger(newLogger)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// required values are set
	if viperConfig.GetString(nlbLoginName) == "" {
		logging.Fatal("login be set")
	}
	if viperConfig.GetString(nlbPasswordName) == "" {
		logging.Fatal("password must be set")
	}
	if viperConfig.GetString(ipAndPortSearchModeName) != "nope" {
		ipAndPortSlice := strings.Split(viperConfig.GetString(ipAndPortSearchModeName), ":")
		if len(ipAndPortSlice) != 2 {
			logging.Fatalf("wrong ip and port: %v; expected format 1.1.1.1:1111", viperConfig.GetString(ipAndPortSearchModeName))
		}
	}
	if len(viperConfig.GetStringSlice(selectedColumnsName)) != 0 {
		tmpSelectedColumns := viperConfig.GetStringSlice(selectedColumnsName)
		selectedColumns = make([]string, len(tmpSelectedColumns))
		for i, userColumn := range tmpSelectedColumns {
			if !isFindedInDefaultColumns(userColumn, tmpColumnsHeadersForReplace) {
				logging.Fatalf("column %v not supported. Supported names: %v", userColumn, tmpColumnsHeadersForReplace)
			}
			selectedColumns[i] = replacerForUserColumns(userColumn)
		}
	}
}

func isFindedInDefaultColumns(userColumn string, defaultColumnsHeaders []string) bool {
	for _, internalColumn := range defaultColumnsHeaders {
		if internalColumn == userColumn {
			return true
		}
	}
	return false
}

func replacerForUserColumns(userColumn string) string {
	switch userColumn {
	case "SRV-STATE":
		return "SRV STATE"
	case "HC-TYPE":
		return "HC TYPE"
	case "HC-ADDR":
		return "HC ADDR"
	case "HC-TIMERS":
		return "HC TIMERS"
	default:
		return userColumn
	}
}
