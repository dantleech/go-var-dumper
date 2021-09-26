package dump

import (
	"fmt"
	"reflect"
	"strings"
)

type Dumper struct {
    NumericFormat string
    StringFormat string
}

func (f Dumper) ToString(value interface{}) string {
    return f.dumpValue(reflect.ValueOf(value))
}

func (f Dumper) dumpValue(value reflect.Value) string {
    kind := value.Kind()
    switch kind {
    case reflect.Int:
        return f.numeric(fmt.Sprintf("%d", value.Int()))
    case reflect.Float32:
        return f.numeric(fmt.Sprintf("%v", value.Float()))
    case reflect.Float64:
        return f.numeric(fmt.Sprintf("%v", value.Float()))
    case reflect.String:
        return f.string(fmt.Sprintf("%s", value.String()))
    case reflect.Struct:
        ds := dStruct{
            name: value.Type().Name(),
            fields: []dStructField{},
        }
        for i := 0; i < value.NumField(); i++ {
            ds.fields = append(ds.fields, dStructField{
                name: value.Type().Field(i).Name,
                value: f.valueOfField(value.Field(i)),
            })
        }
        return f.formatStruct(ds)
    }
    panic(fmt.Sprintf("Did not know how to format: %s", kind))
}

func (f Dumper) numeric(value string) string {
    return fmt.Sprintf(f.NumericFormat, value)
}
func (f Dumper) string(value string) string {
    return fmt.Sprintf(f.StringFormat, value)
}
func (f Dumper) formatStruct(s dStruct) string {
    out := []string{};
    for _, field := range(s.fields) {
        out = append(out, fmt.Sprintf("%s:%s", field.name, f.dumpValue(field.value)))
    }

    return fmt.Sprintf("%s{%s}", s.name, strings.Join(out, " "))
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
