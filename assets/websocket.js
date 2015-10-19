var sock = new WebSocket(wsuri);

sock.onopen = function() { console.log("connected to " + wsuri); }
sock.onclose = function(e) { console.log("connection closed (" + e.code + ")"); }
sock.onmessage = function(e) {
	var chart_time = $('#time_benchmark').highcharts();
	var chart_mem = $('#mem_benchmark').highcharts();
	
	msg = JSON.parse(e.data);

	if (msg.type === 'status') {
		document.getElementById('status').innerHTML = msg.status;
		if (msg.status === "Finished") {
			document.getElementById('commit_block').style.visibility = "hidden";
		} else {
			date = moment(msg.commit.date).format('YYYY-MM-DD HH:mm:ss');
			document.getElementById('commit').innerHTML = date + ' (' + msg.commit.subject + ')' + '<br />' + 'Hash: ' + msg.commit.hash;
		};
	} else if (msg.type === 'result') {
		result = msg.result;

		$.each(result.set, function(key, value) {
			var bench = value[0];
			date = moment(result.commit.date).format('YYYY-MM-DD HH:mm:ss');

			// time chart data
			{
				item = {
					name: date,
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
					name: date,
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
		});

		// Now, iterate over known series, and insert nulls
		// if values are missing.
		$.each(chart_time.series, function(index, serie) {
			found = false;
			$.each(serie.data, function(idx, item) {
				if (item.name == date) {
					found = true;
					return false;
				}
			});
			if (!found) {
				serie.addPoint({
					name: date,
					y: null,
				}, true, false)
			};
		});

		// The same for memory chart
		// TODO: move to separated function
		$.each(chart_mem.series, function(index, serie) {
			found = false;
			$.each(serie.data, function(idx, item) {
				if (item.name == date) {
					found = true;
					return false;
				}
			});
			if (!found) {
				serie.addPoint({
					name: date,
					y: null,
				}, true, false)
			};
		});
	}
}
