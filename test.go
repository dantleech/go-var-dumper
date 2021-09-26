package main

import "github.com/dantleech/go-var-dump/dump"


type Barfoo struct {
    one int;
    two *Foobar
}

type Foobar struct {
    foobar string;
    barfoo Barfoo;
    baz int
    boo int
    bazboo Barfoo;
}
func main() {
    dump.Dump("FOOBAR")

    f := Foobar{
        foobar: "Hello",
        barfoo: Barfoo{
            one: 12,
        },
        baz: 32,
        boo: 12,
        bazboo: Barfoo{},
    }

    dump.Dump(f)
}
