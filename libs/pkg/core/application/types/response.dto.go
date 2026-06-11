package cdto

type ApplicationResponse[D any, R any] interface {
	GetData() D
	ToGrpcResponse() R
}
