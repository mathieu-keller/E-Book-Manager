import React from 'react';
import {LibraryItemType} from "./LibraryItem.type";
import {useNavigate} from "react-router-dom";

type LibraryItemProps = {
  readonly item: LibraryItemType;
}

const LibraryItem = (props: LibraryItemProps) => {
  const item = props.item;
  const navigator = useNavigate();
  const openItem = () => {
    const name = item.name;
    let route = item.type === "collection" ? name : `${item.id}-${name}`;
    navigator(`${item.type}/${route}`);
  };

  return (
    <div onClick={openItem} className="m-5 flex max-w-sm flex-col shadow">
      <img src={`data:image/png;base64,${item.cover}`} alt={`cover picture of ${item.name}`}/>
      <h1 className="text-center break-words text-2xl font-bold">{item.name}</h1>
    </div>
  );
};

export default LibraryItem;
