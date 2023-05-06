package main

import (
	"fmt"
	"github.com/teivah/100-go-mistakes/02-code-project-organization/3-init-functions/db"
	"github.com/teivah/100-go-mistakes/02-code-project-organization/3-init-functions/local_in_memory"
	"github.com/teivah/100-go-mistakes/02-code-project-organization/3-init-functions/redis"
)

var a = func() int {
	fmt.Println("variables in main init 1")
	return 0
}()

func init() {
	fmt.Println("function init 2")
}

func init() {
	fmt.Println("function init 3")
}

func main() {
	err := redis.Store("foo", "bar")
	db.NewDB()
	local_in_memory.NewLocalInMemory()
	_ = err
}
