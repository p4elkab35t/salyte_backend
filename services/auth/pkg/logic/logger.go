package logic

import (
	"context"

	// "errors"

	// "os/user"
	// "fmt"
	"time"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/storage/repository"
)

type SecurityLogLogicService struct {
	SecurityLogRepo repository.SecurityLogRepository
}

func NewSecurityLogLogicService(securityLogRepo repository.SecurityLogRepository) *SecurityLogLogicService {
	return &SecurityLogLogicService{
		SecurityLogRepo: securityLogRepo,
	}
}

// TODO: ADD IP SOTRING OPTIONS

func (s *SecurityLogLogicService) CreateSecurityLog(ctx context.Context, user_id, action string) (*models.SecurityLog, error) {
	securityLog := &models.SecurityLog{
		User_id:   user_id,
		Action:    action,
		Timestamp: time.Now(),
	}

	// fmt.Println("securityLog: ", securityLog)

	securityLog, err := s.SecurityLogRepo.CreateSecurityLog(ctx, securityLog)
	if err != nil {
		// fmt.Println("error: ", err)
		return nil, err
	}
	return securityLog, nil
}

func (s *SecurityLogLogicService) GetAllSecurityLogsByUserID(ctx context.Context, user_id string) ([]*models.SecurityLog, error) {
	securityLogs, err := s.SecurityLogRepo.GetAllSecurityLogsByUserID(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return securityLogs, nil
}

func (s *SecurityLogLogicService) GetSecurityLogByID(ctx context.Context, id string) (*models.SecurityLog, error) {
	securityLog, err := s.SecurityLogRepo.GetSecurityLogByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return securityLog, nil
}
