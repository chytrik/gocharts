// This file is automatically generated by qtc from "extensions.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line reports/extensions_by_month/templates/extensions.qtpl:1
package templates

//line reports/extensions_by_month/templates/extensions.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line reports/extensions_by_month/templates/extensions.qtpl:1
import "github.com/grokify/go-rickshaw/reports/extensions_by_month"

//line reports/extensions_by_month/templates/extensions.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line reports/extensions_by_month/templates/extensions.qtpl:2
func StreamRickshawExtensionsReport(qw422016 *qt422016.Writer, data rickshawextensions.TemplateData) {
	//line reports/extensions_by_month/templates/extensions.qtpl:2
	qw422016.N().S(`
<!DOCTYPE html>
<html>
<head>
	<link type="text/css" rel="stylesheet" href="https://ajax.googleapis.com/ajax/libs/jqueryui/1.8/themes/base/jquery-ui.css">
	<link type="text/css" rel="stylesheet" href="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:7
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:7
	qw422016.N().S(`/src/css/graph.css">
	<link type="text/css" rel="stylesheet" href="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:8
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:8
	qw422016.N().S(`/src/css/detail.css">
	<link type="text/css" rel="stylesheet" href="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:9
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:9
	qw422016.N().S(`/src/css/legend.css">
	<link type="text/css" rel="stylesheet" href="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:10
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:10
	qw422016.N().S(`/examples/css/extensions.css?v=2">

	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:12
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:12
	qw422016.N().S(`/vendor/d3.v3.js"></script>

	<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.6.2/jquery.min.js"></script>
	<script>
		jQuery.noConflict();
	</script>

	<script src="https://ajax.googleapis.com/ajax/libs/jqueryui/1.8.15/jquery-ui.min.js"></script>

	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:21
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:21
	qw422016.N().S(`/src/js/Rickshaw.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:22
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:22
	qw422016.N().S(`/src/js/Rickshaw.Class.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:23
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:23
	qw422016.N().S(`/src/js/Rickshaw.Compat.ClassList.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:24
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:24
	qw422016.N().S(`/src/js/Rickshaw.Graph.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:25
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:25
	qw422016.N().S(`/src/js/Rickshaw.Graph.Renderer.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:26
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:26
	qw422016.N().S(`/src/js/Rickshaw.Graph.Renderer.Area.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:27
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:27
	qw422016.N().S(`/src/js/Rickshaw.Graph.Renderer.Line.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:28
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:28
	qw422016.N().S(`/src/js/Rickshaw.Graph.Renderer.Bar.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:29
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:29
	qw422016.N().S(`/src/js/Rickshaw.Graph.Renderer.ScatterPlot.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:30
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:30
	qw422016.N().S(`/src/js/Rickshaw.Graph.Renderer.Stack.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:31
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:31
	qw422016.N().S(`/src/js/Rickshaw.Graph.RangeSlider.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:32
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:32
	qw422016.N().S(`/src/js/Rickshaw.Graph.RangeSlider.Preview.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:33
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:33
	qw422016.N().S(`/src/js/Rickshaw.Graph.HoverDetail.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:34
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:34
	qw422016.N().S(`/src/js/Rickshaw.Graph.Annotate.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:35
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:35
	qw422016.N().S(`/src/js/Rickshaw.Graph.Legend.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:36
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:36
	qw422016.N().S(`/src/js/Rickshaw.Graph.Axis.Time.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:37
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:37
	qw422016.N().S(`/src/js/Rickshaw.Graph.Behavior.Series.Toggle.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:38
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:38
	qw422016.N().S(`/src/js/Rickshaw.Graph.Behavior.Series.Order.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:39
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:39
	qw422016.N().S(`/src/js/Rickshaw.Graph.Behavior.Series.Highlight.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:40
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:40
	qw422016.N().S(`/src/js/Rickshaw.Graph.Smoother.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:41
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:41
	qw422016.N().S(`/src/js/Rickshaw.Fixtures.Time.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:42
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:42
	qw422016.N().S(`/src/js/Rickshaw.Fixtures.Time.Local.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:43
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:43
	qw422016.N().S(`/src/js/Rickshaw.Fixtures.Number.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:44
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:44
	qw422016.N().S(`/src/js/Rickshaw.Fixtures.RandomData.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:45
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:45
	qw422016.N().S(`/src/js/Rickshaw.Fixtures.Color.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:46
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:46
	qw422016.N().S(`/src/js/Rickshaw.Color.Palette.js"></script>
	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:47
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:47
	qw422016.N().S(`/src/js/Rickshaw.Graph.Axis.Y.js"></script>

	<script src="`)
	//line reports/extensions_by_month/templates/extensions.qtpl:49
	qw422016.E().S(data.RickshawURL)
	//line reports/extensions_by_month/templates/extensions.qtpl:49
	qw422016.N().S(`/examples/js/extensions.js"></script>

	<link type="text/css" rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/datatables/1.9.4/css/jquery.dataTables.min.css">
	<script src="https://cdnjs.cloudflare.com/ajax/libs/datatables/1.9.4/jquery.dataTables.min.js"></script>
</head>
<body>

`)
	//line reports/extensions_by_month/templates/extensions.qtpl:56
	qw422016.N().S(data.HeaderHTML)
	//line reports/extensions_by_month/templates/extensions.qtpl:56
	qw422016.N().S(`

<div id="content">

	<form id="side_panel">
		<h1>`)
	//line reports/extensions_by_month/templates/extensions.qtpl:61
	qw422016.E().S(data.ReportName)
	//line reports/extensions_by_month/templates/extensions.qtpl:61
	qw422016.N().S(`</h1>
		<div>Link: `)
	//line reports/extensions_by_month/templates/extensions.qtpl:62
	qw422016.N().S(data.ReportLink)
	//line reports/extensions_by_month/templates/extensions.qtpl:62
	qw422016.N().S(`</div>
		<section><div id="legend"></div></section>
		<section>
			<div id="renderer_form" class="toggler">
				<input type="radio" name="renderer" id="area" value="area" checked>
				<label for="area">area</label>
				<input type="radio" name="renderer" id="bar" value="bar">
				<label for="bar">bar</label>
				<input type="radio" name="renderer" id="line" value="line">
				<label for="line">line</label>
				<input type="radio" name="renderer" id="scatter" value="scatterplot">
				<label for="scatter">scatter</label>
			</div>
		</section>
		<section>
			<div id="offset_form">
				<label for="stack">
					<input type="radio" name="offset" id="stack" value="zero" checked>
					<span>stack</span>
				</label>
				<label for="stream">
					<input type="radio" name="offset" id="stream" value="wiggle">
					<span>stream</span>
				</label>
				<label for="pct">
					<input type="radio" name="offset" id="pct" value="expand">
					<span>pct</span>
				</label>
				<label for="value">
					<input type="radio" name="offset" id="value" value="value">
					<span>value</span>
				</label>
			</div>
			<div id="interpolation_form">
				<label for="cardinal">
					<input type="radio" name="interpolation" id="cardinal" value="cardinal" checked>
					<span>cardinal</span>
				</label>
				<label for="linear">
					<input type="radio" name="interpolation" id="linear" value="linear">
					<span>linear</span>
				</label>
				<label for="step">
					<input type="radio" name="interpolation" id="step" value="step-after">
					<span>step</span>
				</label>
			</div>
		</section>
		<section>
			<h6>Smoothing</h6>
			<div id="smoother"></div>
		</section>
		<section></section>
	</form>

	<div id="chart_container">
		<div id="chart"></div>
		<div id="timeline"></div>
		<div id="preview"></div>
	</div>

</div>

<script>

// set up our data series with 150 random data points
var seriesData = [ [], [], [], [], [], [], [], [], [] ];
var random = new Rickshaw.Fixtures.RandomData(150);

for (var i = 0; i < 150; i++) {
	random.addData(seriesData);
}

var palette = new Rickshaw.Color.Palette( { scheme: 'classic9' } );

// instantiate our graph!

var series = `)
	//line reports/extensions_by_month/templates/extensions.qtpl:139
	qw422016.N().V(string(data.FormattedDataJSON()))
	//line reports/extensions_by_month/templates/extensions.qtpl:139
	qw422016.N().S(`;

for (i = 0; i < series.length; i++) {
    series[i]["color"] = palette.color();
} 

var graph = new Rickshaw.Graph( {
	element: document.getElementById("chart"),
	width: 900,
	height: 500,
	renderer: 'area',
	stroke: true,
	preserve: true,
	series: series
} );

graph.render();

var preview = new Rickshaw.Graph.RangeSlider( {
	graph: graph,
	element: document.getElementById('preview'),
} );

var hoverDetail = new Rickshaw.Graph.HoverDetail( {
	graph: graph,
	xFormatter: function(x) {
		return new Date(x * 1000).toString();
	}
} );

var annotator = new Rickshaw.Graph.Annotate( {
	graph: graph,
	element: document.getElementById('timeline')
} );

var legend = new Rickshaw.Graph.Legend( {
	graph: graph,
	element: document.getElementById('legend')

} );

var shelving = new Rickshaw.Graph.Behavior.Series.Toggle( {
	graph: graph,
	legend: legend
} );

var order = new Rickshaw.Graph.Behavior.Series.Order( {
	graph: graph,
	legend: legend
} );

var highlighter = new Rickshaw.Graph.Behavior.Series.Highlight( {
	graph: graph,
	legend: legend
} );

var smoother = new Rickshaw.Graph.Smoother( {
	graph: graph,
	element: document.querySelector('#smoother')
} );

var ticksTreatment = 'glow';

var xAxis = new Rickshaw.Graph.Axis.Time( {
	graph: graph,
	ticksTreatment: ticksTreatment,
	timeFixture: new Rickshaw.Fixtures.Time.Local()
} );

xAxis.render();

var yAxis = new Rickshaw.Graph.Axis.Y( {
	graph: graph,
	tickFormat: Rickshaw.Fixtures.Number.formatKMBT,
	ticksTreatment: ticksTreatment
} );

yAxis.render();

var controls = new RenderControls( {
	element: document.querySelector('form'),
	graph: graph
} );

// add some data every so often

var messages = [
	"Changed home page welcome message",
	"Minified JS and CSS",
	"Changed button color from blue to green",
	"Refactored SQL query to use indexed columns",
	"Added additional logging for debugging",
	"Fixed typo",
	"Rewrite conditional logic for clarity",
	"Added documentation for new methods"
];

setInterval( function() {
	random.removeData(seriesData);
	random.addData(seriesData);
	graph.update();

}, 3000 );

function addAnnotation(force) {
	if (messages.length > 0 && (force || Math.random() >= 0.95)) {
		annotator.add(seriesData[2][seriesData[2].length-1].x, messages.shift());
		annotator.update();
	}
}

addAnnotation(true);
setTimeout( function() { setInterval( addAnnotation, 6000 ) }, 6000 );

var previewXAxis = new Rickshaw.Graph.Axis.Time({
	graph: preview.previews[0],
	timeFixture: new Rickshaw.Fixtures.Time.Local(),
	ticksTreatment: ticksTreatment
});

previewXAxis.render();

</script>

`)
	//line reports/extensions_by_month/templates/extensions.qtpl:263
	if data.IncludeDataTable {
		//line reports/extensions_by_month/templates/extensions.qtpl:263
		qw422016.N().S(`

<style>
body,h3,th {font-family:Arial,sans-serif}
</style>

<h3>Data</h3>

`)
		//line reports/extensions_by_month/templates/extensions.qtpl:272
		tableHeader, tableData := data.TableData()

		//line reports/extensions_by_month/templates/extensions.qtpl:273
		qw422016.N().S(`

<table id="myTable" style="border:1px #aaa solid;font-size:80%">
<thead>
`)
		//line reports/extensions_by_month/templates/extensions.qtpl:277
		for _, thValue := range tableHeader {
			//line reports/extensions_by_month/templates/extensions.qtpl:277
			qw422016.N().S(`
  <th>`)
			//line reports/extensions_by_month/templates/extensions.qtpl:278
			qw422016.E().S(thValue)
			//line reports/extensions_by_month/templates/extensions.qtpl:278
			qw422016.N().S(`</th>
`)
			//line reports/extensions_by_month/templates/extensions.qtpl:279
		}
		//line reports/extensions_by_month/templates/extensions.qtpl:279
		qw422016.N().S(`
</thead><tbody>
`)
		//line reports/extensions_by_month/templates/extensions.qtpl:281
		for _, dtRow := range tableData {
			//line reports/extensions_by_month/templates/extensions.qtpl:281
			qw422016.N().S(`
  <tr>
  `)
			//line reports/extensions_by_month/templates/extensions.qtpl:283
			for _, dtValue := range dtRow {
				//line reports/extensions_by_month/templates/extensions.qtpl:283
				qw422016.N().S(`
    <td>`)
				//line reports/extensions_by_month/templates/extensions.qtpl:284
				qw422016.E().S(dtValue)
				//line reports/extensions_by_month/templates/extensions.qtpl:284
				qw422016.N().S(`</td>
  `)
				//line reports/extensions_by_month/templates/extensions.qtpl:285
			}
			//line reports/extensions_by_month/templates/extensions.qtpl:285
			qw422016.N().S(`
  </tr>
`)
			//line reports/extensions_by_month/templates/extensions.qtpl:287
		}
		//line reports/extensions_by_month/templates/extensions.qtpl:287
		qw422016.N().S(`
</tbody>
</table>

<script>
jQuery(document).ready(function(){
    jQuery('#myTable').dataTable({
        "iDisplayLength": 100
    });
});
</script>

`)
		//line reports/extensions_by_month/templates/extensions.qtpl:299
	}
	//line reports/extensions_by_month/templates/extensions.qtpl:299
	qw422016.N().S(`

</body>
</html>
`)
//line reports/extensions_by_month/templates/extensions.qtpl:303
}

//line reports/extensions_by_month/templates/extensions.qtpl:303
func WriteRickshawExtensionsReport(qq422016 qtio422016.Writer, data rickshawextensions.TemplateData) {
	//line reports/extensions_by_month/templates/extensions.qtpl:303
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line reports/extensions_by_month/templates/extensions.qtpl:303
	StreamRickshawExtensionsReport(qw422016, data)
	//line reports/extensions_by_month/templates/extensions.qtpl:303
	qt422016.ReleaseWriter(qw422016)
//line reports/extensions_by_month/templates/extensions.qtpl:303
}

//line reports/extensions_by_month/templates/extensions.qtpl:303
func RickshawExtensionsReport(data rickshawextensions.TemplateData) string {
	//line reports/extensions_by_month/templates/extensions.qtpl:303
	qb422016 := qt422016.AcquireByteBuffer()
	//line reports/extensions_by_month/templates/extensions.qtpl:303
	WriteRickshawExtensionsReport(qb422016, data)
	//line reports/extensions_by_month/templates/extensions.qtpl:303
	qs422016 := string(qb422016.B)
	//line reports/extensions_by_month/templates/extensions.qtpl:303
	qt422016.ReleaseByteBuffer(qb422016)
	//line reports/extensions_by_month/templates/extensions.qtpl:303
	return qs422016
//line reports/extensions_by_month/templates/extensions.qtpl:303
}
