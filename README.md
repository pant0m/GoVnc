# GoVnc

GoVnc是使用go+阿里云oss实现对屏幕（截屏）进行监控。截屏后进行实时上传操作，从而达到屏幕监控的效果。
如果要长期监控可以加入到开启启动项，计划任务等；

## 更新 2022年12月8日 11点59分
加入键盘鼠标在活动状态下进行监控。

**仅供学习使用，请勿用作非法用途！**

**详细使用**：https://www.yuque.com/pantom/web/txcveqczgc3h8at2

## 工具使用

### 下载：
```
git clone https://github.com/pantom2077/GoVnc.git
cd .\GoVnc\
go build
```

### 参数

```bash

    /$$$$$$            /$$    /$$
    /$$__  $$          | $$   | $$
   | $$  \__/  /$$$$$$ | $$   | $$ /$$$$$$$   /$$$$$$$
   | $$ /$$$$ /$$__  $$|  $$ / $$/| $$__  $$ /$$_____/
   | $$|_  $$| $$  \ $$ \  $$ $$/ | $$  \ $$| $$      
   | $$  \ $$| $$  | $$  \  $$$/  | $$  | $$| $$      
   |  $$$$$$/|  $$$$$$/   \  $/   | $$  | $$|  $$$$$$$
        \______/  \______/     \_/    |__/  |__/ \_______/
        Author:tom v1.2

Usage of C:\Users\xx\AppData\Local\Temp\go-build3697669349\b001\exe\main.exe:
  -auto
        Screenshot according to the operation of mouse and keyboard (default true) #默认开启
  -o string
        format: bucketName:accessKeyId:accessKeySecret:endpoint
  -s int
        Interval time second (default 10) # 默认 10
```

### 效果
键盘活动时候进行截屏
![image](https://user-images.githubusercontent.com/118233720/206353777-86b24c79-a547-4880-ad79-5c9b8aa16ef4.png)

定时回传
![image](https://user-images.githubusercontent.com/118233720/206353989-06733f60-b76c-41c5-8434-7f4f11311bbe.png)


### 查看图片

![image](https://user-images.githubusercontent.com/118233720/206354074-9741a285-abf1-4913-86d5-a48e22b0c76a.png)

