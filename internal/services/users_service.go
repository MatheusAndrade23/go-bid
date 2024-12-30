package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matheusandrade23/go-bid/internal/store/pgstore"
	"golang.org/x/crypto/bcrypt"
)

var ErrDuplicatedEmailOrPassword = errors.New("invalid username or email")

type UserService struct{
	pool *pgxpool.Pool
	queries *pgstore.Queries
}

func NewUserService(pool *pgxpool.Pool) UserService {
	return UserService {
		pool: pool,
		queries: pgstore.New(pool),
	}
}


func (us *UserService) CreateUser(ctx context.Context, userName, email, password, bio string) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return uuid.UUID{}, err
	}

	args := pgstore.CreateUserParams{
		UserName: userName,
		Email: email,
		Bio: bio,
		PasswordHash: hash,
	}

	id, err := us.queries.CreateUser(ctx, args)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "235" {
			return uuid.UUID{}, ErrDuplicatedEmailOrPassword
		}

		return uuid.UUID{}, err
	}

	return id, nil
}