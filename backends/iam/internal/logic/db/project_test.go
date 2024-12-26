package db_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	domain_db "github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/iam/internal/logic/db"
)

func TestFindProjects(t *testing.T) {
	t.Parallel()
	conf := db.GetDefaultConfig()

	t.Run("find", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db := db.New(&conf)
		mock := db.MustOpenMock()

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "hoge")
		mock.ExpectQuery("^SELECT id,name FROM `projects` WHERE deleted = 0$").WillReturnRows(rows)

		projects, err := db.FindProjects(ctx, &domain_db.FindProjectsInput{})
		assert.NoError(t, err)
		assert.Equal(t,
			[]model.Project{{ID: 1, Name: "hoge"}},
			projects)
	})

	t.Run("find by id", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		db := db.New(&conf)
		mock := db.MustOpenMock()

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "hoge")
		mock.ExpectQuery("^SELECT id,name FROM `projects` WHERE deleted = 0 AND id = \\?$").
			WithArgs(1).
			WillReturnRows(rows)

		projects, err := db.FindProjects(ctx, &domain_db.FindProjectsInput{ID: 1})
		assert.NoError(t, err)
		assert.Equal(t,
			[]model.Project{{ID: 1, Name: "hoge"}},
			projects)
	})
}

func TestProjectSenario(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	conf := db.GetDefaultConfig()
	conf.Config.DBName = "TestProjectSenario"
	db := db.New(&conf)
	db.MustRecreateDatabase()
	db.MustOpen()
	db.MustMigrate()

	var expectedProjects []model.Project
	item1, err := db.AddProject(ctx, &model.Project{Name: "hoge"})
	assert.NoError(t, err)

	expectedProjects = append(expectedProjects, *item1)
	items, err := db.FindProjects(ctx, &domain_db.FindProjectsInput{})
	assert.NoError(t, err)
	assert.Equal(t, expectedProjects, items)

	err = db.DeleteProject(ctx, item1.ID)
	assert.NoError(t, err)

	items, err = db.FindProjects(ctx, &domain_db.FindProjectsInput{})
	assert.NoError(t, err)
	assert.Equal(t, 0, len(items))
}
