import React from 'react';

type ItemCardProps = {
  readonly cover: string;
  readonly name: string;
  readonly onClick: () => void;
}

const ItemCard = ({cover, name, onClick}: ItemCardProps): JSX.Element => {
  return (
    <div onClick={onClick} className="cursor-pointer m-5 p-3 flex hover:p-1 hover:my-3 max-w-sm flex-col">
      <img src={`data:image/jpeg;base64,${cover}`} alt={`cover picture of ${name}`}/>
      <h1 className="text-center break-words text-2xl font-bold">{name}</h1>
    </div>
  );
};

export default ItemCard;
