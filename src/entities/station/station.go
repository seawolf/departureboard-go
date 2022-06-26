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
	// EdgePlatforms holds the string denoting the platforms edge name in mutations.
	EdgePlatforms = "platforms"
	// Table holds the table name of the station in the database.
	Table = "stations"
	// PlatformsTable is the table that holds the platforms relation/edge.
	PlatformsTable = "platforms"
	// PlatformsInverseTable is the table name for the Platform entity.
	// It exists in this package in order to avoid circular dependency with the "platform" package.
	PlatformsInverseTable = "platforms"
	// PlatformsColumn is the table column denoting the platforms relation/edge.
	PlatformsColumn = "station_platforms"
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
