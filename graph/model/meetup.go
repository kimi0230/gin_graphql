package model

import (
	"fmt"
	"gin_graphql/app/models"
)

func GetMeetups(filter *MeetupFilter, limit, offset *int) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	query := models.DB.Model(&meetups).Order("id DESC")
	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}

	if limit != nil {
		query = query.Limit(*limit)
	}

	if offset != nil {
		query = query.Offset(*offset)
	}

	if err := query.Find(&meetups).Error; err != nil {
		return nil, err
	}

	return meetups, nil
}
