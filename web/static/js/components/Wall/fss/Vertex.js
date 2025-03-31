import FSSVector3 from "./Vector3.js";

/**
 *
 * @param {number} [x]
 * @param {number} [y]
 * @param {number} [z]
 */
function Vertex(x, y, z) {
  this.position = FSSVector3.create(x, y, z);
}

Vertex.prototype = {
  /**
   *
   * @param {number} x
   * @param {number} y
   * @param {number} z
   */
  setPosition: function (x, y, z) {
    FSSVector3.set(this.position, x, y, z);
    return this;
  },
};

export default Vertex;
