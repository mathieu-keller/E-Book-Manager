import {Book} from "./Book";

export type Author = {
  readonly id: number;
  readonly name: string;
  readonly books: Book[];
}
