package controller

//微信异步通知
//
//type (
//	WechatNotifyController struct {
//	}
//)
//
//type CDATA struct {
//	Text string `xml:",cdata"`
//}
//type PaySuccessResponse struct {
//	ReturnCode CDATA `xml:"return_code"`
//	ReturnMsg  CDATA `xml:"return_msg"`
//}
//
////微信支付成功通知
//func (w WechatNotifyController) PaySuccessNotify(c *gin.Context) {
//	var (
//		req     wechat.NotifyResult
//		res     PaySuccessResponse
//		payConf model.PaymentConfig
//		user    userModel.User
//	)
//	res.ReturnCode = CDATA{Text: "SUCCESS"}
//	res.ReturnMsg = CDATA{Text: "OK"}
//
//	if err := c.BindXML(&req); err != nil {
//		fmt.Println("PaySuccessNotify Error:" + err.Error())
//		//解析失败
//
//		c.XML(http.StatusOK, res)
//		return
//	}
//
//	//TODO:打印调试
//	srt, _ := json.Marshal(req)
//	fmt.Println(string(srt))
//
//	if req.ReturnCode != "SUCCESS" {
//		//存入数据库和日志
//		c.XML(http.StatusOK, res)
//		return
//	}
//
//	//获取微信的商户key
//	err := payConf.GetPaymentConfig(db.MysqlConn(), paymentEnum.PaymentTypeWeChatPay)
//	if err != nil || payConf.ID <= 0 {
//		fmt.Println("PaySuccessNotify Error:" + err.Error())
//		c.XML(http.StatusOK, res)
//		return
//	}
//	//验证签名
//	if !wechat.VerifySign(req, payConf.Key) {
//		fmt.Println("Verify Sign Error")
//		c.XML(http.StatusOK, res)
//		return
//	}
//
//	//用户信息校验
//	err = user.GetUser(db.MysqlConn(), req.OpenID)
//	if err != nil || user.ID <= 0 {
//		fmt.Println("PaySuccessNotify Error:" + err.Error())
//		c.XML(http.StatusOK, res)
//		return
//	}
//
//	//金额转换 微信这里返回的金额以分为单位
//	totalFee := float64(req.TotalFee) / float64(100)
//	cashFee := float64(req.CashFee) / float64(100)
//	timeEnd, _ := time.Parse("20060102150405", req.TimeEnd)
//
//	m := payment.PaymentNotifyResultModel{
//		AppId:         req.AppID,
//		MchId:         req.MchID,
//		PaymentType:   paymentEnum.PaymentTypeWeChatPay,
//		TransactionNo: req.OutTradeNo,
//		OutTradeNo:    req.TransactionID,
//		BankType:      req.BankType,
//		UserId:        user.ID,
//		TotalFee:      totalFee,
//		CashFee:       cashFee,
//		TimeEnd:       timeEnd,
//	}
//
//	err, trade := payment.PaymentNotifyResult(db.MysqlConn(), &m)
//	if err != nil {
//		fmt.Println("PaySuccessNotify Error:" + err.Error())
//		c.XML(http.StatusOK, res)
//		return
//	}
//
//	//更新订单状态
//	oReq := order.OrderPaymentSuccessUpdateModel{
//		BusinessNo:   trade.BusinessNo,
//		BusinessType: trade.BusinessType,
//	}
//
//	err = order.OrderPaymentSuccessUpdate(db.MysqlConn(), &oReq)
//	if err != nil {
//		fmt.Println("PaySuccessNotify Error:" + err.Error())
//		c.XML(http.StatusOK, res)
//		return
//	}
//
//	c.XML(http.StatusOK, res)
//	return
//}
//
////支付结果同步
//func (w WechatNotifyController) PayResultSync(c *gin.Context) {
//	//获取支付处理中的订单状态 超过1分钟没有更改的
//	//查询这个 交易信息 状态
//	//调用微信接口 查询状态 更新状态
//	var (
//		err       error
//		orderSubs []model2.OrderSub
//		payConf   model.PaymentConfig
//	)
//
//	err = db.MysqlConn().
//		Joins("join order_trades on order_trades.order_sub_no=order_subs.order_sub_no").
//		Where(`order_subs.created_at<=DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? MINUTE)
//				    and order_trades.deleted_at IS NULL `,
//			"-1").Find(&orderSubs).Error
//
//	if err != nil && err != gorm.ErrRecordNotFound {
//		fmt.Println("PayResultSync Error" + err.Error())
//	}
//
//	if len(orderSubs) <= 0 {
//		return
//	}
//
//	//获取微信的商户key
//	err = payConf.GetPaymentConfig(db.MysqlConn(), paymentEnum.PaymentTypeWeChatPay)
//	if err != nil || payConf.ID <= 0 {
//		fmt.Println("PaySuccessNotify Error:" + err.Error())
//		return
//	}
//
//	//查询订单的交易流水 查询交易表的状态
//	for i := 0; i < len(orderSubs); i++ {
//		var trade model.Transaction
//		err = db.MysqlConn().
//			Joins("inner join order_trades on order_trades.trade_no=transactions.trade_no").
//			Where("order_trades.order_sub_no=? and order_trades.deleted_at is null and transactions.deleted_at is null",
//				orderSubs[i].OrderSubNo).
//			First(&trade).Error
//		if err != nil {
//			fmt.Println("PayResultSync Error" + err.Error())
//		} else {
//			//判断 trade 状态 交易成功的 更新订单支付状态
//			if trade.Status == int(paymentEnum.TradeStatusSucceed) {
//				tx := db.MysqlConn().Begin()
//				orderSubs[i].PayStatus = int(orderEnum.PaySuccess)
//				orderSubs[i].DeliveryStatus = int(orderEnum.DeliveryWait)
//
//				err = tx.Save(&orderSubs[i]).Error
//
//				if err != nil {
//					tx.Rollback()
//					fmt.Println("PayResultSync Error" + err.Error())
//				}
//				tx.Commit()
//			} else if trade.Status == int(paymentEnum.TradeStatusWaitPay) {
//				//调用微信接口查询交易订单信息
//				pyOrderQueryParams := wechat.PayOrderQueryParams{
//					AppId: wechat.CDATA{Text: payConf.AppId},
//					MchId: wechat.CDATA{Text: payConf.MchId},
//				}
//				if len(trade.OutTradeNo) > 0 {
//					pyOrderQueryParams.TransactionId = &wechat.CDATA{Text: trade.OutTradeNo}
//				} else {
//					pyOrderQueryParams.OutTradeNo = &wechat.CDATA{Text: trade.TradeNo}
//				}
//				err, result := wechat.PayOrderQuery(&pyOrderQueryParams, payConf.Key)
//
//				if err != nil {
//					fmt.Println("PayResultSync Error" + err.Error())
//				}
//
//				tx := db.MysqlConn().Begin()
//
//				if result.TradeState == "SUCCESS" {
//					//支付成功 更新各相记录表
//					//更新 trade 状态
//					trade.Status = int(paymentEnum.TradeStatusSucceed)
//					if len(trade.OutTradeNo) <= 0 {
//						trade.OutTradeNo = result.TransactionID
//					}
//					err = tx.Save(&trade).Error
//					if err != nil {
//						tx.Rollback()
//						fmt.Println("PayResultSync err" + err.Error())
//						return
//					}
//
//					//更新 订单状态
//					orderSubs[i].PayStatus = int(orderEnum.PaySuccess)
//					orderSubs[i].DeliveryStatus = int(orderEnum.DeliveryWait)
//					err = tx.Save(&orderSubs[i]).Error
//					if err != nil {
//						tx.Rollback()
//						fmt.Println("PayResultSync err" + err.Error())
//						return
//					}
//
//					//创建交易详情记录表
//					tradeDetail := model.TransactionDetail{
//						TransactionId: trade.ID,
//						OutTradeNo:    result.TransactionID,
//						BankType:      result.BankType,
//					}
//					err = tx.Create(&tradeDetail).Error
//					if err != nil {
//						tx.Rollback()
//						fmt.Println("PayResultSync err" + err.Error())
//						return
//					}
//					tx.Commit()
//
//				} else {
//					//支付失败 记录原因 更新各相记录表
//					//更新 trade 状态
//					trade.Status = int(paymentEnum.TradeStatusFail)
//					err = tx.Save(&trade).Error
//					if err != nil {
//						tx.Rollback()
//						fmt.Println("PayResultSync err" + err.Error())
//						return
//					}
//
//					//更新 订单状态
//					orderSubs[i].PayStatus = int(orderEnum.PayFail)
//					err = tx.Save(&orderSubs[i]).Error
//					if err != nil {
//						tx.Rollback()
//						fmt.Println("PayResultSync err" + err.Error())
//						return
//					}
//
//					resultSrt, _ := json.Marshal(result)
//
//					//创建交易日志 记录 接口返回信息
//					tradeRecord := model.TransactionRecord{
//						TransactionId: trade.ID,
//						Events:        "微信支付订单查询 失败",
//						Result:        string(resultSrt),
//					}
//					err = tx.Create(&tradeRecord).Error
//					if err != nil {
//						tx.Rollback()
//						fmt.Println("PayResultSync err" + err.Error())
//						return
//					}
//					tx.Commit()
//				}
//
//			}
//
//		}
//	}
//
//}

////微信支付通知结果处理
//func wechatPaySuccessNotifyProcess(req *wechat.NotifyResult) {
//	var (
//		payConf model.PaymentConfig
//	)
//}
