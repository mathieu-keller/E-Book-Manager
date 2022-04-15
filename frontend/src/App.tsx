import React from 'react';
import {Route, Routes} from 'react-router-dom';
import Library from "./Library/Library";
import Collection from "./Collection/Collection";
import Header from "./Header/Header";
import Book from "./Book/Book";
import Search from "./Search/Search";

const App = (): JSX.Element => {

  return (
    <div>
      <Header/>
      <Routes>
        <Route path="/" element={<Library/>}/>
        <Route path="books" element={<Library/>}/>
        <Route path="collection/:title" element={<Collection/>}/>
        <Route path="book/:title" element={<Book/>}/>
        <Route path="search" element={<Search/>}/>
      </Routes>
    </div>
  );
};

export default App;
