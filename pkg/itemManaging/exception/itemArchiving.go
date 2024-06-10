package exception

import "fmt"

type ItemArchiving struct {
	ItemID uint64
}

func (e *ItemArchiving) Error() string {
	return fmt.Sprintf("Archving item id: %d failed", e.ItemID)
}
