package main

import (
	"context"
	"github.com/fatih/color"
	"net/http"
)

type Result struct {
	url    string
	err    error
	status int
}

func healthCheck(ctx context.Context, url string, results chan Result) {

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		results <- Result{url: url, err: err}
		return
	}
	defer resp.Body.Close()
	results <- Result{url: url, err: nil, status: resp.StatusCode}

}

func printer(res Result) {
	if res.err != nil {
		color.Red("The health check for %s failed with error: %v\n", res.url, res.err)
		return
	}
	color.Green("The health check for %s was successful with status code: %d\n", res.url, res.status)
}
