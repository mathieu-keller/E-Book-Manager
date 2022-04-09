import React from 'react';
import {Route, Routes} from 'react-router-dom';
import Library from "./Library/Library";
import Collection from "./Collection/Collection";

const App = (): JSX.Element => {

  return (
    <>
      <Routes>
        <Route path="/" element={<Library/>}/>
        <Route path="books" element={<Library/>}/>
        <Route path="collection/:name" element={<Collection/>}/>
      </Routes>
    </>
  );
};

export default App;
