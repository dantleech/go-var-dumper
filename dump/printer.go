package dump

import (
	"fmt"
	"strings"
)

func NewPlainPrinter() formatPrinter {
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
