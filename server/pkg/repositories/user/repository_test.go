package repository

import (
	"context"
	"testing"

	"github.com/Taehoya/pocket-mate/pkg/utils/testdb"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {
	db, err := testdb.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewUserRepository(db)

	t.Run("Successfully save user", func(t *testing.T) {
		defer testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "./teardown_test.sql")

		nickname := "test-nickname"

		ctx := context.TODO()
		err := repository.SaveUser(ctx, nickname)
		assert.NoError(t, err)

		user, err := repository.GetUser(ctx, nickname)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, nickname, user.NickName())
	})
}

func TestGetUser(t *testing.T) {
	db, err := testdb.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewUserRepository(db)

	t.Run("Successfully get user", func(t *testing.T) {
		defer testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "setup_test.sql")

		nickname := "test-nickname"

		ctx := context.TODO()

		user, err := repository.GetUser(ctx, nickname)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, nickname, user.NickName())
	})
}

func TestGetUserById(t *testing.T) {
	db, err := testdb.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewUserRepository(db)

	t.Run("Successfully get user by id", func(t *testing.T) {
		defer testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "setup_test.sql")

		id := 1
		expectedNickName := "test-nickname"

		ctx := context.TODO()

		user, err := repository.GetUserById(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedNickName, user.NickName())
		assert.Equal(t, id, user.ID())
	})
}
