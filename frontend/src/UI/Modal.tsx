import React from 'react';
import Button from "./Button";

type ModalProps = {
  readonly title: string;
  readonly children?: React.ReactNode;
  readonly footer?: React.ReactNode;
  readonly onClose: () => void;
}
const Modal = (props: ModalProps) => {
  return (
    <>
      <div onClick={props.onClose} className="fixed top-0 left-0 right-0 bottom-0 z-40 bg-gray-700/75"/>
      <div
        className="dark:bg-slate-900 dark:text-slate-300 bg-slate-50 text-slate-800 z-50 flex flex-col fixed max-h-[50%] w-2/4 left-1/4 top-1/4 my-auto flex justify-center items-center">
        <h1 className="border-b-2 w-full flex justify-center">{props.title}</h1>
        <div className="word-wrap-break-word max-h-[80%]  overflow-y-auto ">
          {props.children}
        </div>
        <footer className="border-t-2 w-full flex justify-center">
          {props.footer !== undefined ? props.footer : <Button onClick={props.onClose}>Close</Button>}
        </footer>
      </div>
    </>
  );
};

export default Modal;
