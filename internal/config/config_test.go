package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("Use environment variables when config files are not found", func(t *testing.T) {
		t.Setenv("APP_DATABASE_USERNAME", "testuser")
		t.Setenv("APP_DATABASE_PASSWORD", "testpassword")
		t.Setenv("APP_DATABASE_HOST", "testhostname")
		t.Setenv("APP_DATABASE_PORT", "1111")
		t.Setenv("APP_DATABASE_NAME", "testdbname")
		t.Setenv("APP_APPLICATION_HOST", "1.1.1.1")
		t.Setenv("APP_APPLICATION_PORT", "2222")

		got, err := GetConfiguration("wrong_folder")
		if err != nil {
			t.Fatal(err)
		}

		want := &Settings{
			ApplicationSettings: ApplicationSettings{
				Port: 2222,
				Host: "1.1.1.1",
			},
			DatabaseSettings: DatabaseSettings{
				Username:     "testuser",
				Password:     "testpassword",
				Port:         1111,
				Host:         "testhostname",
				DatabaseName: "testdbname",
			},
		}

		assert.Equal(t, want, got)
	})

	t.Run("Use config files if they exist", func(t *testing.T) {
		t.Setenv("APP_ENVIRONMENT", "local")
		got, err := GetConfiguration("./config_test")

		if err != nil {
			t.Fatal(err)
		}

		want := &Settings{
			ApplicationSettings: ApplicationSettings{
				Port: 4444,
				Host: "127.0.0.1",
			},
			DatabaseSettings: DatabaseSettings{
				Username:     "postgres",
				Password:     "password",
				Port:         5432,
				Host:         "127.0.0.1",
				DatabaseName: "newsletter",
			},
		}

		assert.Equal(t, want, got)
	})
}
