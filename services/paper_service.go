package services

import (
	"github.com/bagus2x/sirius-c/domain"
)

// PaperService -
type PaperService struct {
	paperRepository domain.PaperRepository
}

// NewPaperService -
func NewPaperService(paperRepository domain.PaperRepository) domain.PaperService {
	return &PaperService{paperRepository}
}

// Create -
func (ps PaperService) Create(p *domain.Paper) error {
	return ps.paperRepository.InsertOne(p)
}

// FindByID -
func (ps PaperService) FindByID(id string) (res *domain.Paper, err error) {
	return ps.paperRepository.FindByID(id)
}
