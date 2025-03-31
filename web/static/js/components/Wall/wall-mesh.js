import FSS from "./fss.js";
import XMath from "./fss/XMath.js";
import { scene } from "./wall-scene.js";
import { renderer } from "./wall-renderer.js";
import { geometry, setGeometry, mesh, setMesh, material, setMaterial } from "./wall-globals.js";

//------------------------------
// Mesh Properties
//------------------------------
let MESH = {
  width: 1.2,
  height: 1.2,
  depth: 10,
  segments: 16,
  slices: 8,
  xRange: 0.8,
  yRange: 0.1,
  zRange: 1.0,
  ambient: "#000000",
  diffuse: "#FD0211",
  speed: 0.0009,
};

function createMesh(override) {
  scene.remove(mesh);
  renderer.clear();

  setGeometry(new FSS.Plane(MESH.width * renderer.width, MESH.height * renderer.height, MESH.segments, MESH.slices));
  setMaterial(new FSS.Material(MESH.ambient, override || window.ASMeshDiffuse || MESH.diffuse));
  setMesh(new FSS.Mesh(geometry, material));

  scene.add(mesh);

  // Augment vertices for animation
  var v, vertex;
  for (v = geometry.vertices.length - 1; v >= 0; v--) {
    vertex = geometry.vertices[v];
    vertex.anchor = FSS.Vector3.clone(vertex.position);
    vertex.step = FSS.Vector3.create(
      XMath.randomInRange(0.2, 1.0),
      XMath.randomInRange(0.2, 1.0),
      XMath.randomInRange(0.2, 1.0)
    );
    vertex.time = XMath.randomInRange(0, XMath.PIM2);
  }
}

export { createMesh, MESH };
