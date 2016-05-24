// Listener simply echos any requests it gets for now.
package main

import (
	"fmt"
	"time"
)

func main() {

	const (
		name = "onest"
		addr = "192.168.1.13:3265"
		key  = "ExKspLt9qo5HC59QerE_Squ2iCxSo_TjxonXhxGAQ8Q="
	)

	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Println(now.Format(time.RFC3339))
	fmt.Println(now.Format(time.RFC3339Nano))
}
