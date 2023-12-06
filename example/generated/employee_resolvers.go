package generated

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"example/dataloaders"
	"example/models"
)

// Company is the resolver for the company field.
func (r *employeeResolver) Company(ctx context.Context, obj *models.Employee) (*models.Company, error) {
	return dataloaders.GetCompanyLoader(ctx).Load(obj.CompanyID)
}

// Tasks is the resolver for the tasks field.
func (r *employeeResolver) Tasks(ctx context.Context, obj *models.Employee) ([]*models.Task, error) {
	return dataloaders.GetTaskLoader(ctx).Load(obj.ID)
}

// Employee returns EmployeeResolver implementation.
func (r *Resolver) Employee() EmployeeResolver { return &employeeResolver{r} }

type employeeResolver struct{ *Resolver }