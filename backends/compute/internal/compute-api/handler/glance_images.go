package handler

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/compute/internal/libs/iam_auth"
	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

// Hop-by-hop headers. These are removed when sent to the backend.
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html
var hopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te", // canonicalized version of "TE"
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func delHopHeaders(header http.Header) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
}

func appendHostToXForwardHeader(header http.Header, host string) {
	// If we aren't the first proxy retain prior
	// X-Forwarded-For information as a comma+space
	// separated list and fold multiple headers into one.
	if prior, ok := header["X-Forwarded-For"]; ok {
		host = strings.Join(prior, ", ") + ", " + host
	}
	header.Set("X-Forwarded-For", host)
}

func proxy(ectx echo.Context) error {
	targetURL := "http://localhost:9292" // プロキシ先のサーバーURL

	// リクエストのパスを追加
	target, err := url.Parse(targetURL)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, "Invalid target URL")
	}
	target.Path = ectx.Request().URL.Path
	target.RawQuery = ectx.QueryString()

	// 新しいリクエストを作成
	req, err := http.NewRequest(ectx.Request().Method, "http://localhost:9292/v2/images", ectx.Request().Body)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, "Failed to create request")
	}

	// 元のリクエストヘッダーをコピー
	// for key, values := range ectx.Request().Header {
	// 	for _, value := range values {
	// 		req.Header.Add(key, value)
	// 	}
	// }

	// curl http://localhost:9292/v2/images -H "X-Identity-Status: Confirmed" -H "x-user-domain-id: default" -H "x-project-domain-id: default" -H "x-user-id: admin" -H "x-project-id: admin" -H "x-roles: admin" -H "X-IS-ADMIN-PROJECT: true"
	req.Header.Add("x-identity-status", "Confirmed")
	req.Header.Add("x-user-domain-id", "default")
	req.Header.Add("x-project-domain-id", "default")
	req.Header.Add("x-user-id", "admin")
	req.Header.Add("x-project-id", "admin")
	req.Header.Add("x-roles", "admin")
	req.Header.Add("x-is-admin-project", ectx.Request().Host)

	// HTTPクライアントでリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, "Failed to send request to target server")
	}
	defer resp.Body.Close()

	// レスポンスヘッダーをコピー
	for key, values := range resp.Header {
		for _, value := range values {
			ectx.Response().Header().Add(key, value)
		}
	}

	// ステータスコードとボディを返す
	ectx.Response().WriteHeader(resp.StatusCode)
	_, err = io.Copy(ectx.Response(), resp.Body)
	if err != nil {
		return ectx.String(http.StatusInternalServerError, "Failed to copy response body")
	}

	return nil
}

func (self *Handler) GetGlanceImages(ectx echo.Context) error {
	return proxy(ectx)
	// ctx := iam_auth.WithEchoContext(ectx)
	// resp := map[string]string{}
}

func (self *Handler) GetGlanceImageByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)
	resp := map[string]string{}
	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) CreateGlanceImage(ectx echo.Context) error {
	ctx := iam_auth.WithEchoContext(ectx)
	resp := map[string]string{}
	return tlog.BindEchoOK(ctx, ectx, resp)
}

func (self *Handler) UpdateGlanceImageByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)
	return tlog.BindEchoNoContent(ctx, ectx)
}

func (self *Handler) DeleteGlanceImageByID(ectx echo.Context, id string) error {
	ctx := iam_auth.WithEchoContext(ectx)
	return tlog.BindEchoNoContent(ctx, ectx)
}
