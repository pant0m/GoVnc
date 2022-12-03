package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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

	t := time.NewTicker(time.Duration(Seconds) * time.Second) //定时器
	defer t.Stop()
	for {
		<-t.C
		screen()
		uploadoss(str1[3], str1[1], str1[2], str1[0], dir)
	}
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
