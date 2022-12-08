package gowinkey

import (
	"fmt"
	"time"
)

func init() {
	events := Listen()
	// time1 := 10
	for e := range events {

		if e.State != 0 {
			fmt.Println("截图1")
			fmt.Println("上传1", e)
			time.Sleep(10 * time.Second)
		} else {
			fmt.Println("截图2", e)
			fmt.Println("上传2", e)
			time.Sleep(10 * time.Second)

		}

	}
}
