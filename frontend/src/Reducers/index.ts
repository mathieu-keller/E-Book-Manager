import {CombinedState, combineReducers, Reducer} from 'redux';
import {LibraryItemReducer} from "./LibraryItemReducer";
import {AppStore} from "../Store/Store.types";
import {CollectionReducer} from "./CollectionReducer";
import {ApplicationReducer} from "./HeaderReducer";

const reducers: Reducer<CombinedState<AppStore>> = combineReducers({
  libraryItems: LibraryItemReducer.reducer,
  collections: CollectionReducer.reducer,
  application: ApplicationReducer.reducer
});

export default reducers;
