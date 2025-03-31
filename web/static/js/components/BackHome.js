import React from "react";
import PropTypes from "prop-types";
import { FaHome } from "react-icons/fa";

import "./styles.css";

/**
 * @param {object} props
 * @param {Function} props.action
 * @param {number} props.offset
 */
export const BackHome = ({ action, offset = 100 }) => {
  return (
    <div
      id="backHome"
      className="backHome"
      style={{ marginTop: `${offset}px` }}
      onClick={() => {
        action();
      }}
    >
      <FaHome />
    </div>
  );
};

BackHome.propTypes = {
  action: PropTypes.func.isRequired,
  offset: PropTypes.number,
};
