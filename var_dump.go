package dump

import (
	"fmt"
	"reflect"
)


type Dumper struct {
    formatter formatter
}

type formatter interface {
    formatNumeric(value string) string
    formatString(value string) string
    formatStruct(d Dumper, s dStruct) string
}

func (f Dumper) ToString(value interface{}) string {
    return f.dumpValue(reflect.ValueOf(value))
}

func (d Dumper) dumpValue(value reflect.Value) string {
    kind := value.Kind()
    switch kind {
    case reflect.Int:
        return d.formatter.formatNumeric(fmt.Sprintf("%d", value.Int()))
    case reflect.Float32:
        return d.formatter.formatNumeric(fmt.Sprintf("%v", value.Float()))
    case reflect.Float64:
        return d.formatter.formatNumeric(fmt.Sprintf("%v", value.Float()))
    case reflect.String:
        return d.formatter.formatString(fmt.Sprintf("%s", value.String()))
    case reflect.Struct:
        ds := dStruct{
            name: value.Type().Name(),
            fields: []dStructField{},
        }
        for i := 0; i < value.NumField(); i++ {
            ds.fields = append(ds.fields, dStructField{
                name: value.Type().Field(i).Name,
                value: d.valueOfField(value.Field(i)),
            })
        }
        return d.formatter.formatStruct(d, ds)
    }
    panic(fmt.Sprintf("Did not know how to format: %s", kind))
}

func (f Dumper) valueOfField(v reflect.Value) reflect.Value {
    if v.Kind() == reflect.Interface && !v.IsNil() {
        return v.Elem()
    }
    return v
}

type dStruct struct {
    name string;
    fields []dStructField;
}

type dStructField struct {
    name string;
    value reflect.Value;
}
