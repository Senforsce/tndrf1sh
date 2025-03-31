/* eslint-disable */
import { getLetterASelector, getLetterSSelector, getCircleSelector } from "./LogoSelectors.mjs";
import getCenter from "./GetCenter.mjs";
import updateAll from "./LogoGauge.mjs";

export default function init(top, left) {
    let w = window || {};
    const root = document.body;
    root.style.innerWidth = window.innerWidth
    root.style.innerHeight = window.innerHeight

    const svgEl = d3.select("#logo").select("svg");
    const logo = document.querySelector("#logo");
    const logoSVG = document.querySelector("#logo svg");
    const letterA = getLetterASelector(svgEl);
    const letterS = getLetterSSelector(svgEl);
    const circle = getCircleSelector(svgEl);

    if (logo && logoSVG) {
      logo.style.marginLeft = "-215px";
      logoSVG.style.left = w.innerWidth / 2 - 340;
    }

    window.addEventListener("resize", function () {
      logoSVG.style.left = w.innerWidth / 2 - 340;
    });

    //ROTATIONS CENTER POINTS
    const a_l_center = getCenter("g#Left");
    const a_r_center = getCenter("g#Right");
    // const s_o_center = getCenter("g#Outer_S");
    // const s_o_t_center = getCenter("g#OS_Top");
    // const s_o_m_center = getCenter("g#OS_Middle");
    // const s_o_b_center = getCenter("g#OS_Bottom");
    // const s_i_center = getCenter("g#Inner_S");
    // const s_i_t_center = getCenter("g#IS_Top");
    // const s_i_m_center = getCenter("g#IS_Middle");
    // const s_i_b_center = getCenter("g#IS_Bottom");

    updateAll(letterA, letterS, circle, a_l_center, a_r_center);
  }

