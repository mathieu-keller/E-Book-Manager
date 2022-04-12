import {LibraryItemType} from "../Library/LibraryItem.type";
import {LibraryItemStore} from "../Store/Store.types";
import {createSlice, PayloadAction} from "@reduxjs/toolkit";

const stringSort = (a: string, b: string) => {
  let fa = a.toLowerCase(),
    fb = b.toLowerCase();

  if (fa < fb) {
    return -1;
  }
  if (fa > fb) {
    return 1;
  }
  return 0;
};
export const initialState: LibraryItemStore = {items: []};
export const LibraryItemReducer = createSlice({
  name: 'library',
  initialState,
  reducers: {
    set: (state, action: PayloadAction<LibraryItemType[]>): void => {
      state.items = action.payload.sort((a, b) =>
        stringSort(a.name, b.name));
    },
    add: (state, action: PayloadAction<LibraryItemType>): void => {
      state.items = [...state.items, action.payload].sort((a, b) =>
        stringSort(a.name, b.name));
    },
    update: (state, action: PayloadAction<LibraryItemType>): void => {
      let index = state.items.findIndex(item => item.id === action.payload.id && item.itemType === action.payload.itemType);
      state.items[index] = action.payload;
    },
  },
});
