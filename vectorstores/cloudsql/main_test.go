package cloudsql_test

import (
	"os"
	"testing"

	"gitlab.com/shidfar/langchaingo/internal/testutil/testctr"
)

func TestMain(m *testing.M) {
	testctr.EnsureTestEnv()
	os.Exit(m.Run())
}
