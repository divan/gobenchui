package main

// HighchartsData holds series data in format
// compatible with highcharts js.
// To be used with html templates.
type HighchartsData struct {
	Series []*Serie `json:"series,omitempty"`
}

// Serie is a single serie object.
// Name should be the name of benchmark.
type Serie struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Data []*Point `json:"data"`
}

// Point represents single data point.
// Name should be the name/id of commit.
type Point struct {
	Name  string   `json:"name"`
	Value *float64 `json:"y"`
}

// AddResult adds and converts benchmark set into
// highcharts-compatible representation of series/points.
//
// typ defines which result goes to this serie: "time" or "memory"
func (d *HighchartsData) AddResult(b BenchmarkSet, typ string) {
	// we currently use commit date as a point
	// X value, stick for it for a while
	pointName := b.Commit.Date.Format("2006-01-02 15:04:05")

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
		for name, _ := range b.Set {
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
