// Code generated by entc, DO NOT EDIT.

package callingpoint

const (
	// Label holds the string label denoting the callingpoint type in the database.
	Label = "calling_point"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the callingpoint in the database.
	Table = "calling_points"
)

// Columns holds all SQL columns for callingpoint fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}