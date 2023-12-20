package domain

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5/pgtype"
	db "project-management/db/sqlc"
	"time"
)

type ProjectCreateRequest struct {
	UserId int `json:"userId"`
	Price  int `json:"price"`
	Paid   int `json:"paid"`
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
	UserId      int64     `json:"userId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Paid        int       `json:"paid"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}

// UC Layer and Repo Layer

type ProjectUseCase interface {
	CreateANewProject(ctx context.Context, projectCreate ProjectCreate) (ProjectResponse, error)
	ListAllProjects(ctx context.Context) ([]ProjectResponse, error)
	ListAllProjectsByUserId(ctx context.Context, userId int) ([]ProjectResponse, error)
	ListAProject(ctx context.Context, id int) (ProjectResponse, error)
	ListAProjectByUserId(ctx context.Context, userId int, projectId int) (ProjectResponse, error)
	UpdateAProjectName(ctx context.Context, updateProjectName ProjectUpdateName) (ProjectResponse, error)
	UpdateAProjectPaid(ctx context.Context, updateProjectPaid ProjectUpdatePaid) (ProjectResponse, error)
	UpdateAProjectTimeWorking(ctx context.Context, updateProjectTimeWorking ProjectUpdateTimeWorking) (ProjectResponse, error)
}

type ProjectRepository interface {
	CreateAProject(ctx context.Context, projectCreate ProjectCreate) (*db.Project, error)
	ListAllProjects(ctx context.Context) ([]db.Project, error)
	ListAllProjectsByUserId(ctx context.Context, userId int) ([]db.Project, error)
	ListAProject(ctx context.Context, id int) (*db.Project, error)
	ListAProjectByUserId(ctx context.Context, userId int, projectId int) (*db.Project, error)
	UpdateAProjectName(ctx context.Context, updateProjectName ProjectUpdateName) (*db.Project, error)
	UpdateAProjectPaid(ctx context.Context, updateProjectPaid ProjectUpdatePaid) (*db.Project, error)
	UpdateAProjectTimeWorking(ctx context.Context, updateProjectTimeWorking ProjectUpdateTimeWorking) (*db.Project, error)
}

// Utils

type ProjectCreate struct {
	UserId int `json:"userId"`
	Price  int `json:"price"`
	Paid   int `json:"paid"`
}

func (p *ProjectCreate) MapProjectCreateRequestToProjectCreate(data ProjectCreateRequest) {
	p.UserId = data.UserId
	p.Price = data.Price
	p.Paid = data.Paid
}

type ProjectUpdateName struct {
	Id          int
	UserId      int
	Name        pgtype.Text
	Description pgtype.Text
}

func (p *ProjectUpdateName) MapProjectUpdateRequestToProjectUpdate(projectId int, userId int, data ProjectUpdateNameRequest) {
	p.Id = projectId
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
	Id        int
	UserId    int
	StartTime pgtype.Timestamptz
	EndTime   pgtype.Timestamptz
}

func (p *ProjectUpdateTimeWorking) MapProjectUpdateTimeWorkingRequestToProjectUpdateTimeWorking(projectId int, userId int, data ProjectUpdateTimeWorkingRequest) error {
	p.Id = projectId
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
	Id     int
	UserId int
	Paid   pgtype.Int4
	Status db.NullProjectStatus
}

func (p *ProjectUpdatePaid) MapProjectUpdatePaidRequestToProjectUpdatePaid(projectId int, userId int, data ProjectUpdatePaidRequest) {
	p.Id = projectId
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
