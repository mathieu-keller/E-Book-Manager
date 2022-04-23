import {defineStore} from 'pinia';
import type {LibraryItemType} from "@/Library/LibraryItem.type";

export type LibraryItemStoreType = {
  items: LibraryItemType[];
  page: number;
  allItemsLoaded: boolean;
};
export const LibraryStore = defineStore<'library', LibraryItemStoreType, {}, {
  addAll: (items: LibraryItemType[]) => void;
  setPage: (page: number) => void;
  setAllLoaded: (allLoaded: boolean) => void;
  update: (libraryItem: LibraryItemType) => void;
}>({
  id: 'library',
  state: () => ({
    items: [],
    page: 1,
    allItemsLoaded: false
  }),
  actions: {
    addAll(items: LibraryItemType[]) {
      this.items = [...this.items, ...items];
    },
    setPage(page: number) {
      this.page = page;
    },
    setAllLoaded(allLoaded: boolean) {
      this.allItemsLoaded = allLoaded;
    },
    update(libraryItem: LibraryItemType) {
      const index = this.items.findIndex(item => item.id === libraryItem.id && item.itemType === libraryItem.itemType);
      this.items[index] = libraryItem;
    },
  }
});
