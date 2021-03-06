module github.com/whileW/enze-file

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.2
	github.com/jinzhu/gorm v1.9.16
	github.com/pkg/errors v0.8.1
	github.com/qiniu/api.v7 v7.2.5+incompatible
	github.com/qiniu/x v7.0.8+incompatible // indirect
	github.com/whileW/enze-global v0.0.0-20201118032215-4e42d0ecd29c
	qiniupkg.com/x v7.0.8+incompatible // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
