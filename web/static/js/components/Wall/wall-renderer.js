import FSS from "./fss.js";
import { RENDER, WEBGL, CANVAS, SVG } from "./configuration.js";

let renderer = null;
let webglRenderer = null;
let canvasRenderer = null;
let svgRenderer = null;

function createRenderer(container, output) {
  webglRenderer = new FSS.WebGLRenderer();
  canvasRenderer = new FSS.CanvasRenderer();
  svgRenderer = new FSS.SVGRenderer();
  setRenderer(RENDER.renderer, container, output);
}

function setRenderer(index, container, output) {
  if (renderer) {
    output.removeChild(renderer.element);
  }
  switch (index) {
    case WEBGL:
      renderer = webglRenderer;
      break;
    case CANVAS:
      renderer = canvasRenderer;
      break;
    case SVG:
      renderer = svgRenderer;
      break;
  }
  renderer.setSize(container.offsetWidth, container.offsetHeight);
  output.appendChild(renderer.element);
}

export { createRenderer, setRenderer, renderer };
