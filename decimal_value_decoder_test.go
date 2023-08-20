package customregistry

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func noerr(t *testing.T, err error) {
	if err != nil {
		t.Helper()
		t.Errorf("Unexpected error: (%T)%v", err, err)
		t.FailNow()
	}
}

func TestDecimalValueDecoder(t *testing.T) {
	sampleDecimal, _ := primitive.ParseDecimal128("1")
	oneDecimal, _ := decimal.NewFromString("1")
	// zeroDecimal, _ := decimal.NewFromString("0")

	decoder := bsoncodec.ValueDecoderFunc(DecimalDecodeValue)
	testCases := []struct {
		name     string
		val      interface{}
		bsontype bsontype.Type
		want     decimal.Decimal
	}{
		{
			"Valid Int32",
			int32(1),
			bson.TypeInt32,
			oneDecimal,
		},
		{
			"Valid Int64",
			int64(1),
			bson.TypeInt64,
			oneDecimal,
		},
		{
			"Valid Double",
			float64(1),
			bson.TypeDouble,
			oneDecimal,
		},
		{
			"Valid String",
			"1",
			bson.TypeString,
			oneDecimal,
		},
		{
			"Valid Decimal128",
			sampleDecimal,
			bson.TypeDecimal128,
			oneDecimal,
		},
	}

	registry := bson.NewRegistry()
	Register(registry)

	for _, tt := range testCases {
		llvr := &bsonrwtest.ValueReaderWriter{BSONType: tt.bsontype, Return: tt.val}
		val := reflect.New(reflect.TypeOf(decimal.Decimal{})).Elem()

		err := decoder.DecodeValue(bsoncodec.DecodeContext{Registry: registry}, llvr, val)
		noerr(t, err)

		got := val.Interface().(decimal.Decimal)
		if !cmp.Equal(got, tt.want) {
			t.Fatalf("got %+v, want %+v", got, tt.want)
		}
	}
}
