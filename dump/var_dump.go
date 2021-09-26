package dump

import (
	"fmt"
	"reflect"
)


type Dumper struct {
    printer printer
}

type printer interface {
    formatNumeric(value string) string
    formatString(value string) string
    formatStruct(d Dumper, ctx context, s dStruct) string
}

func (f Dumper) ToString(value interface{}) string {
    return f.dumpValue(newContext(), reflect.ValueOf(value))
}

func (d Dumper) dumpValue(ctx context, value reflect.Value) string {
    kind := value.Kind()
    switch kind {
    case reflect.Int:
        return d.printer.formatNumeric(fmt.Sprintf("%d", value.Int()))
    case reflect.Float32:
        return d.printer.formatNumeric(fmt.Sprintf("%v", value.Float()))
    case reflect.Float64:
        return d.printer.formatNumeric(fmt.Sprintf("%v", value.Float()))
    case reflect.String:
        return d.printer.formatString(fmt.Sprintf("%s", value.String()))
    case reflect.Ptr:
        ctx.incPointer(value.Pointer())
        return d.dumpValue(ctx, d.valueOfField(value.Elem()))
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
        return d.printer.formatStruct(d, ctx, ds)
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

func newContext() context {
    return context{
        pointers: make(map[uintptr]int),
    }
}

type context struct {
    pointers map[uintptr]int
}

func (c *context) incPointer(ptr uintptr) {
    if _, ok := c.pointers[ptr]; ok {
        c.pointers[ptr]++
        return;
    }

    c.pointers[ptr] = 1
}
