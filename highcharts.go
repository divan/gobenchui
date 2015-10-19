package main

import (
	"fmt"
)

// HighchartsData holds series data in format
// compatible with highcharts js.
// To be used with html templates.
type HighchartsData struct {
	Categories []Commit `json:"categories,omitempty"`
	Series     []*Serie `json:"series,omitempty"`
}

// Serie is a single serie object.
// Name should be the name of benchmark.
type Serie struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Data []*Point `json:"data"`
}

// Point represents single data point in highchart.js.
// Name should be the name/id of commit.
type Point struct {
	Name   string   `json:"name"`
	Value  *float64 `json:"y"`
	Marker *Marker  `json:"marker,omitempty"`
}

// Marker is an icon marker for points.
type Marker struct {
	Symbol string `json:"symbol,omitempty"`
}

// xvalue mirrors xvalue() func in js code,
// it formats commit to serve as an X value for
// charts
func xvalue(commit Commit) string {
	date := commit.Date.Format("06-01-02")
	var hash string
	if len(commit.Hash) > 6 {
		hash = commit.Hash[0:6]
	}
	return fmt.Sprintf("%s/%s", date, hash)
}

// AddResult adds and converts benchmark set into
// highcharts-compatible representation of series/points.
//
// typ defines which result goes to this serie: "time" or "memory"
func (d *HighchartsData) AddResult(b BenchmarkSet, typ string) {
	pointName := xvalue(b.Commit)

	findSerie := func(name string) *Serie {
		if d.Series == nil {
			return nil
		}
		for _, s := range d.Series {
			if s.Name == name {
				return s
			}
		}
		return nil
	}

	// Add error values for all series on error
	if b.Error != nil {
		for _, serie := range d.Series {
			value := 0.0

			// choose different icons for build error and panic error
			symbol := "url(/static/warning.png)"
			if er, ok := b.Error.(*RunError); ok {
				if er.Type == PanicErr {
					symbol = "url(/static/panic.png)"
				}
			}
			point := &Point{
				Name:  pointName,
				Value: &value,
				Marker: &Marker{
					Symbol: symbol,
				},
			}
			serie.Data = append(serie.Data, point)
		}
		return
	}

	for name, bench := range b.Set {
		serie := findSerie(name)
		if serie == nil {
			serie = &Serie{ID: name, Name: name}
			d.Series = append(d.Series, serie)
		}

		point := &Point{
			Name:  pointName,
			Value: nil,
		}
		switch typ {
		case "time":
			point.Value = &(bench[0].NsPerOp)
		case "memory":
			val := float64(bench[0].AllocedBytesPerOp)
			point.Value = &val
		}
		serie.Data = append(serie.Data, point)
	}

	// Now, iterate over series and add null values
	// for this commit if no benchmarks were conducted.
	for _, serie := range d.Series {
		var found bool
		for name := range b.Set {
			if name == serie.Name {
				found = true
				break
			}
		}

		if !found {
			point := &Point{
				Name:  pointName,
				Value: nil,
			}
			serie.Data = append(serie.Data, point)
		}
	}
}
