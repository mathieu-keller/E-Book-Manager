import React from 'react';
import defaultCover from '../../public/default/cover.jpg';

type ItemCardProps = {
  readonly cover: string | null;
  readonly name: string;
  readonly onClick: () => void;
  readonly itemCount?: number | null;
}

const ItemCard = ({cover, name, onClick, itemCount}: ItemCardProps): JSX.Element => {
  return (
    <div onClick={onClick} className="relative cursor-pointer m-3 p-2 flex h-max max-w-sm flex-col">
      {itemCount !== undefined && itemCount !== null ?
        <div className="absolute left-5 top-0 text-5xl bg-red-700 text-white rounded-b-full">{itemCount}</div>
        :
        null
      }
      <img className="hover:pb-3 hover:mt-0 hover:mb-3 p-0 my-3" src={cover !== null ? `data:image/jpeg;base64,${cover}` : defaultCover}
           alt={`cover picture of ${name}`}/>
      <h1 className="text-center break-words text-2xl font-bold">{name}</h1>
    </div>
  );
};

export default ItemCard;
