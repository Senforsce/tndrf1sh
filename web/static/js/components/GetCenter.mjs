export default function getCenter(el) {
    const coords = document.querySelector(el).getBBox();
    const centerX = coords.x + coords.width / 2;
    const centerY = coords.y + coords.height / 2;
  
    return centerX + " " + centerY;
  }
  