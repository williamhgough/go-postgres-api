package cache

import (
	"net/http"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

var codec *cache.Codec

func init() {
	r := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	codec = &cache.Codec{
		Redis: r,
		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
}

// Fetch takes a request and uses the URL as a key
// to retrieve data from cache
func Fetch(r *http.Request) (interface{}, error) {
	var data interface{}
	err := codec.Get(r.URL.Path, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Store uses the request URL as a key to set given data
// inside redis cache.
func Store(r *http.Request, data interface{}) error {
	err := codec.Set(&cache.Item{
		Key:        r.URL.Path,
		Object:     data,
		Expiration: 0,
	})
	if err != nil {
		return err
	}
	return nil
}
