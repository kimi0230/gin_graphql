package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"gin_graphql/app/models"
	"gin_graphql/graph/generated"
	"gin_graphql/graph/model"
	"time"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	// panic(fmt.Errorf("not implemented"))
	// meetup 關聯 user
	user := new(models.User)

	for _, u := range users {
		if u.ID == obj.UserID {
			user = u
			break
		}
	}
	if user == nil {
		return nil, errors.New("user wit id not exist")
	}
	return user, nil
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
	var meetups models.Meetup
	return meetups.GetMeetup()
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var m []*models.Meetup
	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			m = append(m, meetup)
		}
	}
	return m, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

var meetups = []*models.Meetup{
	{
		BaseModel:   models.BaseModel{ID: 1},
		Name:        "a meetup",
		Description: "one",
		UserID:      1,
	},
	{
		BaseModel:   models.BaseModel{ID: 2},
		Name:        "second meetup",
		Description: "two",
		UserID:      2,
	},
}
var users = []*models.User{
	{
		BaseModel: models.BaseModel{ID: 1},
		Username:  "kimi",
	},
	{
		BaseModel: models.BaseModel{ID: 2},
		Username:  "Imik",
	},
}
