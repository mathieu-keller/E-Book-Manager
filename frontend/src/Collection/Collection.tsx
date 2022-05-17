import {Component, createSignal, onMount, Show} from "solid-js";
import ItemGrid from "../UI/ItemGrid";
import {CollectionType} from "./Collection.type";
import {BookType} from "../Book/Book.type";
import {COLLECTION_API} from "../Api/Api";
import Rest from "../Rest";
import {useParams} from "solid-app-router";

const Collection: Component = () => {
  const [collection, setCollection] = createSignal<CollectionType | null>(null);
  const path = useParams<{ readonly collection: string }>();
  const getCollection = async (): Promise<CollectionType> => {
    const response = await Rest.get<CollectionType>(COLLECTION_API(path.collection));
    return response.data;
  };

  onMount(() => {
    getCollection().then(data => setCollection(data));
    console.log("mount!");
    console.log(path.collection);
  });


  return (
    <Show when={collection() !== null} fallback={<h1>Loading....</h1>}>
      <ItemGrid
        items={collection()!.books.map((book: BookType) => ({
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
