package tx

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestRepository_BeginTx(t *testing.T) {
	type mockFields struct {
		sql sqlmock.Sqlmock
	}
	type args struct {
		ctx  context.Context
		opts *sql.TxOptions
	}

	mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	defer func() {
		_ = mockDB.Close()
	}()

	mockGORM, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}))
	assert.NoError(t, err)

	tests := []struct {
		name    string
		args    args
		mock    func(mock mockFields)
		wantErr error
	}{
		{
			name: "case success",
			args: args{
				ctx:  context.Background(),
				opts: &sql.TxOptions{},
			},
			mock: func(mock mockFields) {
				mock.sql.ExpectBegin()
			},
		},
		{
			name: "case error - begin fail",
			args: args{
				ctx: context.Background(),
			},
			wantErr: assert.AnError,
			mock: func(mock mockFields) {
				mock.sql.ExpectBegin().WillReturnError(assert.AnError)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockFields := mockFields{
				sql: mockSQL,
			}
			test.mock(mockFields)
			repository := &Repository{
				db: mockGORM,
			}

			_, err := repository.BeginTx(test.args.ctx, test.args.opts)
			assert.Equal(t, test.wantErr, err)
			assert.NoError(t, mockSQL.ExpectationsWereMet())
		})
	}
}

func TestRepository_Commit(t *testing.T) {
	type mockFields struct {
		sql sqlmock.Sqlmock
	}
	type args struct {
		ctx context.Context
	}

	mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	defer func() {
		_ = mockDB.Close()
	}()

	mockGORM, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}))
	assert.NoError(t, err)

	tests := []struct {
		name    string
		args    args
		mock    func(mock mockFields)
		wantErr error
	}{
		{
			name: "case success",
			args: args{
				ctx: context.Background(),
			},
			mock: func(mock mockFields) {
				mock.sql.ExpectBegin()
				mock.sql.ExpectCommit()
			},
		},
		{
			name: "case error - commit fail",
			args: args{
				ctx: context.Background(),
			},
			wantErr: assert.AnError,
			mock: func(mock mockFields) {
				mock.sql.ExpectBegin()
				mock.sql.ExpectCommit().WillReturnError(assert.AnError)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockFields := mockFields{
				sql: mockSQL,
			}
			test.mock(mockFields)
			repository := &Repository{
				db: mockGORM,
			}

			tx := mockGORM.Begin()
			err := repository.Commit(test.args.ctx, tx)
			assert.Equal(t, test.wantErr, err)
			assert.NoError(t, mockSQL.ExpectationsWereMet())
		})
	}
}

func TestRepository_Rollback(t *testing.T) {
	type mockFields struct {
		sql sqlmock.Sqlmock
	}
	type args struct {
		ctx context.Context
	}

	mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	defer func() {
		_ = mockDB.Close()
	}()

	mockGORM, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}))
	assert.NoError(t, err)

	tests := []struct {
		name    string
		args    args
		mock    func(mock mockFields)
		wantErr error
	}{
		{
			name: "case success",
			args: args{
				ctx: context.Background(),
			},
			mock: func(mock mockFields) {
				mock.sql.ExpectBegin()
				mock.sql.ExpectRollback()
			},
		},
		{
			name: "case error - rollback fail",
			args: args{
				ctx: context.Background(),
			},
			wantErr: assert.AnError,
			mock: func(mock mockFields) {
				mock.sql.ExpectBegin()
				mock.sql.ExpectRollback().WillReturnError(assert.AnError)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockFields := mockFields{
				sql: mockSQL,
			}
			test.mock(mockFields)
			repository := &Repository{
				db: mockGORM,
			}

			tx := mockGORM.Begin()
			err := repository.Rollback(test.args.ctx, tx)
			assert.Equal(t, test.wantErr, err)
			assert.NoError(t, mockSQL.ExpectationsWereMet())
		})
	}
}
