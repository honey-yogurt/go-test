package version

import "github.com/zhiguogg/go-test/mock/version/spider"

func GetGoVersion(s spider.Spider) string {
	// 此时没有spider的实例，怎么测试呢
	body := s.GetBody()
	return body
}
