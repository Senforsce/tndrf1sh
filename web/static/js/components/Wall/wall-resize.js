import FSS from "./fss.js";
import { renderer } from "./wall-renderer.js";
import { createMesh } from "./wall-mesh.js";
import { center } from "./wall-globals.js";

export default function resize(width, height) {
  let w = width || 1024;
  let h = height || 768;
  renderer.setSize(w, h);
  FSS.Vector3.set(center, renderer.halfWidth, renderer.halfHeight);
  createMesh();
}
