package usecases

import (
	"errors"

	"github.com/hallex-abreu/landing-page-test/api/application/repositories"
	"github.com/hallex-abreu/landing-page-test/api/domain"
)

type RegisterLeadRequest struct {
	Name  string
	Email string
	Phone string
}

func RegisterLead(registerLeadRequest RegisterLeadRequest, leadRepository repositories.LeadRepository) (*domain.Lead, error) {
	exist_lead_with_email := leadRepository.FindLeadByEmail(registerLeadRequest.Email)

	if exist_lead_with_email != nil {
		return nil, errors.New("Já existe um lead com esse email")
	}

	lead := &domain.Lead{
		Id:    0,
		Name:  registerLeadRequest.Name,
		Email: registerLeadRequest.Email,
		Phone: registerLeadRequest.Phone,
	}

	lead = leadRepository.RegisterLead(*lead)

	if lead == nil {
		return nil, errors.New("Erro ao tentar criar lead")
	}

	return lead, nil
}
