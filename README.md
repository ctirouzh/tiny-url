# Tiny-URL
This is a simple URL shortener service using golang, gin framework, cassandra database, and redis cache. You can find more details in https://tinyurl.com/bj4bne99.

# Features

:heavy_check_mark: Authentication using jwt. <br/>
:heavy_check_mark: Authenticated users can generate a shorter and unique alias of a given URL. <br/>
:heavy_check_mark: Authenticated users can see the details of their generated URL. <br/>   
:heavy_check_mark: Given a short link, redirect users to the original link. <br/>
:heavy_check_mark: Accessible through REST APIs by other services. <br/>
:heavy_check_mark: Default url ttl in cassandra configuration. <br/>
:heavy_check_mark: Users can delete their urls. <br/>
:heavy_check_mark: List links of a user. <br/>
:x: Link expiration after its standard default timespan. <br/>
:x: Custom expiration time. <br/> 
:x: Custom short link.  <br/>
:x: Use an API developer key to throttle users based on their allocated quota in createUrl() api.<br/>
:x: Limit users via their api_dev_key to a certain number of URL creations and redirections per some time period (which may be set to a different duration per developer key). <br/>
:question: Key Generation Service (KGS) instead of github.com/teris-io/shortid package <br/>
:x: Data partitioning and replication (Hash Based Partitioning). <br/>
:x: Cache eviction policy (LRU, Linked Hash Map). <br/>
:x: Load balancer (LB): See https://tinyurl.com/bdfnc9pk. <br/>
:x: Purging or DB cleanup: See https://tinyurl.com/56tje6tt. <br/>
:x: Telemetry: See https://tinyurl.com/4zrpbupd. <br/>
:x: public/private permission level for each URL in database. <br/>
:x: Dockerfile and docker compose. <br/>

# Installation

Install Docker & Go (>1.18)

```bash
make cassandra \
make redis \
make db-start \
make cache-start \
make db-migrate \
make server
```
You can stop database container using the following commands

```bash
make db-stop
```
To stop Cache

```bash
make cache-stop
```
To drop database you can use this command

```bash
make db-drop
```

# Refrences

1. https://medium.com/easyread/golang-clean-archithecture-efd6d7c43047
2. https://intersog.com/blog/how-to-write-a-custom-url-shortener-using-golang-and-redis/
3. https://github.com/tushar9989/url-short
4. https://www.educative.io/courses/grokking-the-system-design-interview/m2ygV4E81AR
5. https://github.com/quyenphamkhac/go-tinyurl
6. https://www.geeksforgeeks.org/system-design-url-shortening-service/
7. https://github.com/teris-io/shortid
