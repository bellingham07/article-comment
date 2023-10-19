dev:
	go run ./api/hibug-track.go -f ./api/etc/hibug-track.yaml

build:
	go clean && go mod tidy &&  GOOS=linux GOARCH=amd64  go build -o hibug-track  -ldflags="-X main.BuildStamp=`date +%Y-%m-%d.%H:%M:%S`" ./api/hibug-track.go


docker-build:
	docker build -t hibug-track:1.0 -f Dockerfile .

gen-api:
	goctl api go -api ./api/desc/all.api -dir ./api --style=go_zero --home="./template"

#goctl 生成表结构对于的gorm代码
gen-model:
	goctl model pg datasource \
    -url="host=1.2.3.4 user=user password=pwd dbname=dbname port=5432 sslmode=disable TimeZone=Asia/Shanghai" \
     -table="*" \
     -dir="./api/internal/model" \
     -home="./template/gorm"

# mongo
gen-mongo:
	goctl model mongo -e -dir /model/user -t edgeinstance


gen-model-sql:
	goctl model pg ddl --src ./resources/ddl.sql -dir .