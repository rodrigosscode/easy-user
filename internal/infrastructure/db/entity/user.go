package db

type (
	User struct {
		ID    string
		Name  string
		Email string
		Age   int
	}
)

// func (p *Pedido) BeforeCreate(tx *gorm.DB) error {
// 	uuid := uuid.New().String()
// 	p.Id = uuid
// 	return nil
// }
