package main

import (
	"fmt"
	"u-akihiro/openbd/client"
)

func main() {
	resp, _ := client.FetchBookInfo([]string{"9784048686051", "9784048686051", "9784048686051"})
	fmt.Println(resp)
}
