import React, {useEffect} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {CollectionType} from "./Collection.type";
import ItemCard from "../UI/ItemCard";
import {BookType} from "../Book/Book.type";
import {useDispatch, useSelector} from "react-redux";
import {AppStore} from "../Store/Store.types";
import {CollectionReducer} from "../Reducers/CollectionReducer";
import {ApplicationReducer} from "../Reducers/HeaderReducer";

const Collection = (): JSX.Element => {
  const {name} = useParams<{ name: string }>();
  if (name === undefined) {
    throw new Error("name path param missing!");
  }
  const collection = useSelector((store: AppStore): BookType[] => store.collections[name]);

  const getCollection = async (): Promise<CollectionType> => {
    const response = await fetch('/collection?name=' + name);
    return response.json();
  };

  const dispatch = useDispatch();
  useEffect((): void => {
    if (collection === undefined) {
      getCollection()
        .then((c: CollectionType): void => {
          dispatch(CollectionReducer.actions.set({collection: c.name, books: c.books}));
        });
    }
    dispatch(ApplicationReducer.actions.setHeaderText(name));
  }, [name]);

  const navigator = useNavigate();
  const openItem = (book: BookType): void => {
    navigator(`/book/${book.title}`);
  };

  if (collection === undefined) {
    return <div>loading...</div>;
  }
  return (
    <div>
      <div className="flex flex-wrap flex-row">
        {collection.map((book: BookType): JSX.Element => <ItemCard
          key={book.id}
          name={book.title}
          cover={book.cover}
          id={book.id}
          type="book"
          onClick={(): void => openItem(book)}
        />)}
      </div>
    </div>
  );
};

export default Collection;
