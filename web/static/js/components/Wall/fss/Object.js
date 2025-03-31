import FSSVector3 from "./Vector3.js";

function FSSObject() {
  this.position = FSSVector3.create();
}

FSSObject.prototype = {
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

export default FSSObject;
