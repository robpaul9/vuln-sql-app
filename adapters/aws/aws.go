package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	log "github.com/robpaul9/golog"
)

type Config struct {
	Logger log.Logger
}

type Service struct {
	*Config
	Context context.Context
	SSMVC   *SSM
}

type SSM struct {
	client ssmiface.SSMAPI
}

type Param struct {
	Name           string
	WithDecryption bool
	ssmsvc         *SSM
}

func Sessions() (*session.Session, error) {
	sess, err := session.NewSession()
	svc := session.Must(sess, err)
	return svc, err
}
func New(config *Config) *Service {
	sess, err := Sessions()
	if err != nil {
		panic(err)
	}
	ssmsvc := &SSM{ssm.New(sess)}

	return &Service{
		Config:  config,
		Context: context.Background(),
		SSMVC:   ssmsvc,
	}
}

func (s *SSM) Param(name string, decryption bool) *Param {
	return &Param{
		Name:           name,
		WithDecryption: decryption,
		ssmsvc:         s,
	}
}

func (p *Param) GetValue() (string, error) {
	ssmsvc := p.ssmsvc.client
	parameter, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           &p.Name,
		WithDecryption: &p.WithDecryption,
	})
	if err != nil {
		return "", err
	}
	value := *parameter.Parameter.Value
	return value, nil
}
