package httpc

import (
	"article-comment/api/internal/common/errorx"
	"article-comment/api/internal/common/jsonc"
	"fmt"
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/avast/retry-go/v4"
)

func Stringify(v interface{}) string {
	b, err := jsonc.Marshal(v)

	if err != nil {
		return "{}"
	}
	return string(b)
}

type FetchConf struct {
	Method      string
	URL         string
	Data        any
	Headers     map[string]string
	AccessToken string
}

const maxAttempt = 3

// Fetch 发起http请求, json data 和 query string
func Fetch(c *FetchConf) (contents []byte, err error) {
	err = retry.Do(func() error {
		body := ""
		if c.Method == http.MethodPost {
			body = Stringify(c.Data)
		}

		request, err := http.NewRequest(c.Method, c.URL, strings.NewReader(body))
		if err != nil {
			return errorx.BadRequest("创建HTTP请求异常").
				WithMetadata(errorx.Metadata{"req": c}).
				WithError(err)
		}
		if c.Method == http.MethodPost {
			request.Header.Add("Content-Type", "application/json;charset=utf-8")
		}

		for k, v := range c.Headers {
			request.Header.Add(k, v)
		}
		resp, err := http.DefaultClient.Do(request)

		if err != nil {
			httpCode := http.StatusInternalServerError
			// 检查 resp 是否为 nil，确保不会在 defer 函数中发生 panic
			if resp != nil {
				httpCode = resp.StatusCode
			}
			return errorx.New("http-Client", httpCode, "发起HTTP请求异常").
				WithMetadata(errorx.Metadata{"req": c}).
				WithError(err)
		}

		defer func() {
			if err = resp.Body.Close(); err != nil {
				zap.L().Error("Close HTTP 响应异常", zap.Error(err), zap.Any("入参", c))
				return
			}
		}()

		if contents, err = io.ReadAll(resp.Body); err != nil {
			return errorx.Internal(err, "读取HTTP响应异常").WithMetadata(errorx.Metadata{"req": c})
		}

		return nil
	},
		retry.Attempts(maxAttempt), // 重试3次
		retry.OnRetry(func(n uint, err error) {
			zap.L().Error(fmt.Sprintf("HTTP请求失败, 当前重试次数: %d", n),
				zap.Error(err), zap.Any("入参", c))

		}),
	)

	return
}

// FormFetch 表单提交, form data
func FormFetch(c *FetchConf) (contents []byte, err error) {
	err = retry.Do(func() error {
		strData, ok := c.Data.(string)
		if !ok {
			return errorx.BadRequest("请求参数类型异常").
				WithMetadata(errorx.Metadata{"data:": c.Data}).
				WithError(err)
		}
		request, err := http.NewRequest(c.Method, c.URL, strings.NewReader(strData))
		if err != nil {
			return errorx.BadRequest("创建HTTP请求异常").
				WithMetadata(errorx.Metadata{"req": c}).
				WithError(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		for k, v := range c.Headers {
			request.Header.Add(k, v)
		}

		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			return errorx.New("http-Client", resp.StatusCode, "发起HTTP请求异常").
				WithMetadata(errorx.Metadata{"req": c}).
				WithError(err)
		}

		defer func() {
			if err = resp.Body.Close(); err != nil {
				zap.L().Error("Close HTTP 响应异常", zap.Error(err), zap.Any("入参", c))
				return

			}
		}()

		if contents, err = io.ReadAll(resp.Body); err != nil {
			return errorx.Internal(err, "读取HTTP响应异常").WithMetadata(errorx.Metadata{"req": c})
		}

		return nil
	},
		retry.Attempts(maxAttempt), // 重试3次
		retry.OnRetry(func(n uint, err error) {
			zap.L().Error(fmt.Sprintf("HTTP请求失败, 当前重试次数: %d", n),
				zap.Error(err), zap.Any("入参", c))

		}),
	)

	return
}
