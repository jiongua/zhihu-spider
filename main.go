package main

import (
	"fmt"
	"log"
	"sync"
	"zhihu-spider/crawl"
)

func main() {
	topics := map[string]int{
		"科技": 19556664,
		//"互联网": 19550517,
		//"体育": 19554827,
	}
	//questionChan := make(chan crawl.QuestionFetch)
	answerChan := make(chan []crawl.AnswerFetch)
	commentChan := make(chan []crawl.CommentFetch)
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			for fetches := range answerChan {
				for _, fetch := range fetches {
					fetch.Fetch()
				}
			}
			log.Println("没有新问题了")
		}()
	}


	for i := 0; i < 4; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for fetches := range commentChan {
				for _, fetch := range fetches {
					fetch.Fetch()
				}
			}
			log.Println("没有新回答了")
		}()
	}

	go func() {
		for name, id := range topics {
			qf := crawl.QuestionFetch{
				TopicID:   id,
				TopicName: name,
				Questions: nil,
				NextURL:   fmt.Sprintf(crawl.TopicURL, id, 10, 0),
				AnswerCh:  answerChan,
				CommentCh: commentChan,
			}
			qf.Fetch()
		}
		close(answerChan)
		log.Println("没有新话题了")
	}()

	var done = make(chan struct{})
	go func() {
		wg1.Wait()
		close(commentChan)
		wg2.Wait()
		close(done)
	}()

	<-done

}
