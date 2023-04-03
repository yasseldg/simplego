package sNet

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/yasseldg/simplego/sLog"
)

type Header struct {
	Key   string
	Value string
}

type Headers []Header

type Conf struct {
	Container  string
	IsExternal bool
	Port       int
	PathPrefix string
	Headers    Headers
}

func GetServiceUrl(service *Conf) string {
	if service.IsExternal {
		return fmt.Sprintf("https://%s%s", service.Container, service.PathPrefix)
	}
	return fmt.Sprintf("http://%s:%d%s", service.Container, service.Port, service.PathPrefix)
}

func Call(service *Conf, method string, params string, body io.Reader) ([]byte, error) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	url := GetServiceUrl(service)
	if len(params) > 0 {
		url = fmt.Sprintf("%s/%s", url, params)
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Call: http.NewRequest(): %s ", err)
	}

	for _, header := range service.Headers {
		request.Header.Add(header.Key, header.Value)
	}

	sLog.Info("Call: request: %v ", request)

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Call: client.Do(request): %s ", err)
	}
	sLog.Info("Call: resp: %v ", resp)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Call: StatusCode: %d ", resp.StatusCode)
	}

	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Call: ioutil.ReadAll(resp.Body): %s ", err)
	}
	sLog.Debug("Call: body: %s ", string(resp_body))

	return resp_body, nil
}
