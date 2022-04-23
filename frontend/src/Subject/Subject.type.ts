import type {BookType} from "@/Book/Book.type";

export type Subject = {
  readonly id: number;
  readonly name: string;
  readonly books: BookType[];
}
