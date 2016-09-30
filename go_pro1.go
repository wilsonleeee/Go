// GO 语言学习day 1
// 2016/9/30
//
// go语言是编译型语言 脚本化 简洁
//《Go并发编程实战》
// Go的优势就是比Python更快

package main //命令源文件必须在这里声明自己属于main包

import ( //引入了代码包fmt和runtime
	"fmt"
	"runtime"
)

func init() {
	fmt.Printf("Map:%v\n", m) //格式化在打印
	//通过调用runtime包的代码获取当前机器运行的操作系统以及计算架构
	//而后通过fmt包的Sprintf方法进行字符串格式化并赋值给变量info
	info = fmt.Sprintf("OS:%s,Arch:%s,runtime.GOOS,runtime.GOARCH")
}

var m map[int]string = map[int]string{1: "A", 2: "B", 3: "C"}

//非局部变量，map类型，已被显式赋值

var info string

func main() { //命令源码文件必须有的入口函数
	fmt.Println(info) //打印变量info
}
