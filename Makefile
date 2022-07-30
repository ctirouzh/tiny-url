server:
	go run .

cassandra: 
	sudo docker run --name tinyurl-cassandra-db -p 7000:7000 -p 7001:7001 \
	-p 7199:7199 -p 9160:9160 -p 9042:9042 -d cassandra:latest

redis: 
	sudo docker run --name  tinyurl-redis-cache -p 6379:6379 -d redis:latest
	
db-start: 
	sudo docker start tinyurl-cassandra-db

db-stop:
	sudo docker stop tinyurl-cassandra-db

db-migrate: 
	sudo docker cp ./storage/migration/create_db.cql tinyurl-cassandra-db:/
	sudo docker exec -d tinyurl-cassandra-db cqlsh localhost -f /create_db.cql

db-drop: 
	sudo docker cp ./storage/migration/drop_db.cql tinyurl-cassandra-db:/
	sudo docker exec -d tinyurl-cassandra-db cqlsh localhost -f /drop_db.cql

cache-start: 
	sudo start tinyurl-redis-cache

cache-stop:
	sudo stop tinyurl-redis-cache

.PHONY: server cassandra redis