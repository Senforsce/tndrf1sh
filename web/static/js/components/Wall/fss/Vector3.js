import FSSArray from "./array.js";

const Vector3 = {
  /**
   *
   * @param {number} [x]
   * @param {number} [y]
   * @param {number} [z]
   */
  create: function (x, y, z) {
    let vector = new FSSArray(3);
    this.set(vector, x, y, z);
    return vector;
  },
  /**
   *
   * @param {number[]} a
   */
  clone: function (a) {
    var vector = this.create();
    this.copy(vector, a);
    return vector;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} x
   * @param {number} y
   * @param {number} z
   */
  set: function (target, x = 0, y = 0, z = 0) {
    target[0] = x;
    target[1] = y;
    target[2] = z;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} x
   */
  setX: function (target, x) {
    target[0] = x || 0;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} y
   */
  setY: function (target, y) {
    target[1] = y || 0;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} z
   */
  setZ: function (target, z) {
    target[2] = z || 0;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   */
  copy: function (target, a) {
    target[0] = a[0];
    target[1] = a[1];
    target[2] = a[2];
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
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   * @param {Array<any>|Float32Array} b
   */
  addVectors: function (target, a, b) {
    target[0] = a[0] + b[0];
    target[1] = a[1] + b[1];
    target[2] = a[2] + b[2];
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} s
   */
  addScalar: function (target, s) {
    target[0] += s;
    target[1] += s;
    target[2] += s;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   */
  subtract: function (target, a) {
    target[0] -= a[0];
    target[1] -= a[1];
    target[2] -= a[2];
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   * @param {Array<any>|Float32Array} b
   */
  subtractVectors: function (target, a, b) {
    target[0] = a[0] - b[0];
    target[1] = a[1] - b[1];
    target[2] = a[2] - b[2];
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} s
   */
  subtractScalar: function (target, s) {
    target[0] -= s;
    target[1] -= s;
    target[2] -= s;
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   */
  multiply: function (target, a) {
    target[0] *= a[0];
    target[1] *= a[1];
    target[2] *= a[2];
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   * @param {Array<number>|Float32Array} b
   */
  multiplyVectors: function (target, a, b) {
    target[0] = a[0] * b[0];
    target[1] = a[1] * b[1];
    target[2] = a[2] * b[2];
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
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   */
  divide: function (target, a) {
    target[0] /= a[0];
    target[1] /= a[1];
    target[2] /= a[2];
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   * @param {Array<any>|Float32Array} b
   */
  divideVectors: function (target, a, b) {
    target[0] = a[0] / b[0];
    target[1] = a[1] / b[1];
    target[2] = a[2] / b[2];
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} s
   */
  divideScalar: function (target, s) {
    if (s !== 0) {
      target[0] /= s;
      target[1] /= s;
      target[2] /= s;
    } else {
      target[0] = 0;
      target[1] = 0;
      target[2] = 0;
    }
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   */
  cross: function (target, a) {
    var x = target[0];
    var y = target[1];
    var z = target[2];
    target[0] = y * a[2] - z * a[1];
    target[1] = z * a[0] - x * a[2];
    target[2] = x * a[1] - y * a[0];
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {Array<any>|Float32Array} a
   * @param {Array<any>|Float32Array} b
   */
  crossVectors: function (target, a, b) {
    target[0] = a[1] * b[2] - a[2] * b[1];
    target[1] = a[2] * b[0] - a[0] * b[2];
    target[2] = a[0] * b[1] - a[1] * b[0];
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
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} min
   * @param {number} max
   */
  limit: function (target, min, max) {
    var length = this.length(target);
    if (min !== null && length < min) {
      this.setLength(target, min);
    } else if (max !== null && length > max) {
      this.setLength(target, max);
    }
    return this;
  },
  /**
   *
   * @param {Array<any>|Float32Array} a
   * @param {Array<any>|Float32Array} b
   */
  dot: function (a, b) {
    return a[0] * b[0] + a[1] * b[1] + a[2] * b[2];
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   */
  normalise: function (target) {
    return this.divideScalar(target, this.length(target));
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   */
  negate: function (target) {
    return this.multiplyScalar(target, -1);
  },
  /**
   *
   * @param {Array<any>|Float32Array} a
   * @param {Array<any>|Float32Array} b
   */
  distanceSquared: function (a, b) {
    var dx = a[0] - b[0];
    var dy = a[1] - b[1];
    var dz = a[2] - b[2];
    return dx * dx + dy * dy + dz * dz;
  },
  /**
   *
   * @param {Array<any>|Float32Array} a
   * @param {Array<any>|Float32Array} b
   */
  distance: function (a, b) {
    return Math.sqrt(this.distanceSquared(a, b));
  },
  /**
   *
   * @param {Array<any>|Float32Array} a
   */
  lengthSquared: function (a) {
    return a[0] * a[0] + a[1] * a[1] + a[2] * a[2];
  },
  /**
   *
   * @param {Array<any>|Float32Array} a
   */
  length: function (a) {
    return Math.sqrt(this.lengthSquared(a));
  },
  /**
   *
   * @param {Array<any>|Float32Array} target
   * @param {number} l
   */
  setLength: function (target, l) {
    var length = this.length(target);
    if (length !== 0 && l !== length) {
      this.multiplyScalar(target, l / length);
    }
    return this;
  },
};

export default Vector3;
