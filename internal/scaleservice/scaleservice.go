package scaleservice

import (
	"context"

	pb "github.com/sheetpilot/sheet-pilot-proto/scaleservice"
	"github.com/sirupsen/logrus"
)

type ScaleService struct {
	client pb.ScaleServiceClient
}

func New(log *logrus.Entry, client pb.ScaleServiceClient) (*ScaleService, error) {
	return &ScaleService{client: client}, nil
}

func (s *ScaleService) SendScaleRequest(ctx context.Context) (*pb.ScaleResponse, error) {
	req := &pb.ScaleRequest{}

	response, err := s.client.ScaleServiceRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
