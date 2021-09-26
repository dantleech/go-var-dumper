package dump

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test1 struct {
    Field1 string
    Field2 int
    private string
}

func TestDump(t *testing.T) {
    t.Run("Scalars", func(t *testing.T) {
        require.Equal(t, "n:12", testDump(12))
        require.Equal(t, "n:12.123", testDump(12.123))
        require.Equal(t, "s:hello", testDump("hello"))
    })

    t.Run("Structs", func(t *testing.T) {
        s := test1{Field1: "hello", Field2: 23}
        require.Equal(t, "test1{Field1:s:hello Field2:n:23 private:s:}", testDump(s))
    })
}

func testDump(value interface{}) string {
    dumper := Dumper{
        formatter: PlainFormatter{
            NumericFormat: "n:%s",
            StringFormat: "s:%s",
        },
    }
    return dumper.ToString(value)
}
