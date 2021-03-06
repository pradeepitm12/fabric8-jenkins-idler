// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "auth": search Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-auth/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-auth
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// UsersSearchPath computes a request path to the users action of search.
func UsersSearchPath() string {

	return fmt.Sprintf("/api/search/users")
}

// Search by fullname
func (c *Client) UsersSearch(ctx context.Context, path string, q string, pageLimit *int, pageOffset *string) (*http.Response, error) {
	req, err := c.NewUsersSearchRequest(ctx, path, q, pageLimit, pageOffset)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUsersSearchRequest create the request corresponding to the users action endpoint of the search resource.
func (c *Client) NewUsersSearchRequest(ctx context.Context, path string, q string, pageLimit *int, pageOffset *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("q", q)
	if pageLimit != nil {
		tmp56 := strconv.Itoa(*pageLimit)
		values.Set("page[limit]", tmp56)
	}
	if pageOffset != nil {
		values.Set("page[offset]", *pageOffset)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
