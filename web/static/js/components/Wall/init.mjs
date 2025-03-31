import initialise from "./wall-init.js"

export default function init() {
     window.ASMeshDiffuse = "#555555";

    setTimeout(() => {
        initialise(document.getElementById('wall'), document.getElementById('output'));
    }, 100);
}