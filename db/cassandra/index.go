package cassandra

import (
	"os"

	"github.com/gocql/gocql"

	"GO_PROJECT/logger"
	"GO_PROJECT/model"
)

var CassandraSession *gocql.Session

func ConnectCassandraDb() {

	cfg := model.CassandraCfg{
		Host:     os.Getenv("CASSANDRA_HOST"),
		Port:     os.Getenv("CASSANDRA_PORT"),
		Keyspace: os.Getenv("CASSANDRA_KEYSPACE"),
	}

	cluster := gocql.NewCluster(cfg.Host + ":" + cfg.Port)
	cluster.Keyspace = cfg.Keyspace

	var err error
	CassandraSession, err = cluster.CreateSession()
	if err != nil {
		logger.Log.Fatal("Error in Connecting Cassandra Db: ", err)
	}

	logger.Log.Info("Cassandra Db Connected Successfully")
}
