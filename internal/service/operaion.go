package service

import (
	"context"

	v1 "github.com/BradleyZhang/ecommerce-review-o/api/gen/go/operation/v1"
	"github.com/BradleyZhang/ecommerce-review-o/internal/biz"
)

type OperationService struct {
	v1.UnimplementedOperationServiceServer
	uc *biz.OperationUsecase
}

func NewOperationService(uc *biz.OperationUsecase) *OperationService {
	return &OperationService{uc: uc}
}

func (s *OperationService) AuditReview(ctx context.Context, req *v1.AuditReviewRequest) (*v1.AuditReviewResponse, error) {
	err := s.uc.AuditReview(ctx, &biz.AuditReviewParam{
		ReviewID:  req.GetReviewId(),
		Status:    int(req.GetStatus()),
		OpReason:  req.OpReason,
		OpRemarks: req.GetOpRemarks(),
		OpUser:    req.GetOpUser(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.AuditReviewResponse{}, nil
}

func (s *OperationService) AuditAppeal(ctx context.Context, req *v1.AuditAppealRequest) (*v1.AuditAppealResponse, error) {
	err := s.uc.AuditAppeal(ctx, &biz.AuditAppealParam{
		AppealID:  req.GetAppealId(),
		ReviewID:  req.GetReviewId(),
		Status:    int(req.GetStatus()),
		OpReason:  req.GetOpReason(),
		OpRemarks: req.GetOpRemarks(),
		OpUser:    req.GetOpUser(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.AuditAppealResponse{}, nil
}
