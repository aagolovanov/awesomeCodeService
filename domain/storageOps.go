package domain

import (
	"context"
	"fmt"
	"time"
)

func (d *Domain) saveCode(c *RequestWithCode) error {
	ctx := context.Background()

	data := map[string]string{
		"code":     fmt.Sprintf("%04d", c.Code),
		"attempts": "0",
	}

	err := d.Storage.SetData(ctx, c.RequestId, data)
	if err != nil {
		d.logg.Printf("Error while creating new K:V pair in DB: %v\n", err)
	}

	err = d.Storage.SetExpire(ctx, c.RequestId, time.Duration(d.Config.TTL)*time.Second)
	if err != nil {
		d.logg.Printf("Error while setting expiration new K:V pair in DB: %v\n", err) // maybe add val removing after that
	}

	return err
}

func (d *Domain) getAttempts() int {

}
