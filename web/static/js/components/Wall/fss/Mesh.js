import FSSObject from "./Object.js";
import FSSGeometry from "./Geometry.js";
import FSSMaterial from "./Material.js";
import FSSVector3 from "./Vector3.js";
import FSSVector4 from "./Vector4.js";

const Options = {
  FRONT: 0,
  BACK: 1,
  DOUBLE: 2,
  SVGNS: "http://www.w3.org/2000/svg",
};

/**
 *
 * @param {FSSGeometry} geometry
 * @param {FSSMaterial} material
 */
export default function FSSMesh(geometry, material) {
  FSSObject.call(this);
  this.geometry = geometry || new FSSGeometry();
  this.material = material || new FSSMaterial();
  this.side = Options.FRONT;
  this.visible = true;
}

FSSMesh.prototype = Object.create(FSSObject.prototype);

/**
 *
 * @param {*} lights
 * @param {*} calculate
 */
FSSMesh.prototype.update = function (lights, calculate) {
  var t, triangle, l, light, illuminance;

  // Update Geometry
  this.geometry.update();

  // Calculate the triangle colors
  if (calculate) {
    // Iterate through Triangles
    for (t = this.geometry.triangles.length - 1; t >= 0; t--) {
      triangle = this.geometry.triangles[t];

      // Reset Triangle Color
      FSSVector4.set(triangle.color.rgba);

      // Iterate through Lights
      for (l = lights.length - 1; l >= 0; l--) {
        light = lights[l];

        // Calculate Illuminance
        FSSVector3.subtractVectors(light.ray, light.position, triangle.centroid);
        FSSVector3.normalise(light.ray);
        illuminance = FSSVector3.dot(triangle.normal, light.ray);
        if (this.side === Options.FRONT) {
          illuminance = Math.max(illuminance, 0);
        } else if (this.side === Options.BACK) {
          illuminance = Math.abs(Math.min(illuminance, 0));
        } else if (this.side === Options.DOUBLE) {
          illuminance = Math.max(Math.abs(illuminance), 0);
        }

        // Calculate Ambient Light
        FSSVector4.multiplyVectors(this.material.slave.rgba, this.material.ambient.rgba, light.ambient.rgba);
        FSSVector4.add(triangle.color.rgba, this.material.slave.rgba);

        // Calculate Diffuse Light
        FSSVector4.multiplyVectors(this.material.slave.rgba, this.material.diffuse.rgba, light.diffuse.rgba);
        FSSVector4.multiplyScalar(this.material.slave.rgba, illuminance);
        FSSVector4.add(triangle.color.rgba, this.material.slave.rgba);
      }

      // Clamp & Format Color
      FSSVector4.clamp(triangle.color.rgba, 0, 1);
    }
  }
  return this;
};
