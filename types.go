package timing

import (
	"time"
)

type response struct {
	StatusCode int
	Body       []byte
}

type Task struct {
	Self      string    `json:"self"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Duration  int       `json:"duration"`
	Project   Project   `json:"project"`
	Title     string    `json:"title"`
	Notes     string    `json:"notes"`
	IsRunning bool      `json:"is_running"`
}

type Project struct {
	Self              string   `json:"self"`
	Title             string   `json:"title"`
	TitleChain        []string `json:"title_chain"`
	Color             string   `json:"color"`
	ProductivityScore int      `json:"productivity_score"`
	IsArchived        bool     `json:"is_archived"`
	Parent            *Project `json:"parent"`
	Children          *Project `json:"children"`
}

type Links struct {
	First string      `json:"first"`
	Last  string      `json:"last"`
	Prev  interface{} `json:"prev"`
	Next  interface{} `json:"next"`
}
