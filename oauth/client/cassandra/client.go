package cassandra

import (
	"os"

	"github.com/DiptoChakrabarty/go-mvcs/logger"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

var (
	session *gocql.Session
)

func getenvValue(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("unable to load env file", err)
	}
	return os.Getenv(key)
}

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: getenvValue("DBUSERNAME"), Password: getenvValue("PASSWORD")}
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		logger.Error("unable to create Cluster session", err)
		panic(err)
	}
	defer session.Close()

	// creating keyspace
	err = session.Query("CREATE KEYSPACE IF NOT EXISTS oauth WITH REPLICATION= {'class': 'NetworkTopologyStrategy'};").Exec()
	if err != nil {
		logger.Error("unable to create keyspace", err)
		panic(err)
	}
	cluster.Keyspace = "oauth"

	// creating table
	err = session.Query("CREATE TABLE IF NOT EXISTS oauth.users (name text, access_token text, id int), PRIMARY KEY (id));").Exec()
	if err != nil {
		logger.Error("unable to identify table", err)
		panic(err)
	}
}

func GetDBSession() *gocql.Session {
	return session
}
