import React, {useEffect} from 'react';
import {LibraryItemType} from "./LibraryItem.type";
import {useDispatch, useSelector} from "react-redux";
import {AppStore} from "../Store/Store.types";
import {LibraryItemReducer} from "../Reducers/LibraryItemReducer";
import {ApplicationReducer} from "../Reducers/HeaderReducer";
import ItemsGrid from "../UI/ItemsGrid";
import {useNavigate} from "react-router-dom";
import Rest from "../Rest";

const Library = (): JSX.Element => {
  const items = useSelector<AppStore, LibraryItemType[]>((store): LibraryItemType[] => store.libraryItems.items);

  const getLibraryItems = async (): Promise<void> => {
    const response = await Rest.get<LibraryItemType[]>('/all');
    dispatch(LibraryItemReducer.actions.set(response.data));
  };
  const dispatch = useDispatch();
  useEffect((): void => {
    dispatch(ApplicationReducer.actions.reset());
    if (items.length === 0) {
      getLibraryItems();
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
