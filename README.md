# go-postgres-api
Simple Golang API using docker to provide a Postgres DB & redis for endpoint caching.

To run in development mode with hot reload:
`make run-dev`

To run a production build:
`make run-prod`

#### Requirements
- Docker
- Docker Compose

#### TODO:
- [ ] Implement redis endpoint caching
- [ ] Use `sync.Pool` for db conns
- [ ] Finish Product CRUD utility
- [ ] Add unit tests
- [ ] Add benchmark tests & profile performance
