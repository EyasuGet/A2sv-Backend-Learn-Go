package main

import (
	"github.com/zaahidali/task_manager_api/router"
)

func main() {
	r := router.CreateROuter()
	r.Run(":8080")
}