package entities


type MasterProductPtXyz struct {
	Master
	CompanyName    string      `db:"company_name" json:"companyName"`
	CompanyCategory string     `db:"company_category" json:"companyCategory"`
	OTR            float64     `db:"otr" json:"otr"`
	AdminFee       float64     `db:"admin_fee" json:"adminFee"`
	AssetName      string      `db:"asset_name" json:"assetName"`
	Price          float64     `db:"price" json:"price"`
	Stock          int         `db:"stock" json:"stock"`
	ContactNumber  string      `db:"contact_number" json:"contactNumber"`
}