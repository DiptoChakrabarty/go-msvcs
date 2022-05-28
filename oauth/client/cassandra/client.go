package cassandra

import (
	"errors"
	"os"

	"github.com/DiptoChakrabarty/go-mvcs/logger"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

var (
	cluster *gocql.ClusterConfig
)

func getenvValue(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Unable to load env file", err)
	}
	return os.Getenv(key)
}

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: getenvValue("username"), Password: getenvValue("password")}
	cluster.Consistency = gocql.Quorum

}

func GetDBSession() (*gocql.Session, error) {
	session, err := cluster.CreateSession()
	if err != nil {
		logger.Error("Unable to create Cluster session", err)
		return nil, errors.New("Unable to create session")
	}
	defer session.Close()

	// creating keyspace
	err = session.Query("CREATE KEYSPACE IF NOT EXISTS oauth WITH REPLICATION= {'class': 'NetworkTopologyStrategy'};").Exec()
	if err != nil {
		logger.Error("Unable to get keyspace", err)
		return nil, errors.New("Unable to get keyspace")
	}
	cluster.Keyspace = "oauth"

	// creating table
	err = session.Query("CREATE TABLE IF NOT EXISTS oauth.users (name text, access_token text, id int), PRIMARY KEY (id));").Exec()
	if err != nil {
		logger.Error("Unable to identify table", err)
		return nil, errors.New("Unable to get table")
	}
	return session, nil
}
