package main

import (
	"data"
	"fmt"
)

func main() {
	pubList, _ := data.LoadPubs()

	fmt.Println("PubList %v", pubList)
}
