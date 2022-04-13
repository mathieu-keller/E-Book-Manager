import {ApplicationStore} from "../Store/Store.types";
import {createSlice, PayloadAction} from "@reduxjs/toolkit";

export const initialState: ApplicationStore = {headerText: 'Manager'};
export const ApplicationReducer = createSlice({
  name: 'application',
  initialState,
  reducers: {
    setHeaderText: (state, action: PayloadAction<string>): void => {
      state.headerText = action.payload;
    },
    reset: (state): void => {
      state.headerText = initialState.headerText;
    },
  },
});
