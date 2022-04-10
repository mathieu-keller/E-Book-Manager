import {LibraryItemType} from "../Library/LibraryItem.type";

export type LibraryItemStore = { readonly items: LibraryItemType[] };


export type AppStore = {
  readonly libraryItems: LibraryItemStore;
}
