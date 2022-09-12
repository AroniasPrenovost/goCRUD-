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
