import React from 'react';
import {HashRouter} from "react-router-dom";
import App from "./App";
import {createRoot} from 'react-dom/client';
import '../public/style.css';

const container = document.getElementById('root');
if (container === null) {
  window.alert("no root container found!");
} else {
  const root = createRoot(container);
  root.render(
    <HashRouter>
      <App/>
    </HashRouter>);
}
