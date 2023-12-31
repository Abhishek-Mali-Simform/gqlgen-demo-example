package generated

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"example/dataloaders"
	"example/models"
)

// Employees is the resolver for the employees field.
func (r *taskResolver) Employees(ctx context.Context, obj *models.Task) ([]*models.Employee, error) {
	return dataloaders.GetEmployeeLoader(ctx, dataloaders.EmployeeLoaderByTaskKey).Load(obj.ID)
}

// Task returns TaskResolver implementation.
func (r *Resolver) Task() TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }
