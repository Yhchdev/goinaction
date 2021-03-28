package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTrem string) {
	//1.获取数据源
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	wg := sync.WaitGroup{}
	wg.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			// 面向接口编程
			Match(matcher, feed, searchTrem, results)
			wg.Done()
		}(matcher, feed)

	}

	// 监听查询是否结束
	go func() {
		wg.Wait()
		close(results)
	}()

	//显示查询结果
	Display(results)
}

// 注册匹配器
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "Matcher")
	matchers[feedType] = matcher
}
