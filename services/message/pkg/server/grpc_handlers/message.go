package grpc_handler

import (
	"context"

	// "fmt"
	"github.com/google/uuid"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/logic"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/models"
	proto "github.com/p4elkab35t/salyte_backend/services/message/pkg/server/proto"
)

type MessageHandler struct {
	messageService *logic.MessageService
	proto.UnimplementedMessagingServiceServer
}

func NewMessageHandler(messageService *logic.MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

func (h *MessageHandler) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	newMessage := &models.Message{
		Content:  req.Content,
		ChatID:   uuid.MustParse(req.ChatId),
		SenderID: uuid.MustParse(req.SenderId),
	}

	message, err := h.messageService.SendMessage(ctx, newMessage, newMessage.SenderID)
	if err != nil {
		return &proto.SendMessageResponse{Status: 1}, err
	}
	if message == nil {
		return &proto.SendMessageResponse{Status: 1}, nil
	}

	protoMessage := &proto.Message{
		Id:        message.ID.String(),
		Content:   message.Content,
		ChatId:    message.ChatID.String(),
		SenderId:  message.SenderID.String(),
		CreatedAt: message.CreatedAt.String(),
		// Add other fields as necessary
	}
	return &proto.SendMessageResponse{Status: 0, Message: protoMessage}, nil
}

func (h *MessageHandler) EditMessage(ctx context.Context, req *proto.EditMessageRequest) (*proto.EditMessageResponse, error) {
	message, err := h.messageService.GetMessageByID(ctx, uuid.MustParse(req.MessageId), uuid.MustParse(req.UserId))
	if err != nil {
		return &proto.EditMessageResponse{Status: 1}, err
	}
	if message == nil {
		return &proto.EditMessageResponse{Status: 1}, nil
	}

	editedMessage := &models.Message{
		ID:       message.ID,
		Content:  req.NewContent,
		ChatID:   message.ChatID,
		SenderID: uuid.MustParse(req.UserId),
	}

	err = h.messageService.EditMessage(ctx, message.ID, editedMessage.Content, editedMessage.SenderID)
	if err != nil {
		return &proto.EditMessageResponse{Status: 1}, err
	}
	if message == nil {
		return &proto.EditMessageResponse{Status: 1}, nil
	}

	protoMessage := &proto.Message{
		Id:        message.ID.String(),
		Content:   message.Content,
		ChatId:    message.ChatID.String(),
		SenderId:  message.SenderID.String(),
		CreatedAt: message.CreatedAt.String(),
		// Add other fields as necessary
	}
	return &proto.EditMessageResponse{Status: 0, Message: protoMessage}, nil
}
