import { Component, createSignal, onCleanup, onMount, Show } from 'solid-js';
import ItemGrid from '../UI/ItemGrid';
import { LibraryItemType } from './LibraryItem.type';
import { LIBRARY_API } from '../Api/Api';
import Rest from '../Rest';

const Library: Component = () => {
  const [libraryItems, setLibraryItems] = createSignal<LibraryItemType[]>([]);
  const [loading, setLoading] = createSignal<boolean>(false);
  const [allLoaded, setAllLoaded] = createSignal<boolean>(false);
  const [page, setPage] = createSignal<number>(1);

  onMount(() => {
    window.addEventListener('scroll', shouldLoadNextPage);
    loadLibraryItems();
  });

  const loadLibraryItems = () => {
    setLoading(true);
    getLibraryItems(page()).then(r => {
      if (r.length > 0) {
        setPage(page() + 1);
        setLoading(false);
        setLibraryItems((prev) => [...prev, ...r]);
        window.setTimeout(() => shouldLoadNextPage(), 50);
      } else if (r.length === 0 || r.length > 32) {
        setAllLoaded(true);
      }
    });
  };

  const getLibraryItems = async (page: number): Promise<LibraryItemType[]> => {
    const response = await Rest.get<LibraryItemType[]>(LIBRARY_API(page));
    return response.data;
  };

  const shouldLoadNextPage = (): void => {
    const element = document.querySelector('#loading-trigger');
    const position = element?.getBoundingClientRect();
    if (position !== undefined && !loading() && position.top >= 0 && position.bottom <= window.innerHeight) {
      loadLibraryItems();
    }
  };

  onCleanup(() => {
    window.removeEventListener('scroll', shouldLoadNextPage);
  });

  return (
    <>
      <ItemGrid
        items={libraryItems()}
      />
      <Show when={!allLoaded()}>
        <div id="loading-trigger" onClick={loadLibraryItems} class="m-5 border cursor-pointer text-center text-5xl">Load More</div>
      </Show>
    </>
  );
};

export default Library;
