module github.com/micro/services/helloworld/web

go 1.14

require (
	github.com/micro/go-micro/v2 v2.7.1-0.20200523154723-bd049a51e637
	github.com/micro/services/helloworld v0.0.0-20200424130444-5e7802513f8a
)

replace github.com/micro/go-micro/v2 => ../../../go-micro
