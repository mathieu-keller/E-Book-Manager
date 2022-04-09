import React from 'react';
import {Route, Routes} from 'react-router-dom';
import Library from "./Library/Library";

const App = (): JSX.Element => {

  return (
    <>
      <Routes>
        <Route path="/" element={<Library/>}/>
        <Route path="books" element={<Library/>}/>
      </Routes>
    </>
  );
};

export default App;
