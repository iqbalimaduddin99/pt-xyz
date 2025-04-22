package entities

import "github.com/google/uuid"

type TransactionTable struct {
	Master
	ConsumerID          uuid.UUID  `db:"consumer_id" json:"consumerId"`
	CompanyID           int         `db:"company_id" json:"companyId"`
	ExternalTransactionID string     `db:"external_transaction_id" json:"externalTransactionId"`
	CompanyName         string      `db:"company_name" json:"companyName"`
	CompanyCategory     string      `db:"company_category" json:"companyCategory"`
	ContactNumber       string      `db:"contact_number" json:"contactNumber"`
	AdminFee            float64     `db:"admin_fee" json:"adminFee"`
	TotalPrice          float64     `db:"total_price" json:"totalPrice"`
}

type TransactionProduct struct {
	Master
	TransactionID    uuid.UUID  `db:"transaction_id" json:"transactionId"`
	CompanyID        int         `db:"company_id" json:"companyId"`
	ProductCompanyID uuid.UUID   `db:"product_company_id" json:"productCompanyId"`
	OTR              float64     `db:"otr" json:"otr"`
	AssetName        string      `db:"asset_name" json:"assetName"`
	Price            float64     `db:"price" json:"price"`
}


type TransactionTableReq struct {
	CompanyID            int                `db:"company_id" json:"companyId"`
	ExternalTransactionID string            `db:"external_transaction_id" json:"externalTransactionId"`
	CompanyName          string             `db:"company_name" json:"companyName"`
	CompanyCategory      string             `db:"company_category" json:"companyCategory"`
	ContactNumber        string             `db:"contact_number" json:"contactNumber"`
	AdminFee             float64            `db:"admin_fee" json:"adminFee"`
	TotalPrice           float64            `db:"total_price" json:"totalPrice"`
	IsExternalCompany    bool     			`json:"isExternalCompany"`
	TransactionProducts  []TransactionProduct `json:"transactionProducts"`
}

type TransactionProductReq struct {
	CompanyID        int         `db:"company_id" json:"companyId"`
	ProductCompanyID int         `db:"product_company_id" json:"productCompanyId"`
	OTR              float64     `db:"otr" json:"otr"`
	AssetName        string      `db:"asset_name" json:"assetName"`
	Price            float64     `db:"price" json:"price"`
}
