package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mercadolibre/fury_post-compra-dojo-graphql/cmd/api/graph/generated"
	"github.com/mercadolibre/fury_post-compra-dojo-graphql/cmd/api/graph/model"
	"github.com/mercadolibre/fury_post-compra-dojo-graphql/cmd/api/graph/rest"
)

func (r *queryResolver) Claim(ctx context.Context, claimID int) (*model.Claim, error) {
	return rest.GetClaim(claimID)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
