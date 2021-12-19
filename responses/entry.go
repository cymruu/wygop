package responses

type Entry struct {
	ID            int64          `json:"id"`
	Date          string         `json:"date"`
	Body          string         `json:"body"`
	Author        Author         `json:"author"`
	Blocked       bool           `json:"blocked"`
	Favorite      bool           `json:"favorite"`
	VoteCount     int64          `json:"vote_count"`
	Comments      []EntryComment `json:"comments"`
	CommentsCount int64          `json:"comments_count"`
	Status        string         `json:"status"`
	Embed         *EntryEmbed    `json:"embed,omitempty"`
	UserVote      int64          `json:"user_vote"`
	Survey        *Survey        `json:"survey,omitempty"`
	App           *string        `json:"app,omitempty"`
}

type EntryComment struct {
	ID           int64   `json:"id"`
	Author       Author  `json:"author"`
	Date         string  `json:"date"`
	Body         string  `json:"body"`
	Blocked      bool    `json:"blocked"`
	Favorite     bool    `json:"favorite"`
	VoteCount    int64   `json:"vote_count"`
	Status       string  `json:"status"`
	UserVote     int64   `json:"user_vote"`
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
	ID         int64   `json:"id"`
	Answer     string  `json:"answer"`
	Count      int64   `json:"count"`
	Percentage float64 `json:"percentage"`
}
