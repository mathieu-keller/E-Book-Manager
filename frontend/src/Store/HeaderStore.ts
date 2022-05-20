import { createStore } from 'solid-js/store';

type HeaderStoreType = {
  readonly title: string;
}

const defaultTitleName = 'E-Book Manager';

const initialState: HeaderStoreType = {
  title: defaultTitleName
};

export const [headerStore, setHeaderStore] = createStore<HeaderStoreType>(initialState);

export const setHeaderTitle = (title: string) => {
  document.title = title;
  setHeaderStore({ title });
};

export const resetHeaderTitle = () => {
  document.title = defaultTitleName;
  setHeaderStore({ title: defaultTitleName });
};
