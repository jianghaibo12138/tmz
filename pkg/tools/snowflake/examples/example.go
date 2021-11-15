package main

import (
	"fmt"
	"stp_go/universal/tools/snowflake"
)

func main() {
	sf := snowflake.InitSnowFlake(0)
	sfId := sf.NextVal()
	fmt.Println(sfId)
}
