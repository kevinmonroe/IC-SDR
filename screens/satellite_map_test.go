package screens

import (
	"testing"

	"go-zero/internal/satellite"
)

func TestSatelliteMapCatalogSearchFiltersNameNORADAndGroup(t *testing.T) {
	v := &satelliteMap{snapshot: satellite.Snapshot{Satellites: []satellite.State{
		{Name: "ISS (ZARYA)", NORAD: 25544, Group: "Space Stations"},
		{Name: "QO-100", NORAD: 43700, Group: "Radioaficionados"},
	}}}
	for _, query := range []string{"zarya", "25544", "espaciales"} {
		v.query = query
		_, groups := v.grouped()
		if len(groups["Space Stations"]) != 1 || len(groups["Amateur Radio"]) != 0 {
			t.Fatalf("query %q returned %#v", query, groups)
		}
	}
}
