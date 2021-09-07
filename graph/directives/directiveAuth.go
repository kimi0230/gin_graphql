package directives

import (
	"context"
	"gin_graphql/graph"
	"gin_graphql/graph/middleware"
	"gin_graphql/graph/model"

	"github.com/99designs/gqlgen/graphql"
)

const CurrentUserKey = "currentUser"

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	ctxUserID := ctx.Value(CurrentUserKey)
	if ctxUserID == nil {
		return nil, graph.ErrUnauthenticated
	}
	return next(ctx)
}

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
	if user, err := middleware.GetCurrentUserFromCTX(ctx); err == nil {
		if user.Role != role.String() {
			return nil, graph.ErrUnauthenticated
		}
	} else {
		return nil, graph.ErrUnauthenticated
	}

	return next(ctx)
}
