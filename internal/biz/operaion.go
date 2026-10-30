package biz

import (
	"context"

	"github.com/BradleyZhang/ecommerce-review-o/common"
)

type OperationRepo interface {
	AuditReview(ctx context.Context, param *AuditReviewParam) error
	AuditAppeal(ctx context.Context, param *AuditAppealParam) error
}

type OperationUsecase struct {
	repo OperationRepo
}

func NewOperationUsecase(repo OperationRepo) *OperationUsecase {
	return &OperationUsecase{
		repo: repo,
	}
}
func (uc *OperationUsecase) AuditReview(ctx context.Context, param *AuditReviewParam) error {
	common.DebugLogf("[biz] AuditReview ,param:%v", param)
	return uc.repo.AuditReview(ctx, param)
}
func (uc *OperationUsecase) AuditAppeal(ctx context.Context, param *AuditAppealParam) error {
	common.DebugLogf("[biz] AuditAppeal,param: %v", param)
	return uc.repo.AuditAppeal(ctx, param)
}

// AuditReviewParam 审核评价的参数
type AuditReviewParam struct {
	ReviewID  int64
	Status    int
	OpReason  string
	OpRemarks string
	OpUser    string
}

// AuditAppealParam 审核申诉的参数
type AuditAppealParam struct {
	AppealID  int64
	ReviewID  int64
	StoreID   int64
	Status    int
	OpReason  string
	OpRemarks string
	OpUser    string
}
