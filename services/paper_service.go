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

// GetOneByID -
func (ps PaperService) GetOneByID(id string) (res map[string]interface{}, err error) {
	raw, err := ps.paperRepository.GetPaper(id)
	if err != nil {
		return nil, err
	}
	if len(raw) == 0 {
		return nil, domain.ErrPaperIDNotFound
	}
	res = raw[0]
	return
}

// PushExamResult -
func (ps PaperService) PushExamResult(id string, rst *domain.Result) (resid string, err error) {
	return ps.paperRepository.PushExamResult(id, rst)
}

// GetExamResult -
func (ps PaperService) GetExamResult(id string, resid string) (res map[string]interface{}, err error) {
	raw, err := ps.paperRepository.GetExamResult(id, resid)
	if err != nil {
		return nil, err
	}
	if len(raw) == 0 {
		return nil, domain.ErrPaperIDNotFound
	}
	res = raw[0]
	return
}
