package main

import (
	"strings"
	"fmt"
)

func main()  {
	s := "1lis1t_etcd_host1"

	idx := strings.Index(s,"list_")
	fmt.Println(idx)
}
