package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type User struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Email        string    `db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `db:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
	UserStatus   int       `db:"user_status" json:"user_status" validate:"required,len=1"` // 0 == blocked, 1 == active: 사용자 권한 이외에 block 상태가 되면 등록한 책을 수정 및 삭제할 수 없는데 문제가 있을 수 있다. 권한 문제로 해결?
	UserRole     string    `db:"user_role" json:"user_role" validate:"required,lte=25"`
}
