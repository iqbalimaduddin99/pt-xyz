package entities

type Consumer struct {
	Master    
    KTP           string  `db:"KTP"`
    UserName      string  `db:"user_name"`
    Password      string  `db:"password"`
    FullName      string  `db:"full_name"`
    LegalName     string  `db:"legal_name"`
    BornLocation  string  `db:"born_location"`
    BornDate      string  `db:"born_date"`
    PhotoKTP      string  `db:"photo_KTP"`      
    SelfiePhoto   string  `db:"selfie_photo"`   
    Salary        float64 `db:"salary"`
}


type ReqConsumer struct {
    KTP           string  `db:"KTP" json:"KTP"`
    UserName      string  `db:"user_name" json:"userName"`
    Password      string  `db:"password" json:"password"`
    FullName      string  `db:"full_name" json:"fullName"`
    LegalName     string  `db:"legal_name" json:"legalName"`
    BornLocation  string  `db:"born_location" json:"bornLocation"`
    BornDate      string  `db:"born_date" json:"bornDate"`
    PhotoKTP      string  `db:"photo_KTP" json:"photoKTP"`
    SelfiePhoto   string  `db:"selfie_photo" json:"selfiePhoto"`
    Salary        float64 `db:"salary" json:"salary"`
}

type LoginRequest struct {
    UserName      string  `db:"user_name" json:"userName"`
    Password      string  `db:"password" json:"password"`
    IsAdmin       bool  `json:"isAdmin"`
}


