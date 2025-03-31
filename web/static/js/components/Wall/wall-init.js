import { createRenderer } from "./wall-renderer.js";
import { createScene } from "./wall-scene.js";
import { createMesh } from "./wall-mesh.js";
import { createLights } from "./wall-lights.js";
import { addEventListeners } from "./wall-events.js";
import resize from "./wall-resize.js";
import animate from "./wall-animate.js";

/**
 *
 * @param {*} wallElement
 * @param {*} outputElement
 */
export default function initialise(wallElement, outputElement) {
  createRenderer(wallElement, outputElement);
  createScene();
  createMesh();
  createLights();
  addEventListeners(wallElement);
  resize(window.innerWidth, window.innerHeight);
  animate();
}
