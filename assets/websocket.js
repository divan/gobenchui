var sock = new WebSocket(wsuri);

sock.onopen = function() { console.log("connected to " + wsuri); }
sock.onclose = function(e) { console.log("connection closed (" + e.code + ")"); }
sock.onmessage = function(e) {
	var chart_time = $('#time_benchmark').highcharts();
	var chart_mem = $('#mem_benchmark').highcharts();
	var chart_alloc = $('#alloc_benchmark').highcharts();
	
	msg = JSON.parse(e.data);

	if (msg.type === 'status') {
		date = moment(msg.commit.date).format('YYYY-MM-DD HH:mm:ss');
		name = xvalue(msg.commit);
		document.getElementById('status').innerHTML = msg.status;

		// hide loader if not 'in progress'
		if (msg.status !== 'In progress') {
			document.getElementById('status_image').style.visibility = "hidden";
		};

		// add markers on error
		if (msg.error !== undefined) {
			// TODO: these icons somehow doesn't work with exported highcharts.js png/jpg
			// find out why (already spent 2 hours) and replace with working ones.
			marker = 'url(/static/warning.png)';
			if (msg.error.Type === 'panic') {
				marker = 'url(/static/panic.png)';
			}

			item = {
				name: name,
				y: 0,
				marker: { symbol: marker }
			};

			$.each(chart_time.series, function(index, serie) {
				serie.addPoint(item, true, false);
			});
			$.each(chart_mem.series, function(index, serie) {
				serie.addPoint(item, true, false);
			});
			$.each(chart_alloc.series, function(index, serie) {
				serie.addPoint(item, true, false);
			});
		};
		if (msg.status === "Finished") {
			document.getElementById('commit_block').style.visibility = "hidden";
		} else {
			document.getElementById('commit').innerHTML = date + ' (' + msg.commit.subject + ')' + '<br />' + 'Hash: ' + msg.commit.hash;
		};
	} else if (msg.type === 'result') {
		result = msg.result;

		$.each(result.set, function(key, value) {
			var bench = value[0];
			date = moment(result.commit.date).format('YYYY-MM-DD HH:mm:ss');
			name = xvalue(result.commit);

			// time chart data
			{
				item = {
					name: name,
					y: bench.NsPerOp,
				};

				series = chart_time.get(bench.Name);
				if (series) { // series already exists
					series.addPoint(item, true, false);
				} else { //  new series
					chart_time.addSeries({
						data: [item],
						id: bench.Name,
						name: bench.Name,
					});
				}
			}

			// memory chart data
			{
				item = {
					name: name,
					y: bench.AllocedBytesPerOp,
				};

				series = chart_mem.get(bench.Name);
				if (series) { // series already exists
					series.addPoint(item, true, false);
				} else { //  new series
					chart_mem.addSeries({
						data: [item],
						id: bench.Name,
						name: bench.Name,
					});
				}
			}

			// allocs chart data
			{
				item = {
					name: name,
					y: bench.AllocsPerOp,
				};

				series = chart_alloc.get(bench.Name);
				if (series) { // series already exists
					series.addPoint(item, true, false);
				} else { //  new series
					chart_alloc.addSeries({
						data: [item],
						id: bench.Name,
						name: bench.Name,
					});
				}
			}
		});

		// Now, iterate over known series, and insert nulls
		// if values are missing.
		$.each(chart_time.series, function(index, serie) {
			found = false;
			$.each(serie.data, function(idx, item) {
				if (item.name == name) {
					found = true;
					return false;
				}
			});
			if (!found) {
				serie.addPoint({
					name: name,
					y: null,
				}, true, false)
			};
		});

		// The same for memory chart
		// TODO: move to separated function
		$.each(chart_mem.series, function(index, serie) {
			found = false;
			$.each(serie.data, function(idx, item) {
				if (item.name == name) {
					found = true;
					return false;
				}
			});
			if (!found) {
				serie.addPoint({
					name: name,
					y: null,
				}, true, false)
			};
		});

		$.each(chart_alloc.series, function(index, serie) {
			found = false;
			$.each(serie.data, function(idx, item) {
				if (item.name == name) {
					found = true;
					return false;
				}
			});
			if (!found) {
				serie.addPoint({
					name: name,
					y: null,
				}, true, false)
			};
		});
	}
}
