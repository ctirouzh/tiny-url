package storage

import (
	"fmt"
	"log"
	"sync"

	"github.com/ctirouzh/tiny-url/config"
	"github.com/gocql/gocql"
)

type cassandra struct {
	Session *gocql.Session
}

var (
	cassandraInstance *cassandra
	cassandraOnce     sync.Once
)

func GetCassandraInstance(cfg config.Cassandra) *cassandra {
	if cassandraInstance == nil {
		cassandraOnce.Do(func() {
			fmt.Println("[storage][cassandra]--> Creating single cassandra instance...")
			cluster := gocql.NewCluster(cfg.Host)
			cluster.Keyspace = cfg.KeySpace
			session, err := cluster.CreateSession()
			if err != nil {
				log.Fatal("[storage][cassndra]-->", err)
			}
			cassandraInstance = &cassandra{Session: session}
		})
	} else {
		fmt.Println("[storage][cassandra]--> cassandra instance already created.")
	}
	return cassandraInstance
}
