import FSSGeometry from "./Geometry.js";
import FSSVertex from "./Vertex.js";
import FSSTriangle from "./Triangle.js";

/**
 *
 * @param {number} [width]
 * @param {number} [height]
 * @param {number} [segments]
 * @param {number} [slices]
 */
function Plane(width = 100, height = 100, segments = 4, slices = 4) {
  FSSGeometry.call(this);
  this.width = width;
  this.height = height;
  this.segments = segments;
  this.slices = slices;
  this.segmentWidth = this.width / this.segments;
  this.sliceHeight = this.height / this.slices;

  // Cache Variables
  let x,
    y,
    v0,
    v1,
    v2,
    v3,
    t0,
    t1,
    vertex,
    vertices = [],
    offsetX = this.width * -0.5,
    offsetY = this.height * 0.5;

  // Add Vertices
  for (x = 0; x <= this.segments; x++) {
    vertices.push([]);
    for (y = 0; y <= this.slices; y++) {
      vertex = new FSSVertex(offsetX + x * this.segmentWidth, offsetY - y * this.sliceHeight);
      vertices[x].push(vertex);
      this.vertices.push(vertex);
    }
  }

  // Add Triangles
  for (x = 0; x < this.segments; x++) {
    for (y = 0; y < this.slices; y++) {
      v0 = vertices[x + 0][y + 0];
      v1 = vertices[x + 0][y + 1];
      v2 = vertices[x + 1][y + 0];
      v3 = vertices[x + 1][y + 1];
      t0 = new FSSTriangle(v0, v1, v2);
      t1 = new FSSTriangle(v2, v1, v3);
      this.triangles.push(t0, t1);
    }
  }
}

Plane.prototype = Object.create(FSSGeometry.prototype);

export default Plane;
