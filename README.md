# Go http server template for microservices

 Setup go using setup-go.sh

### Go package management commands
 1. Create go.mod folder with module path  <br>
 ```go mod init <module path>```
 2. Build and install the program <br>
 ```go install <module path>```  <br>
 ```go install .```
 3. Build internal packages by moving into the directory and running ```go build```
 4. After adding a remote package run ```go mod tidy``` to download the required packages and adding requirements to the go.mod file.

 > Use ```go clean -modcache``` to clean the cache for the downloaded packages in $GOPATH/pkg/mod

 5. Use vscode's generate unit test functionality to generate unit tests and run the tests under a package using ```go test``` 1. Run ```go run .```

### http testing endpoints
 1. Curl write - ```curl -H "Content-Type: application/json; charset=UTF-8" -d '{ "pageName":"hellopie", "pageData":"U29tZXRpbWVzIGkgYW0gaW5zYW5lCg==" }' http://localhost:8080/page/write/hellopie```
 2. Curl read - ```curl http://localhost:8080/page/read/hellopie```
 3. Encode to base64 for pageData ```echo "text text" | openssl base64```
 4. Decode to string ```echo aG93IGFyZSB5b3UgZG9pbmc | base64 --decode```

### Docker commands
 Setup docker using setup-docker.sh
 1. ```docker build -t http-server-template .```
 2. List all the images -> ```docker image ls```
 3. Run in attached mode - ```docker run -it --rm -p 8080:8080 --name running-http-golang-server http-server-template```
 4. Run in detached mode -> ```docker run -d --rm -p 8080:8080 --name running-http-golang-server http-server-template```
 5. Ssh into docker -> ```docker exec -it running-http-golang-server /bin/bash```

### Run and test your service using curl

### Benchmarking
 1. register/http.go has some basic benchmark code written
 2. Benchmark for 5s - `go test -bench . -benchmem -benchtime 5s github.com/dragod812/go-http-server-template/register`
 3. Benchmark 100 iterations - `go test -bench . -benchmem -benchtime 50x github.com/dragod812/go-http-server-template/register`
 4. Benchmark and generate a CPU profile - `go test -bench ReadPageHttpHandler -benchmem -benchtime 5s -cpuprofile pprof.cpu github.com/dragod812/go-http-server-template/register`
 5. Benchmark and generate a heap profile - `go test -bench ReadPageHttpHandler -benchmem -benchtime 5s -memprofile pprof.mem github.com/dragod812/go-http-server-template/register`
****
### TODO
 - Database - Mongodb, postgreSQL, rethinkdb
    - Lets start with rethinkdb then PostgreSQL
 - Cache
    - Redis is the obvious choice
 - Dependency Injection
    - Create two branches one - uberfx, google wire
 - Logging and metrics
   - ELK Logging
   - prometheus metrics
 - Profiling
    - CPU, Memory - pprof
 - Kafka message queue production and consumption
 - AWS or GCP cloud deployment
****

### Go building Services
 1.


### Effective Go Tips
 1.
