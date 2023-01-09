package query

import (
	"github.com/Kapakac/crud-kafka/internal/pkg/database"
)

type Finder struct {
	Arr          interface{}
	Start, Count int
}

func (f Finder) All() interface{} {
	database.DB.Limit(f.Count).Offset(f.Start).Find(&f.Arr)

	return f.Arr
}
