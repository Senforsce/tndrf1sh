// @ts-nocheck
import FSS from "./fss.js";

var now,
  start = Date.now();
const center = FSS.Vector3.create();
const attractor = FSS.Vector3.create();

let renderer, scene, mesh, geometry, material;
let webglRenderer, canvasRenderer, svgRenderer;
let gui, autopilotController;

function setNow(newNow) {
  now = newNow;
}

function setGeometry(newGeometry) {
  geometry = newGeometry;
}

function setMesh(newMesh) {
  mesh = newMesh;
  return newMesh;
}

function setMaterial(newMaterial) {
  material = newMaterial;
}

export {
  now,
  start,
  center,
  attractor,
  renderer,
  scene,
  mesh,
  geometry,
  setGeometry,
  setMesh,
  setMaterial,
  material,
  webglRenderer,
  canvasRenderer,
  svgRenderer,
  gui,
  autopilotController,
  setNow,
};
