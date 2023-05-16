package repositories

import "github.com/hallex-abreu/landing-page-test/api/domain"

type LeadRepository interface {
	FindAllLeads() ([]*domain.Lead, error)
	FindLeadByEmail(email string) *domain.Lead
	RegisterLead(lead domain.Lead) *domain.Lead
}
