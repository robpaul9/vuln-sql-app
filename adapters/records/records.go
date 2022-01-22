package records

import (
	"context"
	"fmt"
	"time"

	log "github.com/robpaul9/golog"
	"github.com/robpaul9/vulnsqlapp/adapters/db"
)

type Config struct {
	Logger   log.Logger
	Database *db.DB
}

type Service struct {
	*Config
	Context context.Context
}

type User struct {
	ID        int32
	Pass      string
	Role      string
	Email     string
	CreatedAt time.Time
}

func New(config *Config) *Service {
	return &Service{
		Config:  config,
		Context: context.Background(),
	}
}

func (s *Service) GetUser(email string) ([]User, error) {

	statement := fmt.Sprintf("SELECT * FROM users WHERE email='%s';", email)
	fmt.Println(statement)

	rows, err := s.Database.Connection.Query(s.Context, statement)
	if err != nil {
		return nil, err
	}

	users := []User{}
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Pass, &u.Role, &u.Email, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
