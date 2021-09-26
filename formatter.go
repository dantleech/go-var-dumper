package dump

import (
	"fmt"
	"strings"
)

type PlainFormatter struct {
    NumericFormat string
    StringFormat string
}

func (f PlainFormatter) formatNumeric(value string) string {
    return fmt.Sprintf(f.NumericFormat, value)
}
func (f PlainFormatter) formatString(value string) string {
    return fmt.Sprintf(f.StringFormat, value)
}
func (f PlainFormatter) formatStruct(d Dumper, s dStruct) string {
    out := []string{};
    for _, field := range(s.fields) {
        out = append(out, fmt.Sprintf("%s:%s", field.name, d.dumpValue(field.value)))
    }

    return fmt.Sprintf("%s{%s}", s.name, strings.Join(out, " "))
}
