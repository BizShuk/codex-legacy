package hide

type Twitter struct {
	following map[int]map[int]bool
	// recent10 map[int]*tweet
	tweets map[int]*tweet
}

type tweet struct {
	id   int
	time int
	prev *tweet
	next *tweet
}

func Constructor() Twitter {
	return Twitter{
		following: map[int]map[int]bool{},
		// recent10:map[int]*tweet{},
		tweets: map[int]*tweet{},
	}
}

var count int

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if this.tweets[userId] == nil {
		this.tweets[userId] = &tweet{}
	}
	newTweet := &tweet{id: tweetId, time: count}
	count += 1
	head := this.tweets[userId]
	next := head.next

	head.next, newTweet.prev = newTweet, head
	if next != nil {
		newTweet.next, next.prev = next, newTweet
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	followeeTweets := make([]*tweet, 0)

	for followeeId := range this.following[userId] {
		followeeTweets = append(followeeTweets, this.tweets[followeeId])
	}

	result := make([]int, 0)

	for i := 0; i < 10; i++ {
		latest := 0
		latestIdx := 0
		var latestTweet *tweet

		for i, tweet := range followeeTweets {
			if latest < tweet.time {
				latest = tweet.time
				latestTweet = tweet
				latestIdx = i
			}
		}
		followeeTweets[latestIdx] = followeeTweets[latestIdx].next

		result = append(result, latestTweet.id)
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.following[followerId] == nil {
		this.following[followerId] = map[int]bool{}
	}

	this.following[followerId][followeeId] = true

}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.following[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
