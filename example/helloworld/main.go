package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"flag"

	"github.com/go-redis/redis"
)

var redisCli *redis.Client
var ip string
var startTime string
var totalCount int64

func init() {
	log.Println("init start")
	redisCli = redis.NewClient(&redis.Options{
		Addr:     "redis-srv:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := redisCli.Ping().Result()
	log.Println("redis", pong, err)
	ip = GetIP()
	log.Println("IP", ip)
	startTime = time.Now().Format("2006-01-02 15:04:05.999")
	log.Println("Time:", startTime)
	log.Println("init end")
}

type Hello struct {
}

func (h Hello) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	totalCount++
	count, err := redisCli.HIncrBy("hello", ip, 1).Result()
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
		return
	}
	bf := bytes.Buffer{}
	bf.WriteString("helloworld\nStartTime:")
	bf.WriteString(startTime)
	bf.WriteString("\nTime:")
	bf.WriteString(time.Now().Format("2006-01-02 15:04:05.999"))
	bf.WriteString("\nIP:")
	bf.WriteString(ip)
	bf.WriteString("\nCount:")
	bf.WriteString(strconv.FormatInt(count, 10))
	bf.WriteString("\nTotalCount:")
	bf.WriteString(strconv.FormatInt(totalCount, 10))
	resp.Write(bf.Bytes())
}
func main() {
	var port int
	flag.IntVar(&port, "port", 80, "启动端口")
	flag.Parse()
	log.Printf("app start with port:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), Hello{})
	if err != nil {
		log.Fatalf("err:%s\n", err)
	}
}

func GetIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return ""
}
