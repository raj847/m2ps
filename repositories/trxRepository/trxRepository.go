package trxRepository

import (
	"m2ps/models"
	"m2ps/repositories"
)

type trxRepository struct {
	RepoDB repositories.Repository
}

func NewTrxRepository(repoDB repositories.Repository) trxRepository {
	return trxRepository{
		RepoDB: repoDB,
	}
}

func (ctx trxRepository) CreateTrxInquiry(trx *models.InquryTrx) (id int, err error) {

	query := `INSERT INTO trx (
		doc_no, doc_date, ext_doc_no,checkin_datetime,
		checkout_datetime, checkin_time, checkout_time, duration_time, ou_id,
		ou_code, ou_name, ou_sub_branch_id, ou_sub_branch_code, ou_sub_branch_name,
		merchant_key, product_id, product_code, product_name, price,
		base_time, progressive_time, progressive_price, is_pct, progressive_pct,
		max_price, is_24h, overnight_time, overnight_price, grace_period,
		drop_off_time, service_fee, grand_total, log_trans, payment_method,
		mdr, mid, tid, response_trx_code, status,
		status_desc, vehicle_number_in, vehicle_number_out, ext_local_datetime, settlement_datetime,
		deduct_datetime, path_image_in, path_image_out, created_at, created_by,
		updated_at, updated_by, payment_ref_doc_no, ref_doc_no, flg_sync_ops
	) VALUES (
		$1, $2, $3, $4, $5,
		$6, $7, $8, $9, $10,
		$11, $12, $13, $14, $15,
		$16, $17, $18, $19, $20,
		$21, $22, $23, $24, $25,
		$26, $27, $28, $29, $30,
		$31, $32, $33, $34, $35,
		$36, $37, $38, $39, $40,
		$41, $42, $43, $44, $45,
		$46, $47, $48, $49, $50,
		$51,$52,$53,$54
	) RETURNING id`

	if ctx.RepoDB.DB != nil {
		err = ctx.RepoDB.DB.QueryRowContext(ctx.RepoDB.Context, query,
			trx.DocNo, trx.DocDate, trx.ExtDocNo, trx.CheckInDatetime,
			trx.CheckOutDatetime, trx.CheckInTime, trx.CheckOutTime, trx.DurationTime, trx.OuID,
			trx.OuCode, trx.OuName, trx.OuSubBranchID, trx.OuSubBranchCode, trx.OuSubBranchName,
			trx.MerchantKey, trx.ProductID, trx.ProductCode, trx.ProductName, trx.Price,
			trx.BaseTime, trx.ProgressiveTime, trx.ProgressivePrice, trx.IsPct, trx.ProgressivePct,
			trx.MaxPrice, trx.Is24H, trx.OvernightTime, trx.OvernightPrice, trx.GracePeriod,
			trx.DropOffTime, trx.ServiceFee, trx.GrandTotal, trx.LogTrx, trx.PaymentMethod,
			trx.Mdr, trx.Mid, trx.Tid, trx.ResponseTrxCode, trx.Status,
			trx.StatusDesc, trx.VehicleNumberIn, trx.VehicleNumberOut, trx.ExtLocalDatetime, trx.SettlementDatetime,
			trx.DeductDatetime, trx.PathImageIn, trx.PathImageOut, trx.CreatedAt, trx.CreatedBy,
			trx.UpdatedAt, trx.UpdatedBy, trx.PaymentRefDocNo, trx.RefDocNo, trx.FlgRepeat).Scan(&id)
	} else {
		err = ctx.RepoDB.DB.QueryRowContext(ctx.RepoDB.Context, query,
			trx.DocNo, trx.DocDate, trx.ExtDocNo, trx.IdempotencyKey, trx.CheckInDatetime,
			trx.CheckOutDatetime, trx.CheckInTime, trx.CheckOutTime, trx.DurationTime, trx.OuID,
			trx.OuCode, trx.OuName, trx.OuSubBranchID, trx.OuSubBranchCode, trx.OuSubBranchName,
			trx.MerchantKey, trx.ProductID, trx.ProductCode, trx.ProductName, trx.Price,
			trx.BaseTime, trx.ProgressiveTime, trx.ProgressivePrice, trx.IsPct, trx.ProgressivePct,
			trx.MaxPrice, trx.Is24H, trx.OvernightTime, trx.OvernightPrice, trx.GracePeriod,
			trx.DropOffTime, trx.ServiceFee, trx.GrandTotal, trx.LogTrx, trx.PaymentMethod,
			trx.Mdr, trx.Mid, trx.Tid, trx.ResponseTrxCode, trx.Status,
			trx.StatusDesc, trx.VehicleNumberIn, trx.VehicleNumberOut, trx.ExtLocalDatetime, trx.SettlementDatetime,
			trx.DeductDatetime, trx.PathImageIn, trx.PathImageOut, trx.CreatedAt, trx.CreatedBy,
			trx.UpdatedAt, trx.UpdatedBy, trx.PaymentRefDocNo, trx.RefDocNo, trx.FlgRepeat).Scan(&id)
	}

	if err != nil {
		return id, err
	}

	return id, nil
}

func (ctx trxRepository) CreateTrxExt(trx *models.TrxExt) (id int, err error) {
	query := `INSERT INTO trx_ext (
		trx_id, bank_ref_no, card_type, card_pan, last_balance,
		current_balance, member_code, member_name, member_type, card_number_uuid,
		username, shift_code, created_at, created_by, updated_at, 
		updated_by
	) VALUES (
		$1, $2, $3, $4, $5,
		$6, $7, $8, $9, $10,
		$11, $12, $13, $14, $15,
		$16
	) RETURNING trx_id`

	if ctx.RepoDB.DB != nil {
		err = ctx.RepoDB.DB.QueryRowContext(ctx.RepoDB.Context, query,
			trx.TrxId, trx.BankRefNo, trx.CardType, trx.CardPan, trx.LastBalance,
			trx.CurrentBalance, trx.MemberCode, trx.MemberName, trx.MemberType, trx.CardNumberUuid,
			trx.Username, trx.ShiftCode, trx.CreatedAt, trx.CreatedBy, trx.UpdatedAt,
			trx.UpdatedBy,
		).Scan(&id)
	} else {
		err = ctx.RepoDB.DB.QueryRowContext(ctx.RepoDB.Context, query,
			trx.TrxId, trx.BankRefNo, trx.CardType, trx.CardPan, trx.LastBalance,
			trx.CurrentBalance, trx.MemberCode, trx.MemberName, trx.MemberType, trx.CardNumberUuid,
			trx.Username, trx.ShiftCode, trx.CreatedAt, trx.CreatedBy, trx.UpdatedAt,
			trx.UpdatedBy,
		).Scan(&id)
	}

	if err != nil {
		return id, err
	}

	return id, nil
}

func (ctx trxRepository) CreateTrxOu(trx *models.TrxOu) (id int, err error) {
	query := `INSERT INTO trx_ou (
		trx_id, ou_id, ou_code, ou_name, ou_branch_id,
		ou_branch_code, ou_branch_name, ou_sub_branch_id, ou_sub_branch_code, ou_sub_branch_name,
		created_by, updated_by
	) VALUES (
		$1, $2, $3, $4, $5,
		$6, $7, $8, $9, $10,
		$11, $12
	) RETURNING trx_id`

	if ctx.RepoDB.DB != nil {
		err = ctx.RepoDB.DB.QueryRowContext(ctx.RepoDB.Context, query,
			trx.TrxID, trx.OuID, trx.OuCode, trx.OuName, trx.OuBranchID,
			trx.OuBranchCode, trx.OuBranchName, trx.OuSubBranchID, trx.OuSubBranchCode, trx.OuSubBranchName,
			trx.CreatedBy, trx.UpdatedBy,
		).Scan(&id)
	} else {
		err = ctx.RepoDB.DB.QueryRowContext(ctx.RepoDB.Context, query,
			trx.TrxID, trx.OuID, trx.OuCode, trx.OuName, trx.OuBranchID,
			trx.OuBranchCode, trx.OuBranchName, trx.OuSubBranchID, trx.OuSubBranchCode, trx.OuSubBranchName,
			trx.CreatedBy, trx.UpdatedBy,
		).Scan(&id)
	}

	if err != nil {
		return id, err
	}

	return id, nil
}
