import React from 'react';

type ItemCardProps = {
  readonly cover: string;
  readonly name: string;
  readonly onClick: () => void;
}

const ItemCard = ({cover, name, onClick}: ItemCardProps) => {
  return (
    <div onClick={onClick} className="m-5 flex max-w-sm flex-col shadow bg-slate-100 dark:bg-slate-800">
      <img src={`data:image/png;base64,${cover}`} alt={`cover picture of ${name}`}/>
      <h1 className="text-center break-words text-2xl font-bold">{name}</h1>
    </div>
  );
};

export default ItemCard;
