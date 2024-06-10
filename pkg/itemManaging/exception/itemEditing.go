package exception

import "fmt"

type ItemEditing struct {
	ItemId uint64
}

func (e *ItemEditing) Error() string {
	return fmt.Sprintf("Editing item id: %d failed", e.ItemId)
}
