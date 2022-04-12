import React, {FormEvent} from 'react';
import Modal from "../UI/Modal";
import Button, {PrimaryButton} from "../UI/Button";
import {useDispatch, useSelector} from "react-redux";
import {AppStore} from "../Store/Store.types";
import {BookType} from "../Book/Book.type";
import {LibraryItemType} from "../Library/LibraryItem.type";
import {LibraryItemReducer} from "../Reducers/LibraryItemReducer";
import {CollectionReducer} from "../Reducers/CollectionReducer";
import {CollectionType} from "../Collection/Collection.type";

type UploadProps = {
  readonly onClose: () => void;
}

const Upload = (props: UploadProps): JSX.Element => {

  const libraryItems = useSelector((store: AppStore) => store.libraryItems);
  const collections = useSelector((store: AppStore) => store.collections);
  const dispatch = useDispatch();
  const uploadBooks = async (data: FormData): Promise<void> => {
    await fetch('/upload/multi',
      {
        method: 'POST',
        body: data
      }
    );
    location.reload();
  };
  const uploadBook = async (data: FormData): Promise<void> => {
    const response = await fetch('/upload',
      {
        method: 'POST',
        body: data
      }
    );
    const book = await response.json() as BookType;
    if (book.collectionId === 0) {
      const lib: LibraryItemType = {
        id: book.id,
        name: book.name,
        itemType: 'book',
        cover: book.cover,
        bookCount: 1
      };
      dispatch(LibraryItemReducer.actions.add(lib));
    } else {
      const col = Object.entries(collections).map(a => a[1]).flat().find(b => b.collectionId === book.collectionId);
      if (col !== undefined) {
        const collectionResponse = await fetch(`/collection/${book.collectionId}`);
        const collection = await collectionResponse.json() as CollectionType;
        dispatch(CollectionReducer.actions.set({collection: collection.name, books: collection.books}));
      }
      const lib = libraryItems.items.find(i => i.id === book.collectionId && i.itemType === 'collection');
      const libraryResponse = await fetch(`/library/${book.collectionId}`);
      const library = await libraryResponse.json() as LibraryItemType;
      if (lib === undefined) {
        dispatch(LibraryItemReducer.actions.add(library));
      } else {
        dispatch(LibraryItemReducer.actions.update(library));
      }
    }
  };

  return (
    <Modal
      onClose={props.onClose}
      title="Upload E-Book"
      footer={
        <div className="flex justify-around w-full">
          <PrimaryButton type="submit" form="upload-epub">Upload</PrimaryButton>
          <PrimaryButton type="submit" form="multi-upload-epub">Multi Upload</PrimaryButton>
          <Button onClick={props.onClose}>Close</Button>
        </div>
      }>
      <form
        id="upload-epub"
        onSubmit={(e: FormEvent<HTMLFormElement>): void => {
          e.preventDefault();
          const form = new FormData(e.currentTarget);
          uploadBook(form)
            .then(() => props.onClose())
            .catch((e: string): void => console.error(e));
        }}
      >
        <input type="file" accept="application/epub+zip" name="myFile"/>
      </form>
      <form
        id="multi-upload-epub"
        onSubmit={(e: FormEvent<HTMLFormElement>): void => {
          e.preventDefault();
          const form = new FormData(e.currentTarget);
          uploadBooks(form)
            .then(() => props.onClose())
            .catch((e: string): void => console.error(e));
        }}
      >
        <h1>Debug Multi</h1>
        <input type="file" accept="application/epub+zip" name="myFiles" multiple/>
      </form>
    </Modal>
  );
};

export default Upload;
