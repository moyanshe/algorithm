package main

import (
	"algorithm/classics/hashring"
	"fmt"
)

func main() {
	cHashRing := hashring.NewConsistent()

	for i := 0; i < 5; i++ {
		si := fmt.Sprintf("%d", i)
		cHashRing.Add(hashring.NewNode(i, "172.18.1."+si, 8080, "host_"+si, 1))
	}

	for k, v := range cHashRing.Nodes {
		fmt.Println("Hash:", k, " IP:", v.Ip)
	}

	ipMap := make(map[string]int, 0)
	for i := 0; i < 1000000; i++ {
		si := fmt.Sprintf("key%d", i)
		k := cHashRing.Get(si)
		if _, ok := ipMap[k.Ip]; ok {
			ipMap[k.Ip]++
		} else {
			ipMap[k.Ip] = 1
		}
	}

	for k, v := range ipMap {
		fmt.Println("Node IP:", k, " count:", v)
	}

}
