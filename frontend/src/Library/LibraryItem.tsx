import React from 'react';
import {LibraryItemType} from "./LibraryItem.type";

type LibraryItemProps = {
  readonly item: LibraryItemType;
}

const LibraryItem = (props: LibraryItemProps) => {
  const item = props.item;
  return (
    <>
      <h1>{item.name}</h1>
      <img src={`data:image/png;base64,${item.cover}`} alt={`cover picture of ${name}`} style={{maxHeight: '30rem'}}/>
    </>
  );
};

export default LibraryItem;
