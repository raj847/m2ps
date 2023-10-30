package trxService

import (
	"m2ps/constans"
	"m2ps/helpers"
	"m2ps/models"
	"m2ps/services"
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
	var tReq models.Date
	var result models.Response
	// var DataPsql *models.Trx
	if err := helpers.BindValidateStruct(c, tReq); err != nil {
		result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, result)
	}
	data, _, err := svc.Service.TrxMongoRepo.GetData(tReq.Start, tReq.End)
	if err = helpers.BindValidateStruct(c, tReq); err != nil {
		result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, result)
	}
	for _, v := range data {
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
			ProductID:          v.MemberData.ProductId,
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
			LogTrx:             constans.EMPTY_VALUE,
			PaymentMethod:      constans.EMPTY_VALUE,
			Mdr:                constans.EMPTY_VALUE_INT,
			Mid:                constans.EMPTY_VALUE,
			Tid:                constans.EMPTY_VALUE,
			ResponseTrxCode:    constans.EMPTY_VALUE,
			Status:             constans.EMPTY_VALUE,
			StatusDesc:         constans.EMPTY_VALUE,
			VehicleNumberIn:    v.VehicleNumberIn,
			VehicleNumberOut:   v.VehicleNumberOut,
			ExtLocalDatetime:   v.ExtLocalDatetime,
			SettlementDatetime: nil,
			DeductDatetime:     nil,
			PathImageIn:        constans.EMPTY_VALUE,
			PathImageOut:       constans.EMPTY_VALUE,
			CreatedAt:          v.DocDate,
			CreatedBy:          "ADMIN",
			UpdatedAt:          v.DocDate,
			UpdatedBy:          "ADMIN",
			PaymentRefDocNo:    constans.EMPTY_VALUE,
			RefDocNo:           constans.EMPTY_VALUE,
			FlgRepeat:          v.TrxInvoiceItem[0].FlgRepeat,
		}

		id, err := svc.Service.TrxRepo.CreateTrxInquiry(DataPsql)
		if err = helpers.BindValidateStruct(c, tReq); err != nil {
			result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
			return c.JSON(http.StatusBadRequest, result)
		}
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
		_, err = svc.Service.TrxRepo.CreateTrxExt(DataExt)
		if err = helpers.BindValidateStruct(c, tReq); err != nil {
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
		_, err = svc.Service.TrxRepo.CreateTrxOu(DataOu)
		if err = helpers.BindValidateStruct(c, tReq); err != nil {
			result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
			return c.JSON(http.StatusBadRequest, result)
		}
	}

	return c.JSON(http.StatusOK, &data)
}
