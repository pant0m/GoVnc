package main

import (
	"flag"
	"fmt"
	"image/png"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-co-op/gocron"
	"github.com/vova616/screenshot"
)

func screen() {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	f, err := os.Create("./ss.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	f.Close()
}

func uploadoss(Endpoint string, AccessKeyId string, AccessKeySecret string, bucketName string, LocalFile string) { //上传到阿里云oss
	var hostname string

	hostname, _ = os.Hostname()
	// IP, err = GetOutBoundIP()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		fmt.Print(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Print(err)
	}

	now := strconv.FormatInt(time.Now().Unix(), 10)
	myobject := hostname + "_" + now + ".jpg"
	// 上传文件。
	err = bucket.PutObjectFromFile(myobject, LocalFile)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + ": upload " + myobject + " succeeded")
	}
}

func uploadingfile(str1 []string, Seconds int) {

	directory := "./ss.png"
	dir := strings.Replace(directory, "\\", "/", -1)
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(timezone)

	// 每秒执行一次
	s.Every(Seconds).Seconds().Do(func() {
		screen()                                           //截屏
		uploadoss(str1[3], str1[1], str1[2], str1[0], dir) //上传
	})
	s.StartBlocking()

}

func GetOutBoundIP() (ip string, err error) { //获取本地使用的IP地址
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func main() {

	var osskey string
	var Seconds int

	flag.StringVar(&osskey, "o", "", "format: bucketName:accessKeyId:accessKeySecret:endpoint")
	flag.IntVar(&Seconds, "s", 30, "second")
	flag.Parse()
	str1 := strings.Split(osskey, ":")
	uploadingfile(str1, Seconds)
}
