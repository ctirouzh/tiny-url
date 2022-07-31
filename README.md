# Tiny-URL
This is a simple URL shortener service using golang, gin framework, cassandra database, and redis cache. You can find more details in https://tinyurl.com/bj4bne99.

# Features

- [ ] Authentication using jwt.
- [ ] Authenticated users can generate a shorter and unique alias of a given URL.
- [ ] Authenticated users can see the details of thier generated URL.   
- [ ] Given a short link, redirect users to the original link. 
- [ ] Accessibe through REST APIs by other services.
- [x] Add deleteUrl(api_dev_key, url_key) api.
- [x] Link expiration after a stanadard default timespan.
- [x] Custom expiration time.  
- [x] Custom short link.  
- [x] Use an API developer key to throttle users based on their allocated quota in createUrl() api.
- [x] Limit users via their api_dev_key to a certain number of URL creations and redirections per some time period (which may be set to a different duration per developer key).
- [x] Key Generation Service (KGS) instead of github.com/teris-io/shortid package?
- [x] Data partitioning and replication (Hash Based Partitioning)
- [x] Cache eviction policy (LRU, Linked Hash Map)
- [x] Load balancer (LB): See https://tinyurl.com/bdfnc9pk
- [x] Purging or DB cleanup: See https://tinyurl.com/56tje6tt
- [x] Telemetry: See https://tinyurl.com/4zrpbupd 
- [x] public/private permission level for each URL in database.
- [x] Dockerize the app. 

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

- [ ] https://medium.com/easyread/golang-clean-archithecture-efd6d7c43047
- [ ] https://intersog.com/blog/how-to-write-a-custom-url-shortener-using-golang-and-redis/
- [ ] https://github.com/tushar9989/url-short
- [ ] https://www.educative.io/courses/grokking-the-system-design-interview/m2ygV4E81AR
- [ ] https://github.com/quyenphamkhac/go-tinyurl
- [ ] https://www.geeksforgeeks.org/system-design-url-shortening-service/
- [ ] https://github.com/teris-io/shortid