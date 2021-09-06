package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"gin_graphql/app/models"
	"gin_graphql/graph/generated"
	"gin_graphql/graph/middleware"
	"gin_graphql/graph/model"
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	var user models.User
	return user.GetUserByID(obj.UserID)
}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	var user *models.User
	_, err := user.GetUserByEmail(input.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	_, err = user.GetUserByAccount(input.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrAccountUsed
	}

	user = &models.User{
		Account: input.Account,
		Email:   input.Email,
	}

	if input.Password != input.ConfirmPassword {
		return nil, ErrWrongPassword
	}
	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, ErrUnkown
	}

	_, err = user.Create(user)
	if err != nil {
		log.Printf("error while create user: %v", err)
		return nil, ErrUnkown
	}

	// TODO: 產生 Token >>>
	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week
	authToken := &model.AuthToken{
		AccessToken: "gogopowerkimi",
		ExpiredAt:   expiredAt,
	}
	// 產生 Token <<<

	return &model.AuthResponse{
		AuthToken: authToken,
		User:      user,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	var user *models.User
	user, err := user.GetUserByEmail(input.Email)
	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredentials
	}

	// TODO: 取得 Token >>>
	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week
	authToken := &model.AuthToken{
		AccessToken: "gogopowerkimi",
		ExpiredAt:   expiredAt,
	}
	// 產生 Token <<<

	return &model.AuthResponse{
		AuthToken: authToken,
		User:      user,
	}, nil
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, ErrUnauthenticated
	}

	if len(input.Name) < 3 {
		return nil, ErrNameNotLongEnough
	}
	if len(input.Description) < 3 {
		return nil, ErrDescriptionNotLongEnough
	}
	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}
	return meetup.Create(meetup)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*models.Meetup, error) {
	// fmt.Println("kkikimimimi")
	var meetup models.Meetup
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	inputObj := map[string]interface{}{
		"name":        *input.Name,
		"description": *input.Description,
	}

	return meetup.Update(idInt, inputObj)
}

func (r *mutationResolver) DeleteMeetUp(ctx context.Context, id string) (bool, error) {
	var meetups models.Meetup
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return false, err
	}
	return meetups.Delete(idInt)
}

func (r *queryResolver) Meetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
	return model.GetMeetups(filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	var users models.User
	return users.Get()
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
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
