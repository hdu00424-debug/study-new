package main

//常量
import (
	"fmt"
	"math"
)

const s string = "hello"

func main() {
	fmt.Println(s)
	const n = 5000000000000000 //这是一个牛逼的地方，如果用  var n int 来定义那只能到32位会溢出
	//但是这个const 是一个无类型的定义所以他可以自由长度，因为它是无类型的
	fmt.Println(n)
	const d = 3e20 / n //3e20 = 3 × 10²⁰ = 300000000000000000000科学计数法
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
