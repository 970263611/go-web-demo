# go-web-demo
windows部署
go build main.go
会将main.go涉及到的文件都编译打包好，放在当前目录下，文件名为xxx.exe（windows平台下默认编译为exe文件，可修改）
linux部署 必须用windows的cmd，不能使用powershell或者git bash 和 cmder等工具
set GOARCH=amd64
set GOOS=linux
go build
将该文件放入linux系统某个文件夹下，chmod 773 [文件名] 赋予文件权限，./xxx 命令即可执行文件，不需要go的任何依赖，就可以直接运行了。