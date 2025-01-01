// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// TodosFilterGet invokes GET /todos/filter operation.
	//
	// Filter TODOs based on specific conditions.
	//
	// GET /todos/filter
	TodosFilterGet(ctx context.Context, params TodosFilterGetParams) (TodosFilterGetRes, error)
	// TodosGet invokes GET /todos operation.
	//
	// Get a list of all TODOs.
	//
	// GET /todos
	TodosGet(ctx context.Context) ([]TodosGetOKItem, error)
	// TodosIDCompletePost invokes POST /todos/{id}/complete operation.
	//
	// Mark a TODO as completed.
	//
	// POST /todos/{id}/complete
	TodosIDCompletePost(ctx context.Context, params TodosIDCompletePostParams) (TodosIDCompletePostRes, error)
	// TodosIDDelete invokes DELETE /todos/{id} operation.
	//
	// Delete a TODO.
	//
	// DELETE /todos/{id}
	TodosIDDelete(ctx context.Context, params TodosIDDeleteParams) (TodosIDDeleteRes, error)
	// TodosIDGet invokes GET /todos/{id} operation.
	//
	// Get TODO details by ID.
	//
	// GET /todos/{id}
	TodosIDGet(ctx context.Context, params TodosIDGetParams) (TodosIDGetRes, error)
	// TodosIDPatch invokes PATCH /todos/{id} operation.
	//
	// Update a TODO.
	//
	// PATCH /todos/{id}
	TodosIDPatch(ctx context.Context, request *TodosIDPatchReq, params TodosIDPatchParams) (TodosIDPatchRes, error)
	// TodosIDPriorityPatch invokes PATCH /todos/{id}/priority operation.
	//
	// Change the priority of a TODO.
	//
	// PATCH /todos/{id}/priority
	TodosIDPriorityPatch(ctx context.Context, request *TodosIDPriorityPatchReq, params TodosIDPriorityPatchParams) (TodosIDPriorityPatchRes, error)
	// TodosIDReopenPost invokes POST /todos/{id}/reopen operation.
	//
	// Reopen a completed TODO.
	//
	// POST /todos/{id}/reopen
	TodosIDReopenPost(ctx context.Context, params TodosIDReopenPostParams) (TodosIDReopenPostRes, error)
	// TodosIDTagsDelete invokes DELETE /todos/{id}/tags operation.
	//
	// Remove a tag from a TODO.
	//
	// DELETE /todos/{id}/tags
	TodosIDTagsDelete(ctx context.Context, params TodosIDTagsDeleteParams) (TodosIDTagsDeleteRes, error)
	// TodosIDTagsPost invokes POST /todos/{id}/tags operation.
	//
	// Add a tag to a TODO.
	//
	// POST /todos/{id}/tags
	TodosIDTagsPost(ctx context.Context, request *TodosIDTagsPostReq, params TodosIDTagsPostParams) (TodosIDTagsPostRes, error)
	// TodosPost invokes POST /todos operation.
	//
	// Create a new TODO.
	//
	// POST /todos
	TodosPost(ctx context.Context, request *TodosPostReq) (TodosPostRes, error)
	// UsersIDDelete invokes DELETE /users/{id} operation.
	//
	// Delete a user and all related data.
	//
	// DELETE /users/{id}
	UsersIDDelete(ctx context.Context, params UsersIDDeleteParams) (UsersIDDeleteRes, error)
	// UsersIDGet invokes GET /users/{id} operation.
	//
	// Get user details by ID.
	//
	// GET /users/{id}
	UsersIDGet(ctx context.Context, params UsersIDGetParams) (UsersIDGetRes, error)
	// UsersPost invokes POST /users operation.
	//
	// Register a new user.
	//
	// POST /users
	UsersPost(ctx context.Context, request *UsersPostReq) (UsersPostRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// TodosFilterGet invokes GET /todos/filter operation.
//
// Filter TODOs based on specific conditions.
//
// GET /todos/filter
func (c *Client) TodosFilterGet(ctx context.Context, params TodosFilterGetParams) (TodosFilterGetRes, error) {
	res, err := c.sendTodosFilterGet(ctx, params)
	return res, err
}

func (c *Client) sendTodosFilterGet(ctx context.Context, params TodosFilterGetParams) (res TodosFilterGetRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/todos/filter"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TodosFilterGetOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/todos/filter"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "status" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "status",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Status.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "priority" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "priority",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Priority.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTodosFilterGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TodosGet invokes GET /todos operation.
//
// Get a list of all TODOs.
//
// GET /todos
func (c *Client) TodosGet(ctx context.Context) ([]TodosGetOKItem, error) {
	res, err := c.sendTodosGet(ctx)
	return res, err
}

func (c *Client) sendTodosGet(ctx context.Context) (res []TodosGetOKItem, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/todos"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TodosGetOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/todos"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTodosGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TodosIDCompletePost invokes POST /todos/{id}/complete operation.
//
// Mark a TODO as completed.
//
// POST /todos/{id}/complete
func (c *Client) TodosIDCompletePost(ctx context.Context, params TodosIDCompletePostParams) (TodosIDCompletePostRes, error) {
	res, err := c.sendTodosIDCompletePost(ctx, params)
	return res, err
}

func (c *Client) sendTodosIDCompletePost(ctx context.Context, params TodosIDCompletePostParams) (res TodosIDCompletePostRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/todos/{id}/complete"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TodosIDCompletePostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/todos/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/complete"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTodosIDCompletePostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TodosIDDelete invokes DELETE /todos/{id} operation.
//
// Delete a TODO.
//
// DELETE /todos/{id}
func (c *Client) TodosIDDelete(ctx context.Context, params TodosIDDeleteParams) (TodosIDDeleteRes, error) {
	res, err := c.sendTodosIDDelete(ctx, params)
	return res, err
}

func (c *Client) sendTodosIDDelete(ctx context.Context, params TodosIDDeleteParams) (res TodosIDDeleteRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/todos/{id}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TodosIDDeleteOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/todos/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTodosIDDeleteResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TodosIDGet invokes GET /todos/{id} operation.
//
// Get TODO details by ID.
//
// GET /todos/{id}
func (c *Client) TodosIDGet(ctx context.Context, params TodosIDGetParams) (TodosIDGetRes, error) {
	res, err := c.sendTodosIDGet(ctx, params)
	return res, err
}

func (c *Client) sendTodosIDGet(ctx context.Context, params TodosIDGetParams) (res TodosIDGetRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/todos/{id}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TodosIDGetOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/todos/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTodosIDGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TodosIDPatch invokes PATCH /todos/{id} operation.
//
// Update a TODO.
//
// PATCH /todos/{id}
func (c *Client) TodosIDPatch(ctx context.Context, request *TodosIDPatchReq, params TodosIDPatchParams) (TodosIDPatchRes, error) {
	res, err := c.sendTodosIDPatch(ctx, request, params)
	return res, err
}

func (c *Client) sendTodosIDPatch(ctx context.Context, request *TodosIDPatchReq, params TodosIDPatchParams) (res TodosIDPatchRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("PATCH"),
		semconv.HTTPRouteKey.String("/todos/{id}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TodosIDPatchOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/todos/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PATCH", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeTodosIDPatchRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTodosIDPatchResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TodosIDPriorityPatch invokes PATCH /todos/{id}/priority operation.
//
// Change the priority of a TODO.
//
// PATCH /todos/{id}/priority
func (c *Client) TodosIDPriorityPatch(ctx context.Context, request *TodosIDPriorityPatchReq, params TodosIDPriorityPatchParams) (TodosIDPriorityPatchRes, error) {
	res, err := c.sendTodosIDPriorityPatch(ctx, request, params)
	return res, err
}

func (c *Client) sendTodosIDPriorityPatch(ctx context.Context, request *TodosIDPriorityPatchReq, params TodosIDPriorityPatchParams) (res TodosIDPriorityPatchRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("PATCH"),
		semconv.HTTPRouteKey.String("/todos/{id}/priority"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TodosIDPriorityPatchOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/todos/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/priority"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PATCH", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeTodosIDPriorityPatchRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTodosIDPriorityPatchResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TodosIDReopenPost invokes POST /todos/{id}/reopen operation.
//
// Reopen a completed TODO.
//
// POST /todos/{id}/reopen
func (c *Client) TodosIDReopenPost(ctx context.Context, params TodosIDReopenPostParams) (TodosIDReopenPostRes, error) {
	res, err := c.sendTodosIDReopenPost(ctx, params)
	return res, err
}

func (c *Client) sendTodosIDReopenPost(ctx context.Context, params TodosIDReopenPostParams) (res TodosIDReopenPostRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/todos/{id}/reopen"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TodosIDReopenPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/todos/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/reopen"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTodosIDReopenPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TodosIDTagsDelete invokes DELETE /todos/{id}/tags operation.
//
// Remove a tag from a TODO.
//
// DELETE /todos/{id}/tags
func (c *Client) TodosIDTagsDelete(ctx context.Context, params TodosIDTagsDeleteParams) (TodosIDTagsDeleteRes, error) {
	res, err := c.sendTodosIDTagsDelete(ctx, params)
	return res, err
}

func (c *Client) sendTodosIDTagsDelete(ctx context.Context, params TodosIDTagsDeleteParams) (res TodosIDTagsDeleteRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/todos/{id}/tags"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TodosIDTagsDeleteOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/todos/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/tags"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "tagId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "tagId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.IntToString(params.TagId))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTodosIDTagsDeleteResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TodosIDTagsPost invokes POST /todos/{id}/tags operation.
//
// Add a tag to a TODO.
//
// POST /todos/{id}/tags
func (c *Client) TodosIDTagsPost(ctx context.Context, request *TodosIDTagsPostReq, params TodosIDTagsPostParams) (TodosIDTagsPostRes, error) {
	res, err := c.sendTodosIDTagsPost(ctx, request, params)
	return res, err
}

func (c *Client) sendTodosIDTagsPost(ctx context.Context, request *TodosIDTagsPostReq, params TodosIDTagsPostParams) (res TodosIDTagsPostRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/todos/{id}/tags"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TodosIDTagsPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/todos/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/tags"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeTodosIDTagsPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTodosIDTagsPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TodosPost invokes POST /todos operation.
//
// Create a new TODO.
//
// POST /todos
func (c *Client) TodosPost(ctx context.Context, request *TodosPostReq) (TodosPostRes, error) {
	res, err := c.sendTodosPost(ctx, request)
	return res, err
}

func (c *Client) sendTodosPost(ctx context.Context, request *TodosPostReq) (res TodosPostRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/todos"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TodosPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/todos"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeTodosPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTodosPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UsersIDDelete invokes DELETE /users/{id} operation.
//
// Delete a user and all related data.
//
// DELETE /users/{id}
func (c *Client) UsersIDDelete(ctx context.Context, params UsersIDDeleteParams) (UsersIDDeleteRes, error) {
	res, err := c.sendUsersIDDelete(ctx, params)
	return res, err
}

func (c *Client) sendUsersIDDelete(ctx context.Context, params UsersIDDeleteParams) (res UsersIDDeleteRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/users/{id}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UsersIDDeleteOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/users/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeUsersIDDeleteResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UsersIDGet invokes GET /users/{id} operation.
//
// Get user details by ID.
//
// GET /users/{id}
func (c *Client) UsersIDGet(ctx context.Context, params UsersIDGetParams) (UsersIDGetRes, error) {
	res, err := c.sendUsersIDGet(ctx, params)
	return res, err
}

func (c *Client) sendUsersIDGet(ctx context.Context, params UsersIDGetParams) (res UsersIDGetRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/users/{id}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UsersIDGetOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/users/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeUsersIDGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UsersPost invokes POST /users operation.
//
// Register a new user.
//
// POST /users
func (c *Client) UsersPost(ctx context.Context, request *UsersPostReq) (UsersPostRes, error) {
	res, err := c.sendUsersPost(ctx, request)
	return res, err
}

func (c *Client) sendUsersPost(ctx context.Context, request *UsersPostReq) (res UsersPostRes, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/users"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UsersPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/users"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUsersPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeUsersPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
