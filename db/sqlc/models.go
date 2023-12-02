// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type AccountStatus string

const (
	AccountStatusPending   AccountStatus = "pending"
	AccountStatusActivated AccountStatus = "activated"
)

func (e *AccountStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccountStatus(s)
	case string:
		*e = AccountStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for AccountStatus: %T", src)
	}
	return nil
}

type NullAccountStatus struct {
	AccountStatus AccountStatus `json:"account_status"`
	Valid         bool          `json:"valid"` // Valid is true if AccountStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccountStatus) Scan(value interface{}) error {
	if value == nil {
		ns.AccountStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AccountStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccountStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AccountStatus), nil
}

type AccountType string

const (
	AccountTypeClient AccountType = "client"
	AccountTypeAdmin  AccountType = "admin"
)

func (e *AccountType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccountType(s)
	case string:
		*e = AccountType(s)
	default:
		return fmt.Errorf("unsupported scan type for AccountType: %T", src)
	}
	return nil
}

type NullAccountType struct {
	AccountType AccountType `json:"account_type"`
	Valid       bool        `json:"valid"` // Valid is true if AccountType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccountType) Scan(value interface{}) error {
	if value == nil {
		ns.AccountType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AccountType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccountType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AccountType), nil
}

type ProjectStatus string

const (
	ProjectStatusRegistering ProjectStatus = "registering"
	ProjectStatusProgressing ProjectStatus = "progressing"
	ProjectStatusFinished    ProjectStatus = "finished"
)

func (e *ProjectStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProjectStatus(s)
	case string:
		*e = ProjectStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ProjectStatus: %T", src)
	}
	return nil
}

type NullProjectStatus struct {
	ProjectStatus ProjectStatus `json:"project_status"`
	Valid         bool          `json:"valid"` // Valid is true if ProjectStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullProjectStatus) Scan(value interface{}) error {
	if value == nil {
		ns.ProjectStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ProjectStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullProjectStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ProjectStatus), nil
}

type Project struct {
	ID          int32       `json:"id"`
	UserProfile pgtype.Int4 `json:"user_profile"`
	Name        string      `json:"name"`
	// Which technologies and algorithms used
	Description string `json:"description"`
	// Price of the project
	Price int32 `json:"price"`
	// How much money client paid
	Paid      int32            `json:"paid"`
	Status    ProjectStatus    `json:"status"`
	StartTime pgtype.Timestamp `json:"start_time"`
	EndTime   pgtype.Timestamp `json:"end_time"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type UserAccount struct {
	UserID    int32            `json:"user_id"`
	Username  string           `json:"username"`
	Password  string           `json:"password"`
	Type      AccountType      `json:"type"`
	Status    AccountStatus    `json:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type UserProfile struct {
	ID        int32            `json:"id"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
