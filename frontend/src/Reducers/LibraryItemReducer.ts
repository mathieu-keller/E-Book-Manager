import {LibraryItemType} from "../Library/LibraryItem.type";
import {LibraryItemStore} from "../Store/Store.types";
import {createSlice, PayloadAction} from "@reduxjs/toolkit";

const stringSort = (a: string, b: string): number => {
  const fa = a.toLowerCase(),
    fb = b.toLowerCase();

  if (fa < fb) {
    return -1;
  }
  if (fa > fb) {
    return 1;
  }
  return 0;
};
export const initialState: LibraryItemStore = {items: [], page: 0, allItemsLoaded: false};
export const LibraryItemReducer = createSlice({
  name: 'library',
  initialState,
  reducers: {
    set: (state, action: PayloadAction<LibraryItemType[]>): void => {
      state.items = action.payload.sort((a, b): number =>
        stringSort(a.title, b.title));
    },
    addAll: (state, action: PayloadAction<LibraryItemType[]>): void => {
      state.items = [...state.items, ...action.payload];
    },
    add: (state, action: PayloadAction<LibraryItemType>): void => {
      state.items = [...state.items, action.payload].sort((a, b): number =>
        stringSort(a.title, b.title));
    },
    setPage: (state, action: PayloadAction<number>): void => {
      state.page = action.payload;
    },
    setAllLoaded: (state, action: PayloadAction<boolean>): void => {
      state.allItemsLoaded = action.payload;
    },
    update: (state, action: PayloadAction<LibraryItemType>): void => {
      const index = state.items.findIndex((item): boolean => item.id === action.payload.id && item.itemType === action.payload.itemType);
      state.items[index] = action.payload;
    },
  },
});
