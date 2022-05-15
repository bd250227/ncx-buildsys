package main

import (
	"fmt"
	"time"

	builduuid "buildsys/internal/build-uuid"
)

func main() {
	for {
		fmt.Println("hello world: " + builduuid.GetUUID())
		time.Sleep(1 * time.Second)
	}
}
