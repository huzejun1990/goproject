package main

import "fmt" //fmt 包中提供格式化，输出，输入的函数
func main() {

	fmt.Println("tom\tjack")
	fmt.Println("hello\nworld")
	fmt.Println("F:\\360Downloads\\go系列视频\\视频-1（更新）\\视频")
	fmt.Println("tom说\"i like you\"")

	// \r 回车，从当前行的最前面开始输出，覆盖掉以前的内容
	fmt.Println("天龙八部雪山飞狐\r张三厉害")

	fmt.Println("helloworldhelloworldhelloworld\n", "helloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworld\n", "helloworldhelloworldhelloworldhelloworldhelloworld\n")
	//var num = 2 + 4 * 5
}
