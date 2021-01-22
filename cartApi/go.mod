module github.com/siliconvalley001/micro_demo001/cartApi

go 1.15

require (
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v3 v3.0.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

//replace (
//	cart  => github.com/siliconvalley001/micro_demo001/cart
//)
