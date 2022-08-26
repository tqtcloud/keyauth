package impl

import (
	"context"
	"github.com/infraboard/mcube/exception"
	"github.com/tqtcloud/keyauth/apps/endpoint"
)

// save 入库mongo  的实现
func (s *service) save(ctx context.Context, set *endpoint.EndpiontSet) error {
	// s.col.InsertMany()
	if _, err := s.col.InsertMany(ctx, set.ToDocs()); err != nil {
		return exception.NewInternalServerError("inserted service %s endpoint document error, %s",
			set.Service, err)
	}
	return nil
}
