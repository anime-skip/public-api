package entities

type Show struct {
	BaseEntity
	Name         string
	OriginalName *string
	Website      *string
	Image        *string
}
