package main

import (
	redis2 "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"	
	"fmt"
)

var(
	RedisClient *redis2.Client
	RedisPool *redis.Pool
)

func RedisInit(addr string){

	RedisClient = redis2.NewClient(&redis2.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	fmt.Println(pong, err)

	/*
	value, error := RedisClient.SAdd("testset", 1232232).Result()//RedisClient.Set("testa", 'a', 0).Result()
	if error != nil{
	
	}
	fmt.Println(value)*/

	RedisPool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 6000,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}


	conn := GetConn()
	defer (*conn).Close()

	r, err := (*conn).Do("ping")
	if err != nil {

	}	
	fmt.Println("r :", r)
}

func GetConn() *redis.Conn{
	conn := RedisPool.Get()
	return &conn
}