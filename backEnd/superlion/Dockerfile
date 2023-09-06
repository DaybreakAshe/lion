
FROM golang:latest    
#创建一个文件夹
RUN mkdir /superlion
# 移动到指定的目录
WORKDIR /superlion
# 将当前的目录的文件都拷贝到/app目录
ADD ./* /superlion
#构建go二进制文件
RUN rm -rf go.mod && go mod init superlion && go get github.com/gin-gonic/gin && go build -o main ./main.go ./router.go
# 输出的端口号
EXPOSE 8080        
# 执行,有点像shell
CMD /superlion/main
 
