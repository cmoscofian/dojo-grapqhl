package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mercadolibre/fury_post-compra-dojo-graphql/cmd/api/graph/generated"
	"github.com/mercadolibre/fury_post-compra-dojo-graphql/cmd/api/graph/model"
	"github.com/mercadolibre/fury_post-compra-dojo-graphql/cmd/api/graph/rest"
)

func (r *claimResolver) Order(ctx context.Context, obj *model.Claim) (*model.Order, error) {
	return rest.GetOrder(obj.ResourceID)
}

func (r *orderResolver) Shipment(ctx context.Context, obj *model.Order) (*model.Shipment, error) {
	return rest.GetShipment(obj.Shipping.ID)
}

func (r *queryResolver) Claim(ctx context.Context, claimID int) (*model.Claim, error) {
	return rest.GetClaim(claimID)
}

func (r *queryResolver) Order(ctx context.Context, orderID int) (*model.Order, error) {
	return rest.GetOrder(orderID)
}

// Claim returns generated.ClaimResolver implementation.
func (r *Resolver) Claim() generated.ClaimResolver { return &claimResolver{r} }

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type claimResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Shipment(ctx context.Context, shipmentID int) (*model.Shipment, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *claimResolver) Shipment(ctx context.Context, obj *model.Claim) (*model.Shipment, error) {
	panic(fmt.Errorf("not implemented"))
}
