package spider

//go:generate mockgen -destination spider/mock/mock_spider.go -package spider github.com/zhiguogg/go-test/mock/version/spider Spider

type Spider interface {
	GetBody() string
}
