import FSSCanvasRenderer from "./fss/CanvasRenderer.js";
import FSSRenderer from "./fss/Renderer.js";
import FSSUtils from "./fss/utils.js";
import FSSVector3 from "./fss/Vector3.js";
import FSSVector4 from "./fss/Vector4.js";
import FSSColor from "./fss/Color.js";
import FSSObject from "./fss/Object.js";
import FSSLight from "./fss/Light.js";
import FSSVertex from "./fss/Vertex.js";
import FSSTriangle from "./fss/Triangle.js";
import FSSGeometry from "./fss/Geometry.js";
import FSSPlane from "./fss/Plane.js";
import FSSMaterial from "./fss/Material.js";
import FSSMesh from "./fss/Mesh.js";
import FSSScene from "./fss/Scene.js";
import FSSWebGLRenderer from "./fss/WebGlRenderer.js";
import FSSSVGRenderer from "./fss/SVGRenderer.js";

const FSS = {
  Array: Float32Array,
  Utils: FSSUtils,
  Vector3: FSSVector3,
  Vector4: FSSVector4,
  Color: FSSColor,
  Object: FSSObject,
  Light: FSSLight,
  Vertex: FSSVertex,
  Triangle: FSSTriangle,
  Geometry: FSSGeometry,
  Plane: FSSPlane,
  Material: FSSMaterial,
  Mesh: FSSMesh,
  Scene: FSSScene,
  Renderer: FSSRenderer,
  CanvasRenderer: FSSCanvasRenderer,
  WebGLRenderer: FSSWebGLRenderer,
  SVGRenderer: FSSSVGRenderer,
};

export default FSS;
