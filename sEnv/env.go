package sEnv

import (
	"os"

	"github.com/yasseldg/simplego/sConv"
	"github.com/yasseldg/simplego/sStr"

	"gopkg.in/yaml.v3"
)

// Get
func Get(env_name, defaults string) string {

	env, ok := os.LookupEnv(env_name)
	if ok {
		return env
	}
	return defaults
}

// GetSlice
func GetSlice(env_name string, defaults ...string) []string {

	string_values := Get(env_name, "")

	if len(string_values) > 0 {
		return sStr.SplitString(string_values, ",")
	}
	return defaults
}

// GetSliceInt
func GetSliceInt(env_name string, defaults []int) (res []int) {

	values := GetSlice(env_name)
	if len(values) > 0 {
		for _, v := range values {
			res = append(res, sConv.GetInt(v))
		}
		return
	}
	return defaults
}

// LoadYaml
func LoadYaml(file_name string, obj interface{}) error {

	data, err := os.ReadFile(file_name)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}
