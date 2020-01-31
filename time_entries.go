package timing

import (
	"net/url"
	"strconv"
	"time"
)

const timeEntriesURL = baseURL + "api/v1/time-entries"

const (
	kStartDateMin         = "start_date_min"
	kStartDateMax         = "start_date_max"
	kProjects             = "projects[]"
	kSearchQuery          = "search_query"
	kIsRunning            = "is_running"
	kIncludeProjectData   = "include_project_data"
	kIncludeChildProjects = "include_child_projects"

	IsRunningDefault = 0
	IsRunningTrue    = 1
	IsRunningFalse   = 2
)

type TimeEntriesRequest struct {
	StartDateMin         time.Time
	StartDateMax         time.Time
	Projects             []string
	SearchQuery          string
	IsRunning            int // IsRunningDefault, IsRunningDefault, IsRunningFalse
	IncludeProjectData   bool
	IncludeChildProjects bool
}

// easyjson:json
type TimeEntriesResponse struct {
	Data  []Task `json:"data"`
	Links Links  `json:"links"`
	Meta  struct {
		CurrentPage int    `json:"current_page"`
		From        int    `json:"from"`
		LastPage    int    `json:"last_page"`
		Path        string `json:"path"`
		PerPage     int    `json:"per_page"`
		To          int    `json:"to"`
		Total       int    `json:"total"`
	} `json:"meta"`
}

func (r *TimeEntriesResponse) parse(response *response) {
	if response == nil {
		return
	}

	_ = r.UnmarshalJSON(response.Body)
}

func TimeEntries(r *TimeEntriesRequest) (*TimeEntriesResponse, error) {
	v := url.Values{}

	if !r.StartDateMin.IsZero() {
		v.Set(kStartDateMin, r.StartDateMin.String())
	}
	if !r.StartDateMax.IsZero() {
		v.Set(kStartDateMax, r.StartDateMax.String())
	}
	v[kProjects] = r.Projects
	if len(r.SearchQuery) != 0 {
		v.Set(kSearchQuery, r.SearchQuery)
	}
	switch r.IsRunning {
	case IsRunningTrue:
		v.Set(kIsRunning, "true")
	case IsRunningFalse:
		v.Set(kIsRunning, "false")
	}

	v.Set(kIncludeProjectData, strconv.FormatBool(r.IncludeProjectData))
	v.Set(kIncludeChildProjects, strconv.FormatBool(r.IncludeChildProjects))

	resp, err := request("GET", timeEntriesURL+"?"+v.Encode(), "")

	if err != nil {
		return nil, err
	}

	res := new(TimeEntriesResponse)
	res.parse(resp)

	return res, nil
}
