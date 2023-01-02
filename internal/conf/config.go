package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type (
	etcd struct {
		Scheme string   `yaml:"scheme"`
		Hosts  []string `yaml:"hosts"`
	}

	rpcSvc struct {
		UserRPC svc `yaml:"userRPC"`
	}

	svc struct {
		Hosts []string `yaml:"hosts"`
		Name  string   `yaml:"name"`
	}
)

var (
	_c struct {
		Etcd   etcd   `yaml:"etcd"`
		RPCSvc rpcSvc `yaml:"rpcSvc"`
	}
)

func InitYaml() {
	data, err := os.ReadFile("./internal/conf/etcd-test.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &_c); err != nil {
		panic(err)
	}

	for _, h := range _c.Etcd.Hosts {
		fmt.Println(h)
	}

}

func Etcd() *etcd {
	return &_c.Etcd
}

func RPCSvc() *rpcSvc {
	return &_c.RPCSvc
}
