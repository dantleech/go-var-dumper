package dump

import "reflect"

type Dumper struct {
    formatter formatter;
}

func (d Dumper) ToString(value interface{}) string {
    return d.formatter.Format(
        loadData(
            context{},
            reflect.ValueOf(value),
        ),
    )
}

func loadData(ctx context, value reflect.Value) data {
    return data{
        value: value,
    }
}

type formatter interface {
    Format(data data) string;
}

type data struct {
    value reflect.Value;
}

type context struct {
    depth int;
}
