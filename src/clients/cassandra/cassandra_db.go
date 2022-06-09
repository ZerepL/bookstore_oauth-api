package cassandra

import (
	"github.com/gocql/gocql"
	"os"
)

var (
	cluster *gocql.ClusterConfig
	session *gocql.Session
)

const (
	cassandraOAuthHost     = "cassandra_oauth_host"
	cassandraOAuthKeyspace = "cassandra_oauth_keyspace"
)

var (
	host     = os.Getenv(cassandraOAuthHost)
	keyspace = os.Getenv(cassandraOAuthKeyspace)
)

func init() {
	// Connect to Cassandra cluster
	cluster = gocql.NewCluster(host)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
