package time_test

import (
	"fmt"
	"testing"

	"github.com/lxkedinh/stampy/time"
)

var unixSnowflakeTests = []struct {
	input    string
	expected int64
}{
	{"175928847299117063", 1462015105796},
}

func TestTimeFromSnowflake(t *testing.T) {
	for _, test := range unixSnowflakeTests {
		testName := fmt.Sprintf("Testing snowflake %s", test.input)
		t.Run(testName, func(t *testing.T) {
			got := time.TimeFromSnowflake(test.input)
			if got.UnixMilli() != test.expected {
				t.Fatalf("Testing snowflake %s failed, got %d but expected %d", test.input, got.Unix(), test.expected)
			}
		})
	}
}