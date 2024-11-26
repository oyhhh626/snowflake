package main

import (
	"fmt"
	snowflake "snowflake_demo/pkg"
)

func main() {
	if err := snowflake.Init("2024-11-01", 1); err != nil {
		fmt.Printf("init failed, err:%#v\n", err)
		return
	}
	for i := 0; i < 10; i++ {
		id := snowflake.GenID()
		fmt.Println(id)
	}

}
