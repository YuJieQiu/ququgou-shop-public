package model

import (
	"github.com/ququgou-shop/library/base_model"
	"time"
)

//TODO:订单售后 退款申请
type OrderReturnsApply struct {
	base_model.IDAutoModel
	UserId        uint64                    `json:"userId" gorm:"column:user_id"`
	OrderId       uint64                    `json:"orderId" gorm:"column:order_id"`
	OrderMasterId uint64                    `json:"orderMasterId" gorm:"column:order_master_id;"`
	OrderDetailId uint64                    `json:"orderDetailId" gorm:"column:order_detail_id"`
	ReturnNo      string                    `json:"returnNo" gorm:"column:return_no"`           //售后流水
	MerId         uint64                    `json:"merId" gorm:"column:mer_id"`                 //商家id
	Type          int                       `json:"type" gorm:"column:type"`                    //类型id
	ProductStatus int                       `json:"productStatus" gorm:"column:product_status"` //商品状态
	ProductNumber int                       `json:"productNumber" gorm:"column:product_number"` //商品数量
	Reason        string                    `json:"reason" gorm:"column:reason"`                //原因
	ReasonImages  base_model.ImageJsonModel `json:"reasonImages" gorm:"column:reason_images"`   //原因图片
	Status        int                       `json:"status" gorm:"column:status"`                //'审核状态 -1 拒绝 0 未审核 1审核通过'
	AuditTime     time.Time                 `json:"auditTime" gorm:"column:audit_time"`         //审核时间
	AuditReason   time.Time                 `json:"auditReason" gorm:"column:audit_reason"`     //审核时间
	Remark        string                    `json:"remark" gorm:"column:remark"`                //备注
	base_model.TimeAllModel
}

// Set table name
func (OrderReturnsApply) TableName() string {
	return "order_returns_apply"
}
