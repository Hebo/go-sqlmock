package sqlmock

import (
	"database/sql/driver"
	"testing"
	"time"
)

func TestQueryExpectationArgComparison(t *testing.T) {
	e := &queryBasedExpectation{}
	e.args = []driver.Value{5, "str"}

	against := []driver.Value{5}

	if e.argsMatches(against) {
		t.Error("arguments should not match, since the size is not the same")
	}

	against = []driver.Value{3, "str"}
	if e.argsMatches(against) {
		t.Error("arguments should not match, since the first argument (int value) is different")
	}

	against = []driver.Value{5, "st"}
	if e.argsMatches(against) {
		t.Error("arguments should not match, since the second argument (string value) is different")
	}

	against = []driver.Value{5, "str"}
	if !e.argsMatches(against) {
		t.Error("arguments should match, but it did not")
	}

	e.args = []driver.Value{5, time.Now()}

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	tm, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")

	against = []driver.Value{5, tm}
	if !e.argsMatches(against) {
		t.Error("arguments should match (time will be compared only by type), but it did not")
	}
}
