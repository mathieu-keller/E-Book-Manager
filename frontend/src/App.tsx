import React from 'react';
import {Route, Routes} from 'react-router-dom';
import Books from "./Books/Books";

const App = (): JSX.Element => {

  return (
    <>
      <Routes>
        <Route path="/" element={<Books/>}/>
        <Route path="books" element={<Books/>}/>
      </Routes>
    </>
  );
};

export default App;
