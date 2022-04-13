import React, {useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {BookType} from "./Book.type";
import {useSelector} from "react-redux";
import {AppStore, CollectionStore} from "../Store/Store.types";
import defaultCover from '../../public/default/cover.jpg';

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
    <>
      <h1>{book.title}</h1>
      <img
        src={book.cover !== null ? `data:image/jpeg;base64,${book.cover}` : defaultCover}
        alt={`cover picture of ${book.title}`}
      />
      {book.authors}
      <a href={`/download/${book.id}`} download={`${book.title}.epub`}>Download</a>
    </>
  );
};

export default Book;
