package slrequest
import (
	"finalproject/db/queries/count"
	"fmt"
)
type SlDeleteRequest struct {
    ID      uint
    Version int 
    
}
func (s *SlDeleteRequest) Validate() error {
	count, err := count.CountBySlrefrences(&s.ID)
	if err != nil {
		return nil
	}
	if *count > 0 {
		return fmt.Errorf("SL with ID %d cannot be deleted because it is referenced in %d voucher items", s.ID, *count)
	}
    return nil
}
