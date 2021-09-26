package dump

import (
	"fmt"
	"reflect"
	"strings"
)

type PlainFormatter struct {
}

func (f PlainFormatter) Format(data data) string {
    return fmt.Sprintf("%#v", data.value)
}

type AnsiFormatter struct {
    NumericFormat string
    StringFormat string
}

func (f AnsiFormatter) Format(d data) string {
    kind := d.value.Kind()
    switch kind {
    case reflect.Int:
        return f.numeric(fmt.Sprintf("%d", d.value.Int()))
    case reflect.Float32:
        return f.numeric(fmt.Sprintf("%v", d.value.Float()))
    case reflect.Float64:
        return f.numeric(fmt.Sprintf("%v", d.value.Float()))
    case reflect.String:
        return f.string(fmt.Sprintf("%s", d.value.String()))
    case reflect.Struct:
        out := []string{}
        for i := 0; i < d.value.NumField(); i++ {
            field := d.value.Field(i);
            formattedValue := f.Format(
                data{
                    value: f.valueOfField(field),
                },
            )

            out = append(out, fmt.Sprintf(
                "%s:%s",
                d.value.Type().Field(i).Name,
                formattedValue,
            ))
        }
        return fmt.Sprintf("%s{%s}", d.value.Type().Name(), strings.Join(out, " "))
    }
    panic(fmt.Sprintf("Did not know how to format: %s", kind))
}

func (f AnsiFormatter) numeric(value string) string {
    return fmt.Sprintf(f.NumericFormat, value)
}
func (f AnsiFormatter) string(value string) string {
    return fmt.Sprintf(f.StringFormat, value)
}
func (f AnsiFormatter) valueOfField(v reflect.Value) reflect.Value {
    if v.Kind() == reflect.Interface && !v.IsNil() {
        return v.Elem()
    }
    return v
}
