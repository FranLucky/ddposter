package ddposter

type FeedCard struct{
	Links []Link `json:"links"`
}

func (f FeedCard) TypeString() string {
	return "feedCard"
}

