package client

import (
	"fmt"
	"testing"
)

func TestFetchBookInfo(t *testing.T) {
	i := []string{"9784048686051", "9784048686051", "9784048686051"}
	resp, _ := FetchBookInfo(i)
	fmt.Println(resp)
}
