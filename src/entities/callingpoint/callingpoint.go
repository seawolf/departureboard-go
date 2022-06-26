// Code generated by entc, DO NOT EDIT.

package callingpoint

import (
	"time"
)

const (
	// Label holds the string label denoting the callingpoint type in the database.
	Label = "calling_point"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldArrivalTime holds the string denoting the arrival_time field in the database.
	FieldArrivalTime = "arrival_time"
	// FieldDepartureTime holds the string denoting the departure_time field in the database.
	FieldDepartureTime = "departure_time"
	// EdgePlatform holds the string denoting the platform edge name in mutations.
	EdgePlatform = "platform"
	// EdgeService holds the string denoting the service edge name in mutations.
	EdgeService = "service"
	// Table holds the table name of the callingpoint in the database.
	Table = "calling_points"
	// PlatformTable is the table that holds the platform relation/edge.
	PlatformTable = "calling_points"
	// PlatformInverseTable is the table name for the Platform entity.
	// It exists in this package in order to avoid circular dependency with the "platform" package.
	PlatformInverseTable = "platforms"
	// PlatformColumn is the table column denoting the platform relation/edge.
	PlatformColumn = "platform_calling_points"
	// ServiceTable is the table that holds the service relation/edge.
	ServiceTable = "calling_points"
	// ServiceInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServiceInverseTable = "services"
	// ServiceColumn is the table column denoting the service relation/edge.
	ServiceColumn = "service_calling_points"
)

// Columns holds all SQL columns for callingpoint fields.
var Columns = []string{
	FieldID,
	FieldArrivalTime,
	FieldDepartureTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "calling_points"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"platform_calling_points",
	"service_calling_points",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultArrivalTime holds the default value on creation for the "arrival_time" field.
	DefaultArrivalTime time.Time
	// DefaultDepartureTime holds the default value on creation for the "departure_time" field.
	DefaultDepartureTime time.Time
)
