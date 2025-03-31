import { setNow, start } from "./wall-globals.js";
import update from "./wall-update.js";
import render from "./wall-render.js";

export default function animate() {
  setNow(Date.now() - start);
  update();
  render();
  requestAnimationFrame(animate);
}
