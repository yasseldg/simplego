package sMongo

import (
	"fmt"

	"github.com/yasseldg/simplego/sEnv"
	"github.com/yasseldg/simplego/sLog"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectionParams, Databases access data
type ConnectionParams struct {
	Environment      string `yaml: "environment"`
	Username         string `yaml: "username"`
	Password         string `yaml: "password"`
	Protocol         string `yaml: "protocol"`
	Host             string `yaml: "host"`
	Port             string `yaml: "port"`
	AuthDatabase     string `yaml: "auth_atabase"`
	DirectConnection bool   `yaml: "direct_connection"`
	Tls              bool   `yaml: "tls"`
	ReadPreference   string `yaml: "read_preference"`
	AuthMechanism    string `yaml: "auth_mechanism"`
}

func defaultConnectionParams() *ConnectionParams {
	return &ConnectionParams{
		Environment:  "read",
		Username:     "devRead",
		Password:     "devReadPass",
		Host:         "mongodb",
		Port:         "27017",
		AuthDatabase: "admin"}
}

// GetConnection, Databases access data predefined
func getConnection(name string) *ConnectionParams {
	var conn ConnectionParams
	err := sEnv.LoadYaml(fmt.Sprint(".env/", name, ".mongodb"), &conn)
	if err != nil {
		sLog.Error("GetConnectionUri: conn is nil, using default ConnectionParams READ")

		return defaultConnectionParams()
	}
	return &conn
}

// GetConnectionUri, return (Uri, Credentials)
func getConnectionUri(conn *ConnectionParams) (string, options.Credential) {

	optCredential := options.Credential{AuthSource: conn.AuthDatabase, Username: conn.Username, Password: conn.Password}

	if conn.DirectConnection && conn.Tls && len(conn.ReadPreference) > 0 {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/?directConnection=%t&tls=%t&readPreference=%s",
			conn.Username, conn.Password, conn.Host, conn.Port, conn.DirectConnection, conn.Tls, conn.ReadPreference), optCredential
	}

	if len(conn.ReadPreference) > 0 && len(conn.AuthMechanism) > 0 && conn.DirectConnection {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/?readPreference=%s&authMechanism=%s&directConnection=%t",
			conn.Username, conn.Password, conn.Host, conn.Port, conn.ReadPreference, conn.AuthMechanism, conn.DirectConnection), optCredential
	}

	if conn.Environment == "prod" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s", conn.Username, conn.Password, conn.Host, conn.Port), optCredential
	}

	return fmt.Sprintf("mongodb://%s:%s", conn.Host, conn.Port), optCredential
}
