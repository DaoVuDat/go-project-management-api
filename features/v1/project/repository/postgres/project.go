package postgresproject

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"project-management/common"
	db "project-management/db/sqlc"
	"project-management/domain"
	"time"
)

type projectRepo struct {
	appContext common.AppContext
	connPool   *pgxpool.Pool
}

func NewPostgresProjectRepo(appCtx common.AppContext) domain.ProjectRepository {
	return &projectRepo{
		appContext: appCtx,
		connPool:   appCtx.Pool,
	}
}

func (repo *projectRepo) CreateAProject(ctx context.Context, projectCreate domain.ProjectCreate) (*db.Project, error) {
	queries := db.New(repo.connPool)
	project, err := queries.CreateProject(ctx, db.CreateProjectParams{
		UserProfile: pgtype.Int8{
			Int64: int64(projectCreate.UserId),
			Valid: true,
		},
		Price: int32(projectCreate.Price),
		Paid:  int32(projectCreate.Paid),
	})
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (repo *projectRepo) ListAllProjects(ctx context.Context) ([]db.Project, error) {
	queries := db.New(repo.connPool)
	projects, err := queries.ListProjects(ctx)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (repo *projectRepo) ListAllProjectsByUserId(ctx context.Context, userId int) ([]db.Project, error) {
	queries := db.New(repo.connPool)
	projects, err := queries.ListProjectsByUserId(ctx, pgtype.Int8{
		Int64: int64(userId),
		Valid: true,
	})

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (repo *projectRepo) ListAProject(ctx context.Context, id int) (*db.Project, error) {
	queries := db.New(repo.connPool)
	project, err := queries.GetProject(ctx, int64(id))

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (repo *projectRepo) ListAProjectByUserId(ctx context.Context, userId int, projectId int) (*db.Project, error) {
	queries := db.New(repo.connPool)
	project, err := queries.GetAProjectByUserId(ctx, db.GetAProjectByUserIdParams{
		ID: int64(projectId),
		UserProfile: pgtype.Int8{
			Int64: int64(userId),
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (repo *projectRepo) UpdateAProjectName(ctx context.Context, updateProjectName domain.ProjectUpdateName) (*db.Project, error) {
	queries := db.New(repo.connPool)
	project, err := queries.UpdateProjectName(ctx, db.UpdateProjectNameParams{
		ID: int64(updateProjectName.Id),
		UserProfile: pgtype.Int8{
			Int64: int64(updateProjectName.UserId),
			Valid: true,
		},
		UpdatedAt:   time.Now(),
		Name:        updateProjectName.Name,
		Description: updateProjectName.Description,
	})
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (repo *projectRepo) UpdateAProjectPaid(ctx context.Context, updateProjectPaid domain.ProjectUpdatePaid) (*db.Project, error) {
	queries := db.New(repo.connPool)
	project, err := queries.UpdateProjectPaid(ctx, db.UpdateProjectPaidParams{
		ID: int64(updateProjectPaid.Id),
		UserProfile: pgtype.Int8{
			Int64: int64(updateProjectPaid.Id),
			Valid: true,
		},
		UpdatedAt: time.Now(),
		Paid:      updateProjectPaid.Paid,
		Status:    updateProjectPaid.Status,
	})

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (repo *projectRepo) UpdateAProjectTimeWorking(ctx context.Context, updateProjectTimeWorking domain.ProjectUpdateTimeWorking) (*db.Project, error) {
	queries := db.New(repo.connPool)
	project, err := queries.UpdateProjectTimeWorking(ctx, db.UpdateProjectTimeWorkingParams{
		ID: int64(updateProjectTimeWorking.Id),
		UserProfile: pgtype.Int8{
			Int64: int64(updateProjectTimeWorking.UserId),
			Valid: true,
		},
		UpdatedAt: time.Now(),
		StartTime: updateProjectTimeWorking.StartTime,
		EndTime:   updateProjectTimeWorking.EndTime,
	})

	if err != nil {
		return nil, err
	}

	return &project, nil
}
