package handler

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/compute/internal/compute-api/config"
	"github.com/syunkitada/stadyapp/backends/compute/internal/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

const HeaderXForwardedFor = "x-forwarded-for"

func appendHostToXForwardHeader(header http.Header, host string) {
	forwardedFor := header.Get(HeaderXForwardedFor)
	if forwardedFor != "" {
		header.Set(HeaderXForwardedFor, forwardedFor+","+host)
	} else {
		header.Set(HeaderXForwardedFor, host)
	}
}

func copyHeader(src http.Header, dst http.Header) {
	dst.Set("content-type", src.Get("content-type"))
}

func proxy(ectx echo.Context, proxy config.Proxy) error {
	ctx := iam_auth.WithEchoContext(ectx)
	authContext, err := iam_auth.GetAuthContext(ctx)
	if err != nil {
		return tlog.BindEchoError(ctx, ectx, err)
	}

	srcReq := ectx.Request()

	// Copy the request body
	reqBody := []byte{}
	if srcReq.Body != nil {
		reqBody, _ = io.ReadAll(srcReq.Body)
	}
	srcReq.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	// Copy the request without headers
	target, err := url.Parse(proxy.URL)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, "Invalid target URL")
	}
	target.Path = strings.Replace(srcReq.URL.Path, proxy.OldBasePath, proxy.NewBasePath, 1)
	target.RawQuery = ectx.QueryString()

	// Create a new request
	req, err := http.NewRequest(srcReq.Method, target.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return ectx.String(http.StatusInternalServerError, "Failed to create request")
	}

	if clientIP, _, err := net.SplitHostPort(ectx.Request().RemoteAddr); err == nil {
		appendHostToXForwardHeader(req.Header, clientIP)
	}

	req.Header.Set("x-auth-token", srcReq.Header.Get("x-auth-token"))
	req.Header.Set("x-service-token", srcReq.Header.Get("x-service-token"))

	iam_auth.AddAuthHeader(req, authContext)
	copyHeader(srcReq.Header, req.Header)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, "Failed to send request to target server")
	}
	defer resp.Body.Close()

	// Copy response header
	for key, values := range resp.Header {
		for _, value := range values {
			ectx.Response().Header().Add(key, value)
		}
	}

	// Copy status code and response body
	ectx.Response().WriteHeader(resp.StatusCode)
	_, err = io.Copy(ectx.Response(), resp.Body)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, "Failed to copy response body")
	}

	return nil
}

func (self *Handler) GetGlanceImages(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Glance)
}

func (self *Handler) GetGlanceImageByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Glance)
}

func (self *Handler) CreateGlanceImage(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Glance)
}

func (self *Handler) UpdateGlanceImageByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Glance)
}

func (self *Handler) DeleteGlanceImageByID(ectx echo.Context, id string) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Glance)
}

func (self *Handler) GetGlanceSchemasImage(ectx echo.Context) error {
	return proxy(ectx, self.conf.Compute.ProxyCatalog.Glance)
}
