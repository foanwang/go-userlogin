package until

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)
var setting = GetYamlConfig("./conf/config.yaml")

func GetYamlConfig(path string) map[interface{}]interface{}{
	fmt.Printf("init setting...\n")
	data, err := ioutil.ReadFile(path)
	m := make(map[interface{}]interface{})
	if err != nil {
		LogErr(err)
	}
	err = yaml.Unmarshal([]byte(data), &m)
	return m
}

func GetElement(key string)string {
	if value,ok:=setting[key];ok {

		return fmt.Sprint(value)
	}

	Log("can't find the config file")
	return ""
}
