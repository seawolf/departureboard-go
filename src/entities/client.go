// Code generated by entc, DO NOT EDIT.

package entities

import (
	"context"
	"fmt"
	"log"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/migrate"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/callingpoint"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/day"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/platform"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/service"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/station"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/toc"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CallingPoint is the client for interacting with the CallingPoint builders.
	CallingPoint *CallingPointClient
	// Day is the client for interacting with the Day builders.
	Day *DayClient
	// Platform is the client for interacting with the Platform builders.
	Platform *PlatformClient
	// Service is the client for interacting with the Service builders.
	Service *ServiceClient
	// Station is the client for interacting with the Station builders.
	Station *StationClient
	// TOC is the client for interacting with the TOC builders.
	TOC *TOCClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CallingPoint = NewCallingPointClient(c.config)
	c.Day = NewDayClient(c.config)
	c.Platform = NewPlatformClient(c.config)
	c.Service = NewServiceClient(c.config)
	c.Station = NewStationClient(c.config)
	c.TOC = NewTOCClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("entities: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("entities: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		CallingPoint: NewCallingPointClient(cfg),
		Day:          NewDayClient(cfg),
		Platform:     NewPlatformClient(cfg),
		Service:      NewServiceClient(cfg),
		Station:      NewStationClient(cfg),
		TOC:          NewTOCClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		CallingPoint: NewCallingPointClient(cfg),
		Day:          NewDayClient(cfg),
		Platform:     NewPlatformClient(cfg),
		Service:      NewServiceClient(cfg),
		Station:      NewStationClient(cfg),
		TOC:          NewTOCClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CallingPoint.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.CallingPoint.Use(hooks...)
	c.Day.Use(hooks...)
	c.Platform.Use(hooks...)
	c.Service.Use(hooks...)
	c.Station.Use(hooks...)
	c.TOC.Use(hooks...)
}

// CallingPointClient is a client for the CallingPoint schema.
type CallingPointClient struct {
	config
}

// NewCallingPointClient returns a client for the CallingPoint from the given config.
func NewCallingPointClient(c config) *CallingPointClient {
	return &CallingPointClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `callingpoint.Hooks(f(g(h())))`.
func (c *CallingPointClient) Use(hooks ...Hook) {
	c.hooks.CallingPoint = append(c.hooks.CallingPoint, hooks...)
}

// Create returns a create builder for CallingPoint.
func (c *CallingPointClient) Create() *CallingPointCreate {
	mutation := newCallingPointMutation(c.config, OpCreate)
	return &CallingPointCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CallingPoint entities.
func (c *CallingPointClient) CreateBulk(builders ...*CallingPointCreate) *CallingPointCreateBulk {
	return &CallingPointCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CallingPoint.
func (c *CallingPointClient) Update() *CallingPointUpdate {
	mutation := newCallingPointMutation(c.config, OpUpdate)
	return &CallingPointUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CallingPointClient) UpdateOne(cp *CallingPoint) *CallingPointUpdateOne {
	mutation := newCallingPointMutation(c.config, OpUpdateOne, withCallingPoint(cp))
	return &CallingPointUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CallingPointClient) UpdateOneID(id int) *CallingPointUpdateOne {
	mutation := newCallingPointMutation(c.config, OpUpdateOne, withCallingPointID(id))
	return &CallingPointUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CallingPoint.
func (c *CallingPointClient) Delete() *CallingPointDelete {
	mutation := newCallingPointMutation(c.config, OpDelete)
	return &CallingPointDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CallingPointClient) DeleteOne(cp *CallingPoint) *CallingPointDeleteOne {
	return c.DeleteOneID(cp.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CallingPointClient) DeleteOneID(id int) *CallingPointDeleteOne {
	builder := c.Delete().Where(callingpoint.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CallingPointDeleteOne{builder}
}

// Query returns a query builder for CallingPoint.
func (c *CallingPointClient) Query() *CallingPointQuery {
	return &CallingPointQuery{
		config: c.config,
	}
}

// Get returns a CallingPoint entity by its id.
func (c *CallingPointClient) Get(ctx context.Context, id int) (*CallingPoint, error) {
	return c.Query().Where(callingpoint.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CallingPointClient) GetX(ctx context.Context, id int) *CallingPoint {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CallingPointClient) Hooks() []Hook {
	return c.hooks.CallingPoint
}

// DayClient is a client for the Day schema.
type DayClient struct {
	config
}

// NewDayClient returns a client for the Day from the given config.
func NewDayClient(c config) *DayClient {
	return &DayClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `day.Hooks(f(g(h())))`.
func (c *DayClient) Use(hooks ...Hook) {
	c.hooks.Day = append(c.hooks.Day, hooks...)
}

// Create returns a create builder for Day.
func (c *DayClient) Create() *DayCreate {
	mutation := newDayMutation(c.config, OpCreate)
	return &DayCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Day entities.
func (c *DayClient) CreateBulk(builders ...*DayCreate) *DayCreateBulk {
	return &DayCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Day.
func (c *DayClient) Update() *DayUpdate {
	mutation := newDayMutation(c.config, OpUpdate)
	return &DayUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DayClient) UpdateOne(d *Day) *DayUpdateOne {
	mutation := newDayMutation(c.config, OpUpdateOne, withDay(d))
	return &DayUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DayClient) UpdateOneID(id int) *DayUpdateOne {
	mutation := newDayMutation(c.config, OpUpdateOne, withDayID(id))
	return &DayUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Day.
func (c *DayClient) Delete() *DayDelete {
	mutation := newDayMutation(c.config, OpDelete)
	return &DayDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DayClient) DeleteOne(d *Day) *DayDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DayClient) DeleteOneID(id int) *DayDeleteOne {
	builder := c.Delete().Where(day.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DayDeleteOne{builder}
}

// Query returns a query builder for Day.
func (c *DayClient) Query() *DayQuery {
	return &DayQuery{
		config: c.config,
	}
}

// Get returns a Day entity by its id.
func (c *DayClient) Get(ctx context.Context, id int) (*Day, error) {
	return c.Query().Where(day.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DayClient) GetX(ctx context.Context, id int) *Day {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DayClient) Hooks() []Hook {
	return c.hooks.Day
}

// PlatformClient is a client for the Platform schema.
type PlatformClient struct {
	config
}

// NewPlatformClient returns a client for the Platform from the given config.
func NewPlatformClient(c config) *PlatformClient {
	return &PlatformClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `platform.Hooks(f(g(h())))`.
func (c *PlatformClient) Use(hooks ...Hook) {
	c.hooks.Platform = append(c.hooks.Platform, hooks...)
}

// Create returns a create builder for Platform.
func (c *PlatformClient) Create() *PlatformCreate {
	mutation := newPlatformMutation(c.config, OpCreate)
	return &PlatformCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Platform entities.
func (c *PlatformClient) CreateBulk(builders ...*PlatformCreate) *PlatformCreateBulk {
	return &PlatformCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Platform.
func (c *PlatformClient) Update() *PlatformUpdate {
	mutation := newPlatformMutation(c.config, OpUpdate)
	return &PlatformUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlatformClient) UpdateOne(pl *Platform) *PlatformUpdateOne {
	mutation := newPlatformMutation(c.config, OpUpdateOne, withPlatform(pl))
	return &PlatformUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlatformClient) UpdateOneID(id int) *PlatformUpdateOne {
	mutation := newPlatformMutation(c.config, OpUpdateOne, withPlatformID(id))
	return &PlatformUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Platform.
func (c *PlatformClient) Delete() *PlatformDelete {
	mutation := newPlatformMutation(c.config, OpDelete)
	return &PlatformDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PlatformClient) DeleteOne(pl *Platform) *PlatformDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PlatformClient) DeleteOneID(id int) *PlatformDeleteOne {
	builder := c.Delete().Where(platform.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlatformDeleteOne{builder}
}

// Query returns a query builder for Platform.
func (c *PlatformClient) Query() *PlatformQuery {
	return &PlatformQuery{
		config: c.config,
	}
}

// Get returns a Platform entity by its id.
func (c *PlatformClient) Get(ctx context.Context, id int) (*Platform, error) {
	return c.Query().Where(platform.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlatformClient) GetX(ctx context.Context, id int) *Platform {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStation queries the station edge of a Platform.
func (c *PlatformClient) QueryStation(pl *Platform) *StationQuery {
	query := &StationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(platform.Table, platform.FieldID, id),
			sqlgraph.To(station.Table, station.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, platform.StationTable, platform.StationColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PlatformClient) Hooks() []Hook {
	return c.hooks.Platform
}

// ServiceClient is a client for the Service schema.
type ServiceClient struct {
	config
}

// NewServiceClient returns a client for the Service from the given config.
func NewServiceClient(c config) *ServiceClient {
	return &ServiceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `service.Hooks(f(g(h())))`.
func (c *ServiceClient) Use(hooks ...Hook) {
	c.hooks.Service = append(c.hooks.Service, hooks...)
}

// Create returns a create builder for Service.
func (c *ServiceClient) Create() *ServiceCreate {
	mutation := newServiceMutation(c.config, OpCreate)
	return &ServiceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Service entities.
func (c *ServiceClient) CreateBulk(builders ...*ServiceCreate) *ServiceCreateBulk {
	return &ServiceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Service.
func (c *ServiceClient) Update() *ServiceUpdate {
	mutation := newServiceMutation(c.config, OpUpdate)
	return &ServiceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServiceClient) UpdateOne(s *Service) *ServiceUpdateOne {
	mutation := newServiceMutation(c.config, OpUpdateOne, withService(s))
	return &ServiceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServiceClient) UpdateOneID(id int) *ServiceUpdateOne {
	mutation := newServiceMutation(c.config, OpUpdateOne, withServiceID(id))
	return &ServiceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Service.
func (c *ServiceClient) Delete() *ServiceDelete {
	mutation := newServiceMutation(c.config, OpDelete)
	return &ServiceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ServiceClient) DeleteOne(s *Service) *ServiceDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ServiceClient) DeleteOneID(id int) *ServiceDeleteOne {
	builder := c.Delete().Where(service.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServiceDeleteOne{builder}
}

// Query returns a query builder for Service.
func (c *ServiceClient) Query() *ServiceQuery {
	return &ServiceQuery{
		config: c.config,
	}
}

// Get returns a Service entity by its id.
func (c *ServiceClient) Get(ctx context.Context, id int) (*Service, error) {
	return c.Query().Where(service.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServiceClient) GetX(ctx context.Context, id int) *Service {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ServiceClient) Hooks() []Hook {
	return c.hooks.Service
}

// StationClient is a client for the Station schema.
type StationClient struct {
	config
}

// NewStationClient returns a client for the Station from the given config.
func NewStationClient(c config) *StationClient {
	return &StationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `station.Hooks(f(g(h())))`.
func (c *StationClient) Use(hooks ...Hook) {
	c.hooks.Station = append(c.hooks.Station, hooks...)
}

// Create returns a create builder for Station.
func (c *StationClient) Create() *StationCreate {
	mutation := newStationMutation(c.config, OpCreate)
	return &StationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Station entities.
func (c *StationClient) CreateBulk(builders ...*StationCreate) *StationCreateBulk {
	return &StationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Station.
func (c *StationClient) Update() *StationUpdate {
	mutation := newStationMutation(c.config, OpUpdate)
	return &StationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StationClient) UpdateOne(s *Station) *StationUpdateOne {
	mutation := newStationMutation(c.config, OpUpdateOne, withStation(s))
	return &StationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StationClient) UpdateOneID(id int) *StationUpdateOne {
	mutation := newStationMutation(c.config, OpUpdateOne, withStationID(id))
	return &StationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Station.
func (c *StationClient) Delete() *StationDelete {
	mutation := newStationMutation(c.config, OpDelete)
	return &StationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StationClient) DeleteOne(s *Station) *StationDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StationClient) DeleteOneID(id int) *StationDeleteOne {
	builder := c.Delete().Where(station.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StationDeleteOne{builder}
}

// Query returns a query builder for Station.
func (c *StationClient) Query() *StationQuery {
	return &StationQuery{
		config: c.config,
	}
}

// Get returns a Station entity by its id.
func (c *StationClient) Get(ctx context.Context, id int) (*Station, error) {
	return c.Query().Where(station.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StationClient) GetX(ctx context.Context, id int) *Station {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlatforms queries the platforms edge of a Station.
func (c *StationClient) QueryPlatforms(s *Station) *PlatformQuery {
	query := &PlatformQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(station.Table, station.FieldID, id),
			sqlgraph.To(platform.Table, platform.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, station.PlatformsTable, station.PlatformsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StationClient) Hooks() []Hook {
	return c.hooks.Station
}

// TOCClient is a client for the TOC schema.
type TOCClient struct {
	config
}

// NewTOCClient returns a client for the TOC from the given config.
func NewTOCClient(c config) *TOCClient {
	return &TOCClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `toc.Hooks(f(g(h())))`.
func (c *TOCClient) Use(hooks ...Hook) {
	c.hooks.TOC = append(c.hooks.TOC, hooks...)
}

// Create returns a create builder for TOC.
func (c *TOCClient) Create() *TOCCreate {
	mutation := newTOCMutation(c.config, OpCreate)
	return &TOCCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TOC entities.
func (c *TOCClient) CreateBulk(builders ...*TOCCreate) *TOCCreateBulk {
	return &TOCCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TOC.
func (c *TOCClient) Update() *TOCUpdate {
	mutation := newTOCMutation(c.config, OpUpdate)
	return &TOCUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TOCClient) UpdateOne(t *TOC) *TOCUpdateOne {
	mutation := newTOCMutation(c.config, OpUpdateOne, withTOC(t))
	return &TOCUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TOCClient) UpdateOneID(id int) *TOCUpdateOne {
	mutation := newTOCMutation(c.config, OpUpdateOne, withTOCID(id))
	return &TOCUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TOC.
func (c *TOCClient) Delete() *TOCDelete {
	mutation := newTOCMutation(c.config, OpDelete)
	return &TOCDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TOCClient) DeleteOne(t *TOC) *TOCDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TOCClient) DeleteOneID(id int) *TOCDeleteOne {
	builder := c.Delete().Where(toc.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TOCDeleteOne{builder}
}

// Query returns a query builder for TOC.
func (c *TOCClient) Query() *TOCQuery {
	return &TOCQuery{
		config: c.config,
	}
}

// Get returns a TOC entity by its id.
func (c *TOCClient) Get(ctx context.Context, id int) (*TOC, error) {
	return c.Query().Where(toc.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TOCClient) GetX(ctx context.Context, id int) *TOC {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TOCClient) Hooks() []Hook {
	return c.hooks.TOC
}
