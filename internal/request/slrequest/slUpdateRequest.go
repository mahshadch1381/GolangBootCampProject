package slrequest
import (
	"finalproject/db/queries/count"
	"fmt"
)
type SlUpdateRequest struct {
	ID      uint
    Code    string 
    Title   string 
	Version int
	IsDetailable bool
}
func (s *SlUpdateRequest) Validate() error {
	count, err := count.CountBySlrefrences(&s.ID)
	if err != nil {
		return nil
	}
	if *count > 0 {
		return fmt.Errorf("SL with ID %d cannot be updated because it is referenced in %d voucher items", s.ID, *count)
	}
    return nil
}