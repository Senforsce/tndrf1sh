import FSS from "./fss.js";
import XMath from "./fss/XMath.js";
import { MESH } from "./wall-mesh.js";
import { LIGHT } from "./wall-lights.js";
import { scene } from "./wall-scene.js";
import { geometry, center, attractor, now } from "./wall-globals.js";

export default function update() {
  var ox,
    oy,
    oz,
    l,
    light,
    v,
    vertex,
    offset = MESH.depth / 2;

  // Update Bounds
  FSS.Vector3.copy(LIGHT.bounds, center);
  FSS.Vector3.multiplyScalar(LIGHT.bounds, LIGHT.xyScalar);

  // Update Attractor
  FSS.Vector3.setZ(attractor, LIGHT.zOffset);

  // Overwrite the Attractor position
  if (LIGHT.autopilot) {
    ox = Math.sin(LIGHT.step[0] * now * LIGHT.speed);
    oy = Math.cos(LIGHT.step[1] * now * LIGHT.speed);
    FSS.Vector3.set(attractor, LIGHT.bounds[0] * ox, LIGHT.bounds[1] * oy, LIGHT.zOffset);
  }

  // Animate Lights
  for (l = scene.lights.length - 1; l >= 0; l--) {
    light = scene.lights[l];

    // Reset the z position of the light
    FSS.Vector3.setZ(light.position, LIGHT.zOffset);

    // Calculate the force Luke!
    var D = XMath.clamp(FSS.Vector3.distanceSquared(light.position, attractor), LIGHT.minDistance, LIGHT.maxDistance);
    var F = (LIGHT.gravity * light.mass) / D;
    FSS.Vector3.subtractVectors(light.force, attractor, light.position);
    FSS.Vector3.normalise(light.force);
    FSS.Vector3.multiplyScalar(light.force, F);

    // Update the light position
    FSS.Vector3.set(light.acceleration);
    FSS.Vector3.add(light.acceleration, light.force);
    FSS.Vector3.add(light.velocity, light.acceleration);
    FSS.Vector3.multiplyScalar(light.velocity, LIGHT.dampening);
    FSS.Vector3.limit(light.velocity, LIGHT.minLimit, LIGHT.maxLimit);
    FSS.Vector3.add(light.position, light.velocity);
  }

  // Animate Vertices
  for (v = geometry.vertices.length - 1; v >= 0; v--) {
    vertex = geometry.vertices[v];
    ox = Math.sin(vertex.time + vertex.step[0] * now * MESH.speed);
    oy = Math.cos(vertex.time + vertex.step[1] * now * MESH.speed);
    oz = Math.sin(vertex.time + vertex.step[2] * now * MESH.speed);
    FSS.Vector3.set(
      vertex.position,
      MESH.xRange * geometry.segmentWidth * ox,
      MESH.yRange * geometry.sliceHeight * oy,
      MESH.zRange * offset * oz - offset
    );
    FSS.Vector3.add(vertex.position, vertex.anchor);
  }

  // Set the Geometry to dirty
  geometry.dirty = true;
}
