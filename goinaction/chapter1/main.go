package main

import (
	"fmt"
	"go-code/goinaction/chapter1/pubsub"
)

//func main() {
//	search.Run("president")
//}
//
//func init() {
//	log.SetOutput(os.Stdout)
//}


func main() {
	ch := pubsub.GenerateNatural() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = pubsub.PrimeFilter(ch, prime) // 基于新素数构造的过滤器
	}
}