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

var (
	redisCliSS *redis.Client //redis客户端测试有状态服务
	redis_ss   string        //有状态redis服务名
	redisCliD  *redis.Client //redis客户端测试无状态服务
	redis_d    string        //无状态redis服务名
	ip         string        //获取本机IP
	startTime  string        //启动时间
	port       int           //启动端口
)

func init() {
	log.Println("init start")
	flag.IntVar(&port, "port", 80, "启动端口")
	flag.StringVar(&redis_ss, "redis_ss", "redis_ss", "有状态redis服务名")
	flag.StringVar(&redis_d, "redis_d", "redis_d", "无状态redis服务名")
	flag.Parse()

	ip = GetIP()
	log.Println("IP", ip, "redis_ss", redis_ss, "redis_d", redis_d)
	startTime = time.Now().Format("2006-01-02 15:04:05.999")
	log.Println("Time:", startTime)
	log.Println("init end")
	initRedis()
}
func initRedis() {
	redisCliSS = redis.NewClient(&redis.Options{
		Addr:        redis_ss,
		Password:    "", // no password set
		DB:          0,  // use default DB
		DialTimeout: 5 * time.Second,
	})
	pong, err := redisCliSS.Ping().Result()
	log.Println("redis_ss", pong, err)

	redisCliD = redis.NewClient(&redis.Options{
		Addr:        redis_d,
		Password:    "", // no password set
		DB:          0,  // use default DB
		DialTimeout: 5 * time.Second,
	})
	pong, err = redisCliD.Ping().Result()
	log.Println("redis_d", pong, err)
}

func hello(resp http.ResponseWriter, req *http.Request) {
	var (
		count  int64
		count2 int64
		err    error
	)
	if redisCliSS != nil {
		count, err = redisCliSS.HIncrBy("hello", ip, 1).Result()
		if err != nil {
			resp.WriteHeader(500)
			resp.Write([]byte(err.Error()))
			return
		}
	}
	if redisCliD != nil {
		count2, err = redisCliD.HIncrBy("hello", ip, 1).Result()
		if err != nil {
			resp.WriteHeader(500)
			resp.Write([]byte(err.Error()))
			return
		}
	}
	bf := bytes.Buffer{}
	bf.WriteString("helloworld\nStartTime:")
	bf.WriteString(startTime)
	bf.WriteString("\nTime:")
	bf.WriteString(time.Now().Format("2006-01-02 15:04:05.999"))
	bf.WriteString("\nIP:")
	bf.WriteString(ip)
	bf.WriteString("\nCountSS:")
	bf.WriteString(strconv.FormatInt(count, 10))
	bf.WriteString("\nCountD:")
	bf.WriteString(strconv.FormatInt(count2, 10))
	resp.Write(bf.Bytes())
}
func main() {
	log.Printf("app start with port:%d\n", port)
	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}
	http.HandleFunc("/", hello)
	server.ListenAndServe()
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
