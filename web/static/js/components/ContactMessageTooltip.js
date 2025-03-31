import React from "react";
import PropTypes from "prop-types";

/**
 * ContactMessageTooltip Sets a message in a map
 * @param {{name: string, message:string}} props 
 */
export const ContactMessageTooltip = ({name, message}) => {
    return <div>
        <p>{name}</p>
        <p>{message}</p>
    </div>
};

ContactMessageTooltip.propTypes = {
    name: PropTypes.string.isRequired,
    message: PropTypes.string.isRequired
};