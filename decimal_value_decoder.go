package decimalcodec

import (
	"fmt"
	"reflect"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

func DecimalDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	td := reflect.TypeOf(decimal.Decimal{})
	if !val.IsValid() || !val.CanSet() || val.Type() != td {
		return bsoncodec.ValueDecoderError{Name: "DecimalDecodeValue", Types: []reflect.Type{td}, Received: val}
	}

	var d decimal.Decimal
	var err error

	switch vrType := vr.Type(); vrType {
	case bson.TypeDecimal128:
		dec, err := vr.ReadDecimal128()
		i, e, err := dec.BigInt()
		if err != nil {
			return err
		}
		d = decimal.NewFromBigInt(i, int32(e))
	case bson.TypeDouble:
		f64, err := vr.ReadDouble()
		if err != nil {
			return err
		}
		d = decimal.NewFromFloat(f64)
	case bson.TypeInt32:
		i32, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		d = decimal.NewFromInt32(i32)
	case bson.TypeInt64:
		i64, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		d = decimal.NewFromInt(i64)
	case bson.TypeString:
		str, err := vr.ReadString()
		if err != nil {
			return err
		}
		d, err = decimal.NewFromString(str)
	case bson.TypeNull:
		if err = vr.ReadNull(); err != nil {
			return err
		}
	case bson.TypeUndefined:
		if err = vr.ReadUndefined(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot decode %v into a decimal.Decimal", vrType)
	}

	val.Set(reflect.ValueOf(d))

	return nil
}
