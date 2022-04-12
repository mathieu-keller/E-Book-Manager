import React, {useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {BookType} from "./Book.type";
import {useSelector} from "react-redux";
import {AppStore, CollectionStore} from "../Store/Store.types";

const Book = (): JSX.Element => {
  const {title} = useParams();
  const [book, setBook] = useState<BookType | null>(null);

  const getBook = async (): Promise<BookType> => {
    const response = await fetch('/book/' + title);
    return response.json();
  };

  const collections = useSelector((store: AppStore): CollectionStore => store.collections);

  useEffect((): void => {
    const storedBook = Object.values(collections).flat().find((b): boolean => b.name === title);
    if (storedBook !== undefined) {
      setBook(storedBook);
    } else {
      getBook()
        .then((b: BookType): void => setBook(b));
    }
  }, [title]);

  if (book == null) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>{book.name}</h1>
      <img src={`data:image/jpeg;base64,${book.cover}`} alt={`cover picture of ${book.name}`}/>
      {book.author}
      <a href={`/download/${book.id}`} download={`${book.name}.epub`}>Download</a>
    </div>
  );
};

export default Book;
