package directives

import (
	"context"
	"gin_graphql/graph"

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
