package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var ServiceMap = map[string]string{
	"auth":  "http://127.0.0.1:20021",
	"user":  "http://127.0.0.1:20022",
	"chat":  "http://127.0.0.1:20023",
	"group": "http://127.0.0.1:20024",
}

// Data 结构体定义了响应数据的结构，包括状态码、任意类型的数据和消息。
type Data struct {
	Code int    `json:"code"`
	Data any    `json:"data"` // Go 1.18+ 使用 any 表示任意类型
	Msg  string `json:"msg"`
}

// toJson 函数接收 Data 类型的参数，并将其序列化为 JSON 格式的字节数组。
func toJson(data Data) []byte {
	byteData, _ := json.Marshal(data)
	return byteData
}

// handler
func gateway(res http.ResponseWriter, req *http.Request) {
	//匹配请求前缀   /api/user/xxx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrList := regex.FindStringSubmatchIndex(req.URL.Path)
	if len(addrList) != 2 {
		res.Write([]byte("404"))
		return
	}
	service := addrList[1]
	addr, ok := ServiceMap[string(service)] // 根据服务名查找对应的地址。
	if !ok {                                // 如果服务名不存在于ServiceMap中。
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"})) // 返回错误信息。
		return
	}
	url := addr + req.URL.String() // 构造目标URL。

	proxyReq, _ := http.NewRequest(req.Method, url, req.Body) // 创建新的HTTP请求。
	proxyReq.Header = req.Header                              // 复制原始请求头到新的请求中。
	remoteAddr := strings.Split(req.RemoteAddr, ":")          // 分割远程地址以获取IP地址。
	if len(remoteAddr) != 2 {                                 // 如果无法正确分割远程地址。
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"})) // 返回错误信息。
		return
	}
	fmt.Printf("%s %s =>  %s\n", remoteAddr[0], service, url) // 打印转发信息。
	proxyReq.Header.Set("X-Forwarded-For", remoteAddr[0])     // 设置转发头。
	proxyResponse, err := http.DefaultClient.Do(proxyReq)     // 发送请求并接收响应。
	if err != nil {                                           // 如果发送请求时出错。
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"})) // 返回错误信息。
		return
	}
	io.Copy(res, proxyResponse.Body) // 将响应体复制回客户端。
	return
}

var configFile = flag.String("f", "setting.yaml", "the config file")

type Config struct {
	// 配置文件
	Addr string
}

func main() {
	flag.Parse()
	var c Config
	conf.MustLoad(*configFile, &c)
	// 回调函数
	http.HandleFunc("/", gateway)
	fmt.Printf("gateway running %s\n", c.Addr)
	// 绑定服务
	http.ListenAndServe(*configFile, nil)
}
