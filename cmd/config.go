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
	defaultConfigFilePath = "./states.yaml"
	defaultURL            = "http://127.0.0.1:7000/"
	defaultLogin          = ""
	defaultPasword        = ""
	defaultAutoMerge      = false
)

// Config names
const (
	configFilePathName         = "states-config-file-path"
	urlName                    = "nlb-url"
	nlbLoginName               = "nlb-login"
	nlbPasswordName            = "nlb-password"
	autoMergeName              = "auto-merge"
	ipAndPortSearchModeName    = "ip-and-port-search"
	defaultIPAndPortSearchMode = "nope"
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
	pflag.StringP(configFilePathName, "c", defaultConfigFilePath, "Path to config file. Example value: './states.yaml'")

	pflag.StringP(urlName, "u", defaultURL, "nlb tier 1 url")
	pflag.StringP(nlbLoginName, "l", defaultLogin, "Login")
	pflag.StringP(nlbPasswordName, "p", defaultPasword, "Password")
	pflag.BoolP(autoMergeName, "a", defaultAutoMerge, "Auto merge cells")

	pflag.StringP(ipAndPortSearchModeName, "o", defaultIPAndPortSearchMode, "Mode for IP + port search")

	pflag.Parse()
	if err := viperConfig.BindPFlags(pflag.CommandLine); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
}
