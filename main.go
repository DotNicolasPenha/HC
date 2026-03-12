package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "run" {
		fmt.Println("Usage: hca run @caller:action")
		return
	}

	callerArg := os.Args[2]
	if !strings.HasPrefix(callerArg, "@") {
		fmt.Println("Caller name must start with @")
		return
	}

	mainHCA, err := parseMain("main.hca")
	if err != nil {
		fmt.Println("Error reading main.hca:", err)
		return
	}

	parts := strings.Split(callerArg[1:], ":")
	if len(parts) != 2 {
		fmt.Println("Invalid caller format")
		return
	}

	endpointFile := fmt.Sprintf("%s.hca", parts[0])
	callers, err := parseHCAFile(endpointFile)
	if err != nil {
		fmt.Println("Error reading caller file:", err)
		return
	}

	callerName := fmt.Sprintf("%s:%s", parts[0], parts[1])
	caller, ok := callers[callerName]
	if !ok {
		fmt.Println("Caller not found:", callerName)
		return
	}
	reqReturn, err := createRequest(caller, mainHCA)
	if err != nil {
		fmt.Println("Request err: ", err.Error())
		return
	}
	fmt.Println("Status:", reqReturn.Status)
	fmt.Println("Body: ", string(reqReturn.Body))
}
