type Author = {
  readonly id: string;
  readonly name: string;
  readonly books: BookType[];
}

type Subject = {
  readonly name: string;
  readonly books: string;
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
