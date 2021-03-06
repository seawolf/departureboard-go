// Code generated by entc, DO NOT EDIT.

package entities

import (
	"time"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/callingpoint"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/day"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/platform"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/service"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/station"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/toc"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities_schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	callingpointFields := entities_schema.CallingPoint{}.Fields()
	_ = callingpointFields
	// callingpointDescArrivalTime is the schema descriptor for arrival_time field.
	callingpointDescArrivalTime := callingpointFields[0].Descriptor()
	// callingpoint.DefaultArrivalTime holds the default value on creation for the arrival_time field.
	callingpoint.DefaultArrivalTime = callingpointDescArrivalTime.Default.(time.Time)
	// callingpointDescDepartureTime is the schema descriptor for departure_time field.
	callingpointDescDepartureTime := callingpointFields[1].Descriptor()
	// callingpoint.DefaultDepartureTime holds the default value on creation for the departure_time field.
	callingpoint.DefaultDepartureTime = callingpointDescDepartureTime.Default.(time.Time)
	dayFields := entities_schema.Day{}.Fields()
	_ = dayFields
	// dayDescDate is the schema descriptor for date field.
	dayDescDate := dayFields[0].Descriptor()
	// day.DefaultDate holds the default value on creation for the date field.
	day.DefaultDate = dayDescDate.Default.(time.Time)
	platformFields := entities_schema.Platform{}.Fields()
	_ = platformFields
	// platformDescName is the schema descriptor for name field.
	platformDescName := platformFields[0].Descriptor()
	// platform.DefaultName holds the default value on creation for the name field.
	platform.DefaultName = platformDescName.Default.(string)
	serviceFields := entities_schema.Service{}.Fields()
	_ = serviceFields
	// serviceDescUID is the schema descriptor for uid field.
	serviceDescUID := serviceFields[0].Descriptor()
	// service.DefaultUID holds the default value on creation for the uid field.
	service.DefaultUID = serviceDescUID.Default.(string)
	// serviceDescHeadcode is the schema descriptor for headcode field.
	serviceDescHeadcode := serviceFields[1].Descriptor()
	// service.DefaultHeadcode holds the default value on creation for the headcode field.
	service.DefaultHeadcode = serviceDescHeadcode.Default.(string)
	stationFields := entities_schema.Station{}.Fields()
	_ = stationFields
	// stationDescName is the schema descriptor for name field.
	stationDescName := stationFields[0].Descriptor()
	// station.DefaultName holds the default value on creation for the name field.
	station.DefaultName = stationDescName.Default.(string)
	// stationDescCrs is the schema descriptor for crs field.
	stationDescCrs := stationFields[1].Descriptor()
	// station.DefaultCrs holds the default value on creation for the crs field.
	station.DefaultCrs = stationDescCrs.Default.(string)
	tocFields := entities_schema.TOC{}.Fields()
	_ = tocFields
	// tocDescName is the schema descriptor for name field.
	tocDescName := tocFields[0].Descriptor()
	// toc.DefaultName holds the default value on creation for the name field.
	toc.DefaultName = tocDescName.Default.(string)
}
