package domain

import "regexp"

func (d *Domain) VerifyCode(r *RequestWithCode) (*ResponseVerified, error) {
	return nil, nil
}

// wtf this func for what ?
func verifyUuid(uuid string) bool {
	matchUuid, _ := regexp.Match(`^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`, []byte(uuid))
	return matchUuid
}
