// This file was auto-generated by Fern from our API Definition.

package logging

import (
	context "context"
	fmt "fmt"
	generatorexecgo "github.com/fern-api/generator-exec-go"
	core "github.com/fern-api/generator-exec-go/core"
	http "net/http"
)

type Client interface {
	SendUpdate(ctx context.Context, taskId generatorexecgo.TaskId, request []*generatorexecgo.GeneratorUpdate) error
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func (c *client) SendUpdate(ctx context.Context, taskId generatorexecgo.TaskId, request []*generatorexecgo.GeneratorUpdate) error {
	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/generator-logging/%v", taskId)

	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		nil,
		false,
		c.header,
		nil,
	); err != nil {
		return err
	}
	return nil
}
