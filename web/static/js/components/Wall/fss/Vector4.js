import FSSArray from "./array.js";

const Vector4 = {
  /**
   *
   * @param {number} [x]
   * @param {number} [y]
   * @param {number} [z]
   * @param {number} [w]
   */
  create: function (x, y, z, w) {
    var vector = new FSSArray(4);
    this.set(vector, x, y, z, w);
    return vector;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} [x]
   * @param {number} [y]
   * @param {number} [z]
   * @param {number} [w]
   */
  set: function (target, x, y, z, w) {
    target[0] = x || 0;
    target[1] = y || 0;
    target[2] = z || 0;
    target[3] = w || 0;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} [x]
   */
  setX: function (target, x) {
    target[0] = x || 0;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} [y]
   */
  setY: function (target, y) {
    target[1] = y || 0;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} [z]
   */
  setZ: function (target, z) {
    target[2] = z || 0;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} [w]
   */
  setW: function (target, w) {
    target[3] = w || 0;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   */
  add: function (target, a) {
    target[0] += a[0];
    target[1] += a[1];
    target[2] += a[2];
    target[3] += a[3];
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   * @param {Array<any>|Float32Array} b
   */
  multiplyVectors: function (target, a, b) {
    target[0] = a[0] * b[0];
    target[1] = a[1] * b[1];
    target[2] = a[2] * b[2];
    target[3] = a[3] * b[3];
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} s
   */
  multiplyScalar: function (target, s) {
    target[0] *= s;
    target[1] *= s;
    target[2] *= s;
    target[3] *= s;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} value
   */
  min: function (target, value) {
    if (target[0] < value) {
      target[0] = value;
    }
    if (target[1] < value) {
      target[1] = value;
    }
    if (target[2] < value) {
      target[2] = value;
    }
    if (target[3] < value) {
      target[3] = value;
    }
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} value
   */
  max: function (target, value) {
    if (target[0] > value) {
      target[0] = value;
    }
    if (target[1] > value) {
      target[1] = value;
    }
    if (target[2] > value) {
      target[2] = value;
    }
    if (target[3] > value) {
      target[3] = value;
    }
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} min
   * @param {number} max
   */
  clamp: function (target, min, max) {
    this.min(target, min);
    this.max(target, max);
    return this;
  },
};

export default Vector4;
