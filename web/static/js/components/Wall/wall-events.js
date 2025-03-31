import FSS from "./fss.js";
import { LIGHT } from "./wall-lights.js";
import { renderer } from "./wall-renderer.js";
import render from "./wall-render.js";
import resize from "./wall-resize.js";
import { attractor, autopilotController, center } from "./wall-globals.js";

let container = null;

function addEventListeners(containerElement) {
  container = containerElement;
  window.addEventListener("resize", onWindowResize);
  container.addEventListener("click", onMouseClick);
  container.addEventListener("mousemove", onMouseMove);
}

//------------------------------
// Callbacks
//------------------------------
function onMouseClick(event) {
  FSS.Vector3.set(attractor, event.x, renderer.height - event.y);
  FSS.Vector3.subtract(attractor, center);
  LIGHT.autopilot = !LIGHT.autopilot;
  autopilotController.updateDisplay();
}

function onMouseMove(event) {
  FSS.Vector3.set(attractor, event.x, renderer.height - event.y);
  FSS.Vector3.subtract(attractor, center);
}

function onWindowResize() {
  resize(container.offsetWidth, container.offsetHeight);
  render();
}

export { addEventListeners, onMouseClick, onMouseMove, onWindowResize, container };
