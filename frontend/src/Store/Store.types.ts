import {LibraryItemType} from "../Library/LibraryItem.type";
import {BookType} from "../Book/Book.type";

export type LibraryItemStore = {
  readonly items: LibraryItemType[];
  readonly page: number;
  readonly allItemsLoaded: boolean;
};
export type CollectionStore = { readonly [key: string]: BookType[] };
export type ApplicationStore = { readonly headerText: string };


export type AppStore = {
  readonly libraryItems: LibraryItemStore;
  readonly collections: CollectionStore;
  readonly application: ApplicationStore;
}
