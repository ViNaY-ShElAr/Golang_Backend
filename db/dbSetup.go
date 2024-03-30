package db

import (
	"GO_PROJECT/db/cassandra"
	"GO_PROJECT/db/redis"
)

func ConnectDatabases() {
	cassandra.ConnectCassandraDb()
	redis.ConnectRedisDb()
}
