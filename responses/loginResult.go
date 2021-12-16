package responses

type LoginResult struct {
	Profile Profile `json:"profile"`
	UserKey string  `json:"userkey"`
}

type Profile struct {
	Login               string `json:"login"`
	Color               int32  `json:"color"`
	Sex                 Sex    `json:"sex"`
	Avatar              string `json:"avatar"`
	SignupAt            string `json:"signup_at"`
	Background          string `json:"background"`
	About               string `json:"about"`
	LinksAddedCount     uint32 `json:"links_added_count"`
	LinksPublishedCount uint32 `json:"links_published_count"`
	CommentsCount       uint32 `json:"comments_count"`
	Rank                uint32 `json:"rank"`
	Followers           uint32 `json:"followers"`
}
