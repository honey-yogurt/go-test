package version

import (
	"github.com/golang/mock/gomock"
	spider "github.com/zhiguogg/go-test/mock/version/spider/mock"
	"testing"
)

func TestGetGoVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSpider := spider.NewMockSpider(ctrl)
	mockSpider.EXPECT().GetBody().Return("go1.18.1")
	goVer := GetGoVersion(mockSpider)

	if goVer != "go1.18.1" {
		t.Errorf("Get wrong version %s", goVer)
	}
}
