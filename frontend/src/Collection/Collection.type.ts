import {BookType} from "../Book/Book.type";

export type CollectionType = {
  readonly id: number;
  readonly name: string;
  readonly books: BookType[];
}
