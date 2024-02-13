package rules

import (
	"io"
)

type Interface interface {
	ConvertIoReaderToStruct(data io.Reader, model interface{}) (body interface{}, err error)
	GetMock() interface{}
	// Migrate(connection *dynamodb.DynamoDB) error
	Validate(model interface{}) error
}
