package downloader

// 网页下载器池的接口类型
type PageDownloaderPool interface {
	Take() (PageDownloader, error) // 从池中取出一个网页下载器
	Return(dl PageDownloader) error // 把一个网页下载器归还给池
	Total() uint32 // 获得池的总容量
	Used() uint32 // 获得正在被使用的网页下载器的数量
}
