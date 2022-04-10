import React from 'react';
import {HashRouter} from "react-router-dom";
import App from "./App";
import {createRoot} from 'react-dom/client';
import '../public/style.css';
import {Provider} from "react-redux";
import store from "./Store";

const container = document.getElementById('root');
if (container === null) {
  window.alert("no root container found!");
} else {
  const root = createRoot(container);
  root.render(
    <HashRouter>
      <Provider store={store}>
        <App/>
      </Provider>
    </HashRouter>);
}
