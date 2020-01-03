package cli

import (
	"fmt"
	"os"
)

type ConnectVars struct {
	Userid   string
	Password string
	Url      string
}

func GetEnvVars() (ConnectVars, error) {
	var connectVars ConnectVars
	var ok bool
	if connectVars.Userid, ok = os.LookupEnv("SERVICENOWUSER"); !ok {
		return connectVars, fmt.Errorf("environment variable not set:SERVICENOWUSER")
	}
	if connectVars.Password, ok = os.LookupEnv("SERVICENOWPWD"); !ok {
		return connectVars, fmt.Errorf("environment variable not set:SERVICENOWPWD")
	}
	if connectVars.Url, ok = os.LookupEnv("SERVICENOWURL"); !ok {
		return connectVars, fmt.Errorf("environment variable not set:SERVICENOWURL")
	}
	return connectVars, nil
}
