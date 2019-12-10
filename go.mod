module github.com/purpledb/purple-go-client

go 1.13

require (
	github.com/go-resty/resty/v2 v2.0.0
	github.com/purpledb/purple v0.1.6
	github.com/stretchr/testify v1.3.0
	google.golang.org/grpc v1.22.0
)

replace github.com/lucperkins/strato => github.com/purpledb/purple v0.1.2
