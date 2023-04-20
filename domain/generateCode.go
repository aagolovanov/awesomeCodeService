package domain

import (
	"github.com/google/uuid"
	"math/rand"
	"regexp"
)

func (d *Domain) GenerateCode(r *RequestGenerate) (*RequestWithCode, error) {

	if !verifyNumber(r.Number) {
		return nil, badNumberError
	}

	requestId, err := uuid.NewUUID()
	if err != nil {
		return nil, internal
	}

	code := generateRandomCode()

	request := &RequestWithCode{
		RequestId: requestId.String(),
		Code:      code,
	}

	err = d.saveCode(request)
	if err != nil {
		return nil, internal
	}

	return request, nil
}

func verifyNumber(number string) bool {
	result, _ := regexp.Match(`^(\+7|7|8)?[\s\-]?\(?[489][0-9]{2}\)?[\s\-]?[0-9]{3}[\s\-]?[0-9]{2}[\s\-]?[0-9]{2}$`, []byte(number))
	return result // todo handle error correctly
}

func generateRandomCode() int {
	return rand.Intn(9999-1000) + 1000
}
