import { Component, createSignal, onMount, Show } from 'solid-js';
import ItemGrid from '../UI/ItemGrid';
import { CollectionType } from './Collection.type';
import { COLLECTION_API } from '../Api/Api';
import Rest from '../Rest';
import { useParams } from 'solid-app-router';
import { collectionStore, setCollectionStore } from '../Store/CollectionStore';
import { Store } from 'solid-js/store/types/store';

const Collection: Component = () => {
  const path = useParams<{ readonly collection: string }>();
  const getCollection = async (): Promise<CollectionType> => {
    const response = await Rest.get<CollectionType>(COLLECTION_API(path.collection));
    return response.data;
  };

  const [collection, setCollection] = createSignal<Store<CollectionType> | null>(null);

  onMount(() => {
    const storedCollection = collectionStore.find(col => col.title.toLowerCase() === decodeURIComponent(path.collection).toLowerCase());
    if (storedCollection === undefined) {
      getCollection()
        .then(r => {
          setCollectionStore([...collectionStore, r]);
          setCollection(r);
        });
    } else {
      setCollection(storedCollection);
    }
  });

  return (
    <Show when={collection() !== null} fallback={<h1>Loading....</h1>}>
      <ItemGrid
        items={collection()!.books.map((book) => ({
          id: book.id,
          cover: book.cover,
          title: book.title,
          itemType: 'book',
          bookCount: 1
        }))}
      />
    </Show>
  );
};

export default Collection;
