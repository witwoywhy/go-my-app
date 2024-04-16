package app

type WebFrameworkContext interface {
	GetHeader(key string) string
	SetHeader(key, value string)

	BindJSON(obj any) error
	JSON(code int, obj any) error

	Next()

	GetRequest() any
	GetWriter() any
	SetWriter(writer any)
}
