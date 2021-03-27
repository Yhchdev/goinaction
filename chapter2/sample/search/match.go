package search

import "log"

// 匹配结果
type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan *Result) {
	//对于不同的匹配器，使用接口让代码更通用
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) {

	//会一直阻塞直到通道中有数据进来
	//通道被关闭后，退出for 循环
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
