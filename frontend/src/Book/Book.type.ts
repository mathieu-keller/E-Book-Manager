type Author = {
  readonly id: number;
  readonly name: string;
  readonly books: BookType[];
}

type Subject = {
  readonly id: number;
  readonly name: string;
  readonly books: BookType[];
}

export type BookType = {
  readonly id: number;
  readonly title: string;
  readonly published: string;
  readonly language: string;
  readonly subjects: Subject[];
  readonly publisher: string;
  readonly cover: string;
  readonly book: string;
  readonly authors: Author[];
  readonly collectionId: number;
}
