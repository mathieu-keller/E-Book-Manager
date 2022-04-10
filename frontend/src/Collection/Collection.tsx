import React, {useEffect, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {CollectionType} from "./Collection.type";
import ItemCard from "../UI/ItemCard";
import {BookType} from "../Book/Book.type";

const Collection = (): JSX.Element => {
  const {name} = useParams<{ name: string }>();
  const [collection, setCollection] = useState<CollectionType | null>(null);

  const getCollection = async (): Promise<CollectionType> => {
    const response = await fetch('/collection?name=' + name);
    return response.json();
  };

  useEffect((): void => {
    getCollection()
      .then((c: CollectionType): void => setCollection(c));
  }, [name]);

  const navigator = useNavigate();
  const openItem = (book: BookType): void => {
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
        {collection.books.map((book: BookType): JSX.Element => <ItemCard
          key={book.id}
          name={book.name}
          cover={book.cover}
          onClick={(): void => openItem(book)}
        />)}
      </div>
    </div>
  );
};

export default Collection;
