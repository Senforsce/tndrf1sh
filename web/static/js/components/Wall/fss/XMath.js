function XMath() {}

XMath.prototype = Math;

XMath.PIM2 = Math.PI * 2;
XMath.PID2 = Math.PI / 2;
/**
 *
 * @param {number} min
 * @param {number} max
 */
XMath.randomInRange = function (min, max) {
  return min + (max - min) * Math.random();
};
/**
 *
 * @param {number} value
 * @param {number} min
 * @param {number} max
 */
XMath.clamp = function (value, min, max) {
  value = Math.max(value, min);
  value = Math.min(value, max);
  return value;
};

export default XMath;
