package api

import (
	"errors"
	"github.com/aagolovanov/awesomeCodeService/domain"
	//"github.com/google/uuid"
)

type MockDomain struct {
	//codes map[string]int
}

func GetMockDomain() *MockDomain {
	//
	//id0, _ := uuid.NewUUID()
	//
	//codes := map[string]int{
	//	id0.String(): 1234,
	//}

	return &MockDomain{}
}

func (m MockDomain) GenerateCode(req *domain.RequestGenerate) (*domain.RequestWithCode, error) {
	if req.Number == "internal" {
		return nil, domain.Internal
	}
	if req.Number == "custom" {
		return nil, errors.New("customError")
	}
	return &domain.RequestWithCode{}, nil
}

func (m MockDomain) VerifyCode(req *domain.RequestWithCode) (*domain.ResponseVerified, error) {
	if req.RequestId == "internal" {
		return nil, domain.Internal
	}
	if req.RequestId == "custom" {
		return nil, errors.New("customError")
	}
	return &domain.ResponseVerified{}, nil
}

var _ domain.IDomain = (*MockDomain)(nil)
