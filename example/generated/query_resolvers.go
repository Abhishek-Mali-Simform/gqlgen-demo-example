package generated

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"

	"example/models"
)

// Employees is the resolver for the employees field.
func (r *queryResolver) Employees(ctx context.Context, filter *models.FilterInput, page *int) (*models.Pagination, error) {
	return r.Domain.RetrieveEmployees(filter, page)
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, filter *models.FilterInput, page *int) (*models.Pagination, error) {
	return r.Domain.RetrieveTasks(filter, page)
}

// Companies is the resolver for the companies field.
func (r *queryResolver) Companies(ctx context.Context, filter *models.FilterCompanyInput, page *int) (*models.Pagination, error) {
	return r.Domain.RetrieveCompanies(filter, page)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
