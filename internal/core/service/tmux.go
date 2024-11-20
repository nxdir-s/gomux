package service

import (
	"context"

	"github.com/nxdir-s/gomux/internal/ports"
)

type TmuxService struct {
	adapter ports.TmuxPort
}

func NewTmuxService(adapter ports.TmuxPort) (*TmuxService, error) {
	return &TmuxService{
		adapter: adapter,
	}, nil
}

func (s *TmuxService) SessionExists(ctx context.Context) int {
	return s.adapter.HasSession(ctx)
}

func (s *TmuxService) NewSession(ctx context.Context, name string) error {
	return s.adapter.NewSession(ctx, name)
}

func (s *TmuxService) NewWindow(ctx context.Context, cfgIndex int) error {
	return s.adapter.NewWindow(ctx, cfgIndex)
}

func (s *TmuxService) SendKeys(ctx context.Context, name string, args ...string) error {
	return s.adapter.SendKeys(ctx, name, args...)
}

func (s *TmuxService) AttachSession(ctx context.Context) error {
	return s.adapter.AttachSession(ctx)
}

func (s *TmuxService) SelectWindow(ctx context.Context, cfgIndex int) error {
	return s.adapter.SelectWindow(ctx, cfgIndex)
}
