package paginator

import (
	"gorm.io/gorm"
	"math"
)

type Paginator struct {
	Items       any
	Total       int
	PerPage     int
	CurrentPage int
	LastPage    int
}

func (p *Paginator) WithItems(items any) *Paginator {
	p.Items = items
	return p
}

func New(db *gorm.DB, perPage int, page int) *Paginator {
	var count int64
	paginator := Paginator{
		Total:       int(count),
		PerPage:     perPage,
		CurrentPage: page + 1,
	}

	db.Count(&count)
	if perPage > 0 && page > 0 {
		totalPages := int(math.Ceil(float64(count) / float64(perPage)))
		if page < 1 {
			page = 1
		} else if page > totalPages {
			page = totalPages
		}
		offset := (page - 1) * perPage
		db.Offset(offset).Limit(perPage)
		paginator.LastPage = totalPages
	}

	return &paginator
}
