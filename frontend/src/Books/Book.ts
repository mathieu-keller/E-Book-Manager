import {Author} from "./Author";

export type Book = {
  readonly id: number;
  readonly name: string;
  readonly published: string;
  readonly language: string;
  readonly subject: string;
  readonly publisher: string;
  readonly cover: string;
  readonly book: string;
  readonly author: Author[];
  readonly collectionId: number;
}
