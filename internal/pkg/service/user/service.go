package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"app_microservice/internal/app_microservice"
	"app_microservice/internal/pkg/dto"
	"app_microservice/internal/pkg/repository/user"
	"app_microservice/internal/pkg/util"
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

func (s *Service) Validate(ctx context.Context, user dto.User) (result bool, ok error) {

	if message, ok := user.Validate(); !ok {
		return false, errors.New(message)
	}

	_sql, args, err := squirrel.
		StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From("users").
		Where(squirrel.Eq{"email": user.Email}).
		ToSql()
	if err != nil {
		return false, err
	}

	data, err := s.userRepository.Get(ctx, _sql, args...)
	if len(data) == 0 {
		return true, nil
	}

	return result, errors.New("email address already in use by another user")
}

func (s *Service) Create(ctx context.Context, user dto.User) (string, error) {

	if response, ok := s.Validate(ctx, user); !response {
		return "", ok
	}

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

func (s *Service) Login(ctx context.Context, item dto.User) (*dto.User, error) {

	account, err := s.GetByEmail(ctx, item)
	if err != nil {
		return nil, err
	}

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

	return account, nil
}

func (s *Service) GetByEmail(ctx context.Context, item dto.User) (*dto.User, error) {
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

	var accounts []dto.User
	err = util.ToEntity(data, &accounts)
	if err != nil {
		return nil, err
	}

	return &accounts[0], nil
}

func (s *Service) GetById(ctx context.Context, item dto.User) (*dto.User, error) {
	_sql, args, err := squirrel.
		StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From("users").
		Where(squirrel.Eq{"id": item.Id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	data, err := s.userRepository.Get(ctx, _sql, args...)
	if len(data) < 1 {
		return nil, errors.New("no user data")
	}

	if len(data) > 1 {
		message := fmt.Sprintf("not unique user by id: %d", item.Id)
		return nil, errors.New(message)
	}

	var accounts []dto.User
	err = util.ToEntity(data, &accounts)

	if err != nil {
		return nil, err
	}

	return &accounts[0], nil
}

func (s *Service) List(ctx *gin.Context) ([]dto.User, error) {

	_sql, args, err := squirrel.
		StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From("users").
		ToSql()
	if err != nil {
		return nil, err
	}

	data, err := s.userRepository.Get(ctx, _sql, args...)
	if err != nil {
		return nil, err
	}

	var users []dto.User
	err = util.ToEntity(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
