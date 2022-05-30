export type LibraryItemType = {
  readonly id: number;
  readonly isSvg: boolean;
  readonly cover: string | null;
  readonly title: string;
  readonly itemType: 'book' | 'collection';
  readonly bookCount: number;
}
