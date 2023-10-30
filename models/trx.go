package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trx struct {
	Id               primitive.ObjectID      `json:"_id" bson:"_id"`
	DocNo            string                  `json:"docNo" bson:"docNo"`
	DocDate          string                  `json:"docDate" bson:"docDate"`
	CheckInDatetime  string                  `json:"checkInDatetime" bson:"checkInDatetime"`
	CheckOutDatetime string                  `json:"checkOutDatetime" bson:"checkOutDatetime"`
	DeviceIdIn       string                  `json:"deviceIdIn" bson:"deviceIdIn"`
	DeviceId         string                  `json:"device_id" bson:"deviceId"`
	GateIn           string                  `json:"gateIn" bson:"gateIn"`
	GateOut          string                  `json:"gateOut" bson:"gateOut"`
	CardNumberUUIDIn string                  `json:"cardNumberUuidIn" bson:"cardNumberUuidIn"`
	CardNumberIn     string                  `json:"cardNumberIn" bson:"cardNumberIn"`
	CardNumberUUID   string                  `json:"cardNumberUuid" bson:"cardNumberUuid"`
	CardNumber       string                  `json:"cardNumber" bson:"cardNumber"`
	TypeCard         string                  `json:"typeCard" bson:"typeCard"`
	BeginningBalance float64                 `json:"beginningBalance" bson:"beginningBalance"`
	ExtLocalDatetime string                  `json:"extLocalDatetime" bson:"extLocalDatetime"`
	GrandTotal       float64                 `json:"grandTotal" bson:"grandTotal"`
	ProductCode      string                  `json:"productCode" bson:"productCode"`
	ProductName      string                  `json:"productName" bson:"productName"`
	ProductData      string                  `json:"productData" bson:"productData"`
	RequestData      string                  `json:"requestData" bson:"requestData"`
	RequestOutData   string                  `json:"requestOutData" bson:"requestOutData"`
	OuId             int64                   `json:"ouId" bson:"ouId"`
	OuName           string                  `json:"ouName" bson:"ouName"`
	OuCode           string                  `json:"ouCode" bson:"ouCode"`
	OuSubBranchId    int64                   `json:"ouSubBranchId" bson:"ouSubBranchId"`
	OuSubBranchName  string                  `json:"ouSubBranchName" bson:"ouSubBranchName"`
	OuSubBranchCode  string                  `json:"ouSubBranchCode" bson:"ouSubBranchCode"`
	MainOuId         int64                   `json:"mainOuId" bson:"mainOuId"`
	MainOuCode       string                  `json:"mainOuCode" bson:"mainOuCode"`
	MainOuName       string                  `json:"mainOuName" bson:"mainOuName"`
	MemberCode       string                  `json:"memberCode" bson:"memberCode"`
	MemberName       string                  `json:"memberName" bson:"memberName"`
	MemberType       string                  `json:"memberType" bson:"memberType"`
	CheckInTime      int64                   `json:"checkInTime" bson:"checkInTime"`
	CheckOutTime     int64                   `json:"checkOutTime" bson:"checkOutTime"`
	DurationTime     int64                   `json:"durationTime" bson:"durationTime"`
	VehicleNumberIn  string                  `json:"vehicleNumberIn" bson:"vehicleNumberIn"`
	VehicleNumberOut string                  `json:"vehicleNumberOut" bson:"vehicleNumberOut"`
	LogTrans         string                  `json:"logTrans" bson:"logTrans"`
	MerchantKey      string                  `json:"merchantKey" bson:"merchantKey"`
	QrText           string                  `json:"qrText" bson:"qrText"`
	TrxInvoiceItem   []TrxInvoiceItem        `json:"trxInvoiceItem" bson:"trxInvoiceItem"`
	FlagSyncData     bool                    `json:"flagSyncData" bson:"flagSyncData"`
	MemberData       *TrxMember              `json:"memberData" bson:"memberData"`
	TrxAddInfo       *map[string]interface{} `json:"trxAddInfo" bson:"trxAddInfo"`
	FlagTrxFromCloud bool                    `json:"flagTrxFromCloud" bson:"flagTrxFromCloud"`
	IsRsyncDataTrx   bool                    `json:"isRsyncDataTrx" bson:"isRsyncDataTrx"`
	ExcludeSf        bool                    `json:"excludeSf" bson:"excludeSf"`
}

type TrxInvoiceItem struct {
	DocNo                  string  `json:"docNo" bson:"docNo"`
	ProductId              int64   `json:"productId" bson:"productId"`
	ProductCode            string  `json:"productCode" bson:"productCode"`
	ProductName            string  `json:"productName" bson:"productName"`
	IsPctServiceFee        string  `json:"isPctServiceFee" bson:"isPctServiceFee"`
	ServiceFee             float64 `json:"serviceFee" bson:"serviceFee"`
	ServiceFeeMember       float64 `json:"serviceFeeMember" bson:"serviceFeeMember"`
	Price                  float64 `json:"price" bson:"price"`
	BaseTime               int64   `json:"baseTime" bson:"baseTime"`
	ProgressiveTime        int64   `json:"progressiveTime" bson:"progressiveTime"`
	ProgressivePrice       float64 `json:"progressivePrice" bson:"progressivePrice"`
	IsPct                  string  `json:"isPct" bson:"isPct"`
	ProgressivePct         int64   `json:"progressivePct" bson:"progressivePct"`
	MaxPrice               float64 `json:"maxPrice" bson:"maxPrice"`
	Is24H                  string  `json:"is24H" bson:"is24H"`
	OvernightTime          string  `json:"overnightTime" bson:"overnightTime"`
	OvernightPrice         float64 `json:"overnightPrice" bson:"overnightPrice"`
	GracePeriod            int64   `json:"gracePeriod" bson:"gracePeriod"`
	FlgRepeat              string  `json:"flgRepeat" bson:"flgRepeat"`
	TotalAmount            float64 `json:"totalAmount" bson:"totalAmount"`
	TotalProgressiveAmount float64 `json:"totalProgressiveAmount" bson:"totalProgressiveAmount"`
}

type TrxMember struct {
	DocNo              string  `json:"docNo" bson:"docNo"`
	PartnerCode        string  `json:"partnerCode" bson:"partnerCode"`
	FirstName          string  `json:"firstName" bson:"firstName"`
	LastName           string  `json:"lastName" bson:"lastName"`
	RoleType           string  `json:"roleType" bson:"roleType"`
	PhoneNumber        string  `json:"phoneNumber" bson:"phoneNumber"`
	Email              string  `json:"email" bson:"email"`
	Active             string  `json:"active" bson:"active"`
	ActiveAt           string  `json:"activeAt" bson:"activeAt"`
	NonActiveAt        *string `json:"nonActiveAt" bson:"nonActiveAt"`
	OuId               int64   `json:"ouId" bson:"ouId"`
	TypePartner        string  `json:"typePartner" bson:"typePartner"`
	CardNumber         string  `json:"cardNumber" bson:"cardNumber"`
	VehicleNumber      string  `json:"vehicleNumber" bson:"vehicleNumber"`
	RegisteredDatetime string  `json:"registeredDatetime" bson:"registeredDatetime"`
	DateFrom           string  `json:"dateFrom" bson:"dateFrom"`
	DateTo             string  `json:"dateTo" bson:"dateTo"`
	ProductId          int64   `json:"productId" bson:"productId"`
	ProductCode        string  `json:"productCode" bson:"productCode"`
}

type InquryTrx struct {
	DocNo              string  `json:"docNo"`
	DocDate            string  `json:"docDate"`
	ExtDocNo           string  `json:"extDocNo"`
	IdempotencyKey     string  `json:"idempotencyKey"`
	CheckInDatetime    string  `json:"checkInDatetime"`
	CheckOutDatetime   string  `json:"checkOutDatetime"`
	CheckInTime        int64   `json:"checkInTime"`
	CheckOutTime       int64   `json:"checkOutTime"`
	DurationTime       int64   `json:"durationTime"`
	OuID               int64   `json:"ouId"`
	OuCode             string  `json:"ouCode"`
	OuName             string  `json:"ouName"`
	OuSubBranchID      int64   `json:"ouSubBranchId"`
	OuSubBranchCode    string  `json:"ouSubBranchCode"`
	OuSubBranchName    string  `json:"ouSubBranchName"`
	MerchantKey        string  `json:"merchant_key"`
	ProductID          int64   `json:"productId"`
	ProductCode        string  `json:"productCode"`
	ProductName        string  `json:"productName"`
	Price              float64 `json:"price"`
	BaseTime           int     `json:"baseTime"`
	ProgressiveTime    int     `json:"progressiveTime"`
	ProgressivePrice   float64 `json:"progressivePrice"`
	IsPct              string  `json:"isPct"`
	ProgressivePct     int     `json:"progressivePct"`
	MaxPrice           float64 `json:"maxPrice"`
	Is24H              string  `json:"is24H"`
	OvernightTime      string  `json:"overnightTime"`
	OvernightPrice     float64 `json:"overnightPrice"`
	GracePeriod        int     `json:"gracePeriod"`
	DropOffTime        int     `json:"dropOffTime"`
	ServiceFee         float64 `json:"serviceFee"`
	GrandTotal         float64 `json:"grandTotal"`
	LogTrx             string  `json:"logTrx"`
	PaymentMethod      string  `json:"paymentMethod"`
	Mdr                float64 `json:"mdr"`
	Mid                string  `json:"mid"`
	Tid                string  `json:"tid"`
	ResponseTrxCode    string  `json:"responseTrxCode"`
	Status             string  `json:"status"`
	StatusDesc         string  `json:"statusDesc"`
	VehicleNumberIn    string  `json:"vehicleNumberIn"`
	VehicleNumberOut   string  `json:"vehicleNumberOut"`
	ExtLocalDatetime   string  `json:"extLocalDatetime"`
	SettlementDatetime *string `json:"settlementDatetime"`
	DeductDatetime     *string `json:"deductDatetime"`
	PathImageIn        string  `json:"pathImageIn"`
	PathImageOut       string  `json:"pathImageOut"`
	CreatedAt          string  `json:"createdAt"`
	CreatedBy          string  `json:"createdBy"`
	UpdatedAt          string  `json:"updatedAt"`
	UpdatedBy          string  `json:"updatedBy"`
	PaymentRefDocNo    string  `json:"paymentRefDocNo"`
	RefDocNo           string  `json:"refDocNo"`
	FlgRepeat          string  `json:"flgRepeat"`
}

type Date struct {
	Start string `json:"startdate"`
	End   string `json:"enddate"`
}

type TrxExt struct {
	TrxId          int64   `json:"trxId"`
	BankRefNo      string  `json:"bankRefNo"`
	CardType       string  `json:"cardType"`
	CardPan        string  `json:"cardPan"`
	LastBalance    float64 `json:"lastBalance"`
	CurrentBalance float64 `json:"currentBalance"`
	MemberCode     *string `json:"memberCode"`
	MemberName     *string `json:"memberName"`
	MemberType     *string `json:"memberType"`
	CardNumberUuid *string `json:"cardNumberUuid"`
	Username       string  `json:"username"`
	ShiftCode      string  `json:"shiftCode"`
	CreatedAt      string  `json:"createdAt"`
	CreatedBy      string  `json:"createdBy"`
	UpdatedAt      string  `json:"updatedAt"`
	UpdatedBy      string  `json:"updatedBy"`
}

type TrxOu struct {
	TrxID           int64  `json:"trxId" validate:"required"`
	OuID            int64  `json:"ouId" validate:"required"`
	OuCode          string `json:"ouCode" validate:"required"`
	OuName          string `json:"ouName" validate:"required"`
	OuBranchID      int64  `json:"ouBranchId"`
	OuBranchCode    string `json:"ouBranchCode"`
	OuBranchName    string `json:"ouBranchName"`
	OuSubBranchID   int64  `json:"ouSubBranchId"`
	OuSubBranchCode string `json:"ouSubBranchCode"`
	OuSubBranchName string `json:"ouSubBranchName"`
	CreatedBy       string `json:"createdBy"`
	UpdatedBy       string `json:"updatedBy"`
}
