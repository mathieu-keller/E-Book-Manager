import {CombinedState, combineReducers, Reducer} from 'redux';
import {LibraryItemReducer} from "./LibraryItemReducer";
import {AppStore} from "../Store/Store.types";
import {CollectionReducer} from "./CollectionReducer";

const reducers: Reducer<CombinedState<AppStore>> = combineReducers({
  libraryItems: LibraryItemReducer.reducer,
  collections: CollectionReducer.reducer
});

export default reducers;
