import React from 'react';
import ItemCard from "./ItemCard";

type ItemsGridProps<T extends { id: number; title: string; cover: string; itemType?: 'book' | 'collection'; bookCount?: number }> = {
  readonly items: T[];
  readonly onClick: (item: T) => void;
}

function ItemsGrid<T extends { id: number; itemType?: 'book' | 'collection'; title: string; cover: string; bookCount?: number }>(props: ItemsGridProps<T>): JSX.Element {
  return (
    <div className="flex flex-wrap flex-row justify-center">
      {props.items.map((item): JSX.Element => <ItemCard
        key={`${item.id}-${item.title}`}
        name={item.title}
        cover={item.cover}
        itemCount={item.itemType === 'collection' ? item.bookCount : null}
        id={item.id}
        type={item.itemType || 'book'}
        onClick={(): void => props.onClick(item)}
      />)}
    </div>
  );
}

export default ItemsGrid;
