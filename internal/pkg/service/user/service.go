package user

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"

	"app_microservice/internal/app_microservice"
	"app_microservice/internal/pkg/dto"
	"app_microservice/internal/pkg/repository/user"
)

type Token struct {
	jwt.StandardClaims
	UserId uint
}

type Service struct {
	cfg            *app_microservice.Config
	userRepository *user.Repository
}

func NewService(cfg *app_microservice.Config, userRepository *user.Repository) *Service {
	return &Service{
		cfg:            cfg,
		userRepository: userRepository,
	}
}

func (s *Service) Create(ctx context.Context, user dto.User) (string, error) {

	hashedPassword, ok := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if ok != nil {
		return "", ok
	}

	user.Password = string(hashedPassword)

	sql, args, err := squirrel.
		StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		Insert("users").
		Columns("email, password").
		Values(user.Email, user.Password).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", nil
	}

	created, err := s.userRepository.CreateOrUpdate(ctx, sql, args...)
	if err != nil {
		return "", err
	}
	tk := &Token{UserId: created}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	result, ok := token.SignedString([]byte(s.cfg.APIServer.TokenPassword))
	if ok != nil {
		return "", ok
	}

	return result, nil
}

func (s *Service) Login(ctx *gin.Context, item dto.User) (*dto.User, error) {

	_sql, args, err := squirrel.
		StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From("users").
		Where(squirrel.Eq{"email": item.Email}).
		ToSql()
	if err != nil {
		return nil, err
	}

	data, err := s.userRepository.Get(ctx, _sql, args...)
	if len(data) < 1 {
		return nil, errors.New("no user data")
	}
	var account dto.User
	account.Id = uint(data[0]["id"].(int32))
	account.Email = data[0]["email"].(string)
	account.Password = data[0]["password"].(string)

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(item.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	tk := &Token{UserId: account.Id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(s.cfg.APIServer.TokenPassword))
	if err != nil {
		return nil, err
	}
	account.Password = tokenString

	return &account, nil
}
