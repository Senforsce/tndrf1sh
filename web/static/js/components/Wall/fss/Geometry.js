function Geometry() {
    this.vertices = [];
    /** @type {string | any[]} */
    this.triangles = [];
    this.dirty = false;
  }
  
  Geometry.prototype = {
    update: function () {
      if (this.dirty) {
        let t, triangle;
        for (t = this.triangles.length - 1; t >= 0; t--) {
          triangle = this.triangles[t];
          triangle.computeCentroid();
          triangle.computeNormal();
        }
        this.dirty = false;
      }
      return this;
    },
  };
  
  export default Geometry;
  