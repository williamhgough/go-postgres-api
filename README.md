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

#### Performance test with endpoint caching:
I've run a small load test on the initial API using [vegeta](https://github.com/tsenart/vegeta), results shown below.

`echo 'GET http://localhost:8080/api/products' | \
    vegeta attack -rate 100 -duration 5s | vegeta report -reporter text`

```
Requests      [total, rate]            3000, 300.10
Duration      [total, attack, wait]    9.997314574s, 9.996687795s, 626.779µs
Latencies     [mean, 50, 95, 99, max]  438.319µs, 421.598µs, 546.814µs, 707.812µs, 1.710833ms
Bytes In      [total, mean]            963000, 321.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:3000
Error Set:
```
