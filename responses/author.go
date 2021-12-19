package responses

type Author struct {
	Login  string `json:"login"`
	Color  int64  `json:"color"`
	Sex    *Sex   `json:"sex,omitempty"`
	Avatar string `json:"avatar"`
}

type Sex string

const (
	Female Sex = "female"
	Male   Sex = "male"
)
