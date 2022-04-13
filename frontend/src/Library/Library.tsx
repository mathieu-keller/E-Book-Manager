import React, {useEffect} from 'react';
import {LibraryItemType} from "./LibraryItem.type";
import LibraryItem from "./LibraryItem";
import {useDispatch, useSelector} from "react-redux";
import {AppStore} from "../Store/Store.types";
import {LibraryItemReducer} from "../Reducers/LibraryItemReducer";
import {ApplicationReducer} from "../Reducers/HeaderReducer";

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

  return (
    <div className="flex flex-row flex-wrap">
      {items.map((item: LibraryItemType): JSX.Element => <LibraryItem item={item} key={`${item.itemType}-${item.name}`}/>)}
    </div>
  );
};

export default Library;
