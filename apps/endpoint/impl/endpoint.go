package impl

import (
	"context"
	"github.com/tqtcloud/keyauth/apps/endpoint"
)

// RegistryEndpoint 服务功能注册
func (s *service) RegistryEndpoint(ctx context.Context, req *endpoint.EndpiontSet) (*endpoint.RegistryResponse, error) {
	if err := s.save(ctx, req); err != nil {
		return nil, err
	}
	resp := endpoint.NewRegistryResponse()
	return resp, nil
}
