package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func parseHCAFile(filePath string) (map[string]HCACaller, error) {
	contentBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := string(contentBytes)

	callers := make(map[string]HCACaller)

	reCaller := regexp.MustCompile(`@(\w+):(\w+)[\s\S]*?m:\s*(\w+)[\s\S]*?p:\s*(\S+)[\s\S]*?h:\s*(\{[\s\S]*?\})[\s\S]*?j:\s*(\{[\s\S]*?\})`)
	matches := reCaller.FindAllStringSubmatch(content, -1)

	for _, m := range matches {
		callerName := fmt.Sprintf("%s:%s", m[1], m[2])
		headers := make(map[string]string)
		jsonBody := make(map[string]interface{})

		json.Unmarshal([]byte(m[5]), &headers)
		json.Unmarshal([]byte(m[6]), &jsonBody)

		callers[callerName] = HCACaller{
			Method:  strings.ToUpper(m[3]),
			Path:    m[4],
			Headers: headers,
			JSON:    jsonBody,
		}
	}

	return callers, nil
}

func parseMain(filePath string) (HCAMain, error) {
	contentBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return HCAMain{}, err
	}
	content := string(contentBytes)

	reBase := regexp.MustCompile(`baseurl:\s*(\S+)`)
	baseMatch := reBase.FindStringSubmatch(content)

	headers := make(map[string]string)
	reHeaders := regexp.MustCompile(`h:\s*(\{.*\})`)
	headersMatch := reHeaders.FindStringSubmatch(content)
	if len(headersMatch) > 1 {
		json.Unmarshal([]byte(headersMatch[1]), &headers)
	}

	return HCAMain{
		BaseURL: baseMatch[1],
		Headers: headers,
	}, nil
}
