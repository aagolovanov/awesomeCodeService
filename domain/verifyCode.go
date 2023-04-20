package domain

import (
	"context"
	"regexp"
	"time"
)

func (d *Domain) VerifyCode(r *RequestWithCode) (*ResponseVerified, error) {
	/*
		1. check code existence
		2. check attemtps
		3. Check code and ++attempts if not eq
	*/
	if !verifyUuid(r.RequestId) {
		return nil, badUUIDError
	}
	ctx := context.Background()

	if !d.Storage.CheckExist(ctx, r.RequestId) {
		return nil, requestNotExistError
	}

	attempts, err := d.getAttempts(r)
	if err != nil {
		return nil, internal
	}

	if attempts > 3 {
		return nil, attemptsExceededError
	}

	code, err := d.getCode(r)
	if err != nil {
		return nil, internal
	}

	if code != r.Code {
		_ = d.Storage.Increment(ctx, r.RequestId, "attempts")
		return nil, invalidCodeError
	}

	verifiedAt := time.Now().Unix()
	resp := &ResponseVerified{
		VerifiedAt: verifiedAt,
	}

	_ = d.Storage.Delete(ctx, r.RequestId)

	return resp, nil
}

// wtf this func for what ?
func verifyUuid(uuid string) bool {
	matchUuid, _ := regexp.Match(`^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`, []byte(uuid))
	return matchUuid
}
