package helpers

import (
	"math"
	"fmt"
	"simpleapi/models"

	"github.com/jinzhu/gorm"
)

type Filter struct {
	Key   string
	Value string
}

func GetPagination(db *gorm.DB, limit int, page int) (map[string]interface{}, error) {
	count := 0	
	// get total data
	err := db.Model(&models.Message{}).Count(&count).Error

	if err != nil {
		return nil, fmt.Errorf("error get total message, err := %s", err.Error())
	}

	// calculate total page
	totalPage := math.Ceil(float64(count) / float64(limit))
	hasNext := false

	// set has Next
	if float64(page) < totalPage {
		hasNext = true
	}

	pagination := map[string]interface{}{
		"page":        page,
		"total_pages": totalPage,
		"total_items": count,
		"per_page": limit,
		"has_next": hasNext,
		"has_previous": page > 1,
	}

	return pagination, nil
}
