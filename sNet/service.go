package sNet

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/yasseldg/simplego/sConv"
	"github.com/yasseldg/simplego/sEnv"
	"github.com/yasseldg/simplego/sJson"
	"github.com/yasseldg/simplego/sLog"
)

type Header struct {
	Key   string
	Value string
}

type Headers []Header

type Config struct {
	Env        string
	Url        string
	Secure     bool
	Port       int
	Network    string
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
	conf.Env = env
	conf.Update()
	return &conf
}

func (c Config) Log() {
	sLog.Info("%s: %s ", c.Env, c.GetUrl())
}

func (c *Config) Update() {
	c.Url = sEnv.Get(fmt.Sprintf("%s_Url", c.Env), c.Url)
	c.Port = sConv.GetInt(sEnv.Get(fmt.Sprintf("%s_Port", c.Env), sConv.GetStrI(c.Port)))
	c.Secure = sConv.GetBool(sEnv.Get(fmt.Sprintf("%s_Secure", c.Env), sConv.GetStrB(c.Secure)))
	c.Network = sEnv.Get(fmt.Sprintf("%s_Network", c.Env), c.Network)
	c.PathPrefix = sEnv.Get(fmt.Sprintf("%s_Path_Prefix", c.Env), c.PathPrefix)
}

func (c *Config) AddPathPrefixs(path_prefixs ...string) {
	c.PathPrefix = path.Join(c.PathPrefix, path.Join(path_prefixs...))
}

func (c *Config) AddHeaders(headers Headers) {
	c.Headers = append(c.Headers, headers...)
}

func (c Config) GetUri() string {
	uri := c.Url
	if c.Port > 0 {
		uri = fmt.Sprintf("%s:%d", uri, c.Port)
	}
	if len(c.PathPrefix) > 0 {
		uri = path.Join(uri, c.PathPrefix)
	}
	return uri
}

func (c Config) GetUrl() string {
	url := c.GetUri()
	if c.Secure {
		return fmt.Sprintf("https://%s", url)
	}
	return fmt.Sprintf("http://%s", url)
}

func (c Config) GetHandlerPath(handler string) string {
	if len(c.PathPrefix) > 0 {
		return fmt.Sprintf("/%s/%s", c.PathPrefix, handler)
	}
	return fmt.Sprintf("/%s", handler)
}

func (c Config) Call(method string, params string, body io.Reader) ([]byte, error) {

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

	sLog.Debug("Call: request: %v ", request)

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Call: client.Do(request): %s ", err)
	}
	sLog.Debug("Call: resp: %v ", resp)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Call: StatusCode: %d ", resp.StatusCode)
	}

	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Call: ioutil.ReadAll(resp.Body): %s ", err)
	}
	sLog.Debug("Call: resp_body: %s ", string(resp_body))

	return resp_body, nil
}

// default action is "obj"
func (c Config) SendObj(method, action string, body_obj any) ([]byte, error) {
	if len(action) == 0 {
		action = "obj"
	}
	body_str, err := sJson.ToJson(body_obj)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(body): %s", err)
	}
	return c.Call(method, action, strings.NewReader(body_str))
}

// ---- Delete ----

func (c Config) Print(name string) {
	c.Log()
}
