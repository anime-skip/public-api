package models

type Show struct {
	BaseModel
	Name         string
	OriginalName *string
	Website      *string
	Image        *string
}
