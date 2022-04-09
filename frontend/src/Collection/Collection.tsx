import React, {useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {CollectionType} from "./Collection.type";

const Collection = () => {
  const {name} = useParams();
  const [collection, setCollection] = useState<CollectionType | null>(null);

  const getCollection = async () => {
    const response = await fetch('/collection?name=' + name);
    return response.json();
  };

  useEffect(() => {
    if (collection === null) {
      getCollection()
        .then(c => setCollection(c));
    }
  }, [name]);

  if (collection === null) {
    return <div>loading...</div>;
  }
  return (
    <>
      <h1 className="text-center font-bold text-4xl m-5">{collection.name}</h1>
      <hr/>
      <div className="flex flex-wrap flex-row">
        {collection.books.map(book =>
          <div className="m-5 flex max-w-sm flex-col shadow">
            <img src={`data:image/png;base64,${book.cover}`} alt={`cover picture of ${book.name}`}/>
            <h1 className="text-center break-words text-2xl font-bold">{book.name}</h1>
          </div>)}
      </div>
    </>
  );
};

export default Collection;
