package common

import (
	"github.com/asim/go-micro/config"
	"github.com/asim/go-micro/plugins/config/source/consul"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port,10)),
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
		)
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}
	err = config.Load(consulSource)
	return config, err
}