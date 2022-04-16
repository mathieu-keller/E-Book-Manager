import React, {useEffect, useState} from 'react';
import {useLocation, useNavigate} from "react-router-dom";
import {BookType} from "../Book/Book.type";
import ItemCard from "../UI/ItemCard";

const Search = (): JSX.Element => {
  const loc = useLocation();
  const [books, setBooks] = useState<BookType[]>([]);

  const searchBooks = async (search: string): Promise<void> => {
    const response = await fetch('/book' + search);
    if (response.ok) {
      setBooks(await response.json());
    }
  };

  useEffect((): void => {
    if (loc.search.length > 0) {
      searchBooks(loc.search);
    }
  }, [loc.search]);
  const navigator = useNavigate();
  const openItem = (book: BookType): void => {
    navigator(`/book/${book.title}`);
  };
  return (
    <div className="flex flex-wrap flex-row justify-center">
      {books.map((book: BookType): JSX.Element => <ItemCard
        key={book.id}
        name={book.title}
        cover={book.cover}
        id={book.id}
        type="book"
        onClick={(): void => openItem(book)}
      />)}
    </div>
  );
};

export default Search;
