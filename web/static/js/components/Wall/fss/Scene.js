import FSSLight from "./Light.js";
import FSSMesh from "./Mesh.js";

export default function FSSScene() {
  /** @type {FSSMesh[]} */
  this.meshes = [];
  /** @type {FSSLight[]} */
  this.lights = [];
}

FSSScene.prototype = {
  /**
   *
   * @param {FSSLight | FSSMesh} object
   */
  add: function (object) {
    if (object instanceof FSSMesh && !~this.meshes.indexOf(object)) {
      this.meshes.push(object);
    } else if (object instanceof FSSLight && !~this.lights.indexOf(object)) {
      this.lights.push(object);
    }
    return this;
  },
  /**
   *
   * @param {FSSLight | FSSMesh} object
   */
  remove: function (object) {
    if (object instanceof FSSMesh && ~this.meshes.indexOf(object)) {
      this.meshes.splice(this.meshes.indexOf(object), 1);
    } else if (object instanceof FSSLight && ~this.lights.indexOf(object)) {
      this.lights.splice(this.lights.indexOf(object), 1);
    }
    return this;
  },
};
