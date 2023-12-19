package postgresproject

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"project-management/common"
	db "project-management/db/sqlc"
	"project-management/domain"
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

func (repo *projectRepo) ListAllProjects(ctx context.Context) ([]*db.Project, error) {
	panic(1)
}

func (repo *projectRepo) ListAllProjectsByUserId(ctx context.Context, userId int) ([]*db.Project, error) {
	panic(1)
}

func (repo *projectRepo) ListAProject(ctx context.Context, id int) (*db.Project, error) {
	panic(1)
}

func (repo *projectRepo) ListAProjectByUserId(ctx context.Context, userId int, projectId int) (*db.Project, error) {
	panic(1)
}

func (repo *projectRepo) UpdateAProjectName(ctx context.Context, updateProjectName domain.ProjectUpdateName) (*db.Project, error) {
	panic(1)
}

func (repo *projectRepo) UpdateAProjectPaid(ctx context.Context, updateProjectPaid domain.ProjectUpdatePaid) (*db.Project, error) {
	panic(1)
}

func (repo *projectRepo) UpdateAProjectTimeWorking(ctx context.Context, updateProjectTimeWorking domain.ProjectUpdateTimeWorking) (*db.Project, error) {
	panic(1)
}
