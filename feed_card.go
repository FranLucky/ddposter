package ddposter

type FeedCard struct{
	Links []Link `json:"links"`
}

func (f *FeedCard) Post(url string, at *At) error {
	return post(f, "feedCard", url, at)
}

