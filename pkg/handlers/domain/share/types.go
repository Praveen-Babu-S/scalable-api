package share

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Share struct {
	gorm.Model
	NoteID           uint
	RecipientUserID  uint
	SharingUserID    uint
	SharingTimestamp time.Time
}
