// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleApiKeyHeader handles apiKeyHeader security.
	HandleApiKeyHeader(ctx context.Context, operationName string, t ApiKeyHeader) (context.Context, error)
	// HandleApiKeyParam handles apiKeyParam security.
	HandleApiKeyParam(ctx context.Context, operationName string, t ApiKeyParam) (context.Context, error)
	// HandleBearer handles bearer security.
	HandleBearer(ctx context.Context, operationName string, t Bearer) (context.Context, error)
}

func (s *Server) securityApiKeyHeader(ctx context.Context, operationName string, req *http.Request) (context.Context, error) {
	var t ApiKeyHeader
	value := req.Header.Get("apiKey")
	t.APIKey = value
	return s.sec.HandleApiKeyHeader(ctx, operationName, t)
}
func (s *Server) securityApiKeyParam(ctx context.Context, operationName string, req *http.Request) (context.Context, error) {
	var t ApiKeyParam
	value := req.URL.Query().Get("apiKey")
	t.APIKey = value
	return s.sec.HandleApiKeyParam(ctx, operationName, t)
}
func (s *Server) securityBearer(ctx context.Context, operationName string, req *http.Request) (context.Context, error) {
	var t Bearer
	t.Token = req.Header.Get("Authorization")
	return s.sec.HandleBearer(ctx, operationName, t)
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// ApiKeyHeader provides apiKeyHeader security value.
	ApiKeyHeader(ctx context.Context, operationName string) (ApiKeyHeader, error)
	// ApiKeyParam provides apiKeyParam security value.
	ApiKeyParam(ctx context.Context, operationName string) (ApiKeyParam, error)
	// Bearer provides bearer security value.
	Bearer(ctx context.Context, operationName string) (Bearer, error)
}

func (s *Client) securityApiKeyHeader(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.ApiKeyHeader(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"ApiKeyHeader\"")
	}
	req.Header.Set("apiKey", t.APIKey)
	return nil
}
func (s *Client) securityApiKeyParam(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.ApiKeyParam(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"ApiKeyParam\"")
	}
	q := req.URL.Query()
	q.Set("apiKey", t.APIKey)
	req.URL.RawQuery = q.Encode()
	return nil
}
func (s *Client) securityBearer(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.Bearer(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"Bearer\"")
	}
	req.Header.Set("Authorization", t.Token)
	return nil
}