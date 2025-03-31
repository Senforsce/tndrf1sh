function getLetterASelector(svgSelector) {
    let aSelector = svgSelector.select("#Letter_A");
    return {
      selector: aSelector,
      left: aSelector.select("#Left"),
      right: aSelector.select("#Right"),
    };
  }
  
  function getLetterSSelector(svgSelector) {
    let sSelector = svgSelector.select("#Letter_S");
    return {
      selector: sSelector,
      outer: sSelector.select("#Outer_S"),
      inner: sSelector.select("#Inner_S"),
    };
  }
  
  function getCircleSelector(svgSelector) {
    let circleSelector = svgSelector.select("#Circle");
    return {
      selector: circleSelector,
      outer: circleSelector.select("#Outer_Circle"),
      inner: circleSelector.select("#Inner_Circle"),
      mostInner: circleSelector.select("#Most_Inner_Circle"),
    };
  }
  
  export { getLetterASelector, getLetterSSelector, getCircleSelector };
  