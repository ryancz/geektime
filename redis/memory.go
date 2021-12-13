package main

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/redis.v5"
	"math/rand"
	"time"
)

var logger, _ = zap.NewDevelopment()
var log = logger.Sugar()

var letters = []byte("0123456789")

func main() {
	rand.Seed(time.Now().Unix())

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 1,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Errorf("ping failed %v", err)
		return
	}

	valueSizes := []int{10, 20, 50, 100, 200, 500, 1000, 5000}
	for _, valueSize := range valueSizes {
		testMemory(client, valueSize)
		time.Sleep(15 * time.Second)
	}
}

func testMemory(client *redis.Client, valueSize int) {
	info, err := client.Info("memory").Result()
	if err != nil {
		log.Errorf("info memory failed %v", err)
		return
	}
	log.Infof("info memory before test value size %d: %s", valueSize, info)

	for i := 1; i <= 10000; i++ {
		key := fmt.Sprintf("k%05d", i)
		value, err := randNumber(valueSize)
		if err != nil {
			log.Errorf("rand number failed %v", err)
			continue
		}
		_, err = client.Set(key, value, 10*time.Second).Result()
		if err != nil {
			log.Errorf("set key(%s) value(%s) failed %v", key, value, err)
			continue
		}
	}

	info, err = client.Info("memory").Result()
	if err != nil {
		log.Errorf("info memory failed %v", err)
		return
	}
	log.Infof("info memory after test value size %d: %s", valueSize, info)
}

func randNumber(len int) (string, error) {
	bs := make([]byte, len)
	_, err := rand.Read(bs)
	if err != nil {
		return "", errors.Wrap(err, "rand read error")
	}

	ret := make([]byte, len)
	for i, b := range bs {
		ret[i] = letters[b%10]
	}
	return "0x" + string(ret), nil
}
