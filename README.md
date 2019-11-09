# go-basic-web-app
This is going to be web application that can be used as skeleton.

## How to run
Close the repo somewhere outside `GOPATH` and run:
- `make install`
- `make run`

## DB Migrations
Migrations run each time you run server, so you don't have care about that. However,
if you need to run it manually:
- `make migrate:up` - will trigger migrations
- `make migrate:down` - will roll back all migrations (don't run it in production, the data will be lost)

## @todo
- DB
 - several example tables
 - migrations
- cache layer
## @maybe todo
- authorization/authentication
- tests
- roles/access
- more...
