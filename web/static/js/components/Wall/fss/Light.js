import FSSObject from "./Object.js";
import FSSColor from "./Color.js";
import FSSVector3 from "./Vector3.js";

/**
 *
 * @param {string} ambient
 * @param {string} diffuse
 */
function Light(ambient, diffuse) {
  FSSObject.call(this);
  this.ambient = new FSSColor(ambient || "#FFFFFF");
  this.diffuse = new FSSColor(diffuse || "#FFFFFF");
  this.ray = FSSVector3.create();
}

Light.prototype = Object.create(FSSObject.prototype);

export default Light;
