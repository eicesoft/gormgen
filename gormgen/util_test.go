package gormgen

import "testing"

func TestSQLColumnToHumpStyle(t *testing.T) {
	s := SQLColumnToHumpStyle("project_name")
	t.Logf("%s", s)

	d := Capitalize("project_name")
	t.Logf("%s", d)
}
