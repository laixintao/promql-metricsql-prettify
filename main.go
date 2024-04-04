// +build !js

package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/VictoriaMetrics/metricsql"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Printf("Error when read from stdin: %s", err.Error())
	}
	metricsqlString := string(data)
	metricsqlPretty, err := metricsql.Prettify(metricsqlString)
	if err != nil {
		log.Fatalf("Error when Prettify metricsql: %s", err.Error())
	}
	fmt.Println(metricsqlPretty)
}
