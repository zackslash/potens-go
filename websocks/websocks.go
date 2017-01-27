package websocks

import (
	"github.com/cubex/proto-go/sockets"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// SOCKETSGAID is the socket server Global App ID
	SOCKETSGAID = "cubex/sockets"
)

type Handler struct {
	client     sockets.SocketsClient
	connection *grpc.ClientConn
	ctx        context.Context
	appID      string
	vendorID   string
}

func NewHandler(cc *grpc.ClientConn, ctx context.Context, vendor string, appID string) *Handler {
	return &Handler{connection:cc, ctx:ctx, client:sockets.NewSocketsClient(cc), appID:appID, vendorID:vendor}
}

func (h *Handler) Subscribe(socketID string, channelName string) (*sockets.PublishResponse, error) {
	return h.client.Subscribe(h.ctx, &sockets.SubscribeMessage{
		SessionId: socketID,
		Channel:   channelName,
	})
}

func (h *Handler) SendMessage(channelName string, action string, payload string) (*sockets.PublishResponse, error) {
	return h.client.Publish(h.ctx, &sockets.SocketMessage{
		Channel:   channelName,
		Action:    action,
		Payload:   payload,
	})
}
