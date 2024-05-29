package main 

import (
	"fmt"
	"excan/config"
	"excan/config/opts"
)



func main() {
	c := config.NewConfig(
		opts.SetServerHost("myhost"),
	)

	fmt.Printf("%v", c)
}