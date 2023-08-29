package gomarshall

import (
	"encoding/json"
	"reflect"
)

type Options struct {
	IgnoreNil bool
}

var defaultOptions = Options{IgnoreNil: true}

func RawValueToValue(v reflect.Value, opts Options) interface{} {
	switch v.Kind() {
	case reflect.Slice:
		var x []interface{}
		for i := 0; i < v.Len(); i++ {
			x = append(x, RawValueToValue(v.Index(i), opts))
		}
		return x
	case reflect.Pointer:
		return RawValueToValue(v.Elem(), opts)
	case reflect.Bool:
		return v.Bool()
	case reflect.String:
		return v.String()
	case reflect.Int:
		return v.Int()
	case reflect.Int8:
		return v.Int()
	case reflect.Int16:
		return v.Int()
	case reflect.Int32:
		return v.Int()
	case reflect.Int64:
		return v.Int()
	case reflect.Uint:
		return v.Uint()
	case reflect.Uint8:
		return v.Uint()
	case reflect.Uint16:
		return v.Uint()
	case reflect.Uint32:
		return v.Uint()
	case reflect.Uint64:
		return v.Uint()
	case reflect.Float32:
		return v.Float()
	case reflect.Float64:
		return v.Float()
	case reflect.Func:
		return struct{}{}
	case reflect.Map:
		// only map[string]something is supported
		x := map[string]interface{}{}
		iter := v.MapRange()
		for iter.Next() {
			kk := iter.Key()
			vv := iter.Value()
			switch kk.Kind() {
			case reflect.String:
				x[kk.String()] = vv.Interface()
			}
		}
		return x
	case reflect.Struct:
		m, x, y := map[string]interface{}{}, v.Type(), v
		for i := 0; i < x.NumField(); i++ {
			kk := x.Field(i)
			k := kk.Name
			vv := y.Field(i)
			if kk.IsExported() {
				v := RawValueToValue(vv, opts)
				if !opts.IgnoreNil || nil != v {
					m[k] = v
				}
			}
		}
		return m
	case reflect.Interface:
		return RawValueToValue(reflect.ValueOf(v.Interface()), opts)
	default:
		return nil
	}
}

func ValueToMarshallable(s interface{}, opts Options) interface{} {
	return RawValueToValue(reflect.ValueOf(s), opts)
}

func ToJsonBytes(s interface{}, opts Options) ([]byte, error) {
	return json.Marshal(ValueToMarshallable(s, opts))
}

//goland:noinspection GoUnusedExportedFunction
func V(s interface{}) interface{} {
	return ValueToMarshallable(s, defaultOptions)
}
