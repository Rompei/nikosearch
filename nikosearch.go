package nikosearch

type Response struct {
	Meta Meta   `json:"meta"`
	Data []Data `json:"data"`
}

type Meta struct {
	Status     int    `json:"status"`
	ID         string `json:"id"`
	TotalCount int    `json:"totalCount"`
}

type Data struct {
	ContentID              string      `json:"contentId"`
	Title                  string      `json:"title"`
	Description            string      `json:"description"`
	Tags                   string      `json:"tags"`
	CategoryTags           string      `json:"categoryTags"`
	ViewCounter            int         `json:"viewCounter"`
	MyListCounter          int         `json:"myListCounter"`
	CommentCounter         int         `json:"commentCounter"`
	StartTime              string      `json:"startTime"`
	ThumbnailURL           string      `json:"thumbnailUrl"`
	CommunityIcon          string      `json:"communityIcon"`
	ScoreTimeshiftReserved int         `json:"scoreTimeshiftReserved"`
	LiveStatus             interface{} `json:"liveStatus"`
}
