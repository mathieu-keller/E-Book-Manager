import React, {useEffect, useState, useTransition} from 'react';
import {LibraryItemType} from "./LibraryItem.type";
import LibraryItem from "./LibraryItem";

const Library = () => {
  const [items, setItems] = useState<LibraryItemType[]>([]);

  const getLibraryItems = async (): Promise<LibraryItemType[]> => {
    if(items.length === 0) {
      const response = await fetch('/all');
      return response.json();
    }
    return Promise.reject();
  };

  useEffect(() => {
      getLibraryItems().then(res => setItems(res));
  }, []);

  return (
    <div className="flex flex-row flex-wrap">
      {items.map(item => <LibraryItem item={item} key={`${item.id} - ${item.type}`}/>)}
    </div>
  );
};

export default Library;
