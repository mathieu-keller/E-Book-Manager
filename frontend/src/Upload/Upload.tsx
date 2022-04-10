import React from 'react';
import Modal from "../UI/Modal";
import Button, {PrimaryButton} from "../UI/Button";

type UploadProps = {
  readonly onClose: () => void;
}

const Upload = (props: UploadProps) => {
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
        onSubmit={(e) => {
          e.preventDefault();
          const form = new FormData(e.currentTarget);
          fetch('/upload',
            {method: 'POST', body: form})
            .then(() => props.onClose())
            .catch(e => console.error(e));
        }}
      >
        <input type="file" accept="application/epub+zip" name="myFile"/>
      </form>
    </Modal>
  );
};

export default Upload;
