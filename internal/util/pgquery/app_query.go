package pquery

type AppQuery struct {
	Page       int64
	Limit      int64
	CategoryID string
	Status     string
}

// AssignPage ...
func (q *AppQuery) AssignPage() {
	if q.Page > 0 {

	}
}
