package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gin_graphql/graph/generated"
	"gin_graphql/graph/model"
)

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*model.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	// panic(fmt.Errorf("not implemented"))
	return r.user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
