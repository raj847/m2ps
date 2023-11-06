package trxService

import (
	"database/sql"
	"fmt"
	"log"
	"m2ps/constans"
	"m2ps/helpers"
	"m2ps/models"
	"m2ps/services"
	"m2ps/utils"
	"net/http"

	"github.com/labstack/echo"
)

type trxService struct {
	Service services.UsecaseService
}

func NewTrxService(service services.UsecaseService) trxService {
	return trxService{
		Service: service,
	}
}

func (svc *trxService) Create(c echo.Context) error {
	// var tReq models.Date
	var result models.Response
	// var DataPsql *models.Trx
	tReq := new(models.Date)
	if err := helpers.BindValidateStruct(c, tReq); err != nil {
		result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, result)
	}
	data, _, err := svc.Service.TrxMongoRepo.GetData(tReq.Start, tReq.End)
	if err != nil {
		result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, result)
	}
	for _, v := range data {
		var paymentMethod string
		if v.TypeCard == "01" {
			paymentMethod = "PREPAID BCA"
		} else if v.TypeCard == "02" {
			paymentMethod = "PREPAID MANDIRI"
		} else if v.TypeCard == "03" {
			paymentMethod = "PREPAID BNI"
		} else if v.TypeCard == "04" {
			paymentMethod = "PREPAID BRI"
		} else if v.TypeCard == "08" {
			paymentMethod = "TUNAI"
		} else if v.TypeCard == "07" {
			paymentMethod = "QRIS"
		}
		DataPsql := &models.InquryTrx{
			DocNo:              v.DocNo,
			DocDate:            v.DocDate,
			ExtDocNo:           v.DocNo,
			CheckInDatetime:    v.CheckInDatetime,
			CheckOutDatetime:   v.CheckOutDatetime,
			CheckInTime:        v.CheckInTime,
			CheckOutTime:       v.CheckOutTime,
			DurationTime:       v.DurationTime,
			OuID:               v.OuId,
			OuCode:             v.OuCode,
			OuName:             v.OuName,
			OuSubBranchID:      v.OuSubBranchId,
			OuSubBranchCode:    v.OuSubBranchCode,
			OuSubBranchName:    v.OuSubBranchName,
			MerchantKey:        v.MerchantKey,
			ProductID:          v.TrxInvoiceItem[0].ProductId, //hardcode
			ProductCode:        v.ProductCode,
			ProductName:        v.ProductName,
			Price:              v.TrxInvoiceItem[0].Price,
			BaseTime:           int(v.TrxInvoiceItem[0].BaseTime),
			ProgressiveTime:    int(v.TrxInvoiceItem[0].ProgressiveTime),
			ProgressivePrice:   v.TrxInvoiceItem[0].ProgressivePrice,
			IsPct:              v.TrxInvoiceItem[0].IsPct,
			ProgressivePct:     int(v.TrxInvoiceItem[0].ProgressivePct),
			MaxPrice:           v.TrxInvoiceItem[0].MaxPrice,
			Is24H:              v.TrxInvoiceItem[0].Is24H,
			OvernightTime:      v.TrxInvoiceItem[0].OvernightTime,
			OvernightPrice:     v.TrxInvoiceItem[0].OvernightPrice,
			GracePeriod:        int(v.TrxInvoiceItem[0].GracePeriod),
			DropOffTime:        constans.EMPTY_VALUE_INT,
			ServiceFee:         v.TrxInvoiceItem[0].ServiceFee,
			GrandTotal:         v.GrandTotal,
			LogTrx:             v.LogTrans,
			PaymentMethod:      paymentMethod,
			Mdr:                constans.EMPTY_VALUE_INT,
			Mid:                constans.EMPTY_VALUE,
			Tid:                constans.EMPTY_VALUE,
			ResponseTrxCode:    constans.EMPTY_VALUE,
			Status:             constans.SUCCESS_CODE,
			StatusDesc:         constans.SUCCESS,
			VehicleNumberIn:    v.VehicleNumberIn,
			VehicleNumberOut:   v.VehicleNumberOut,
			ExtLocalDatetime:   v.ExtLocalDatetime,
			SettlementDatetime: &v.ExtLocalDatetime,
			DeductDatetime:     &v.ExtLocalDatetime,
			PathImageIn:        constans.EMPTY_VALUE,
			PathImageOut:       constans.EMPTY_VALUE,
			CreatedAt:          v.ExtLocalDatetime,
			CreatedBy:          "ADMIN",
			UpdatedAt:          v.ExtLocalDatetime,
			UpdatedBy:          "ADMIN",
			PaymentRefDocNo:    constans.EMPTY_VALUE,
			RefDocNo:           constans.EMPTY_VALUE,
			FlgRepeat:          v.TrxInvoiceItem[0].FlgRepeat,
		}
		utils.DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {
			id, err := svc.Service.TrxRepo.CreateTrxInquiry(DataPsql, tx)
			if err != nil {
				log.Println("Error AddTrxInquiry : ", err.Error())
				result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
				return c.JSON(http.StatusBadRequest, result)
			}
			log.Println("TRX-ID:", id)
			DataExt := &models.TrxExt{
				TrxId:          int64(id),
				BankRefNo:      constans.EMPTY_VALUE,
				CardType:       v.TypeCard,
				CardPan:        constans.EMPTY_VALUE,
				LastBalance:    constans.EMPTY_VALUE_INT,
				CurrentBalance: constans.EMPTY_VALUE_INT,
				MemberCode:     &v.MemberCode,
				MemberName:     &v.MemberName,
				MemberType:     &v.MemberType,
				CardNumberUuid: &v.CardNumberUUID,
				Username:       constans.EMPTY_VALUE,
				ShiftCode:      constans.EMPTY_VALUE,
				CreatedAt:      v.DocDate,
				CreatedBy:      "ADMIN",
				UpdatedAt:      v.DocDate,
				UpdatedBy:      "ADMIN",
			}
			_, err = svc.Service.TrxRepo.CreateTrxExt(DataExt, tx)
			if err != nil {
				log.Println("Error AddTrxExt : ", err.Error())
				result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
				return c.JSON(http.StatusBadRequest, result)
			}
			DataOu := &models.TrxOu{
				TrxID:           int64(id),
				OuID:            v.MainOuId,
				OuCode:          v.MainOuCode,
				OuName:          v.MainOuName,
				OuBranchID:      v.OuId,
				OuBranchCode:    v.OuCode,
				OuBranchName:    v.OuName,
				OuSubBranchID:   v.OuSubBranchId,
				OuSubBranchCode: v.OuSubBranchCode,
				OuSubBranchName: v.OuSubBranchName,
				CreatedBy:       "ADMIN",
				UpdatedBy:       "ADMIN",
			}
			_, err = svc.Service.TrxRepo.CreateTrxOu(DataOu, tx)
			if err != nil {
				log.Println("Error AddTrxOu : ", err.Error())
				result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
				return c.JSON(http.StatusBadRequest, result)
			}
			if err != nil {
				log.Println("Error DBTransaction:", err.Error())
				return err
			}
			return nil
		})

	}

	return c.JSON(http.StatusOK, &data)
}

func (svc trxService) GetTrx(ctx echo.Context) error {
	var result models.Response
	var payload models.TrxResponsePayload

	request := new(models.TrxFilter)
	if err := helpers.BindValidateStruct(ctx, request); err != nil {
		result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("REQUEST: ", utils.ToString(request))

	if request.CheckOutDatetimeFrom != constans.EMPTY_VALUE {
		request.CheckOutDatetimeFrom = fmt.Sprintf("%s%s", request.CheckOutDatetimeFrom, ":00")
	}

	if request.CheckOutDatetimeTo != constans.EMPTY_VALUE {
		request.CheckOutDatetimeTo = fmt.Sprintf("%s%s", request.CheckOutDatetimeTo, ":59")
	}

	listData, err := svc.Service.TrxRepo.GetTrx(*request)
	if err != nil {
		result = helpers.ResponseJSON(false, constans.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	payload.TrxResponseData = listData
	if request.Draw == 1 {
		summary, err := svc.Service.TrxRepo.GetTrxSummariesAdvance(*request)
		if err != nil {
			result = helpers.ResponseJSON(false, constans.SYSTEM_ERROR_CODE, err.Error(), nil)
			return ctx.JSON(http.StatusBadRequest, result)
		}

		payload.TrxResponseSummaries = summary
		payload.TrxResponseSummaries.TotalNett = payload.TrxResponseSummaries.GrandTotal - payload.TrxResponseSummaries.ServiceFee - payload.TrxResponseSummaries.Mdr
	}

	result = helpers.ResponseJSON(true, constans.SUCCESS_CODE, constans.EMPTY_VALUE, listData)
	return ctx.JSON(http.StatusOK, result)
}
