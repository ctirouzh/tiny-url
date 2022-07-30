package storage

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocql/gocql"
)

type cassandra struct {
	Session *gocql.Session
}

var (
	cassandraInstance *cassandra
	once              sync.Once
)

func GetCassandraInstance(host, keyspace string) *cassandra {
	if cassandraInstance == nil {
		once.Do(func() {
			fmt.Println("[storage][cassandra]--> Creating single cassandra instance...")
			cluster := gocql.NewCluster(host)
			cluster.Keyspace = keyspace
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
