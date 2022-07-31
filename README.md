# Tiny-URL
This is a url shortener service using golang, gin framework, cassandra database, and redis cache. You can find more details in https://tinyurl.com/bj4bne99.

Install Docker & Go (>1.13)

```bash
make cassandra
make redis
make db-start
make cache-start
make db-migrate
make server
```
You can stop database or cache containers using the following commands

```bash
make db-stop
make cache-stop
```
To drop database you can use this command

```bash
make db-drop
```
