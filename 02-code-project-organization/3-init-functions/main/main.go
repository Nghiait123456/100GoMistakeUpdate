package main

import (
	"fmt"
	"github.com/teivah/100-go-mistakes/02-code-project-organization/3-init-functions/db"
	"github.com/teivah/100-go-mistakes/02-code-project-organization/3-init-functions/local_in_memory"

	"github.com/teivah/100-go-mistakes/02-code-project-organization/3-init-functions/redis"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	err := redis.Store("foo", "bar")
	db.NewDB()
	local_in_memory.NewLocalInMemory()
	_ = err
}
