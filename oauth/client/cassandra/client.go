package cassandra

import (
	"log"
	"os"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

func getenvValue(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env file")
	}
	return os.Getenv(key)
}

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: getenvValue("username"), Password: getenvValue("password")}
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()

	// creating keyspace
	err = session.Query("CREATE KEYSPACE IF NOT EXISTS oauth WITH REPLICATION= {'class': 'NetworkTopologyStrategy'};").Exec()
	if err != nil {
		log.Println(err)
		return
	}

	// creating table
	err = session.Query("CREATE TABLE IF NOT EXISTS oauth.users (name text, access_token text, id int), PRIMARY KEY (id));").Exec()
	if err != nil {
		log.Println(err)
		return
	}

}
