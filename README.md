# go-postgres-api
Simple Golang API using docker to provide a Postgres DB & redis for endpoint caching.

#### Usage:

To run in development mode with hot reload:
`make run-dev`

To run a production build:
`make run-prod`

#### Requirements
- Docker
- Docker Compose

#### Working Routes:
```
GET http://localhost:8080/api/products
GET http://localhost:8080/api/products/:id
```

#### TODO:
- [ ] Add more logging
- [x] Implement redis endpoint caching
- [ ] Use `sync.Pool` for db conns
- [ ] Finish Product CRUD utility
- [ ] Add unit tests
- [ ] Add benchmark tests & profile performance

#### Performance testing:
I've run a small load test on the initial API using [vegeta](https://github.com/tsenart/vegeta), results shown below.

`echo 'GET http://localhost:8080/api/products' | \
    vegeta attack -rate 100 -duration 5s | vegeta report -reporter text`

```
Requests      [total, rate]            500, 100.20
Duration      [total, attack, wait]    4.9906915s, 4.99003725s, 654.25µs
Latencies     [mean, 50, 95, 99, max]  838.435µs, 885.092µs, 1.06ms, 1.649987ms, 3.59574ms
Bytes In      [total, mean]            158500, 317.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:500
Error Set:
```
