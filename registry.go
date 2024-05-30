package decimalcodec

import (
	"reflect"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

func Register(registry *bsoncodec.Registry) {
	t := reflect.TypeOf(decimal.Decimal{})
	registry.RegisterTypeEncoder(t, bsoncodec.ValueEncoderFunc(DecimalEncodeValue))
	registry.RegisterTypeDecoder(t, bsoncodec.ValueDecoderFunc(DecimalDecodeValue))
}

func NewDefaultRegistryWithDecimalCodec() *bsoncodec.Registry {
	registry := bson.DefaultRegistry
	Register(registry)

	return registry
}
