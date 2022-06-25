// Code generated by entc, DO NOT EDIT.

package station

const (
	// Label holds the string label denoting the station type in the database.
	Label = "station"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCrs holds the string denoting the crs field in the database.
	FieldCrs = "crs"
	// Table holds the table name of the station in the database.
	Table = "stations"
)

// Columns holds all SQL columns for station fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCrs,
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

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultCrs holds the default value on creation for the "crs" field.
	DefaultCrs string
)
