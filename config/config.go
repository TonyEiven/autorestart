package config

import (
	_ "gopkg.in/yaml.v2"
)

type Authenticator struct {
	Aliyun struct {
		CustomerKey *AK `yaml:"customer-accesskey"`
		ServiceKey  *AK `yaml:"service-accesskey"`
	} `yaml:"aliyun"`

	Service struct {
		TestEnv struct {
			Region        string   `yaml:"region"`
			Servers       []string `yaml:"servers"`
			SourceGateway string   `yaml:"sourceGateway"`
			TargetGateway string   `yaml:"targetGateway"`
			Shadows       []string `yaml:"shadows"`
		} `yaml:"test-environment"`
	} `yaml:"service"`
}

type AK struct {
	AccessKey       string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
}

func (auth *Authenticator) GetAccessKey(env string) string {
	var retvalue string
	switch env {
	case "Customer":
		return auth.Aliyun.CustomerKey.AccessKey
	case "Service":
		return auth.Aliyun.ServiceKey.AccessKey
	default:
		retvalue = "None"
	}
	return retvalue
}

func (auth *Authenticator) GetAccessSecret(env string) string {
	var retvalue string
	switch env {
	case "Customer":
		return auth.Aliyun.CustomerKey.AccessKeySecret
	case "Service":
		return auth.Aliyun.ServiceKey.AccessKeySecret
	default:
		retvalue = "None"
	}
	return retvalue
}

// According to envName return env value
func (auth *Authenticator) GetEnv(spec string) interface{} {
	var retvalue string
	switch spec {
	case "Region":
		return auth.Service.TestEnv.Region
	case "Servers":
		return auth.Service.TestEnv.Servers
	case "SourceGateway":
		return auth.Service.TestEnv.SourceGateway
	case "TargetGateway":
		return auth.Service.TestEnv.TargetGateway
	case "Shadows":
		return auth.Service.TestEnv.Shadows
	default:
		retvalue = "None"
	}
	return retvalue
}
