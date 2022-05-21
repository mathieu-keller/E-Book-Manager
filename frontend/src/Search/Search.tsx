import { createEffect, createSignal, on, onCleanup, onMount, Show } from 'solid-js';
import { setHeaderTitle } from '../Store/HeaderStore';
import Rest from '../Rest';
import { SEARCH_API } from '../Api/Api';
import ItemGrid from '../UI/ItemGrid';
import { searchStore, setSearch, setSearchStore } from '../Store/SearchStore';
import { BookType } from '../Book/Book.type';

const Search = () => {
  const [loading, setLoading] = createSignal<boolean>(false);
  const [searchInput, setSearchInput] = createSignal<string>('');

  onMount(() => {
    setHeaderTitle(`Search: ${searchStore.search}`);
    window.addEventListener('scroll', shouldLoadNextPage);
    setSearchInput(searchStore.search);
    search();
  });

  createEffect(on(() => searchStore.search, (value, prev) => {
    if (prev !== undefined) {
      resetSearch();
    }
  }));

  const resetSearch = () => {
    if (!loading()) {
      setHeaderTitle(`Search: ${searchStore.search}`);
      search();
    } else {
      setTimeout(resetSearch, 200);
    }
  };

  const search = () => {
    if (!searchStore.allLoaded && !loading() && searchStore.search.trim() !== '') {
      setLoading(true);
      getBooks(searchStore.page).then(r => {
        if (r.length > 0) {
          setSearchStore({
            page: searchStore.page + 1,
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
      search();
    }
  };

  const [timer, setTimer] = createSignal<number | null>(null);
  const setSearchValue = (inputValue: string) => {
    setSearchInput(inputValue);
    if (timer() == null) {
      setTimer(setTimeout(() => {
        setSearch(searchInput());
        setTimer(null);
      }, 1000));
    }
  };

  onCleanup(() => {
    window.removeEventListener('scroll', shouldLoadNextPage);
    const timeout = timer();
    if (timeout != null) {
      clearTimeout(timeout);
    }
  });

  return (
    <>
      <input
        class="w-[100%] text-5xl bg-slate-300 dark:bg-slate-700"
        placeholder="Search Books, Authors and Subjects"
        value={searchInput()}
        onInput={e => setSearchValue(e.currentTarget.value)}
      />
      <ItemGrid
        items={searchStore.books.map(book => ({
          id: book.id,
          title: book.title,
          cover: book.cover,
          itemType: 'book',
          bookCount: 1
        }))}
      />
      <Show when={!searchStore.allLoaded && searchStore.search.trim() !== ''}>
        <div
          id="loading-trigger"
          onClick={() => search()}
          class="m-5 border cursor-pointer text-center text-5xl"
        >
          Load More
        </div>
      </Show>
    </>
  );
};

export default Search;
