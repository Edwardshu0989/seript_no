package i18n

import (
	"encoding/json"
	"loanmarket-server/conf"
)

var LanJsonSet map[string]map[string]string

func init() {
	err := json.Unmarshal([]byte(LanJson), &LanJsonSet)
	if err != nil {
		panic(err.Error())
	}
}

func T(lan, key string) string {
	v := LanJsonSet[key][lan]
	if v == "" {
		v = LanJsonSet[key][conf.LanEnvEn]
	}
	return v
}
