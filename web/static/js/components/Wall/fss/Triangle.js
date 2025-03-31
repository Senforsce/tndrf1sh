import FSSVector3 from "./Vector3.js";
import FSSColor from "./Color.js";
import FSSVertex from "./Vertex.js";

/**
 *
 * @param {FSSVertex} a
 * @param {FSSVertex} b
 * @param {FSSVertex} c
 */
function Triangle(a, b, c) {
  this.a = a || new FSSVertex();
  this.b = b || new FSSVertex();
  this.c = c || new FSSVertex();
  this.vertices = [this.a, this.b, this.c];
  this.u = FSSVector3.create();
  this.v = FSSVector3.create();
  this.centroid = FSSVector3.create();
  this.normal = FSSVector3.create();
  this.color = new FSSColor();
  this.polygon = document.createElementNS("http://www.w3.org/2000/svg", "polygon");
  this.polygon.setAttributeNS(null, "stroke-linejoin", "round");
  this.polygon.setAttributeNS(null, "stroke-miterlimit", "1");
  this.polygon.setAttributeNS(null, "stroke-width", "1");
  this.computeCentroid();
  this.computeNormal();
}

Triangle.prototype = {
  computeCentroid: function () {
    this.centroid[0] = this.a.position[0] + this.b.position[0] + this.c.position[0];
    this.centroid[1] = this.a.position[1] + this.b.position[1] + this.c.position[1];
    this.centroid[2] = this.a.position[2] + this.b.position[2] + this.c.position[2];
    FSSVector3.divideScalar(this.centroid, 3);
    return this;
  },
  computeNormal: function () {
    FSSVector3.subtractVectors(this.u, this.b.position, this.a.position);
    FSSVector3.subtractVectors(this.v, this.c.position, this.a.position);
    FSSVector3.crossVectors(this.normal, this.u, this.v);
    FSSVector3.normalise(this.normal);
    return this;
  },
};

export default Triangle;
