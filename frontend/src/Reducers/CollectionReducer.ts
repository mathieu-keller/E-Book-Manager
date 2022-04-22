import {CollectionStore} from "../Store/Store.types";
import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {BookType} from "../Book/Book.type";

export const initialState: CollectionStore = {};
export const CollectionReducer = createSlice({
  name: 'collection',
  initialState,
  reducers: {
    set: (state, action: PayloadAction<{ collection: string, books: BookType[] }>): void => {
      state[action.payload.collection] = action.payload.books
        .sort((a, b): number => a.collectionIndex - b.collectionIndex);
    },
  },
});
