package ghutil

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestNewAuthorizedClientFromEnv_NotSet(t *testing.T) {
	originalVal := os.Getenv("GITHUB_TOKEN")
	os.Setenv("GITHUB_TOKEN", "")
	modifiedVal := os.Getenv("GITHUB_TOKEN")
	c, err := NewAuthorizedClientFromEnv()

	require.Truef(t, c == nil && err != nil, "should fail when GITHUB_TOKEN env value is: %v",
		modifiedVal)
	os.Setenv("GITHUB_TOKEN", originalVal)
}

func TestNewAuthorizedClientFromEnvSet(t *testing.T) {
	token := os.Getenv("GITHUB_TOKEN")
	c, err := NewAuthorizedClientFromEnv()
	require.Truef(t, c != nil && err == nil, "shouldn't fail when GITHUB_TOKEN env value is: %v",
		token)
}

// Assert error is thrown when secret file does not exist
func TestNewAuthorizedClientFromFile_InvalidErrors(t *testing.T) {
	// File does not exist
	var pathError *os.PathError
	var malformedError MalformedFileError
	args := []struct {
		in  string
		out interface{}
		err error
	}{
		{"", nil, pathError},
		{"fileDoesNotExist", nil, pathError},
		{"invalid_secret.json", nil, malformedError},
	}
	for _, arg := range args {
		_, err := NewAuthorizedClientFromSecretFile(arg.in)
		require.Equalf(t, reflect.TypeOf(err), reflect.TypeOf(arg.err),
			"expected an error from the input arg: %v", arg)
	}
}

// Assert no error when secret file is correct
func TestNewAuthorizedClientFromFile_Success(t *testing.T) {
	c, _ := NewAuthorizedClientFromSecretFile("valid_secret.json")
	require.NotEmpty(t, c, "client should be created!")
}
