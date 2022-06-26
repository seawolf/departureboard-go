// Code generated by entc, DO NOT EDIT.

package callingpoint

import (
	"time"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ArrivalTime applies equality check predicate on the "arrival_time" field. It's identical to ArrivalTimeEQ.
func ArrivalTime(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArrivalTime), v))
	})
}

// DepartureTime applies equality check predicate on the "departure_time" field. It's identical to DepartureTimeEQ.
func DepartureTime(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepartureTime), v))
	})
}

// ArrivalTimeEQ applies the EQ predicate on the "arrival_time" field.
func ArrivalTimeEQ(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArrivalTime), v))
	})
}

// ArrivalTimeNEQ applies the NEQ predicate on the "arrival_time" field.
func ArrivalTimeNEQ(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldArrivalTime), v))
	})
}

// ArrivalTimeIn applies the In predicate on the "arrival_time" field.
func ArrivalTimeIn(vs ...time.Time) predicate.CallingPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CallingPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldArrivalTime), v...))
	})
}

// ArrivalTimeNotIn applies the NotIn predicate on the "arrival_time" field.
func ArrivalTimeNotIn(vs ...time.Time) predicate.CallingPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CallingPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldArrivalTime), v...))
	})
}

// ArrivalTimeGT applies the GT predicate on the "arrival_time" field.
func ArrivalTimeGT(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldArrivalTime), v))
	})
}

// ArrivalTimeGTE applies the GTE predicate on the "arrival_time" field.
func ArrivalTimeGTE(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldArrivalTime), v))
	})
}

// ArrivalTimeLT applies the LT predicate on the "arrival_time" field.
func ArrivalTimeLT(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldArrivalTime), v))
	})
}

// ArrivalTimeLTE applies the LTE predicate on the "arrival_time" field.
func ArrivalTimeLTE(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldArrivalTime), v))
	})
}

// DepartureTimeEQ applies the EQ predicate on the "departure_time" field.
func DepartureTimeEQ(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepartureTime), v))
	})
}

// DepartureTimeNEQ applies the NEQ predicate on the "departure_time" field.
func DepartureTimeNEQ(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDepartureTime), v))
	})
}

// DepartureTimeIn applies the In predicate on the "departure_time" field.
func DepartureTimeIn(vs ...time.Time) predicate.CallingPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CallingPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDepartureTime), v...))
	})
}

// DepartureTimeNotIn applies the NotIn predicate on the "departure_time" field.
func DepartureTimeNotIn(vs ...time.Time) predicate.CallingPoint {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CallingPoint(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDepartureTime), v...))
	})
}

// DepartureTimeGT applies the GT predicate on the "departure_time" field.
func DepartureTimeGT(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDepartureTime), v))
	})
}

// DepartureTimeGTE applies the GTE predicate on the "departure_time" field.
func DepartureTimeGTE(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDepartureTime), v))
	})
}

// DepartureTimeLT applies the LT predicate on the "departure_time" field.
func DepartureTimeLT(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDepartureTime), v))
	})
}

// DepartureTimeLTE applies the LTE predicate on the "departure_time" field.
func DepartureTimeLTE(v time.Time) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDepartureTime), v))
	})
}

// HasPlatform applies the HasEdge predicate on the "platform" edge.
func HasPlatform() predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlatformTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlatformTable, PlatformColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlatformWith applies the HasEdge predicate on the "platform" edge with a given conditions (other predicates).
func HasPlatformWith(preds ...predicate.Platform) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlatformInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlatformTable, PlatformColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CallingPoint) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CallingPoint) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CallingPoint) predicate.CallingPoint {
	return predicate.CallingPoint(func(s *sql.Selector) {
		p(s.Not())
	})
}
