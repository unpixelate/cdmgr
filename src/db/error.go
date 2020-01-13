package db

type DbError struct {
	Err string
	Tag string
}

func (e *DbError) Error() string {
	return "Error: " + e.Err + " with tag: " + e.Tag
}

func RaiseErr(text string, tag string) error {
	return &DbError{text, tag}
}

// Error: cdmgr add -t tagMasterData
