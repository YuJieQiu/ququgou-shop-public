package paymentService

import (
	"encoding/json"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/utils"
	"github.com/ququgou-shop/modules/payment/common"
	"github.com/ququgou-shop/modules/payment/model"
	"github.com/ququgou-shop/modules/payment/paymentEnum"
	userModel "github.com/ququgou-shop/modules/user/model"
)

func (m CreatePaymentInfoModel) verify() bool {
	if m.BusinessNo == "" {
		return false
	}
	if m.Amount <= 0 {
		return false
	}
	return true
}

//根据交易流水等信息

//创建支付
//TODO:调用 分布式锁，防止重复交易
func CreatePayment(db *gorm.DB, q *CreatePaymentModel) (error, *OnlinePaymentCreateResult) {
	var (
		err         error
		user        userModel.User
		paymentType model.PaymentType
		trade       *model.Transaction
		res         *OnlinePaymentCreateResult
	)

	//TODO:错误处理
	//defer func() {
	//	if r := recover(); r != nil {
	//		srt, _ := json.Marshal(r)
	//		srts := string(srt)
	//		err = errors.New(fmt.Sprintf("WechatOnlinePaymentCreate Error %v", srts))
	//		return
	//	}
	//}()

	if err != nil {
		return err, nil
	}

	err = db.Where("id=? and status=?", q.PaymentTypeId, 0).First(&paymentType).Error
	if err != nil {
		return err, nil
	}

	err = db.Where("id=?", q.UserId).First(&user).Error
	if err != nil {
		return err, nil
	}

	err, trade = getTransaction(db, q.BusinessNo, q.BusinessType)
	if err != nil {
		return err, nil
	}

	if trade != nil {
		////判断交易状态，如果为等待支付的可以往下走，如果支付成功的不能重复交易
		if trade.Status == int(paymentEnum.TradeStatusSucceed) {
			return errors.New("repetition trade "), nil
		}
	} else {
		//生成交易流水
		tradeNo := common.CreateTradeNo(user.Guid)
		trade = &model.Transaction{
			UserId:           user.ID,
			BusinessNo:       q.BusinessNo,
			BusinessType:     q.BusinessType,
			BusinessTypeCode: paymentEnum.BusinessType(q.BusinessType).String(),
			TradeNo:          tradeNo,
			TradeType:        int(paymentEnum.TradeTypeOrderPay),
			Amount:           q.Amount,
			PaymentTypeId:    paymentType.ID,
			PaymentTypeCode:  paymentType.Code,
			Source:           q.Source,
			Status:           int(paymentEnum.TradeStatusProcessed),
			Note:             q.Note,
		}
		err = db.Create(trade).Error
		if err != nil {
			return err, nil
		}
	}

	err, res = createOnlinePayment(db, trade, &paymentType, &user, q.ClientInfo)
	if err != nil {
		return err, nil
	}

	//更新trade
	if res.Success {
		err = db.Table(trade.TableName()).Where("trade_no=?", trade.TradeNo).
			Updates(model.Transaction{Status: int(paymentEnum.TradeStatusWaitPay)}).Error
	} else {
		err = db.Table(trade.TableName()).Where("trade_no=?", trade.TradeNo).
			Updates(model.Transaction{Status: int(paymentEnum.TradeStatusFail)}).Error
	}
	if err != nil {
		return err, res
	}

	res.TradeNo = trade.TradeNo
	return nil, res
}

//创建在线交易
func createOnlinePayment(db *gorm.DB,
	trade *model.Transaction,
	paymentType *model.PaymentType,
	user *userModel.User,
	clientInfo ClientInfoModel) (error, *OnlinePaymentCreateResult) {

	var (
		err           error
		onlinePayment OnlinePayment
		res           *OnlinePaymentCreateResult
	)

	//调用在线支付 方法，如 微信支付、支付宝支付
	switch paymentEnum.PaymentType(paymentType.Code) {
	case paymentEnum.PaymentTypeWeChatPay:
		onlinePayment = WechatPayment{
			TradeNo:      trade.TradeNo,
			Amount:       trade.Amount,
			WechatOpenId: user.WechatOpenId,
			ClientIp:     clientInfo.Ip,
		}
		break
	}

	var record model.PaymentOnlineRecord
	var resJson string

	//调用在线方法 返回结果
	err, res = onlinePayment.OnlinePaymentCreate(db)
	resByte, _ := json.Marshal(res)
	resJson = string(resByte)

	//记录在线支付信息
	record = model.PaymentOnlineRecord{
		TradeNo:         trade.TradeNo,
		PaymentTypeId:   paymentType.ID,
		PaymentTypeCode: paymentType.Code,
		Success:         res.Success,
		Result:          resJson,
	}

	if err != nil {
		//TODO:支付错误日志记录
	}

	err = db.Create(&record).Error

	//更新 交易表信息
	if err != nil {
		return err, res
	}

	return nil, res

	//if !res.Success {
	//	return errors.New(res.Msg), nil
	//}

}

func getTransaction(db *gorm.DB, businessNo string, businessType int) (error, *model.Transaction) {
	var (
		err   error
		trade model.Transaction
	)
	err = db.Where("business_no=? and business_type=? and status >= ?",
		businessNo,
		businessType,
		int(paymentEnum.TradeStatusProcessed)).
		First(&trade).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err, nil
	}
	if trade.ID == 0 {
		return nil, nil
	}

	return nil, &trade
}

func getPaymentOnlineRecord(db *gorm.DB, tradeNo string) (error, *model.PaymentOnlineRecord) {
	var (
		err    error
		record model.PaymentOnlineRecord
	)
	err = db.Where("trade_no=? and success= ?", tradeNo, utils.BoolToInt(true)).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err, nil
	}

	if record.ID == 0 {
		return nil, nil
	}

	return nil, &record
}
