package handler

import (
	"context"
	"golectro-user/internal/delivery/grpc/interceptor"
	proto "golectro-user/internal/delivery/grpc/proto/address"
	"golectro-user/internal/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddressHandler struct {
	proto.UnimplementedAddressServiceServer
	UseCase *usecase.AddressUseCase
}

func (s *AddressHandler) GetAddress(ctx context.Context, _ *proto.GetAddressRequest) (*proto.GetAddressResponse, error) {
	auth := interceptor.GetUserFromContext(ctx)
	if auth == nil {
		return nil, status.Error(codes.Internal, "auth not found in context")
	}

	user, err := s.UseCase.GetAddressesByUserID(ctx, auth.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user profile")
	}

	addresses := make([]*proto.Address, 0, len(user))
	for _, addr := range user {
		addresses = append(addresses, &proto.Address{
			Id:            addr.ID.String(),
			Label:         addr.Label,
			Recipient:     addr.Recipient,
			Phone:         addr.Phone,
			AddressLine:   addr.AddressLine,
			City:          addr.City,
			Province:      addr.Province,
			PostalCode:    addr.PostalCode,
			IsDefault:     addr.IsDefault,
			Encrypted:     addr.Encrypted,
			CreatedAt:     addr.CreatedAt.Format("2006-01-02T15:04:05+07:00"),
			UpdatedAt:     addr.UpdatedAt.Format("2006-01-02T15:04:05+07:00"),
		})
	}

	return &proto.GetAddressResponse{
		Addresses: addresses,
	}, nil
}
