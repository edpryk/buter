package docs

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Input struct {
	Url        string
	AttackType string
	Headers    string
	// Headers       map[string]string
	ThreadsInTime int
	PayloadFiles
}

/*
Parse flgas validate it
and eather return it or
pritn an error and usage
and exis process
*/
func ParseFlags() Input {
	input := Input{
		// Headers: make(map[string]string),
	}

	flag.Var(&input.PayloadFiles, payloadFlag, payloadUsage)
	flag.StringVar(&input.Url, urlFlag, defaultUrl, urlUsage)
	flag.StringVar(&input.AttackType, attackTypeFlag, defaultAttackType, attackTypeUsage)
	flag.IntVar(&input.ThreadsInTime, threadsFlag, defaultThreads, threadsUseage)
	// headersStr := flag.String(headersFlag, defaultHeaders, headersUsage)
	flag.StringVar(&input.Headers, headersFlag, defaultHeaders, headersUsage)
	flag.Parse()

	d := make(map[string]string)
	if err := json.Unmarshal([]byte(input.Headers), &d); err != nil {
		fmt.Println("Can't parse headers", err)
		os.Exit(1)
	}
	b, _ := json.Marshal(d)

	input.Headers = string(b)

	if err := validateInput(input); err != nil {
		fmt.Println(err.Error())
		flag.Usage()
		os.Exit(1)
	}

	return input
}
