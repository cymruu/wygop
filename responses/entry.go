package responses

type Entry struct {
	ID            uint64         `json:"id"`
	Date          APITime        `json:"date"`
	Body          string         `json:"body"`
	Author        Author         `json:"author"`
	Blocked       bool           `json:"blocked"`
	Favorite      bool           `json:"favorite"`
	VoteCount     uint32         `json:"vote_count"`
	Comments      []EntryComment `json:"comments"`
	CommentsCount uint32         `json:"comments_count"`
	Status        string         `json:"status"`
	Embed         *EntryEmbed    `json:"embed,omitempty"`
	UserVote      uint8          `json:"user_vote"`
	Survey        *Survey        `json:"survey,omitempty"`
	App           *string        `json:"app,omitempty"`
}

type EntryComment struct {
	ID           uint64  `json:"id"`
	Author       Author  `json:"author"`
	Date         APITime `json:"date"`
	Body         string  `json:"body"`
	Blocked      bool    `json:"blocked"`
	Favorite     bool    `json:"favorite"`
	VoteCount    uint32  `json:"vote_count"`
	Status       string  `json:"status"`
	UserVote     uint8   `json:"user_vote"`
	ViolationURL string  `json:"violation_url"`
	App          *string `json:"app,omitempty"`
}

type EntryEmbed struct {
	Type     string  `json:"type"`
	URL      string  `json:"url"`
	Source   string  `json:"source"`
	Preview  string  `json:"preview"`
	Plus18   bool    `json:"plus18"`
	Size     string  `json:"size"`
	Animated bool    `json:"animated"`
	Ratio    float64 `json:"ratio"`
}

type Survey struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	ID         uint64  `json:"id"`
	Answer     string  `json:"answer"`
	Count      uint32  `json:"count"`
	Percentage float64 `json:"percentage"`
}
