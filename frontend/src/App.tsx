import type { Component } from 'solid-js';

import { Route, Routes } from 'solid-app-router';
import Header from './Header/Header';
import { lazy } from 'solid-js';

const Book = lazy(() => import('./Book/Book'));
const Collection = lazy(() => import('./Collection/Collection'));
const Library = lazy(() => import('./Library/Library'));

const App: Component = () => {
  return (
    <>
      <Header/>
      <Routes>
        <Route path="/" element={<Library/>}/>
        <Route path="/collection/:collection" element={<Collection/>}/>
        <Route path="/book/:book" element={<Book/>}/>
      </Routes>
    </>
  );
};

export default App;
