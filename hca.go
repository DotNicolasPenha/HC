package main

type HCAMain struct {
	BaseURL string
	Headers map[string]string
}

type HCACaller struct {
	Method  string
	Path    string
	Headers map[string]string
	JSON    map[string]interface{}
}

type RequestReturn struct {
	Body   []byte
	Status string
}
