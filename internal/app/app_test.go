package app

import (
	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/MarlikAlmighty/library/internal/repository/postgresql"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {

	srv := New()

	tests := []struct {
		name string
		want *Service
	}{
		{"new_app", srv},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_InitConfig(t *testing.T) {

	type fields struct {
		Logger *zap.Logger
		Conf   *config.Config
		Booker postgresql.Booker
		Pool   *pgxpool.Pool
	}

	var f fields

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"init_config", f, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Logger: tt.fields.Logger,
				Conf:   tt.fields.Conf,
				Booker: tt.fields.Booker,
				Pool:   tt.fields.Pool,
			}
			if err := s.InitConfig(); (err != nil) != tt.wantErr {
				t.Errorf("InitConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_InitLogger(t *testing.T) {

	type fields struct {
		Logger *zap.Logger
		Conf   *config.Config
		Booker postgresql.Booker
		Pool   *pgxpool.Pool
	}

	var f fields

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"init_logger", f, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Logger: tt.fields.Logger,
				Conf:   tt.fields.Conf,
				Booker: tt.fields.Booker,
				Pool:   tt.fields.Pool,
			}
			if err := s.InitLogger(); (err != nil) != tt.wantErr {
				t.Errorf("InitLogger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
