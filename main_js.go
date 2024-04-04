package main

import (
	"github.com/VictoriaMetrics/metricsql"
	"syscall/js"
)

type JsReturnValue struct {
    prettyMetricsQL string
    err error
}

func Prettier(this js.Value, p []js.Value) interface{} {
    raw := p[0].String()
	prettyMetricsQL, err := metricsql.Prettify(raw)
	if err != nil {
		return []interface{}{"", err.Error()}
	}
	return []interface{}{prettyMetricsQL, nil}
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("prettier", js.FuncOf(Prettier))

	<-c
}
