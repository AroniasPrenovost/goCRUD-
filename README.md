# REST API example written in Go 

## Requirements

To be able to show the desired features of curl this REST API must match a few
requirements:

* [x] `GET /employees` returns list of employees as JSON
* [x] `GET /employees/{id}` returns details of specific employee as JSON
* [x] `POST /employees` accepts a new employee to be added
* [x] `POST /employees` returns status 415 if content is not `application/json`
* [x] `GET /admin` requires basic auth
* [x] `GET /employees/random` redirects (Status 302) to a random employee

### Data Types

A employee object should look like this:
```json
{
  "id": "someid",
  "email_address": "name of the employee",
  "user_name": "name of the employee",
  "first_name": "the amusement park the ride is in",
  "last_mame": "name of the manufacturer",
  "city": 27,
}
```

### Persistence

Data is stored is MySQL 

### Running Project 

```
$ go run main.go
```

## options: 

Show all items in 'db' 
```
$ curl localhost:8080/employees | jq
```

## deleting a user 
```
curl -X "DELETE" localhost:8080/delete?id=10
```

adding new employees 
```
$ curl localhost:8080/employees -X POST -d '{"name": "As", "inPark": "this is another example of  aplace", "height": 30, "manufacturer": "ANOTHERNEWNAMETHATLOOKSGOOD"}' -H "Content-Type: application/json"
```
```
$ curl localhost:8080/employees -X POST -d '{"user_name": "As", "email_address": "aronp123@hotmail.com", "first_name": "Aron", "last_name": "Prenvoost"}' -H "Content-Type: application/json"
```

### Long-Term Goals 
## Build a real world "production" REST API: 

* [ ] Scalable, must be able to run more than one instance.

* [ ] Dockerized, runnable on minikube.

* [ ] Unit tested, must be able to run "go test ./..." directly from clone.

* [ ] Integration tested, recommend docker-compose.

* [ ] OpenAPI/Swagger (or similar for gRPC) documented.

* [ ] dep vendored, but using the standard library often, instead of piling on dependencies.

* [ ] Authenticated and Authorized via apikeys and/or user accounts.

* [ ] Built and tested via CI: Travis, CircleCi, etc. Recommend Makefile for task documentation.

* [ ] Flag & ENV config, API keys, ports, dev mode, etc.

* [ ] "why" comments, not "what" or "how" which should be clear through func/variable names and godoc comments.

* [ ] Use of Context to limit request time.

* [ ] Leveled JSON logging, logrus or similar.

* [x] Postgres/MySQL, sqlx or an ORM.

* [ ] Redis/memcache for scalable caching.

* [ ] Datadog, New Relic, AppDynamics, etc for monitoring and statistics.

* [ ] Well documented README.md with separate sections for API user and service developer audiences. Maybe even include graphviz or mermaidJS UML diagrams.

* [ ] Clean git history with structured commits and useful messages. No merge master commits.

* [ ] Passing go fmt, go lint, or better, go-metalinter in the CI.
