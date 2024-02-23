package main

import (
	"fmt"
	"log"
	"postfix-log/logparser"
)

func main() {
	lp := logparser.NewLogParser()

	queueItems, err := lp.ReadLog("/Users/sebastiangabbert/Downloads/postfix.log")
	if err != nil {
		log.Fatal(err)
	}

	bad := queueItems.FilterByStatusClass(5)

	fmt.Println(bad)
}
