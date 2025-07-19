package main
/* 这里实际运行网页访问会产生两段日志，因为会网页会产生:

 根路径 (/) 请求：这是你主动访问http://localhost:12345/时触发的请求
 网站图标请求 (/favicon.ico)：浏览器会自动请求这个 URL 获取网站图标
 */
import (
	"fmt"      // 用于格式化输入输出
	"log"      // 用于记录日志信息
	"net/http" // 提供HTTP客户端和服务器实现
	"time"     // 提供时间相关功能
)

func main() {
	// 提示用户访问本地服务器地址
	fmt.Println("Please visit http://127.0.0.1:12345/")
	
	// 注册根路径"/"的处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// 构建响应内容，包含当前时间
		s := fmt.Sprintf("你好，世界 -- Time: %s", time.Now().String())
		
		// 将响应内容写入客户端
		fmt.Fprintf(w, "%v\n", s)
		
		// 同时将响应内容记录到服务器日志
		log.Printf("%v\n", s)
	})
	
	// 启动HTTP服务器，监听本地12345端口
	// nil参数表示使用默认的HTTP请求多路复用器
	if err := http.ListenAndServe(":12345", nil); err != nil {
		// 若服务器启动失败，记录错误并终止程序
		log.Fatal("ListenAndServe:", err)
	}
}