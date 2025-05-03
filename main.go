package main

import(
	"context"
  "fmt"
	"log"
	"github.com/go-redis/redis/v8"
)


var ctx = context.Background()

func main(){
	rdb:= redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:0,
	})

	_, err:= rdb.Ping(ctx).Result()

	if err!=nil{
		log.Fatal("Error connecting:redis",err)
	}

	err:=rdb.Set(ct, "name", "John",0).Err()

	if err!=nil{
		log.Fatal("Set not set :)", err)
	}

	fmt.Println("Redis Set succesfully")

	val, err := rdb.Get(ctx, "name").Result()
	if err!=nil{
		log.Fatal("Set not set :)", err)
	}

	fmt.Printf("The value of name is %s", val)

}
