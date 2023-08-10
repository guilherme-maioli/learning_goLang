package main

import (
	"flag"
	"fmt"
	"stats/sts"
	"stats/utils"
	"strings"
)

func main() {

	op := flag.String("op", "sum", "Tipos de operação para conjunto de números")
	file := flag.String("file", "", "Endereço do arquivo de texto com uma lista de valores.")
	flag.Parse()

	var data []float64

	if *file == "" {
		args := strings.Join(flag.Args(), " ")
		data = utils.TransformData(args, " ")

	} else {
		content := utils.ImportFile(*file)
		data = utils.TransformData(content, "\n")
	}

	switch *op {
	case "sum":
		fmt.Println("Soma: ", sts.SomaFloat64(data))
	case "avg":
		fmt.Println("Média: ", sts.AvgFloat64(data))
	}
}
