package usecases

import (
	"errors"

	"github.com/hallex-abreu/landing-page-test/api/application/repositories"
	"github.com/hallex-abreu/landing-page-test/api/domain"
)

func FindAllLeadsUseCase(leadRepository repositories.LeadRepository) ([]*domain.Lead, error) {
	leads, err := leadRepository.FindAllLeads()

	if err != nil {
		return nil, errors.New("Erro ao tentar listar leads")
	}

	return leads, nil
}
