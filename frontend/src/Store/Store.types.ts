import {LibraryItemType} from "../Library/LibraryItem.type";
import {BookType} from "../Book/Book.type";

export type LibraryItemStore = { readonly items: LibraryItemType[] };
export type CollectionStore = { readonly [key: string]: BookType[] };


export type AppStore = {
  readonly libraryItems: LibraryItemStore;
  readonly collections: CollectionStore;
}
