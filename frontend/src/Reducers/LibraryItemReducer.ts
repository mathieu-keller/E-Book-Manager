// visible for test
import {LibraryItemType} from "../Library/LibraryItem.type";
import {LibraryItemStore} from "../Store/Store.types";
import {createSlice, PayloadAction} from "@reduxjs/toolkit";

export const initialState: LibraryItemStore = {items:[]};
export const LibraryItemReducer = createSlice({
  name: 'library',
  initialState,
  reducers: {
    set: (state, action: PayloadAction<LibraryItemType[]>) => {
      state.items = action.payload;
    },
    add: (state, action: PayloadAction<LibraryItemType>) => {
      state.items = [...state.items, action.payload]
    },
  },
})
