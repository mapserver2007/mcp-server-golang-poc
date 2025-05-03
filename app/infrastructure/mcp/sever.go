package mcp

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/sirupsen/logrus"
)

const (
	ServerName    = "POC"
	ServerVersion = "1.0.0"
)

type Server struct {
	logger *logrus.Logger
}

func NewServer(
	logger *logrus.Logger,
) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) Start() error {
	mcpServer := server.NewMCPServer(
		ServerName,
		ServerVersion,
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	currentTimeTool := mcp.NewTool("current_time",
		mcp.WithDescription("現在時刻を返します"),
	)

	mcpServer.AddTool(currentTimeTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		now := time.Now().In(jst)
		message := fmt.Sprintf("現在の時刻: %s", now.Format("2006-01-02 15:04:05"))
		return mcp.NewToolResultText(message), nil
	})

	if err := server.ServeStdio(mcpServer); err != nil {
		fmt.Printf("サーバーエラー: %v\n", err)
	}

	return nil
}
