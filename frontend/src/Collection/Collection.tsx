import React, {useEffect, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {CollectionType} from "./Collection.type";
import ItemCard from "../UI/ItemCard";
import {BookType} from "../Book/Book.type";

const Collection = () => {
  const {name} = useParams();
  const [collection, setCollection] = useState<CollectionType | null>(null);

  const getCollection = async () => {
    const response = await fetch('/collection?name=' + name);
    return response.json();
  };

  useEffect(() => {
    getCollection()
      .then(c => setCollection(c));
  }, [name]);

  const navigator = useNavigate();
  const openItem = (book: BookType) => {
    navigator(`/book/${book.name}`);
  };

  if (collection === null) {
    return <div>loading...</div>;
  }
  return (
    <div>
      <h1 className="text-center font-bold text-4xl m-5">{collection.name}</h1>
      <hr/>
      <div className="flex flex-wrap flex-row">
        {collection.books.map(book => <ItemCard key={book.id} name={book.name} cover={book.cover} onClick={() => openItem(book)}/>)}
      </div>
    </div>
  );
};

export default Collection;
