# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.16.5 as builder

# 启用go module
ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /home/yongjie/go/src/yingyi.cn/go-universal

COPY . /home/yongjie/go/src/yingyi.cn/go-universal
# 指定OS等，并go build
RUN go mod tidy && go build .

# 由于我不止依赖二进制文件，还依赖views文件夹下的html文件还有assets文件夹下的一些静态文件
# 所以我将这些文件放到了publish文件夹
#RUN mkdir publish && cp main publish &&  cp -r views publish && cp -r assets publish

# 运行阶段指定scratch作为基础镜像
FROM alpine

WORKDIR /home/yongjie/go/src/yingyi.cn/go-universal

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /home/yongjie/go/src/yingyi.cn/go-universal .

# 指定运行时环境变量
ENV GIN_MODE=release \
    PORT=8080

EXPOSE 8080

ENTRYPOINT ["./main"]