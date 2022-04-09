import React, {useEffect, useState} from 'react';
import {LibraryItemType} from "./LibraryItem.type";
import LibraryItem from "./LibraryItem";

const Library = () => {
  const [items, setItems] = useState<LibraryItemType[]>([]);

  const getLibraryItems = async () => {
    const response = await fetch('/all');
    return response.json();
  };

  useEffect(() => {
    getLibraryItems().then(item => setItems(item));
  }, []);

  return (
    <>
      {items.map(item => <LibraryItem item={item} key={item.id + item.type}/>)}
    </>
  );
};

export default Library;
