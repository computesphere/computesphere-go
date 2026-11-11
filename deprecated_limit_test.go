package computesphereapi

import (
	"encoding/json"
	"testing"
)

// The deprecated names keep their old values and still fit ProblemLimit's
// fields, so code written against the old enum types builds and behaves the
// same.
func TestDeprecatedProblemLimitNames(t *testing.T) {
	kinds := map[ProblemLimitKind]string{
		ProblemLimitKindDatabaseSize: "database_size",
		ProblemLimitKindMoney:        "money",
		ProblemLimitKindProjects:     "projects",
		ProblemLimitKindSpherelet:    "spherelet",
		ProblemLimitKindStorage:      "storage",
	}
	units := map[ProblemLimitUnit]string{
		ProblemLimitUnitCents: "cents",
		ProblemLimitUnitCount: "count",
		ProblemLimitUnitGB:    "GB",
		ProblemLimitUnitSize:  "size",
	}
	if len(kinds) != 5 || len(units) != 4 {
		t.Fatalf("constants collide: %d kinds, %d units", len(kinds), len(units))
	}
	for c, want := range kinds {
		if c != want {
			t.Errorf("kind constant = %q, want %q", c, want)
		}
	}
	for c, want := range units {
		if c != want {
			t.Errorf("unit constant = %q, want %q", c, want)
		}
	}

	// Old-style code: typed variables and conversions into the fields.
	var k ProblemLimitKind = ProblemLimitKindMoney
	var u ProblemLimitUnit = ProblemLimitUnitCents
	lim := ProblemLimit{Kind: k, Unit: ProblemLimitUnit(u)}
	if lim.Kind != ProblemLimitKindMoney || lim.Unit != ProblemLimitUnitCents {
		t.Fatalf("got %q/%q", lim.Kind, lim.Unit)
	}

	// A kind the constants don't name still decodes.
	if err := json.Unmarshal([]byte(`{"kind":"brand_new","unit":"widgets"}`), &lim); err != nil {
		t.Fatal(err)
	}
	if lim.Kind != "brand_new" || lim.Unit != "widgets" {
		t.Fatalf("got %q/%q", lim.Kind, lim.Unit)
	}
}
