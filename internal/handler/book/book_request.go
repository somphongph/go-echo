package book

type bookRequest struct {
	Isbn    string `json:"isbn"`
	NameTh  string `json:"nameTh"`
	NameEn  string `json:"nameEn"`
	TitleTh string `json:"titleTh"`
	TitleEn string `json:"titleEn"`
}
