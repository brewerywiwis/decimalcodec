package decimalcodec

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw/bsonrwtest"
)

func compareErrors(err1, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	}

	if err1 == nil || err2 == nil {
		return false
	}

	if err1.Error() != err2.Error() {
		return false
	}

	return true
}
func TestDecimalValueEncoder(t *testing.T) {
	td := reflect.TypeOf(decimal.Decimal{})
	var wrong = func(string, string) string { return "wrong" }

	testCases := []struct {
		name   string
		val    interface{}
		invoke bsonrwtest.Invoked
		err    error
	}{
		{
			"Valid",
			decimal.NewFromInt(1),
			bsonrwtest.WriteDecimal128,
			nil,
		},
		{
			"Invalid",
			wrong,
			bsonrwtest.Nothing,
			bsoncodec.ValueEncoderError{
				Name:     "DecimalEncodeValue",
				Types:    []reflect.Type{td},
				Received: reflect.ValueOf(wrong),
			},
		},
	}

	for _, tt := range testCases {
		llvrw := new(bsonrwtest.ValueReaderWriter)
		llvrw.T = t
		err := bsoncodec.ValueEncoderFunc(DecimalEncodeValue).EncodeValue(bsoncodec.EncodeContext{}, llvrw, reflect.ValueOf(tt.val))
		if !compareErrors(err, tt.err) {
			t.Errorf("Errors do not match. got %v; want %v", err, tt.err)
		}
		invoked := llvrw.Invoked
		if !cmp.Equal(invoked, tt.invoke) {
			t.Errorf("Incorrect method invoked. got %v; want %v", invoked, tt.invoke)
		}
	}
}
