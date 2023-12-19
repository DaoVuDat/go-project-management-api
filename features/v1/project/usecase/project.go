package usecaseproject

import (
	"context"
	"project-management/common"
	"project-management/domain"
)

type projectUC struct {
	appContext        common.AppContext
	projectRepository domain.ProjectRepository
}

func NewProjectUseCase(appCtx common.AppContext, projectRepository domain.ProjectRepository) domain.ProjectUseCase {
	return &projectUC{
		appContext:        appCtx,
		projectRepository: projectRepository,
	}
}

func (p *projectUC) CreateANewProject(ctx context.Context, projectCreate domain.ProjectCreate) (domain.ProjectResponse, error) {
	project, err := p.projectRepository.CreateAProject(ctx, projectCreate)
	if err != nil {
		return domain.ProjectResponse{}, err
	}

	projectResponse := domain.ProjectResponse{
		Id:          project.ID,
		UserId:      project.UserProfile.Int64,
		Name:        project.Name.String,
		Description: project.Description.String,
		Price:       int(project.Price),
		Paid:        int(project.Paid),
		StartTime:   project.StartTime.Time,
		EndTime:     project.EndTime.Time,
	}

	return projectResponse, nil
}

func (p *projectUC) ListAllProjects(ctx context.Context) ([]domain.ProjectResponse, error) {
	panic(1)
}

func (p *projectUC) ListAllProjectsByUserId(ctx context.Context, userId int) ([]domain.ProjectResponse, error) {
	panic(1)
}

func (p *projectUC) ListAProject(ctx context.Context, id int) (domain.ProjectResponse, error) {
	panic(1)
}

func (p *projectUC) ListAProjectByUserId(ctx context.Context, userId int, projectId int) (domain.ProjectResponse, error) {
	panic(1)
}

func (p *projectUC) UpdateAProjectName(ctx context.Context, updateProjectName domain.ProjectUpdateName) (domain.ProjectResponse, error) {
	panic(1)
}

func (p *projectUC) UpdateAProjectPaid(ctx context.Context, updateProjectPaid domain.ProjectUpdatePaid) (domain.ProjectResponse, error) {
	panic(1)
}

func (p *projectUC) UpdateAProjectTimeWorking(ctx context.Context, updateProjectTimeWorking domain.ProjectUpdateTimeWorking) (domain.ProjectResponse, error) {
	panic(1)
}
