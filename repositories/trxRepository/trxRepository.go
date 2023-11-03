package trxRepository

import (
	"database/sql"
	"fmt"
	"m2ps/constans"
	"m2ps/helpers"
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

func (ctx trxRepository) CreateTrxInquiry(trx *models.InquryTrx, tx *sql.Tx) (id int, err error) {

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

	if tx != nil {
		err = tx.QueryRowContext(ctx.RepoDB.Context, query,
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

func (ctx trxRepository) CreateTrxExt(trx *models.TrxExt, tx *sql.Tx) (id int, err error) {
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

	if tx != nil {
		err = tx.QueryRowContext(ctx.RepoDB.Context, query,
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

func (ctx trxRepository) CreateTrxOu(trx *models.TrxOu, tx *sql.Tx) (id int, err error) {
	query := `INSERT INTO trx_ou (
		trx_id, ou_id, ou_code, ou_name, ou_branch_id,
		ou_branch_code, ou_branch_name, ou_sub_branch_id, ou_sub_branch_code, ou_sub_branch_name,
		created_by, updated_by
	) VALUES (
		$1, $2, $3, $4, $5,
		$6, $7, $8, $9, $10,
		$11, $12
	) RETURNING trx_id`

	if tx != nil {
		err = tx.QueryRowContext(ctx.RepoDB.Context, query,
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

func trxDto(rows *sql.Rows) ([]models.TrxResponseData, error) {
	var result []models.TrxResponseData

	for rows.Next() {
		var val models.TrxResponseData
		err := rows.Scan(
			&val.Id,
			&val.DocNo,
			&val.DocDate,
			&val.ExtDocNo,
			&val.OuId,
			&val.OuCode,
			&val.OuName,
			&val.ProductId,
			&val.ProductCode,
			&val.ProductName,
			&val.ServiceFee,
			&val.GrandTotal,
			&val.PaymentMethod,
			&val.Mdr,
			&val.Status,
			&val.StatusDesc,
			&val.SettlementDatetime,
			&val.DeductDatetime,
			&val.PaymentRefDocNo,
			&val.MemberCode,
			&val.MemberName,
			&val.MemberType,
			&val.CardNumberUuid,
			&val.CardPan,
			&val.CheckingDatetime,
			&val.CheckoutDatetime,
			&val.VehicleNumberIn,
			&val.VehicleNumberOut,
			&val.ShiftCode,
			&val.Username,
			&val.Mid,
			&val.Tid,
			&val.Gate,
		)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}

	return result, nil
}

func (ctx trxRepository) GetTrxSummariesAdvance(trx models.TrxFilter) (models.TrxResponseSummaries, error) {
	var args []interface{}
	var summaries models.TrxResponseSummaries

	var query = `
		SELECT COUNT(1) as total_records,
			COALESCE(SUM(grand_total),0) AS sum_grand_total,
			COALESCE(SUM(service_fee),0) AS sum_service_fee,
			COALESCE(SUM(mdr), 0) AS sum_mdr
		FROM vw_trx_list_branch
		WHERE ou_id = ANY(?::bigint[])`

	args = append(args, fmt.Sprintf("%s%s%s", "{", trx.OuList, "}"))

	if trx.DateFrom != constans.EMPTY_VALUE {
		query += ` AND doc_date BETWEEN ? `
		args = append(args, trx.DateFrom)
		query += ` AND ? `
		args = append(args, trx.DateTo)
	}

	if trx.CheckOutDatetimeFrom != constans.EMPTY_VALUE {
		query += ` AND checkout_datetime BETWEEN ? `
		args = append(args, trx.CheckOutDatetimeFrom)
		query += ` AND ? `
		args = append(args, trx.CheckOutDatetimeTo)
	}

	if trx.Username != constans.EMPTY_VALUE {
		query += ` AND username = ? `
		args = append(args, trx.Username)
	}

	if trx.SettlementFrom != constans.EMPTY_VALUE {
		query += ` AND settlement_datetime BETWEEN ? `
		args = append(args, trx.SettlementFrom)
		query += ` AND ? `
		args = append(args, trx.SettlementTo)
	}

	if trx.PaymentMethod != constans.EMPTY_VALUE {
		query += ` AND payment_method = ? `
		args = append(args, trx.PaymentMethod)
	}

	if trx.Status != constans.EMPTY_VALUE {
		query += ` AND status = ? `
		args = append(args, trx.Status)
	}

	if trx.Keyword != constans.EMPTY_VALUE {
		query += ` AND (doc_no ILIKE ? OR ext_doc_no ILIKE ? OR payment_ref_doc_no ILIKE ? OR member_code ILIKE ? OR member_name ILIKE ? OR card_pan ILIKE ? OR card_number_uuid ILIKE ?) `
		args = append(args, "%"+trx.Keyword+"%", "%"+trx.Keyword+"%", "%"+trx.Keyword+"%", "%"+trx.Keyword+"%", "%"+trx.Keyword+"%", "%"+trx.Keyword+"%", "%"+trx.Keyword+"%")
	}

	if trx.VehicleNumber != constans.EMPTY_VALUE {
		query += ` AND (vehicle_number_in ILIKE ? OR vehicle_number_out ILIKE ?) `
		args = append(args, "%"+trx.VehicleNumber+"%", "%"+trx.VehicleNumber+"%")
	}

	newQuery := helpers.ReplaceSQL(query, "?")

	err := ctx.RepoDB.DB.QueryRowContext(ctx.RepoDB.Context, newQuery, args...).Scan(&summaries.TotalRecords, &summaries.GrandTotal, &summaries.ServiceFee, &summaries.Mdr)
	if err != nil {
		return summaries, err
	}

	return summaries, nil

}

func (ctx trxRepository) GetTrx(trx models.TrxFilter) ([]models.TrxResponseData, error) {
	var args []interface{}

	var query = `
			SELECT id, doc_no, doc_date, ext_doc_no, ou_id, ou_code, ou_name, 
				product_id, product_code, product_name, service_fee, grand_total, payment_method, mdr, 
				status, status_desc, settlement_datetime, deduct_datetime, payment_ref_doc_no,
				member_code, member_name, member_type, card_number_uuid, card_pan, checkin_datetime, checkout_datetime,
				vehicle_number_in, vehicle_number_out, shift_code, username, mid, tid, gate
			FROM vw_trx_list_branch WHERE ou_id = ANY(?::bigint[]) `

	args = append(args, fmt.Sprintf("%s%s%s", "{", trx.OuList, "}"))

	if trx.DateFrom != constans.EMPTY_VALUE {
		query += ` AND doc_date BETWEEN ? `
		args = append(args, trx.DateFrom)
		query += ` AND ? `
		args = append(args, trx.DateTo)
	}

	if trx.CheckOutDatetimeFrom != constans.EMPTY_VALUE {
		query += ` AND checkout_datetime >= ? `
		args = append(args, trx.CheckOutDatetimeFrom)
		query += ` AND checkout_datetime <= ? `
		args = append(args, trx.CheckOutDatetimeTo)
	}

	if trx.Username != constans.EMPTY_VALUE {
		query += ` AND username = ? `
		args = append(args, trx.Username)
	}

	if trx.SettlementFrom != constans.EMPTY_VALUE {
		query += ` AND settlement_datetime >= ? `
		args = append(args, trx.SettlementFrom)
		query += ` AND settlement_datetime <= ? `
		args = append(args, trx.SettlementTo)
	}

	if trx.PaymentMethod != constans.EMPTY_VALUE {
		query += ` AND payment_method = ? `
		args = append(args, trx.PaymentMethod)
	}

	if trx.Status != constans.EMPTY_VALUE {
		query += ` AND status = ? `
		args = append(args, trx.Status)
	}

	if trx.Keyword != constans.EMPTY_VALUE {
		query += ` AND (doc_no ILIKE ? OR ext_doc_no ILIKE ? OR payment_ref_doc_no ILIKE ? OR member_code ILIKE ? OR member_name ILIKE ? OR card_pan ILIKE ? OR card_number_uuid ILIKE ?) `
		args = append(args, "%"+trx.Keyword+"%", "%"+trx.Keyword+"%", "%"+trx.Keyword+"%", "%"+trx.Keyword+"%", "%"+trx.Keyword+"%", "%"+trx.Keyword+"%", "%"+trx.Keyword+"%")
	}

	if trx.VehicleNumber != constans.EMPTY_VALUE {
		query += ` AND (vehicle_number_in ILIKE ? OR vehicle_number_out ILIKE ?) `
		args = append(args, "%"+trx.VehicleNumber+"%", "%"+trx.VehicleNumber+"%")
	}

	if trx.AscDesc == "asc" {
		query += ` ORDER BY ` + trx.ColumOrderName + ` ASC `
	} else if trx.AscDesc == "desc" {
		query += ` ORDER BY ` + trx.ColumOrderName + ` DESC `
	} else {
		query += ` ORDER BY doc_date DESC, doc_no DESC `
	}

	query += ` LIMIT ? OFFSET ? `
	args = append(args, trx.Length, trx.Start)

	newQuery := helpers.ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.QueryContext(ctx.RepoDB.Context, newQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return trxDto(rows)
}
