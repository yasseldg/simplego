package sNet

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/yasseldg/simplego/sConv"
	"github.com/yasseldg/simplego/sEnv"
	"github.com/yasseldg/simplego/sLog"
)

type Header struct {
	Key   string
	Value string
}

type Headers []Header

type Config struct {
	Url        string
	Secure     bool
	Port       int
	PathPrefix string
	Headers    Headers
}

// GetConfig
func GetConfig(env string) *Config {
	name := sEnv.Get(env, "DEV")
	var conf Config
	err := sEnv.LoadYaml(fmt.Sprint(".env/services/", name, ".yaml"), &conf)
	if err != nil {
		sLog.Fatal("sNet: getConf: can't load env file %s: %s", name, err)
	}
	conf.Update(env)
	return &conf
}

func (c *Config) Update(env string) {
	c.Url = sEnv.Get(fmt.Sprintf("%s_Url", env), c.Url)
	c.Port = sConv.GetInt(sEnv.Get(fmt.Sprintf("%s_Port", env), sConv.GetStrI(c.Port)))
	c.Secure = sConv.GetBool(sEnv.Get(fmt.Sprintf("%s_Secure", env), sConv.GetStrB(c.Secure)))
	c.PathPrefix = sEnv.Get(fmt.Sprintf("%s_Path_Prefix", env), c.PathPrefix)
}

func (c *Config) GetUrl() string {
	url := c.Url
	if c.Port > 0 {
		url = fmt.Sprintf("%s:%d", url, c.Port)
	}
	if len(c.PathPrefix) > 0 {
		url = fmt.Sprintf("%s/%s", url, c.PathPrefix)
	}
	if c.Secure {
		return fmt.Sprintf("https://%s", url)
	}
	return fmt.Sprintf("http://%s", url)
}

func (c *Config) Call(method string, params string, body io.Reader) ([]byte, error) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	url := c.GetUrl()
	if len(params) > 0 {
		url = fmt.Sprintf("%s/%s", url, params)
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Call: http.NewRequest(): %s ", err)
	}

	for _, header := range c.Headers {
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
