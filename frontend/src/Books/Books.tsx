import React, {useEffect, useState} from 'react';
import {Book} from "./Book";

const Books = (): JSX.Element => {
  const [books, setBooks] = useState<Book[]>([]);
  useEffect(() => {
    fetch('/all')
      .then(r => {
        console.log(r);
        r.json().then(j => {
          console.log(j);
          setBooks(j);
        });
      });
  }, []);
  return (
    <div>
      {books.map(book => <div key={book.id}>
        <h1>{book.name}</h1>
        <img src={`data:image/png;base64,${book.cover}`} alt={`cover picture of ${book.name}`} style={{maxHeight: '30rem'}}/>
      </div>)}
    </div>
  );
};

export default Books;
