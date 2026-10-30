package data

import (
	"context"

	reviewPb "github.com/BradleyZhang/ecommerce-review-o/api/gen/go/review/v1"
	"github.com/BradleyZhang/ecommerce-review-o/common"
	"github.com/BradleyZhang/ecommerce-review-o/internal/biz"
)

type OperationRepo struct {
	data *Data
}

// AuditAppeal implements [biz.OperationRepo].
func (o *OperationRepo) AuditAppeal(ctx context.Context, param *biz.AuditAppealParam) error {
	res, err := o.data.rc.AuditAppeal(ctx, &reviewPb.AuditAppealRequest{
		AppealId:  param.AppealID,
		ReviewId:  param.ReviewID,
		Status:    int32(param.Status),
		OpUser:    param.OpUser,
		OpRemarks: &param.OpRemarks,
	})
	common.DebugLogf("[data]AuditAppeal res:%v", res)
	return err
}

// AuditReview implements [biz.OperationRepo].
func (o *OperationRepo) AuditReview(ctx context.Context, param *biz.AuditReviewParam) error {
	res, err := o.data.rc.AuditReview(ctx, &reviewPb.AuditReviewRequest{
		ReviewId:  param.ReviewID,
		Status:    int32(param.ReviewID),
		OpReason:  param.OpReason,
		OpRemarks: &param.OpRemarks,
		OpUser:    param.OpUser,
	})
	common.DebugLogf("data AuditReview res:%v", res)
	return err
}

func NewOperationRepo(data *Data) biz.OperationRepo {
	return &OperationRepo{
		data: data,
	}
}
