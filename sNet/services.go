package sNet

import (
	"github.com/yasseldg/simplego/sEnv"
	"github.com/yasseldg/simplego/sLog"
)

type Services []*Config

func GetServices(env string, defaults ...string) (objs Services) {
	services := sEnv.GetSlice(env, defaults...)
	for _, service := range services {
		obj := GetConfig(service)
		if obj == nil {
			continue
		}
		obj.Log()
		objs = append(objs, obj)
	}
	return
}

func (s Services) Log() {
	for _, service := range s {
		service.Log()
	}
}

func (s Services) SendObj(body_obj any) {
	for _, service := range s {
		_, err := service.SendObj("POST", body_obj)
		if err != nil {
			sLog.Error("(%s).SendObj: %v  ..  err: %s", service.Url, body_obj, err)
		}
	}
}

func (s Services) AddPathPrefixs(path_prefixs ...string) {
	for _, service := range s {
		service.AddPathPrefixs(path_prefixs...)
	}
}

func (s Services) AddHeaders(headers Headers) {
	for _, service := range s {
		service.AddHeaders(headers)
	}
}
