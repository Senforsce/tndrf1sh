// @ts-nocheck
import React, { useState, useEffect, useRef } from "react";
import PropTypes from "prop-types";
import ReactDOM from "react-dom";
import mapboxgl from "mapbox-gl";

import places from "./messages-mock";
import { BackHome } from "../BackHome/mod";
import { ContactMessageTooltip } from "./Tooltip/mod";
import "./style.css";

//@ts-ignore
mapboxgl.accessToken = process.env.REACT_APP_MAPBOX_TOKEN; //load from server on production
const mapStyles = process.env.REACT_APP_MAPBOX_STYLES;

function encode(data) {
  return Object.keys(data)
    .map((key) => encodeURIComponent(key) + "=" + encodeURIComponent(data[key]))
    .join("&");
}

const handleSubmit = (event) => {
  event.preventDefault();
  const formData = new FormData(event.target);
  let body = {};
  for (let [key, value] of formData.entries()) {
    console.log(key, value);
    body[key] = value;
  }
  console.log({ body });
  fetch("/", {
    method: "POST",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body: encode(body),
  })
    .then(() => alert("Thanks !"))
    .catch((error) => alert(error));
};
//TODO Refactor

const handleSubmitMap = (event) => {
  event.preventDefault();
  const formData = new FormData(event.target);
  let body = {};
  for (let [key, value] of formData.entries()) {
    body[key] = value;
  }
  console.log({ body });
  fetch("/", {
    method: "POST",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body: encode(body),
  })
    .then(() => alert("Thanks !"))
    .catch((error) => alert(error));
};

/**
 *
 * @param {{
 * home: {x:string, y:string},
 * setPosition: Function}
 * } props
 */
export const ContactMap = ({ home, setPosition }) => {
  const [coords, setCoords] = useState({
    lat: 5,
    lng: 34,
    zoom: 8,
  });
  const mapRef = useRef("");
  const tooltipRef = useRef(new mapboxgl.Popup({ offset: 0 }));

  useEffect(() => {
    let map = {};
    if (window.innerWidth < 1024) return;

    if (mapRef.current) {
      map = new mapboxgl.Map({
        container: mapRef.current,
        style: mapStyles,
        center: [-77.04, 38.907],
        // center: [coords.lng, coords.lat],
        zoom: coords.zoom,
      });

      map.on("move", () => {
        setCoords({
          lng: map.getCenter().lng.toFixed(4),
          lat: map.getCenter().lat.toFixed(4),
          zoom: map.getZoom().toFixed(2),
        });
      });

      map.on("load", function () {
        // Add a GeoJSON source containing place coordinates and information.
        map.addSource("places", {
          type: "geojson",
          data: places,
        });

        map.addLayer({
          id: "poi-labels",
          type: "symbol",
          source: "places",
          layout: {
            "text-field": ["get", "description"],
            "text-variable-anchor": ["top", "bottom", "left", "right"],
            "text-radial-offset": 0.5,
            "text-justify": "auto",
            "icon-image": ["concat", ["get", "icon"], "-15"],
          },
        });
      });

      // Create tooltip node
      const tooltipNode = document.createElement("div");
      ReactDOM.render(<ContactMessageTooltip name="Abdoul SY" message="Yatta" />, tooltipNode);

      // Set tooltip on map
      tooltipRef.current.setLngLat({ lon: 24, lat: 28 }).setDOMContent(tooltipNode).addTo(map);
    }

    return () => map.remove();
  }, [mapRef]);

  return (
    <div>
      <BackHome
        offset={50}
        action={() => {
          setPosition(home);
        }}
      />
      <div className="contactFormInbox contactable">
        <h4>Send a Direct Message</h4>
        <form name="contact-direct" method="POST" onSubmit={handleSubmit} netlify>
          <input type="hidden" name="form-name" value="contact-direct" />
          <input type="text" name="fullName" placeholder="Full Name" required />
          <input type="email" name="email" placeholder="Email" required />
          <input type="phone" name="phone" placeholder="phone (optional)" />
          <textarea name="messageInbox" placeholder="enter your message"></textarea>
          <input type="submit" value="Send" />
        </form>
      </div>
      <div className="contactFormMap contactable">
        <h4>Add a Map Note</h4>
        <form name="contact-map" method="POST" onSubmit={handleSubmitMap} netlify>
          <input type="hidden" name="form-name" value="contact-map" />
          <input type="text" name="fullName" placeholder="Name" required />
          <textarea name="message" placeholder="enter your comment" />
          <input type="hidden" name="lat" />
          <input type="hidden" name="lng" />
          <input type="submit" value="Send" />
        </form>
      </div>
      <div ref={mapRef} className="mapContainer" />
    </div>
  );
};

ContactMap.propTypes = {
  home: PropTypes.shape({
    x: PropTypes.string.isRequired,
    y: PropTypes.string.isRequired,
  }).isRequired,
  setPosition: PropTypes.func.isRequired,
};
