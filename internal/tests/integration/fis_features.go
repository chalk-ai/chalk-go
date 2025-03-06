package integration

import (
	"github.com/chalk-ai/chalk-go"
	"time"
)

var InitFeaturesErr error

type ConfirmedFraud struct {
	Id                *string
	FiTransactionId   *string
	TrnDt             *time.Time
	LstUpdDt          *time.Time
	IsFraud           *int64
	Txn               *Transaction
	TerminalId        *string
	Terminal          *Terminal
	TrnAmt            *float64
	CustomerXidHash   *string
	MerId             *string
	Merchant          *Merchant
	MerNm             *string
	MerIdNm           *string
	SicCd             *string
	Sic               *SIC
	CustomerMerchant  *CustomerMerchant
	CustomerMerIdNm   *string
	CustomerSicId     *string
	CustomerSic       *CustomerSIC
	DecisionXcd       *string
	TrnAuthPost       *string
	DecisionXcdFilter *bool
}

type Customer struct {
	Id                             *string
	Transaction                    *[]Transaction    `has_many:"id,customer_xid_hash"`
	ConfirmedFraud                 *[]ConfirmedFraud `has_many:"id,customer_xid_hash"`
	NonMonetary                    *[]NonMonetary    `has_many:"id,customer_xid_hash"`
	IsFrdAcct                      *int64
	FirstFrdDt                     *time.Time
	LastFrdDt                      *time.Time
	WindowedCustomerTxnCount       map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedCustomerTxnAmtAvg      map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerTxnAmtSum      map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerTxnAmtStddev   map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerTxnAmtMin      map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerTxnAmtMax      map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerFrdTxnCount    map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedCustomerFrdTxnAmtSum   map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedNmonSystemMaintenance  map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedNmonCustomerInfoChange map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedNmonAccountInfoChange  map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedNmonPanInfoChange      map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedNmonOtherInfoChange    map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedNmonRecentCount        map[string]*int64   `windows:"1d,2d,1w,30d"`
}

type CustomerMerchant struct {
	Id                                   *string
	Transaction                          *[]Transaction      `has_many:"id,customer_mer_id_nm"`
	ConfirmedFraud                       *[]ConfirmedFraud   `has_many:"id,customer_mer_id_nm"`
	WindowedCustomerMerchantTxnCount     map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedCustomerMerchantTxnAmtAvg    map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerMerchantTxnAmtSum    map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerMerchantTxnAmtMin    map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerMerchantTxnAmtMax    map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerMerchantFrdTxnCount  map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedCustomerMerchantFrdTxnAmtSum map[string]*float64 `windows:"1d,2d,1w,30d"`
}

type CustomerPointOfSale struct {
	Id                                   *string
	Transaction                          *[]Transaction      `has_many:"id,customer_pos_id"`
	WindowedCustomerPointOfSaleTxnCount  map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedCustomerPointOfSaleTxnAmtSum map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerPointOfSaleTxnAmtAvg map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerPointOfSaleTxnAmtMax map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerPointOfSaleTxnAmtMin map[string]*float64 `windows:"1d,2d,1w,30d"`
}

type CustomerSIC struct {
	Id                              *string
	Transaction                     *[]Transaction      `has_many:"id,customer_sic_id"`
	ConfirmedFraud                  *[]ConfirmedFraud   `has_many:"id,customer_sic_id"`
	WindowedCustomerSicTxnCount     map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedCustomerSicTxnAmtSum    map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerSicTxnAmtAvg    map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerSicTxnAmtMin    map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerSicTxnAmtMax    map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerSicFrdTxnCount  map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedCustomerSicFrdTxnAmtSum map[string]*float64 `windows:"1d,2d,1w,30d"`
}

type CustomerTrnAmtRepeated struct {
	Id                                     *string
	Transaction                            *[]Transaction    `has_many:"id,customer_trn_amt_repeated_id"`
	WindowedCustomerTrnAmtRepeatedTxnCount map[string]*int64 `windows:"1d,2d,1w,30d"`
}

type CustomerTrnType struct {
	Id                               *string
	Transaction                      *[]Transaction      `has_many:"id,customer_trn_typ_id"`
	WindowedCustomerTrnTypeTxnCount  map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedCustomerTrnTypeTxnAmtSum map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerTrnTypeTxnAmtAvg map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerTrnTypeTxnAmtMax map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerTrnTypeTxnAmtMin map[string]*float64 `windows:"1d,2d,1w,30d"`
}

type InputTransactions struct {
	Id               *string
	FiTransactionId  *string
	MerId            *string
	CustomerXidHash  *string
	TerminalId       *string
	TrnDt            *time.Time
	CustomerMerId    *string
	CustomerPosId    *string
	CustomerTrnTypId *string
	IsFraud          *int64
	Dummy            *string
}

type Merchant struct {
	Id                                 *string
	Transaction                        *[]Transaction      `has_many:"id,mer_id_nm"`
	ConfirmedFraud                     *[]ConfirmedFraud   `has_many:"id,mer_id_nm"`
	WindowedMerchantTxnCount           map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedMerchantTxnAmtAvg          map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedMerchantTxnAmtSum          map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedMerchantTxnAmtStddev       map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedMerchantCardholderCount    map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedMerchantFrdTxnCount        map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedMerchantFrdTxnAmtSum       map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedMerchantFrdCardholderCount map[string]*int64   `windows:"1d,2d,1w,30d"`
}

type NonMonetary struct {
	Id              *string
	TransactionDttm *time.Time
	CustomerXidHash *string
	NonmonEventXcd  *string
	Kind            *string
}

type SIC struct {
	Id                            *string
	Transaction                   *[]Transaction      `has_many:"id,sic_cd"`
	ConfirmedFraud                *[]ConfirmedFraud   `has_many:"id,sic_cd"`
	WindowedSicTxnCount           map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedSicTxnAmtAvg          map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedSicTxnAmtSum          map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedSicTxnAmtStddev       map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedSicCardholderCount    map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedSicFrdTxnCount        map[string]*int64   `windows:"1d,2d,1w,30d"`
	WindowedSicFrdTxnAmtSum       map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedSicFrdCardholderCount map[string]*float64 `windows:"1d,2d,1w,30d"`
}

type Terminal struct {
	Id *string
}

type TerminalPointOfSale struct {
	Id *string
}

type Transaction struct {
	FiTransactionId                    *string
	TrnAmt                             *float64
	TrnAmtRound                        *float64
	TrnAmtDiscrete                     *string
	TerminalId                         *string
	Terminal                           *Terminal
	TrnDt                              *time.Time
	MerId                              *string
	MerNm                              *string
	Mer                                *Merchant
	MerIdNm                            *string
	CustomerXidHash                    *string
	Customer                           *Customer
	CustomerMerchant                   *CustomerMerchant
	CustomerMerIdNm                    *string
	SicCd                              *string
	Sic                                *SIC
	FrdScor                            *float64
	CrdExpDt                           *time.Time
	CashbackAmt                        *float64
	CusIncome                          *float64
	CusZip5                            *string
	MerZip5                            *string
	IsFraud                            *int64
	Bin                                *string
	UsrInd5                            *string
	ExtScor1                           *string `name:"ext_scor1"`
	UsrDat2                            *string
	TrnTyp                             *string
	TrnPosEntCd                        *string
	PisIssueDt                         *time.Time
	PisPmtIdCount                      *float64
	DecisionXcd                        *string
	DecisionXcdFilter                  *bool
	TrnAuthPost                        *string
	McVisaIdx                          *int64
	McVisaScr                          *int64
	McVisaScrNorm                      *int64
	TrnDayOfWeek                       *int64
	TrnDayOfMonth                      *int64
	TrnHourOfDay                       *int64
	TrnMonthOfYear                     *int64
	TrnUnixSec                         *int64
	TrnHourOfWeek                      *int64
	TrnHourOfMonth                     *int64
	CrdExpUnixSec                      *int64
	CrdExpDeltaSec                     *int64
	PisIssueUnixSec                    *int64
	PisIssueDeltaSec                   *int64
	CustomerPosId                      *string
	CustomerPointOfSale                *CustomerPointOfSale
	TerminalPosId                      *string
	TerminalPointOfSale                *TerminalPointOfSale
	CustomerTrnTypId                   *string
	CustomerTrnType                    *CustomerTrnType
	CustomerSicId                      *string
	CustomerSic                        *CustomerSIC
	AcquirerCntry                      *string
	AcquirerCntryFrdRate               *float64
	AcquirerId                         *string
	AcquirerIdFrdRate                  *float64
	AtcCrd                             *string
	AtcCrdFrdRate                      *float64
	AtcHost                            *string
	AtcHostFrdRate                     *float64
	AtmNetworkId                       *string
	AtmNetworkIdFrdRate                *float64
	AtmProcessingXcd                   *string
	AtmProcessingXcdFrdRate            *float64
	AutRespCd                          *string
	BinFrdRate                         *float64
	CardVerifResults                   *string
	CardVerifResultsFrdRate            *float64
	CrdPort                            *string
	CrdPortFrdRate                     *float64
	CrdPsntInd                         *string
	CrdPsntIndFrdRate                  *float64
	CryptogramValid                    *string
	CryptogramValidFrdRate             *float64
	CusCardTyp                         *string
	CusCardTypFrdRate                  *float64
	CusCntyCd                          *string
	CusCntyCdFrdRate                   *float64
	CusState                           *string
	CusStateFrdRate                    *float64
	CusZip3                            *string
	CusZip3FrdRate                     *float64
	MerCntyCd                          *string
	MerCntyCdFrdRate                   *float64
	MerSt                              *string
	MerStFrdRate                       *float64
	MerStSearch                        *string
	MerStSearchFrdRate                 *float64
	MerZip3                            *string
	MerZip3FrdRate                     *float64
	PadResponse                        *string
	PadResponseFrdRate                 *float64
	PisCardCagetory                    *string
	PisCardCategoryFrdRate             *float64
	PosConditionCd                     *string
	PosConditionCdFrdRate              *float64
	ProcessReasonCd                    *string
	ProcessReasonCdFrdRate             *float64
	RltmReq                            *string
	RltmReqFrdRate                     *float64
	TransCategory                      *string
	TransCategoryFrdRate               *float64
	TrnCvvVrfyCd                       *string
	TrnCvvVrfyCdFrdRate                *float64
	TrnPinVrfyCd                       *string
	TrnPinVrfyCdFrdRate                *float64
	TrnPosEntCdFrdRate                 *float64
	TrnTypFrdRate                      *float64
	CusLat                             *float64
	CusLon                             *float64
	MerLat                             *float64
	MerLon                             *float64
	CustomerMerchantDistance           *float64
	SicCdFrdRate                       *float64
	CustomerTrnAmtRepeatedId           *string
	CustomerTrnAmtRepeated             *CustomerTrnAmtRepeated
	WindowedCustomerMerchantTxnAmtNorm map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerSicTxnAmtNorm      map[string]*float64 `windows:"1d,2d,1w,30d"`
	WindowedCustomerTxnAmtZscore       map[string]*float64 `windows:"1d,2d,1w,30d"`
	Id                                 *string
	CrdExpDelta                        *int64
	CustomerXidEnc                     *string
	McVisaFlg                          *int64
}

var FISFeatures struct {
	ConfirmedFraud         *ConfirmedFraud
	Customer               *Customer
	CustomerMerchant       *CustomerMerchant
	CustomerPointOfSale    *CustomerPointOfSale
	CustomerSic            *CustomerSIC
	CustomerTrnAmtRepeated *CustomerTrnAmtRepeated
	CustomerTrnType        *CustomerTrnType
	InputTransactions      *InputTransactions
	Merchant               *Merchant
	NonMonetary            *NonMonetary
	Sic                    *SIC
	Terminal               *Terminal
	TerminalPointOfSale    *TerminalPointOfSale
	Transaction            *Transaction
}

func init() {
	InitFeaturesErr = chalk.InitFeatures(&FISFeatures)
}
