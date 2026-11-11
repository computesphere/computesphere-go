package computesphereapi

// Hand-written, not generated: oapi-codegen writes only client.gen.go.
//
// ProblemLimit.Kind and ProblemLimit.Unit used to be closed enums with their
// own named types and constants. They are now open strings (a client must
// accept a value it doesn't know), so the generator no longer emits the types
// or the constants. The names below keep code that still uses them building.
// New code compares ProblemLimit.Kind and ProblemLimit.Unit with string
// literals, or with these constants while they last.

// ProblemLimitKind is what a ProblemLimit limits.
//
// Deprecated: ProblemLimit.Kind is a plain string; use string. An alias, so
// values of this type and string mix freely.
type ProblemLimitKind = string

// ProblemLimitUnit is the unit of a ProblemLimit's current, cap and requested.
//
// Deprecated: ProblemLimit.Unit is a plain string; use string. An alias, so
// values of this type and string mix freely.
type ProblemLimitUnit = string

// Values ProblemLimit.Kind has today. The API may add others; a client must
// accept a kind it doesn't know.
const (
	// Deprecated: compare ProblemLimit.Kind with "database_size".
	ProblemLimitKindDatabaseSize ProblemLimitKind = "database_size"
	// Deprecated: compare ProblemLimit.Kind with "money".
	ProblemLimitKindMoney ProblemLimitKind = "money"
	// Deprecated: compare ProblemLimit.Kind with "projects".
	ProblemLimitKindProjects ProblemLimitKind = "projects"
	// Deprecated: compare ProblemLimit.Kind with "spherelet".
	ProblemLimitKindSpherelet ProblemLimitKind = "spherelet"
	// Deprecated: compare ProblemLimit.Kind with "storage".
	ProblemLimitKindStorage ProblemLimitKind = "storage"
)

// Values ProblemLimit.Unit has today. The API may add others; a client must
// accept a unit it doesn't know.
const (
	// Deprecated: compare ProblemLimit.Unit with "cents".
	ProblemLimitUnitCents ProblemLimitUnit = "cents"
	// Deprecated: compare ProblemLimit.Unit with "count".
	ProblemLimitUnitCount ProblemLimitUnit = "count"
	// Deprecated: compare ProblemLimit.Unit with "GB".
	ProblemLimitUnitGB ProblemLimitUnit = "GB"
	// Deprecated: compare ProblemLimit.Unit with "size".
	ProblemLimitUnitSize ProblemLimitUnit = "size"
)
