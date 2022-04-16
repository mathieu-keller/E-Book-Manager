import React, {FormEvent, useState} from 'react';
import Modal from "../UI/Modal";
import Button, {PrimaryButton} from "../UI/Button";
import Rest from "../Rest";

type UploadProps = {
  readonly onClose: () => void;
}

const Upload = (props: UploadProps): JSX.Element => {


  const [maxSize, setMaxSize] = useState<number | null>(null);
  const [current, setCurrent] = useState<number | null>(null);
  const uploadBooks = async (data: FormData): Promise<void> => {
    await Rest.post('/upload/multi', data, {
      onUploadProgress: (e: ProgressEvent): void => {
        setMaxSize(e.total);
        setCurrent(e.loaded);
      }
    });
    location.reload();
  };

  return (
    <Modal
      onClose={props.onClose}
      title="Upload E-Book"
      footer={
        <div className="flex justify-around w-full">
          <PrimaryButton type="submit" form="upload-epub">Upload</PrimaryButton>
          <Button onClick={props.onClose}>Close</Button>
        </div>
      }>
      <form
        id="upload-epub"
        onSubmit={(e: FormEvent<HTMLFormElement>): void => {
          e.preventDefault();
          const form = new FormData(e.currentTarget);
          uploadBooks(form)
            .then((): void => props.onClose())
            .catch((e: string): void => console.error(e));
        }}
      >
        <input type="file" accept="application/epub+zip" name="myFiles" multiple/>
      </form>
      {current !== null && maxSize !== null ?
        <>
          <progress value={current} max={maxSize}/>
          {(Math.round((current / maxSize) * 10000)) / 100}% <br/>
          ({current}/{maxSize})
        </> : null}
    </Modal>
  );
};

export default Upload;
