package grpc_handler

import (
	"context"

	// "fmt"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/logic"
	proto "github.com/p4elkab35t/salyte_backend/services/auth/pkg/server/proto"
)

type SecurityLogHandler struct {
	securityLogLogic *logic.SecurityLogLogicService
	proto.UnimplementedSecurityLogsServiceServer
}

func NewSecurityLogHandler(securityLogLogic *logic.SecurityLogLogicService) *SecurityLogHandler {
	return &SecurityLogHandler{securityLogLogic: securityLogLogic}
}

func (h *SecurityLogHandler) GetSecurityLogsWithUsedID(ctx context.Context, req *proto.GetSecurityLogsByUserIDRequest) (*proto.GetSecurityLogsByUserIDResponse, error) {
	securityLogs, err := h.securityLogLogic.GetAllSecurityLogsByUserID(ctx, req.UserId)

	if err != nil {
		return &proto.GetSecurityLogsByUserIDResponse{Status: 1}, err
	}
	if securityLogs == nil {
		return &proto.GetSecurityLogsByUserIDResponse{Status: 1}, nil
	}

	// parse security logs to proto

	securityLogsProto := make([]*proto.SecurityLog, len(securityLogs))
	for i, securityLog := range securityLogs {
		securityLogsProto[i] = &proto.SecurityLog{
			LogId:  securityLog.Log_id,
			UserId: securityLog.User_id,
			Action: securityLog.Action,
			//IpAddress: securityLog.Ip_address,
			Timestamp: securityLog.Timestamp.String(),
		}
	}

	return &proto.GetSecurityLogsByUserIDResponse{SecurityLogs: securityLogsProto, Status: 0}, nil
}

func (h *SecurityLogHandler) GetSecurityLogWithID(ctx context.Context, req *proto.GetSecurityLogWithIDRequest) (*proto.GetSecurityLogWithIDResponse, error) {
	securityLog, err := h.securityLogLogic.GetSecurityLogByID(ctx, req.LogId)

	if err != nil {
		return &proto.GetSecurityLogWithIDResponse{Status: 1}, err
	}
	if securityLog == nil {
		return &proto.GetSecurityLogWithIDResponse{Status: 1}, nil
	}

	// parse security log to proto

	securityLogProto := &proto.SecurityLog{
		LogId:  securityLog.Log_id,
		UserId: securityLog.User_id,
		Action: securityLog.Action,
		//IpAddress: securityLog.Ip_address,
		Timestamp: securityLog.Timestamp.String(),
	}
	return &proto.GetSecurityLogWithIDResponse{SecurityLog: securityLogProto, Status: 0}, nil
}
