import React, {useEffect, useState} from 'react';
import {LibraryItemType} from "./LibraryItem.type";
import LibraryItem from "./LibraryItem";

const Library = () => {
  const [items, setItems] = useState<LibraryItemType[]>([]);

  const getLibraryItems = async () => {
    if(items.length === 0) {
      const response = await fetch('/all');
      return response.json();
    }
  };

  useEffect(() => {
    getLibraryItems().then(item => setItems(item));
  }, []);
  console.log("RENDER!", items)
  return (
    <div className="flex flex-row flex-wrap">
      {items.map(item => <LibraryItem item={item} key={item.id + item.type}/>)}
    </div>
  );
};

export default Library;
