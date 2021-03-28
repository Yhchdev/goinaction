package matchers

import "sample/search"

type otherMatcher struct{}

func init() {
	var matcher otherMatcher
	search.Register("otherfault", matcher)
}

func (matcher otherMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	//接口的具体实现
	// 1.下载数据源
	// 2.正则匹配
	return nil, nil
}
