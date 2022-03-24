package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/yaronp/go-gql-mac-info/instrumentation"

	"github.com/yaronp/go-gql-mac-info/graph/generated"
	"github.com/yaronp/go-gql-mac-info/graph/model"
)

func (r *queryResolver) SysInfo(ctx context.Context) (*model.SoftwareData, error) {
	sdi := instrumentation.SoftwareData()
	m := model.SoftwareData{
		BootMode:        &sdi.BootMode,
		BootVolume:      &sdi.BootVolume,
		KernelVersion:   &sdi.KernelVersion,
		LocalHostname:   &sdi.LocalHostName,
		OsVersion:       &sdi.OsVersion,
		SecureVM:        &sdi.SecureVm,
		SystemIntegrity: &sdi.SystemIntegrity,
		Uptime:          &sdi.Uptime,
		UserName:        &sdi.UserName,
	}
	return &m, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
