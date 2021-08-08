package paymentService

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/payment/model"
	"github.com/ququgou-shop/modules/payment/paymentEnum"
)

//支付通知结果处理
func PaymentNotifyResult(db *gorm.DB, data *PaymentNotifyResultModel) (error, *model.Transaction) {
	var (
		err         error
		trade       model.Transaction
		tradeDetail model.TransactionDetail
	)

	err = db.Where("trade_no=?", data.TransactionNo).First(&trade).Error
	if err != nil {
		return err, nil
	}

	if trade.Amount != data.TotalFee {
		return errors.New("Amount Is not equal to TotalFee "), nil
	}

	trade.OutTradeNo = data.OutTradeNo
	trade.Status = int(paymentEnum.TradeStatusSucceed)
	trade.CompletionTime = &data.TimeEnd

	tx := db.Begin()
	err = tx.Save(&trade).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	tradeDetail = model.TransactionDetail{
		TransactionId: trade.ID,
		OutTradeNo:    data.OutTradeNo,
		BankType:      data.BankType,
	}

	err = tx.Create(&tradeDetail).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	tx.Commit()

	return nil, &trade
}
