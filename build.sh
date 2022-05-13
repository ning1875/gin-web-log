
CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o gin-web-log main.go

docker build -t registry.cn-beijing.aliyuncs.com/ning1875_haiwai_image/gin-web-log:v1.0   .

# push到阿里云
docker push registry.cn-beijing.aliyuncs.com/ning1875_haiwai_image/gin-web-log:v1.0

