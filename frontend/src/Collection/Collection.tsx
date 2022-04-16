import React, {useEffect} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {CollectionType} from "./Collection.type";
import {BookType} from "../Book/Book.type";
import {useDispatch, useSelector} from "react-redux";
import {AppStore} from "../Store/Store.types";
import {CollectionReducer} from "../Reducers/CollectionReducer";
import {ApplicationReducer} from "../Reducers/HeaderReducer";
import ItemsGrid from "../UI/ItemsGrid";
import Rest from "../Rest";

const Collection = (): JSX.Element => {
  const {title} = useParams<{ title: string }>();
  if (title === undefined) {
    throw new Error("name path param missing!");
  }
  const collection = useSelector((store: AppStore): BookType[] => store.collections[title]);


  const dispatch = useDispatch();
  const getCollection = async (): Promise<void> => {
    const response = await Rest.get<CollectionType>(`/collection?title=${title}`);
    const data = response.data;
    dispatch(CollectionReducer.actions.set({collection: data.title, books: data.books}));
  };


  useEffect((): void => {
    if (collection === undefined) {
      getCollection();
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
      <ItemsGrid<BookType> onClick={(item): void => openItem(item)} items={collection}/>
    </>
  );
};

export default Collection;
