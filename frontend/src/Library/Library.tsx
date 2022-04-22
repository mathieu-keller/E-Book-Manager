import React, {useEffect, useState} from 'react';
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

  const [loadingDiv, setLoadingDiv] = useState<boolean>(false);
  const getLibraryItems = async (page: number): Promise<LibraryItemType[]> => {
    const response = await Rest.get<LibraryItemType[]>(`/all?page=${page}`);
    return response.data;
  };

  const shouldLoadNextPage = (): void => {
    const element = document.querySelector('#loading-trigger');
    const position = element?.getBoundingClientRect();

    if (position !== undefined && position.top >= 0 && position.bottom <= window.innerHeight) {
      setLoadingDiv(true);
    } else {
      setLoadingDiv(false);
    }
  };

  window.addEventListener('scroll', shouldLoadNextPage);
  const page = useSelector<AppStore, number>((store): number => store.libraryItems.page);
  const allLoaded = useSelector<AppStore, boolean>((store): boolean => store.libraryItems.allItemsLoaded);
  const dispatch = useDispatch();
  useEffect((): void => {
    if (loadingDiv && !allLoaded) {
      setLoadingDiv(false);
      getLibraryItems(page + 1)
        .then((r): void => {
          if (r.length === 0) {
            dispatch(LibraryItemReducer.actions.setAllLoaded(true));
          } else {
            dispatch(LibraryItemReducer.actions.addAll(r));
          }
        });
      dispatch(LibraryItemReducer.actions.setPage(page + 1));
    }
  }, [loadingDiv]);

  useEffect((): void => {
    shouldLoadNextPage();
  }, [items]);

  useEffect((): (() => void) => {
    dispatch(ApplicationReducer.actions.reset());
    shouldLoadNextPage();
    return (): void => {
      window.removeEventListener('scroll', shouldLoadNextPage);
    };
  }, []);
  const navigator = useNavigate();
  const openItem = (item: LibraryItemType): void => {
    navigator(`/${item.itemType}/${item.title}`);
  };


  return (
    <>
      <ItemsGrid<LibraryItemType> onClick={(item): void => openItem(item)} items={items}/>
      {allLoaded ? null : <div id="loading-trigger" className="m-5 text-center text-5xl">Loading....</div>}
    </>
  );
};

export default Library;
