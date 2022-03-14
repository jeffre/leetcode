package leetcode1348

type TweetCounts struct {
	tweet []TweetCount
}
type TweetCount struct {
	tweetName string
	time      int
}

func Constructor() TweetCounts {
	return TweetCounts{}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	new := TweetCount{
		tweetName: tweetName,
		time:      time,
	}
	this.tweet = append(this.tweet, new)
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	result := []int{}

	var increment int

	switch freq {
	case "minute":
		increment = 60
	case "hour":
		increment = 3600
	case "day":
		increment = 86400
	}

	for startTime <= endTime {

		// End of period cannot exceed endTime
		endPeriod := startTime + increment - 1
		if endPeriod > endTime {
			endPeriod = endTime
		}

		//fmt.Println("startTime:", startTime, "endPeriod:", endPeriod)

		timeSlotOccurrences := 0

		for _, tweet := range this.tweet {
			if startTime <= tweet.time && tweet.time <= endPeriod && tweetName == tweet.tweetName {
				timeSlotOccurrences++
			}
		}
		result = append(result, timeSlotOccurrences)

		startTime += increment
	}

	return result
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */
