package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gin_graphql/app/models"
	"gin_graphql/graph/generated"
	"gin_graphql/graph/model"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	var user models.User
	return user.GetUserByID(obj.UserID)
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
	var meetups models.Meetup
	return meetups.Get()
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	var users models.User
	return users.Get()
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	// var m []*models.Meetup
	// for _, meetup := range meetups {
	// 	if meetup.UserID == obj.ID {
	// 		m = append(m, meetup)
	// 	}
	// }
	var meetups models.Meetup
	return meetups.GetMeetupsByUser(obj)
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
