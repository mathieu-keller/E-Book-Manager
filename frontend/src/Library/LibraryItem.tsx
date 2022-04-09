import React from 'react';
import {LibraryItemType} from "./LibraryItem.type";

type LibraryItemProps = {
  readonly item: LibraryItemType;
}

const LibraryItem = (props: LibraryItemProps) => {
  const item = props.item;
  return (
    <div className="bg-blue">
      <h1 className="underline font-bold text-2xl">{item.name}</h1>
      <img src={`data:image/png;base64,${item.cover}`} alt={`cover picture of ${name}`} style={{maxHeight: '30rem'}}/>
    </div>
  );
};

export default LibraryItem;
