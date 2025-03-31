
export default function updateAll(letterA, letterS, circle, a_l_center, a_r_center) {
  var radius = 220;
  var minWidth = 50;
  var maxWidth = 50;
  var minAngle = (4 / 8) * Math.PI * 2;
  var maxAngle = (12 / 8) * Math.PI * 2;
  var minColor = "#FFFFFF";
  var maxColor = "#BA6000";
  var minDuration = 1000;
  var maxDuration = 3000;
  var tickThickness = 2;

  var domain = [1, 100];

  var angleScale = d3.scaleLinear().domain(domain).range([minAngle, maxAngle]);
  var radiusScale = d3
    .scaleLinear()
    .domain(domain)
    .range([radius - minWidth, radius - maxWidth]);
  var colorScale = d3.scaleLinear().domain(domain).range([minColor, maxColor]);

  var gauge = circle.selector.select("#gauge").attr("transform", "translate(" + 702 + "," + 383 + ")");

  var arc = d3.arc().innerRadius(radiusScale).outerRadius(radius).startAngle(angleScale).endAngle(angleScale);

  var lastValue = 0;

  letterS.outer.attr("stroke-width", 0);

  letterS.outer.transition().duration(1500).attr("stroke-width", 10);

  letterA.left.attr("transform", "translate(20, -10) rotate(-120 " + a_l_center + ")").attr("stroke-width", 0);
  letterA.right.attr("transform", "translate(0, 0) rotate(0 " + a_r_center + ")").attr("stroke-width", 0);

  letterA.left
    .transition()
    .duration(1500)
    .attr("transform", "translate(20 ,-10) rotate(-198 " + a_l_center + ")")
    .attr("stroke-width", 10);

  letterA.right
    .transition()
    .duration(1500)
    .attr("transform", "translate(0, 15) rotate(15 " + a_r_center + ")")
    .attr("stroke-width", 10);

  function updateLogo() {
    letterA.right.attr("stroke-width", 0);
    letterS.outer.attr("stroke-width", 0);
    letterA.left.attr("stroke-width", 0);

    letterA.left.transition().duration(1500).attr("stroke-width", 10);

    letterS.outer.transition().duration(1500).attr("stroke-width", "10px");

    letterA.right.transition().duration(1500).attr("stroke-width", 10);
  }

  function update(n) {
    var exp = 3;
    var enterDuration = function (d) {
      var s1 = d3.scaleLinear().domain([lastValue, n]).range([0, 1]);
      var s2 = d3.scaleLinear().domain([0, 1]).range([minDuration, maxDuration]);
      return s2(Math.pow(s1(d), exp));
    };

    var exitDuration = function (d) {
      var s1 = d3.scaleLinear().domain([n, lastValue]).range([0, 1]);
      var s2 = d3.scaleLinear().domain([0, 1]).range([maxDuration, minDuration]);
      return s2(Math.pow(s1(d), 1 / exp));
    };

    var ticks = gauge.selectAll(".tick").data(d3.range(1, n), function (d) {
      return d;
    });

    ticks
      .enter()
      .append("path")
      .attr("class", "tick")
      .attr("stroke", colorScale)
      .attr("d", arc)
      .attr("stroke-width", tickThickness)
      .attr("opacity", 0)
      .transition()
      .duration(enterDuration)
      .attr("opacity", 1);

    ticks.exit().transition().delay(exitDuration).remove();

    lastValue = n;
    setTimeout(updateLogo, 2000);
  }

  function tick() {
    var n = 100;
    update(n);
    setTimeout(tick, 5000);
  }

  tick();
}
