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
    UserName      string  `db:"user_name" json:"user_name"`
    Password      string  `db:"password" json:"password"`
    FullName      string  `db:"full_name" json:"full_name"`
    LegalName     string  `db:"legal_name" json:"legal_name"`
    BornLocation  string  `db:"born_location" json:"born_location"`
    BornDate      string  `db:"born_date" json:"born_date"`
    PhotoKTP      string  `db:"photo_KTP" json:"photo_KTP"`
    SelfiePhoto   string  `db:"selfie_photo" json:"selfie_photo"`
    Salary        float64 `db:"salary" json:"salary"`
}

type LoginRequest struct {
    UserName      string  `db:"user_name" json:"user_name"`
    Password      string  `db:"password" json:"password"`
    IsAdmin       bool  `json:"is_admin"`
}


