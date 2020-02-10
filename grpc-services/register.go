package register

import (
	"context"

	"github.com/longbei65/reptilia/protobuf/register"
)

type RegistrationService struct {
}

func (s *RegistrationService) Register(ctx context.Context, request *register.RegisterRequest) (*register.RegisterReply, error) {
	return &register.RegisterReply{
		Success: true,
	}, nil
}
