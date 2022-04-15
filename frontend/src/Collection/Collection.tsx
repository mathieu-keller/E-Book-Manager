import React, {useEffect} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {CollectionType} from "./Collection.type";
import {BookType} from "../Book/Book.type";
import {useDispatch, useSelector} from "react-redux";
import {AppStore} from "../Store/Store.types";
import {CollectionReducer} from "../Reducers/CollectionReducer";
import {ApplicationReducer} from "../Reducers/HeaderReducer";
import ItemsGrid from "../UI/ItemsGrid";

const Collection = (): JSX.Element => {
  const {title} = useParams<{ title: string }>();
  if (title === undefined) {
    throw new Error("name path param missing!");
  }
  const collection = useSelector((store: AppStore): BookType[] => store.collections[title]);

  const getCollection = async (): Promise<CollectionType> => {
    const response = await fetch('/collection?title=' + title);
    return response.json();
  };

  const dispatch = useDispatch();
  useEffect((): void => {
    if (collection === undefined) {
      getCollection()
        .then((c: CollectionType): void => {
          dispatch(CollectionReducer.actions.set({collection: c.title, books: c.books}));
        });
    }
    dispatch(ApplicationReducer.actions.setHeaderText(title));
  }, [title]);

  const navigator = useNavigate();
  const openItem = (book: BookType): void => {
    navigator(`/book/${book.title}`);
  };

  if (collection === undefined) {
    return <div>loading...</div>;
  }
  return (
    <>
      <ItemsGrid<BookType> onClick={(item) => openItem(item)} items={collection}/>
    </>
  );
};

export default Collection;
