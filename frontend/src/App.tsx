import React from 'react';
import {Route, Routes} from 'react-router-dom';
import Library from "./Library/Library";
import Collection from "./Collection/Collection";
import Header from "./Header/Header";

const App = (): JSX.Element => {

  return (
    <div>
      <Header/>
      <Routes>
        <Route path="/" element={<Library/>}/>
        <Route path="books" element={<Library/>}/>
        <Route path="collection/:name" element={<Collection/>}/>
      </Routes>
    </div>
  );
};

export default App;
