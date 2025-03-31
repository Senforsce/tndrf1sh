import XMath from "./fss/XMath.js";
import { LIGHT } from "./wall-lights.js";
import { renderer } from "./wall-renderer.js";
import { scene } from "./wall-scene.js";
import { RENDER, CANVAS, SVG } from "./configuration.js";

export default function render() {
  renderer.render(scene);

  // Draw Lights
  if (LIGHT.draw) {
    var l, lx, ly, light;
    for (l = scene.lights.length - 1; l >= 0; l--) {
      light = scene.lights[l];
      lx = light.position[0];
      ly = light.position[1];
      switch (RENDER.renderer) {
        case CANVAS:
          renderer.context.lineWidth = 0.5;
          renderer.context.beginPath();
          renderer.context.arc(lx, ly, 10, 0, XMath.PIM2);
          renderer.context.strokeStyle = light.ambientHex;
          renderer.context.stroke();
          renderer.context.beginPath();
          renderer.context.arc(lx, ly, 4, 0, XMath.PIM2);
          renderer.context.fillStyle = light.diffuseHex;
          renderer.context.fill();
          break;
        case SVG:
          lx += renderer.halfWidth;
          ly = renderer.halfHeight - ly;
          light.core.setAttributeNS(null, "fill", light.diffuseHex);
          light.core.setAttributeNS(null, "cx", lx);
          light.core.setAttributeNS(null, "cy", ly);
          renderer.element.appendChild(light.core);
          light.ring.setAttributeNS(null, "stroke", light.ambientHex);
          light.ring.setAttributeNS(null, "cx", lx);
          light.ring.setAttributeNS(null, "cy", ly);
          renderer.element.appendChild(light.ring);
          break;
        default:
          return;
      }
    }
  }
}
