import React from 'react';
import {LibraryItemType} from "./LibraryItem.type";
import {useNavigate} from "react-router-dom";
import ItemCard from "../UI/ItemCard";

type LibraryItemProps = {
  readonly item: LibraryItemType;
}

const LibraryItem = (props: LibraryItemProps): JSX.Element => {
  const item = props.item;
  const navigator = useNavigate();
  const openItem = (): void => {
    navigator(`/${item.itemType}/${item.name}`);
  };

  return (<ItemCard type={item.itemType} id={item.id} itemCount={item.itemType === 'collection' ? item.bookCount: null} name={item.name} cover={item.cover} onClick={openItem}/>);
};

export default LibraryItem;
