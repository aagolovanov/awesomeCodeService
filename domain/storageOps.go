package domain

import (
	"context"
	"strconv"
	"time"
)

func (d *Domain) saveCode(c *RequestWithCode) error {
	ctx := context.Background()

	data := map[string]string{
		"code":     strconv.Itoa(c.Code),
		"attempts": "0",
	}

	err := d.Storage.SetData(ctx, c.RequestId, data)
	if err != nil {
		d.Logg.Printf("Error while creating new K:V pair in DB: %v\n", err)
	}

	err = d.Storage.SetExpire(ctx, c.RequestId, time.Duration(d.Config.TTL)*time.Second)
	if err != nil {
		d.Logg.Printf("Error while setting expiration new K:V pair in DB: %v\n", err) // maybe add val removing after that
	}

	return err
}

func (d *Domain) getAttempts(c *RequestWithCode) (ret int, err error) {
	ctx := context.Background()
	data, err := d.Storage.GetAllData(ctx, c.RequestId)
	if err != nil {
		d.Logg.Printf("Error while getting attempts: %v", err)
		return 0, err
	}

	ret, _ = strconv.Atoi(data["attempts"])
	return
}

func (d *Domain) getCode(c *RequestWithCode) (code int, err error) {
	ctx := context.Background()
	data, err := d.Storage.GetAllData(ctx, c.RequestId)
	if err != nil {
		d.Logg.Printf("Error while getting attempts: %v", err)
		return 0, err
	}

	code, _ = strconv.Atoi(data["code"])
	return
}
