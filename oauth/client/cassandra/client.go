package cassandra

import (
	"log"

	"github.com/gocql/gocql"
)

func Connect() {
	cluster := gocql.NewCluster()
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "random", Password: "random"}
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

}
