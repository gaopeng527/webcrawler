package analyzer

import (
	"webcrawler/base"
	"net/http"
)

type ParseResponse func(httpResp *http.Response, respDepth uint32) ([]base.Data, []error)

// 分析器的接口类型
type Analyzer interface {
	Id() uint32 // 获得ID
	Analyze(respParsers []ParseResponse, resp base.Response) ([]base.Data, []error)  // 根据规则分析响应并返回请求和条目
}
