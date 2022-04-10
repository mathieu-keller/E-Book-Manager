import {CombinedState, combineReducers, Reducer} from 'redux';
import {LibraryItemReducer} from "./LibraryItemReducer";
import {AppStore} from "../Store/Store.types";

const reducers: Reducer<CombinedState<AppStore>> = combineReducers({
  libraryItems: LibraryItemReducer.reducer,
});

export default reducers;
