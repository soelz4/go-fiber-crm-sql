package lead

import "gorm.io/gorm"

// Lead Storage
type Store struct {
	db *gorm.DB
}

// Return New Lead Storage
func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}
