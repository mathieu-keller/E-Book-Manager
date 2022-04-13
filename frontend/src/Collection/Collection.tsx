import React, {useEffect} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {CollectionType} from "./Collection.type";
import ItemCard from "../UI/ItemCard";
import {BookType} from "../Book/Book.type";
import {useDispatch, useSelector} from "react-redux";
import {AppStore} from "../Store/Store.types";
import {CollectionReducer} from "../Reducers/CollectionReducer";

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
      <h1 className="text-center font-bold text-4xl m-5">{name}</h1>
      <hr/>
      <div className="flex flex-wrap flex-row">
        {collection.map((book: BookType): JSX.Element => <ItemCard
          key={book.id}
          name={book.title}
          cover={book.cover}
          onClick={(): void => openItem(book)}
        />)}
      </div>
    </div>
  );
};

export default Collection;
