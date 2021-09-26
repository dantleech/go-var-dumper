package dump

import (
	"fmt"
	"reflect"
	"strings"
)

func newAnsiPrinter() formatPrinter {
    return formatPrinter{
        NumericFormat: "\033[1;34m%s\033[0m",
        StringFormat:  "\033[1;36m%s\033[0m",
        StructFormat: "\033[1;33m%s\033[0m{\n%s}",
        StructFieldFormat: "\033[0;36m%s\033[0m:%s\n",
    }
}

func newPlainPrinter() formatPrinter {
    return formatPrinter{
        NumericFormat: `%s`,
        StringFormat: `"%s"`,
        StructFormat: `%s{%s}`,
        StructFieldFormat: `%s:%s`,
    }
}

type formatPrinter struct {
    NumericFormat string
    StringFormat string
    StructFormat string
    StructFieldFormat string
}

func (f formatPrinter) formatNumeric(value string) string {
    return fmt.Sprintf(f.NumericFormat, value)
}
func (f formatPrinter) formatString(value string) string {
    return fmt.Sprintf(f.StringFormat, value)
}
func (f formatPrinter) formatStruct(d Dumper, ctx context, s dStruct) string {
    out := []string{};
    for _, field := range(s.fields) {
        out = append(out, fmt.Sprintf(f.StructFieldFormat, field.name, d.dumpValue(ctx, field.value)))
    }

    return fmt.Sprintf(f.StructFormat, s.name, strings.Join(out, " "))
}
func (f formatPrinter) formatPointer(d Dumper, ctx context, v reflect.Value) string {
    return fmt.Sprintf("*%d#%s", v.Pointer(), d.dumpValue(ctx, v.Elem()))
}
func (f formatPrinter) formatCircularPointer(d Dumper, ctx context, v reflect.Value) string {
    return fmt.Sprintf("*%d#**circular**", v.Pointer())
}
