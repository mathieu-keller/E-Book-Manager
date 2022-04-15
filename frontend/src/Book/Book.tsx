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
      <div className="mt-10 flex justify-center">
        <div className=" grid-cols-2 grid">
          <img
            className="mr-10 float-left"
            src={book.cover !== null ? `data:image/jpeg;base64,${book.cover}` : defaultCover}
            alt={`cover picture of ${book.title}`}
          />
          <div className="float-left grid-cols-1 grid h-max">
            <div className="m-5">
              <h1>Authors:</h1>
              {book.authors.map(author => <Badge key={author.id} onClick={() => console.log(author)} text={author.name}/>)}
            </div>
            <div className="m-5">
              <h1>Subjects:</h1>
              {book.subjects.map(subject => <Badge key={subject.id} onClick={() => console.log(subject)} text={subject.name}/>)}
            </div>
          </div>
          <div className="col-start-1 col-end-3 mt-5 flex justify-self-stretch">
            <LinkButton href={`/download/${book.id}`} download={`${book.title}.epub`}>Download</LinkButton>
          </div>
        </div>
      </div>

    </>
  );
};

export default Book;
