package user

import (
	"context"
	"indico/internal/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	RegisterUser(ctx context.Context, req RegisterUserRequest) (err error)
	LoginUser(ctx context.Context, req LoginUserRequest) (token string, err error)
	GetUserInfo(ctx context.Context, userID string) (user User, err error)
	GetUserAdmin(ctx context.Context, uniqueName string) (user []User, err error)
}

type UserService struct {
	repo repositories.Querier
}

// GetUserAdmin implements IUserService.
func (u *UserService) GetUserAdmin(ctx context.Context, uniqueName string) (user []User, err error) {

	sliceUser, err := u.repo.SelectUserByRole(ctx, repositories.UserRole(uniqueName))
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	for _, v := range sliceUser {

		user = append(user, User{

			UserID:   v.UserID.String(),
			FullName: v.FullName,
			Email:    v.Email,
			Role:     v.Name,
		})
	}

	return
}

// GetUserInfo implements IUserService.
func (u *UserService) GetUserInfo(ctx context.Context, userID string) (user User, err error) {

	currentUser, err := u.repo.SelectOneUserById(ctx, pgtype.UUID{Bytes: uuid.MustParse(userID), Valid: true})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	user = User{
		UserID:   currentUser.UserID.String(),
		FullName: currentUser.FullName,
		Email:    currentUser.Email,
		Role:     currentUser.Name,
	}

	return
}

// RegisterUser implements IUserService.
func (u *UserService) RegisterUser(ctx context.Context, req RegisterUserRequest) (err error) {

	id, err := uuid.NewV7()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	err = u.repo.InsertUser(ctx, repositories.InsertUserParams{
		UserID:   pgtype.UUID{Bytes: id, Valid: true},
		FullName: req.FullName,
		Email:    req.Email,
		Password: string(hashPass),
		RoleID:   pgtype.UUID{Bytes: uuid.MustParse(req.RoleId), Valid: true},
	},
	)

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return
}

// LoginUser implements IUserService.
func (u *UserService) LoginUser(ctx context.Context, req LoginUserRequest) (token string, err error) {

	currentUser, err := u.repo.SelectOneUserByEmail(ctx, req.Email)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	claimsJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"sub": currentUser.UserID.String(),
	})

	token, err = claimsJwt.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return
}

func NewUserService(repo repositories.Querier) IUserService {
	return &UserService{
		repo: repo,
	}
}
