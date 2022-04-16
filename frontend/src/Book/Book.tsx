import React, {useEffect, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {BookType} from "./Book.type";
import {useDispatch, useSelector} from "react-redux";
import {AppStore, CollectionStore} from "../Store/Store.types";
import defaultCover from '../../public/default/cover.jpg';
import {ApplicationReducer} from "../Reducers/HeaderReducer";
import {LinkButton} from "../UI/Button";
import Badge from "../UI/Badge";
import Rest from "../Rest";

const Book = (): JSX.Element => {
  const {title} = useParams();
  const [book, setBook] = useState<BookType | null>(null);

  const getBook = async (): Promise<void> => {
    const response = await Rest.get<BookType>(`/book/${title}`);
    setBook(response.data);
  };
  const dispatch = useDispatch();
  useEffect((): void => {
    if (book !== null) {
      dispatch(ApplicationReducer.actions.setHeaderText(book.title));
    } else {
      dispatch(ApplicationReducer.actions.reset());
    }
  }, [book?.title]);

  const collections = useSelector((store: AppStore): CollectionStore => store.collections);

  useEffect((): void => {
    const storedBook = Object.values(collections).flat().find((b): boolean => b.title === title);
    if (storedBook !== undefined) {
      setBook(storedBook);
    } else {
      getBook();
    }
  }, [title]);

  const navigate = useNavigate();

  if (book == null) {
    return <div>Loading...</div>;
  }

  return (
    <>
      <div className="mt-10 flex justify-center">
        <div className="grid max-w-[80%]">
          <img
            src={book.cover !== null ? `data:image/jpeg;base64,${book.cover}` : defaultCover}
            alt={`cover picture of ${book.title}`}
          />
          <div className="grid-cols-1 grid h-max">
            <div className="m-5">
              <h1>Authors:</h1>
              {book.authors
                .map((author): JSX.Element => <Badge
                  key={author.id}
                  onClick={(): void => navigate(`/search?q=${author.name}`, {state: author.name})}
                  text={author.name}
                />)}
            </div>
            <div className="m-5">
              <h1>Subjects:</h1>
              {book.subjects
                .map((subject): JSX.Element => <Badge
                  key={subject.id}
                  onClick={(): void => navigate(`/search?q=${subject.name}`, {state: subject.name})}
                  text={subject.name}
                />)}
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
