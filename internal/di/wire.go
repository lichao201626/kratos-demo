// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/lichao201626/kratos-demo/internal/dao"
	"github.com/lichao201626/kratos-demo/internal/server/grpc"
	"github.com/lichao201626/kratos-demo/internal/server/http"
	"github.com/lichao201626/kratos-demo/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
