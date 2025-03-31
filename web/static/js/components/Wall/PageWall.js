/* eslint-disable */
import React from "react";
// import { createMesh } from "./wall-mesh";
import initialise from "./wall-init";

const grandInitialise = function (wallElement, outputElement, config) {
  setTimeout(() => {
    initialise(wallElement, outputElement, config);
  }, 5000);

  // setTimeout(() => {
  //   createMesh("#3A0023");
  // }, 10000);
};

export default class Wall extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    if (window.innerWidth > 1024) {
      //Testing overriding the color of the mesh
      window.ASMeshDiffuse = "#555555";
      grandInitialise(this.refs.wall, this.refs.output, this.props.config);
    }
  }

  render() {
    return (
      <div id="wall" ref="wall" className="wall">
        <div id="output" ref="output"></div>
      </div>
    );
  }
}
