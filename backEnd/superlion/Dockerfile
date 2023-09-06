
FROM golang:latest    
#创建一个文件夹
RUN mkdir /app     
# 移动到指定的目录
WORKDIR /app       
# 将当前的目录的文件都拷贝到/app目录
ADD ./* /app
#构建go二进制文件
RUN go mod init && go build -o main ./main.go ./router.go
# 输出的端口号
EXPOSE 8080        
# 执行,有点像shell
CMD /app/main      
 
