// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package indexing provides access to the Indexing API.
//
// For product documentation, see: https://developers.google.com/search/apis/indexing-api/
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/indexing/v3"
//   ...
//   ctx := context.Background()
//   indexingService, err := indexing.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   indexingService, err := indexing.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   indexingService, err := indexing.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package indexing // import "google.golang.org/api/indexing/v3"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "indexing:v3"
const apiName = "indexing"
const apiVersion = "v3"
const basePath = "https://indexing.googleapis.com/"
const mtlsBasePath = "https://indexing.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Submit data to Google for indexing
	IndexingScope = "https://www.googleapis.com/auth/indexing"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/indexing",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.UrlNotifications = NewUrlNotificationsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	UrlNotifications *UrlNotificationsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewUrlNotificationsService(s *Service) *UrlNotificationsService {
	rs := &UrlNotificationsService{s: s}
	return rs
}

type UrlNotificationsService struct {
	s *Service
}

// PublishUrlNotificationResponse: Output for PublishUrlNotification
type PublishUrlNotificationResponse struct {
	// UrlNotificationMetadata: Description of the notification events
	// received for this URL.
	UrlNotificationMetadata *UrlNotificationMetadata `json:"urlNotificationMetadata,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "UrlNotificationMetadata") to unconditionally include in API
	// requests. By default, fields with empty or default values are omitted
	// from API requests. However, any non-pointer, non-interface field
	// appearing in ForceSendFields will be sent to the server regardless of
	// whether the field is empty or not. This may be used to include empty
	// fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "UrlNotificationMetadata")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PublishUrlNotificationResponse) MarshalJSON() ([]byte, error) {
	type NoMethod PublishUrlNotificationResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UrlNotification: `UrlNotification` is the resource used in all
// Indexing API calls. It describes one event in the life cycle of a Web
// Document.
type UrlNotification struct {
	// NotifyTime: Creation timestamp for this notification. Users should
	// _not_ specify it, the field is ignored at the request time.
	NotifyTime string `json:"notifyTime,omitempty"`

	// Type: The URL life cycle event that Google is being notified about.
	//
	// Possible values:
	//   "URL_NOTIFICATION_TYPE_UNSPECIFIED" - Unspecified.
	//   "URL_UPDATED" - The given URL (Web document) has been updated.
	//   "URL_DELETED" - The given URL (Web document) has been deleted.
	Type string `json:"type,omitempty"`

	// Url: The object of this notification. The URL must be owned by the
	// publisher of this notification and, in case of `URL_UPDATED`
	// notifications, it _must_ be crawlable by Google.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NotifyTime") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NotifyTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UrlNotification) MarshalJSON() ([]byte, error) {
	type NoMethod UrlNotification
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UrlNotificationMetadata: Summary of the most recent Indexing API
// notifications successfully received, for a given URL.
type UrlNotificationMetadata struct {
	// LatestRemove: Latest notification received with type `URL_REMOVED`.
	LatestRemove *UrlNotification `json:"latestRemove,omitempty"`

	// LatestUpdate: Latest notification received with type `URL_UPDATED`.
	LatestUpdate *UrlNotification `json:"latestUpdate,omitempty"`

	// Url: URL to which this metadata refers.
	Url string `json:"url,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "LatestRemove") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LatestRemove") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UrlNotificationMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod UrlNotificationMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "indexing.urlNotifications.getMetadata":

type UrlNotificationsGetMetadataCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetMetadata: Gets metadata about a Web Document. This method can
// _only_ be used to query URLs that were previously seen in successful
// Indexing API notifications. Includes the latest `UrlNotification`
// received via this API.
func (r *UrlNotificationsService) GetMetadata() *UrlNotificationsGetMetadataCall {
	c := &UrlNotificationsGetMetadataCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Url sets the optional parameter "url": URL that is being queried.
func (c *UrlNotificationsGetMetadataCall) Url(url string) *UrlNotificationsGetMetadataCall {
	c.urlParams_.Set("url", url)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlNotificationsGetMetadataCall) Fields(s ...googleapi.Field) *UrlNotificationsGetMetadataCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *UrlNotificationsGetMetadataCall) IfNoneMatch(entityTag string) *UrlNotificationsGetMetadataCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UrlNotificationsGetMetadataCall) Context(ctx context.Context) *UrlNotificationsGetMetadataCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UrlNotificationsGetMetadataCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UrlNotificationsGetMetadataCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20211119")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/urlNotifications/metadata")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "indexing.urlNotifications.getMetadata" call.
// Exactly one of *UrlNotificationMetadata or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *UrlNotificationMetadata.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *UrlNotificationsGetMetadataCall) Do(opts ...googleapi.CallOption) (*UrlNotificationMetadata, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &UrlNotificationMetadata{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets metadata about a Web Document. This method can _only_ be used to query URLs that were previously seen in successful Indexing API notifications. Includes the latest `UrlNotification` received via this API.",
	//   "flatPath": "v3/urlNotifications/metadata",
	//   "httpMethod": "GET",
	//   "id": "indexing.urlNotifications.getMetadata",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "url": {
	//       "description": "URL that is being queried.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/urlNotifications/metadata",
	//   "response": {
	//     "$ref": "UrlNotificationMetadata"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/indexing"
	//   ]
	// }

}

// method id "indexing.urlNotifications.publish":

type UrlNotificationsPublishCall struct {
	s               *Service
	urlnotification *UrlNotification
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// Publish: Notifies that a URL has been updated or deleted.
func (r *UrlNotificationsService) Publish(urlnotification *UrlNotification) *UrlNotificationsPublishCall {
	c := &UrlNotificationsPublishCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlnotification = urlnotification
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlNotificationsPublishCall) Fields(s ...googleapi.Field) *UrlNotificationsPublishCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UrlNotificationsPublishCall) Context(ctx context.Context) *UrlNotificationsPublishCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UrlNotificationsPublishCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UrlNotificationsPublishCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20211119")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.urlnotification)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/urlNotifications:publish")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "indexing.urlNotifications.publish" call.
// Exactly one of *PublishUrlNotificationResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *PublishUrlNotificationResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *UrlNotificationsPublishCall) Do(opts ...googleapi.CallOption) (*PublishUrlNotificationResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PublishUrlNotificationResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Notifies that a URL has been updated or deleted.",
	//   "flatPath": "v3/urlNotifications:publish",
	//   "httpMethod": "POST",
	//   "id": "indexing.urlNotifications.publish",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v3/urlNotifications:publish",
	//   "request": {
	//     "$ref": "UrlNotification"
	//   },
	//   "response": {
	//     "$ref": "PublishUrlNotificationResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/indexing"
	//   ]
	// }

}
