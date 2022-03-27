package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/yaronp/go-gql-mac-info/graph/generated"
	"github.com/yaronp/go-gql-mac-info/graph/model"
	"github.com/yaronp/go-gql-mac-info/instrumentation"
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

func (r *queryResolver) SysApps(ctx context.Context) ([]*model.ApplicationsData, error) {
	apps := instrumentation.ApplicationData()
	var appData []*model.ApplicationsData
	for _, v := range apps {
		v1 := v
		appData = append(appData, &model.ApplicationsData{
			Name:         &v1.Name,
			ArchKind:     &v1.ArchKind,
			ObtainedFrom: &v1.ObtainedFrom,
			Path:         &v1.Path,
			// SignedBy     []*string `json:"SignedBy"`
			Version: &v1.Version,
			Info:    &v1.Info,
		})
	}
	return appData, nil
}

func (r *queryResolver) App(ctx context.Context, name *string) (*model.ApplicationsData, error) {
	apps := instrumentation.ApplicationData()
	for _, v := range apps {
		if v.Name == *name {
			v1 := v
			return &model.ApplicationsData{
				Name:         &v1.Name,
				ArchKind:     &v1.ArchKind,
				ObtainedFrom: &v1.ObtainedFrom,
				Path:         &v1.Path,
				// SignedBy     []*string `json:"SignedBy"`
				Version: &v1.Version,
				Info:    &v1.Info,
			}, nil

		}
	}
	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
