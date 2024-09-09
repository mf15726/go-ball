package interfaces

type IRestClient[T any, R any] interface {
	Get(endpoint string, headers map[string]string) (*R, error)
	Post(endpoint string, body T, headers map[string]string) (*R, error)
	Put(endpoint string, body T, headers map[string]string) (*R, error)
	Delete(endpoint string, headers map[string]string) (*R, error)
}
