import React, { useState } from "react";
import { Provider } from "react-redux";
import PageWall from "./components/Wall/PageWall";
import { DevMenu, createPosition } from "./components/DevMenu/mod";
import { Matrix } from "./components/Matrix/Matrix";
import { MobilityConfigSlider } from "./components/MobilityConfigSlider/mod";
import { TextConfigSlider } from "./components/TextConfigSlider/mod";
import store from "./redux/store/store";
import dotenv from "dotenv";

dotenv.config();
const devEnv = process.env.REACT_APP_DEV_ENV;

const INITIAL_POS_X = 1;
const INITIAL_POS_Y = 1;
// TODO: move the logic outside of App
const pagesByName = {
  home: createPosition(INITIAL_POS_X, INITIAL_POS_Y),
  contact: createPosition(2, INITIAL_POS_Y),
  skills: createPosition(INITIAL_POS_X, 0),
  timeline: createPosition(0, INITIAL_POS_Y),
};

// TODO: Prevent Motion Sickeness Effect
const App = () => {
  const home = createPosition(INITIAL_POS_X, INITIAL_POS_Y);
  const sanitizedPage = window.location.pathname.replace(/\W/g, "");
  const currentPage = pagesByName[sanitizedPage] || home;
  const [position, setPosition] = useState(currentPage);
  const [appState, setAppState] = useState(0);
  const [selectedFontSize, selectFontSize] = useState("2em");

  window.onpopstate = () => {
    const sanitizedPage = window.location.hash.replace(/\W/g, "");
    const newPage = pagesByName[sanitizedPage] || home;

    setPosition(newPage);
  };

  return (
    <Provider store={store}>
      <div className="App">
        <PageWall appState={appState} />
        {devEnv && <DevMenu position={position} setPosition={setPosition} />}
        <Matrix
          position={position}
          setPosition={setPosition}
          home={home}
          setAppState={setAppState}
          appFontSize={selectedFontSize}
        />
        <MobilityConfigSlider />
        <TextConfigSlider selectFontSize={selectFontSize} />
      </div>
    </Provider>
  );
};

export default App;
