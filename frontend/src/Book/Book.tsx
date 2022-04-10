import React, {useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {BookType} from "./Book.type";


const Book = (): JSX.Element => {
  const {title} = useParams();
  const [book, setBook] = useState<BookType | null>(null);

  const getCollection = async (): Promise<BookType> => {
    const response = await fetch('/book/' + title);
    return response.json();
  };

  useEffect((): void => {
    getCollection()
      .then((b: BookType): void => setBook(b));
  }, [title]);

  if (book == null) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>{book.name}</h1>
      <img src={`data:image/jpeg;base64,${book.cover}`} alt={`cover picture of ${book.name}`}/>
    </div>
  );
};

export default Book;
