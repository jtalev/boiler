package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) > 3 {
		fmt.Println("not a boiler command")
	}
	if len(os.Args) < 2 {
		fmt.Println("'ello boiler")
	}

	var suffix string
	arg := os.Args[1]
	stop_index := strings.LastIndex(arg, ".")
	if stop_index == -1 {
		fmt.Println("invalid file name")
	} else {
		for i := stop_index; i < len(arg); i++ {
			suffix = suffix + string(arg[i])
		}
	}

	fmt.Println(suffix)

	switch suffix {
	case ".html":
		write_file([]byte(
			`<!DOCTYPE html>\n<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Document</title>
</head>
	<body>

	</body>
</html>`))
	default:
		fmt.Println("filetype not implemented")
	}
}

func write_file(boiler_plate []byte) {
	err := os.WriteFile(os.Args[1], boiler_plate, 0755)
	if err != nil {
		fmt.Printf("error writing file: %s", err)
	}
}
