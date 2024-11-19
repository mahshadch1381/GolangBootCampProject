package dlrequest

import (
	"finalproject/db/queries/count"
	"fmt"
)

type DlDeleteRequest struct {
	ID      uint
	Version int
}

func (d *DlDeleteRequest) Validate() error {
	count, err := count.CountByDlrefrences(&d.ID)
	if err != nil {
		return nil
	}
	if *count > 0 {
		return fmt.Errorf("SL with ID %d cannot be deleted because it is referenced in %d voucher items", d.ID, *count)
	}
    return nil
}
