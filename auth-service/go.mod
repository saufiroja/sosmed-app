module github.com/saufiroja/sosmed-app/auth-service

go 1.22.1

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	github.com/oklog/ulid/v2 v2.1.0
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/crypto v0.23.0
	golang.org/x/oauth2 v0.21.0
	golang.org/x/text v0.15.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2
)

require (
	cloud.google.com/go/compute/metadata v0.3.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240513163218-0867130af1f8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240513163218-0867130af1f8 // indirect
)
