package leetcode1348

import (
	"reflect"
	"testing"
)

func TestLeetChallenge(t *testing.T) {
	t.Run("First test", func(t *testing.T) {
		// "TweetCounts", [], null
		obj := Constructor()

		// "recordTweet", ["tweet3",0], null
		obj.RecordTweet("tweet3", 0)
		// "recordTweet", ["tweet3",60], null
		obj.RecordTweet("tweet3", 60)
		// "recordTweet", ["tweet3",10], null
		obj.RecordTweet("tweet3", 10)

		// "getTweetCountsPerFrequency", ["minute","tweet3",0,59], [2]
		want := []int{2}
		got := obj.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %+v but got %+v", want, got)
		}

		// "getTweetCountsPerFrequency", ["minute","tweet3",0,60], [2,1]
		want = []int{2, 1}
		got = obj.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %+v but got %+v", want, got)
		}

		// "recordTweet", ["tweet3",120], null
		obj.RecordTweet("tweet3", 120)

		// "getTweetCountsPerFrequency", ["hour","tweet3",0,210], [4]
		want = []int{4}
		got = obj.GetTweetCountsPerFrequency("hour", "tweet3", 0, 210)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %+v but got %+v", want, got)
		}
	})

	t.Run("test2", func(t *testing.T) {
		// "TweetCounts", [], null
		obj := Constructor()

		// "recordTweet", ["tweet0", 99], null
		obj.RecordTweet("tweet0", 99)
		// "recordTweet", ["tweet1", 80], null
		obj.RecordTweet("tweet1", 80)
		// "recordTweet", ["tweet1", 80], null
		obj.RecordTweet("tweet2", 29)
		// "recordTweet", ["tweet1", 80], null
		obj.RecordTweet("tweet3", 81)
		// "recordTweet", ["tweet1", 80], null
		obj.RecordTweet("tweet4", 56)

		// "getTweetCountsPerFrequency", ["day","tweet4",56,2667], [1]
		want := []int{1}
		got := obj.GetTweetCountsPerFrequency("day", "tweet4", 56, 2667)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %+v but got %+v", want, got)
		}
	})
}
