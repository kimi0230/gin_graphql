package validator

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"gopkg.in/go-playground/validator.v9"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

func Binding(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
	val, err := next(ctx)
	if err != nil {
		panic(err)
	}
	fieldName := *graphql.GetPathContext(ctx).Field

	err = validate.Var(val, constraint)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errMsg := fmt.Errorf("%s%+v", fieldName, validationErrors)
		return val, errMsg
	}

	return val, nil
}
