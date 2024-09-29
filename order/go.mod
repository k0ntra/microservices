module github.com/k0ntra/microservices/order

go 1.23.1

require (
	github.com/k0ntra/microservices-proto/golang/order v0.0.0-20240929091503-89d6a455a7bc
	github.com/k0ntra/microservices-proto/golang/payment v0.0.0-20240927142127-62730a27aa90
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/grpc v1.67.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
