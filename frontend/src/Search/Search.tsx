import { createEffect, createSignal, on, onCleanup, onMount, Show } from 'solid-js';
import { setHeaderTitle } from '../Store/HeaderStore';
import Rest from '../Rest';
import { SEARCH_API } from '../Api/Api';
import ItemGrid from '../UI/ItemGrid';
import { searchStore, setSearchStore } from '../Store/SearchStore';
import { BookType } from '../Book/Book.type';

const Search = () => {
  const [loading, setLoading] = createSignal<boolean>(false);

  onMount(() => {
    setHeaderTitle(`Search: ${searchStore.search}`);
    window.addEventListener('scroll', shouldLoadNextPage);
    search(searchStore.page, searchStore.allLoaded);
  });

  createEffect(on(() => searchStore.search, () => {
    resetSearch();
  }));

  const resetSearch = () => {
    if (!loading()) {
      setHeaderTitle(`Search: ${searchStore.search}`);
      setSearchStore({
        page: 1,
        allLoaded: false,
        books: []
      });
      search(1, false);
    } else {
      setTimeout(resetSearch, 200);
    }
  };

  const search = (page: number, allLoaded: boolean) => {
    if (!allLoaded && !loading() && searchStore.search.trim() !== '') {
      setLoading(true);
      getBooks(page).then(r => {
        if (r.length > 0) {
          setSearchStore({
            page: page + 1,
            books: [...searchStore.books, ...r]
          });
          window.setTimeout(() => shouldLoadNextPage(), 50);
        } else if (r.length === 0 || r.length > 32) {
          setSearchStore({ allLoaded: true });
        }
        setLoading(false);
      });
    }
  };

  const getBooks = async (page: number): Promise<BookType[]> => {
    const response = await Rest.get<BookType[]>(SEARCH_API(searchStore.search, page));
    return response.data;
  };

  const shouldLoadNextPage = (): void => {
    const element = document.querySelector('#loading-trigger');
    const position = element?.getBoundingClientRect();
    if (position !== undefined && position.top >= 0 && position.bottom <= window.innerHeight) {
      search(searchStore.page, searchStore.allLoaded);
    }
  };

  onCleanup(() => {
    window.removeEventListener('scroll', shouldLoadNextPage);
  });

  return (
    <>
      <ItemGrid
        items={searchStore.books.map(book => ({
          id: book.id,
          title: book.title,
          cover: book.cover,
          itemType: 'book',
          bookCount: 1
        }))}
      />
      <Show when={!searchStore.allLoaded}>
        <div
          id="loading-trigger"
          onClick={() => search(searchStore.page, searchStore.allLoaded)}
          class="m-5 border cursor-pointer text-center text-5xl"
        >
          Load More
        </div>
      </Show>
    </>
  );

};

export default Search;
