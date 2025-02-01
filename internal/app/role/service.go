package role

import (
	"context"
	"indico/internal/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type IRoleService interface {
	AddRole(ctx context.Context, req RegisterRoleRequest) (err error)
	ListRole(ctx context.Context) (role []Role, err error)
}

type RoleService struct {
	repo repositories.Querier
}

// ListRole implements IRoleService.
func (r *RoleService) ListRole(ctx context.Context) (role []Role, err error) {

	dataRole, err := r.repo.SelectRoles(ctx)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	for _, v := range dataRole {

		role = append(role, Role{
			UserRoleID: v.UserRoleID.String(),
			UniqueName: string(v.UniqueName),
			Name:       v.Name,
			Status:     v.Status,
		})
	}

	return
}

// AddRole implements IRoleService.
func (r *RoleService) AddRole(ctx context.Context, req RegisterRoleRequest) (err error) {

	id, err := uuid.NewV7()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	err = r.repo.InsertRole(ctx, repositories.InsertRoleParams{
		UserRoleID: pgtype.UUID{Bytes: id, Valid: true},
		UniqueName: repositories.UserRole(req.UniqueName),
		Name:       req.Name,
	})

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return

}

func NewRoleService(repo repositories.Querier) IRoleService {
	return &RoleService{
		repo: repo,
	}
}
