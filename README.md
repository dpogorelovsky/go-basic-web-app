# go-basic-web-app
Skeleton golang starter application, miniframework that has following features:
- running http server
- 2 controllers/handlers for `user` and `city` entities with CRUD operations
- mysql database implementation(which can be easily replaced with any other DB engine) with 2 example tables(users and cities)
- simple data validation
- pagination
- automigrations (can be also called manually)

## How to run
Close the repo somewhere outside `GOPATH`  
create `.env` from `.env_example`(replace your environment values) in root directory  
run:
- `make install`
- `make run`

## DB Migrations
Migrations run each time you run server, so you don't have care about that. However,
if you need to run it manually:
- `make migrate:up` - will trigger migrations
- `make migrate:down` - will roll back all migrations (don't run it in production, the data will be lost)

## Existing routes
_Getting data_
- `GET http://localhost:9000/users` - returns list of users
- `GET http://localhost:9000/users/{id}` - user by id
- `GET http://localhost:9000/cities` - returns list of cities
- `GET http://localhost:9000/cities/{id}` - city by id  
All above endpoints support pagination parameters: `GET http://localhost:9000/users?limit=2&page=2`

_Creating_
- `POST http://localhost:9000/cities` - adds city, body should be like:
```
{
  "name": "Dublin"
}
```
- `POST http://localhost:9000/users` - adds user, body should be like:
```
{
  "email": "harry.from@dublin",
  "name": "Harry",
  "age": 45,
  "cityId": 1
}
```
_Updating_
- `PUT http://localhost:9000/cities/{id}` - updates whole city model:
```
{
  "id": 1,
  "name": "Dublin"
}
```
- `PUT http://localhost:9000/users/{id}` - updates whole users model:
```
{
  "id": 1,
  "email": "harry.from@dublin",
  "name": "Harry",
  "age": 45,
  "cityId": 1
}
```
_Removing_
- `DELETE http://localhost:9000/cities/{id}` - removes city entry:
- `DELETE http://localhost:9000/users/{id}` - removes user entry:

## @todo
- 
## @maybe todo
- authorization/authentication
- tests
- roles/access
- cache layer
- more...
