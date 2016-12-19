package downloader

import "webcrawler/base"

// 网页下载器的接口类型
type PageDownloader interface {
	Id() uint32 // 获得ID
	Download(req base.Request) (*base.Response, error) // 根据请求下载网页并返回响应
}
