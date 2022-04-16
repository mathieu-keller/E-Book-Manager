import React, {useEffect} from 'react';
import {LibraryItemType} from "./LibraryItem.type";
import {useDispatch, useSelector} from "react-redux";
import {AppStore} from "../Store/Store.types";
import {LibraryItemReducer} from "../Reducers/LibraryItemReducer";
import {ApplicationReducer} from "../Reducers/HeaderReducer";
import ItemsGrid from "../UI/ItemsGrid";
import {useNavigate} from "react-router-dom";

const Library = (): JSX.Element => {
  const items = useSelector<AppStore, LibraryItemType[]>((store): LibraryItemType[] => store.libraryItems.items);

  const getLibraryItems = async (): Promise<LibraryItemType[]> => {
    if (items.length === 0) {
      const response = await fetch('/all');
      return response.json();
    }
    return Promise.reject();
  };
  const dispatch = useDispatch();
  useEffect((): void => {
    dispatch(ApplicationReducer.actions.reset());
    if (items.length === 0) {
      getLibraryItems()
        .then((res: LibraryItemType[]): void => {
          dispatch(LibraryItemReducer.actions.set(res));
        });
    }
  }, []);
  const navigator = useNavigate();
  const openItem = (item: LibraryItemType): void => {
    navigator(`/${item.itemType}/${item.title}`);
  };

  return (
    <>
      <ItemsGrid<LibraryItemType> onClick={(item): void => openItem(item)} items={items}/>
    </>
  );
};

export default Library;
