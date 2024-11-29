package models

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

func MarshalProductId(i ProductId) graphql.Marshaler {
	return marshalIds[ProductId](i)
}

func MarshalOrderId(i OrderId) graphql.Marshaler {
	return marshalIds[OrderId](i)
}

func MarshalUserId(i UserId) graphql.Marshaler {
	return marshalIds[UserId](i)
}

func UnmarshalProductId(v any) (ProductId, error) {
	return unmarshalIds[ProductId](v)
}

func UnmarshalOrderId(v any) (OrderId, error) {
	return unmarshalIds[OrderId](v)
}

func UnmarshalUserId(v any) (UserId, error) {
	return unmarshalIds[UserId](v)
}

func marshalIds[T UserId | ProductId | OrderId](i T) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(int64(i), 10))
	})
}

func unmarshalIds[T UserId | ProductId | OrderId](v any) (T, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return 0, err
		}
		return T(iv), nil
	case int:
		return T(v), nil
	case int64:
		return T(v), nil
	case json.Number:
		iv, err := strconv.ParseInt(string(v), 10, 32)
		if err != nil {
			return 0, err
		}
		return T(iv), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}
