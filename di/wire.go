//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/mapserver2007/mcp-server-golang-poc/app/infrastructure/mcp"
	"github.com/sirupsen/logrus"
)

// var ServerSet = wire.NewSet(
// 	mcp.NewServer,
// )

func NewMcpServer(
	logger *logrus.Logger,
) *mcp.Server {
	wire.Build(
		mcp.NewServer,
	)
	return nil
}
