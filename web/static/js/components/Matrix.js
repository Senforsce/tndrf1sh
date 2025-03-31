import React, { useEffect } from "react";
import PropTypes from "prop-types";
import { ContactMap } from "../ContactMap/mod";
import { Home } from "../Home/mod";
import { Timeline } from "../Timeline/mod";
import { SkillsTree } from "../SkillsTree/SkillsTree";
import * as Styles from "./styled";

/**
 * The matrix
 * Currently not efficient, will try to change the layout depending on
 * The performance of the device and its size
 * @param {object} props
 * @param {{x: string, y:string }} props.position
 * @param {{x:string, y:string}} props.home
 * @param {Function} props.setPosition
 * @param {string} props.appFontSize
 */
export const Matrix = (props) => {
  useEffect(() => {
    console.log({ pos: props.position });
  }, [props.position]);

  return (
    <Styles.StyledMatrixLayoutWrapper>
      <Styles.StyledMatrixLayout position={props.position}>
        <Styles.StyledTopRowFirstCell>1</Styles.StyledTopRowFirstCell>
        <Styles.StyledTopRowSecondCell>
          <SkillsTree />
        </Styles.StyledTopRowSecondCell>
        <Styles.StyledTopRowThirdCell>3</Styles.StyledTopRowThirdCell>
        <Styles.StyledMiddleRowFirstCell>
          <Timeline setPosition={props.setPosition} home={props.home} />
        </Styles.StyledMiddleRowFirstCell>
        <Styles.StyledMiddleRowSecondCell>
          <Home setPosition={props.setPosition} appFontSize={props.appFontSize} />
        </Styles.StyledMiddleRowSecondCell>
        <Styles.StyledMiddleRowThirdCell>
          <ContactMap setPosition={props.setPosition} home={props.home} />
        </Styles.StyledMiddleRowThirdCell>
        <Styles.StyledBottomRowFirstCell>7</Styles.StyledBottomRowFirstCell>
        <Styles.StyledBottomRowSecondCell>8</Styles.StyledBottomRowSecondCell>
        <Styles.StyledBottomRowThirdCell>9</Styles.StyledBottomRowThirdCell>
      </Styles.StyledMatrixLayout>
    </Styles.StyledMatrixLayoutWrapper>
  );
};

Matrix.propTypes = {
  appFontSize: PropTypes.string.isRequired,
  position: PropTypes.shape({
    x: PropTypes.string.isRequired,
    y: PropTypes.string.isRequired,
  }).isRequired,
  home: PropTypes.shape({
    x: PropTypes.string.isRequired,
    y: PropTypes.string.isRequired,
  }).isRequired,
  setPosition: PropTypes.func.isRequired,
};
