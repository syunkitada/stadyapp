package db_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	domain_db "github.com/syunkitada/stadyapp/backends/compute/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/compute/internal/domain/model"
	"github.com/syunkitada/stadyapp/backends/compute/internal/logic/db"
)

func TestFindProjects(t *testing.T) {
	t.Parallel()

	conf := db.GetDefaultConfig()

	t.Run("find", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		db := db.New(&conf)
		mock := db.MustOpenMock(ctx)

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow("uuid1", "hoge")
		mock.ExpectQuery("^SELECT id,name FROM `projects` WHERE deleted = 0$").WillReturnRows(rows)

		projects, err := db.GetProjects(ctx, &domain_db.GetProjectsInput{})
		require.NoError(t, err)
		assert.Equal(t,
			[]model.Project{{ID: "uuid1", Name: "hoge"}},
			projects)
	})

	t.Run("find by id", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		db := db.New(&conf)
		mock := db.MustOpenMock(ctx)

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow("uuid1", "hoge")
		mock.ExpectQuery("^SELECT id,name FROM `projects` WHERE deleted = 0 AND id = \\?$").
			WithArgs(1).
			WillReturnRows(rows)

		projects, err := db.GetProjects(ctx, &domain_db.GetProjectsInput{ID: "uuid1"})
		require.NoError(t, err)
		assert.Equal(t,
			[]model.Project{{ID: "uuid1", Name: "hoge"}},
			projects)
	})
}

func TestProjectSenario(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	conf := db.GetDefaultConfig()
	conf.Config.DBName = "TestProjectSenario"
	db := db.New(&conf)
	db.MustRecreateDatabase(ctx)
	db.MustOpen(ctx)
	db.MustMigrate(ctx)

	var expectedProjects []model.Project

	item1, err := db.CreateProject(ctx, &domain_db.CreateProjectInput{Name: "hoge"})
	require.NoError(t, err)

	expectedProjects = append(expectedProjects, *item1)
	items, err := db.GetProjects(ctx, &domain_db.GetProjectsInput{})
	require.NoError(t, err)
	assert.Equal(t, expectedProjects, items)

	err = db.DeleteProjectByID(ctx, item1.ID)
	require.NoError(t, err)

	items, err = db.GetProjects(ctx, &domain_db.GetProjectsInput{})
	require.NoError(t, err)
	assert.Empty(t, items)
}
