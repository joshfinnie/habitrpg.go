package habitrpg

type Task struct {
	Attribute   string    `json:"attribute"`
	Challenge   Challenge `json:"challenge"`
	DateCreated string    `json:"dateCreated"`
	Down        bool      `json:"down"`
	History     []struct {
		Date  float64 `json:"date"`
		Value float64 `json:"value"`
	} `json:"history"`
	ID       string      `json:"id"`
	Notes    string      `json:"notes"`
	Priority float64     `json:"priority"`
	Tags     interface{} `json:"tags"`
	Text     string      `json:"text"`
	Type     string      `json:"type"`
	Up       bool        `json:"up"`
	Value    float64     `json:"value"`
}

// dateCreatedTime is a convenience wrapper that returns the dateCreated time,
// parsed as a time.Time struct
func (t Task) dateCreatedTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.DateCreated)
}
