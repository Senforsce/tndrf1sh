import React from "react";
import PropTypes from "prop-types";

import Logo from "../Logo/mod";
import { ConnectedSmartParagraph } from "../SmartParagraph/mod";
import { StyledHome, StyledHomeTitle } from "./styled";

/**
 * Home
 * The Home Component:
 * shows the SVG logo, a stylized version of my name
 * and a smart paragraph revealing informations and helping you navigate the site
 * @param {{setPosition: Function, appFontSize: string }} props
 */
export const Home = ({ setPosition, appFontSize }) => {
  return (
    <>
      <StyledHome role="main">
        <Logo />
        <StyledHomeTitle>Abdoul Sy</StyledHomeTitle>
        <ConnectedSmartParagraph setPosition={setPosition} fontSize={appFontSize} />
      </StyledHome>
    </>
  );
};

Home.propTypes = {
  appFontSize: PropTypes.string.isRequired,
  setPosition: PropTypes.func.isRequired,
};
