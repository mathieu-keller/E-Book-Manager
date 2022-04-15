import React, {useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {BookType} from "./Book.type";
import {useDispatch, useSelector} from "react-redux";
import {AppStore, CollectionStore} from "../Store/Store.types";
import defaultCover from '../../public/default/cover.jpg';
import {ApplicationReducer} from "../Reducers/HeaderReducer";
import {LinkButton} from "../UI/Button";
import Badge from "../UI/Badge";

const Book = (): JSX.Element => {
  const {title} = useParams();
  const [book, setBook] = useState<BookType | null>(null);

  const getBook = async (): Promise<BookType> => {
    const response = await fetch('/book/' + title);
    return response.json();
  };
  const dispatch = useDispatch();
  useEffect(() => {
    if (book !== null) {
      dispatch(ApplicationReducer.actions.setHeaderText(book.title));
    } else {
      dispatch(ApplicationReducer.actions.reset());
    }
  }, [book?.title]);

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
      <img
        src={book.cover !== null ? `data:image/jpeg;base64,${book.cover}` : defaultCover}
        alt={`cover picture of ${book.title}`}
      />
      {book.authors.map(author => <Badge key={author.id} onClick={()=> console.log(author)} text={author.name}/>)}
      {book.subjects.map(subject => <Badge key={subject.id} text={subject.name}/>)}
      <LinkButton href={`/download/${book.id}`} download={`${book.title}.epub`}>Download</LinkButton>
    </>
  );
};

export default Book;
