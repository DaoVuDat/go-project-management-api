package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5/pgtype"
	db "project-management/db/sqlc"
	"time"
)

type ProjectCreateRequest struct {
	UserId string `json:"userId"`
	Price  int    `json:"price"`
	Paid   int    `json:"paid"`
}

type ProjectUpdateNameRequest struct {
	Name        *string `json:"startTime,omitempty"`
	Description *string `json:"endTime,omitempty"`
}

type ProjectUpdateTimeWorkingRequest struct {
	StartTime *time.Time `json:"startTime,omitempty"`
	EndTime   *time.Time `json:"endTime,omitempty"`
}

type ProjectUpdatePaidRequest struct {
	Paid   *int              `json:"startTime,omitempty"`
	Status *db.ProjectStatus `json:"endTime,omitempty"`
}

type ProjectResponse struct {
	Id          int64     `json:"id"`
	UserId      string    `json:"userId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Paid        int       `json:"paid"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}

// UC Layer and Repo Layer

// Utils

type ProjectCreate struct {
	UserId      string    `json:"userId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Paid        int       `json:"paid"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}

func (p *ProjectCreate) MapProjectCreateRequestToProjectCreate(data ProjectCreateRequest) {
	p.UserId = data.UserId
	p.Price = data.Price
	p.Paid = data.Paid
}

type ProjectUpdateName struct {
	UserId      int
	Name        pgtype.Text
	Description pgtype.Text
}

func (p *ProjectUpdateName) MapProjectUpdateRequestToProjectUpdate(userId int, data ProjectUpdateNameRequest) {
	p.UserId = userId

	p.Name = pgtype.Text{}
	if data.Name != nil {
		p.Name.String = *data.Name
		p.Name.Valid = data.Name != nil
	}

	p.Description = pgtype.Text{}
	if data.Description != nil {
		p.Description.String = *data.Description
		p.Description.Valid = data.Description != nil
	}

}

type ProjectUpdateTimeWorking struct {
	UserId    int
	StartTime pgtype.Timestamptz
	EndTime   pgtype.Timestamptz
}

func (p *ProjectUpdateTimeWorking) MapProjectUpdateTimeWorkingRequestToProjectUpdateTimeWorking(userId int, data ProjectUpdateTimeWorkingRequest) error {
	p.UserId = userId

	if data.StartTime != nil {
		var startTime pgtype.Timestamptz
		err := startTime.Scan(data.StartTime)
		if err != nil {
			return err
		}
		p.StartTime = startTime
	} else {
		p.StartTime = pgtype.Timestamptz{
			Valid: false,
		}
	}

	if data.EndTime != nil {
		var endTime pgtype.Timestamptz
		err := endTime.Scan(data.EndTime)
		if err != nil {
			return err
		}
		p.EndTime = endTime
	} else {
		p.EndTime = pgtype.Timestamptz{
			Valid: false,
		}
	}

	return nil
}

type ProjectUpdatePaid struct {
	UserId int
	Paid   pgtype.Int4
	Status db.NullProjectStatus
}

func (p *ProjectUpdatePaid) MapProjectUpdatePaidRequestToProjectUpdatePaid(userId int, data ProjectUpdatePaidRequest) {
	p.UserId = userId
	p.Paid = pgtype.Int4{}
	if data.Paid != nil {
		p.Paid.Int32 = int32(*data.Paid)
		p.Paid.Valid = data.Paid != nil
	}

	p.Status = db.NullProjectStatus{}
	if data.Status != nil {
		p.Status.ProjectStatus = *data.Status
		p.Status.Valid = data.Status != nil
	}
}

// Setup Validations

func (req ProjectCreateRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.UserId, validation.Required),
		validation.Field(&req.Price, validation.Min(0)),
		validation.Field(&req.Paid, validation.Min(0)),
	)
}

func (req ProjectUpdateNameRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name,
			validation.When(
				req.Name != nil,
				validation.Length(1, 100),
			),
		),
		validation.Field(&req.Description,
			validation.When(
				req.Description != nil,
				validation.Length(1, 200),
			),
		),
	)
}

func (req ProjectUpdateTimeWorkingRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.StartTime,
			validation.When(
				req.StartTime != nil,
				validation.Date(req.StartTime.String()),
			),
		),
		validation.Field(&req.EndTime,
			validation.When(
				req.EndTime != nil,
				validation.Date(req.EndTime.String()),
			),
		),
	)
}

func (req ProjectUpdatePaidRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Paid,
			validation.When(
				req.Paid != nil,
				validation.Min(1),
			),
		),
		validation.Field(&req.Status,
			validation.When(
				req.Status != nil,
				validation.In(
					db.ProjectStatusRegistering,
					db.ProjectStatusProgressing,
					db.ProjectStatusFinished,
				),
			),
		),
	)
}
