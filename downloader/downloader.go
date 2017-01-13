package downloader

import (
	"webcrawler/base"
	"net/http"
	"webcrawler/middleware"
)

// ID生成器
var downloaderIdGenerator middleware.IdGenerator = middleware.NewIdGenerator()

// 网页下载器的接口类型
type PageDownloader interface {
	Id() uint32 // 获得ID
	Download(req base.Request) (*base.Response, error) // 根据请求下载网页并返回响应
}

// 网页下载器的实现类型
type myPageDownloader struct {
	httpClient http.Client // HTTP客户端
	id uint32  // ID
}

// 生成并返回ID
func genDownloaderId() uint32 {
	return downloaderIdGenerator.GetUint32()
}

// 创建网页下载器
func NewPageDownloader(client *http.Client) PageDownloader {
	id := genDownloaderId()
	if client == nil {
		client = &http.Client{}
	}
	return &myPageDownloader{
		id: id,
		httpClient: *client,
	}
}

func (dl *myPageDownloader) Id() uint32 {
	return dl.id
}

func (dl *myPageDownloader) Download(req base.Request) (*base.Response, error) {
	httpReq := req.HttpReq()
	httpResp, err := dl.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return base.NewResponse(httpResp, req.Depth()), nil
}