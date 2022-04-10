import React, {useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {BookType} from "./Book.type";


const Book = () => {
  const {title} = useParams();
  const [book, setBook] = useState<BookType | null>(null);

  const getCollection = async () => {
    const response = await fetch('/book/' + title);
    return response.json();
  };

  useEffect(() => {
    getCollection()
      .then(c => setBook(c));
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
