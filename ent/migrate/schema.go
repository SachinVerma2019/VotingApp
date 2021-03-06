// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PollsColumns holds the columns for the "polls" table.
	PollsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ownerid", Type: field.TypeInt, Default: -1},
		{Name: "topic", Type: field.TypeString, Default: "unknown"},
		{Name: "options", Type: field.TypeJSON, Nullable: true},
		{Name: "createtime", Type: field.TypeTime},
		{Name: "modifytime", Type: field.TypeTime},
	}
	// PollsTable holds the schema information for the "polls" table.
	PollsTable = &schema.Table{
		Name:       "polls",
		Columns:    PollsColumns,
		PrimaryKey: []*schema.Column{PollsColumns[0]},
	}
	// ResultsColumns holds the columns for the "results" table.
	ResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "userid", Type: field.TypeInt, Default: -1},
		{Name: "pollid", Type: field.TypeInt, Default: -1},
		{Name: "option", Type: field.TypeString, Default: "unknown"},
		{Name: "createtime", Type: field.TypeTime},
		{Name: "modifytime", Type: field.TypeTime},
	}
	// ResultsTable holds the schema information for the "results" table.
	ResultsTable = &schema.Table{
		Name:       "results",
		Columns:    ResultsColumns,
		PrimaryKey: []*schema.Column{ResultsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "email", Type: field.TypeString, Default: "unknown"},
		{Name: "password", Type: field.TypeString, Default: "xxxxx"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PollsTable,
		ResultsTable,
		UsersTable,
	}
)

func init() {
}
