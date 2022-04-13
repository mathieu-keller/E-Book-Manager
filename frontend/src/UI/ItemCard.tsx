import React from 'react';
import defaultCover from '../../public/default/cover.jpg';
import {LinkButton} from "./Button";

type ItemCardProps = {
  readonly cover: string | null;
  readonly name: string;
  readonly onClick: () => void;
  readonly itemCount?: number | null;
  readonly id: number;
  readonly type: string;
}

const ItemCard = ({cover, name, onClick, itemCount, id, type}: ItemCardProps): JSX.Element => {
  return (
    <div className="m-3 p-2 flex h-max max-w-sm flex-col">
      <div onClick={onClick} className="hover:pb-3 cursor-pointer hover:mt-0 hover:mb-3 p-0 my-3 relative">
        {itemCount !== undefined && itemCount !== null ?
          <div className="absolute left-5 top-0 text-5xl bg-red-700 text-white rounded-b-full">{itemCount}</div>
          :
          null
        }
        <img
          src={cover !== null ? `data:image/jpeg;base64,${cover}` : defaultCover}
          alt={`cover picture of ${name}`}
        />
      </div>
      <div>
        <h1
          onClick={onClick}
          className={"cursor-pointer text-center break-words text-2xl font-bold " +
            (type === 'book' ? "float-left w-10/12" : "w-12/12")}
        >
          {name}
        </h1>
        {type === 'book' ? <LinkButton
            download={`${name}.epub`}
            href={`/download/${id}`}
            className="w-2/12 float-right"
          >D
          </LinkButton>
          : null}
      </div>
    </div>
  );
};

export default ItemCard;
