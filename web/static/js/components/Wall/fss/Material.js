import FSSColor from "./Color.js";

/**
 *
 * @param {string} [ambient]
 * @param {string} [diffuse]
 */
export default function Material(ambient = "#444444", diffuse = "#FFFFFF") {
  this.ambient = new FSSColor(ambient);
  this.diffuse = new FSSColor(diffuse);
  this.slave = new FSSColor();
}
