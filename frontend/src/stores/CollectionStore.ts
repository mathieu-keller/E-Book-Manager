import {defineStore} from 'pinia';
import type {BookType} from "@/Book/Book.type";

export type CollectionStoreType = { collections: { [key: string]: BookType[] } };

export const CollectionStore = defineStore<'collection', CollectionStoreType, {}, {
  set: (collection: string, books: BookType[]) => void;
}>({
  id: 'collection',
  state: () => ({collections: {}}),
  actions: {
    set(collection: string, books: BookType[]) {
      this.collections[collection] = books;
    },
  }
});
