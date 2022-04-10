import React from 'react';
import {LibraryItemType} from "./LibraryItem.type";
import {useNavigate} from "react-router-dom";
import ItemCard from "../UI/ItemCard";

type LibraryItemProps = {
  readonly item: LibraryItemType;
}

const LibraryItem = (props: LibraryItemProps) => {
  const item = props.item;
  const navigator = useNavigate();
  const openItem = () => {
    navigator(`/${item.type}/${item.name}`);
  };

  return (<ItemCard name={item.name} cover={item.cover} onClick={openItem}/>);
};

export default LibraryItem;
