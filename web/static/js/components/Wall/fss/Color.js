import FSSUtils from "./utils.js";
import FSSVector4 from "./Vector4.js";
/**
 *
 * @param {string} [hex]
 * @param {any} [opacity]
 */
function Color(hex = "#000000", opacity) {
  this.rgba = FSSVector4.create();
  this.hex = hex;
  this.opacity = FSSUtils.isNumber(opacity) ? opacity : 1;
  this.set(this.hex, this.opacity);
}

Color.prototype = {
  /**
   *
   * @param {string} hex
   * @param {any} opacity
   */
  set: function (hex, opacity) {
    hex = hex.replace("#", "");
    var size = hex.length / 3;
    this.rgba[0] = parseInt(hex.substring(size * 0, size * 1), 16) / 255;
    this.rgba[1] = parseInt(hex.substring(size * 1, size * 2), 16) / 255;
    this.rgba[2] = parseInt(hex.substring(size * 2, size * 3), 16) / 255;
    this.rgba[3] = FSSUtils.isNumber(opacity) ? opacity : this.rgba[3];
    return this;
  },
  /**
   *
   * @param {number} channel
   */
  hexify: function (channel) {
    var hex = Math.ceil(channel * 255).toString(16);
    if (hex.length === 1) {
      hex = "0" + hex;
    }
    return hex;
  },
  /**
   * @returns {string}
   */
  format: function () {
    var r = this.hexify(this.rgba[0]);
    var g = this.hexify(this.rgba[1]);
    var b = this.hexify(this.rgba[2]);
    this.hex = "#" + r + g + b;
    return this.hex;
  },
};

export default Color;
