import { FlowComponent } from 'solid-js';

type ModalProps = {
  readonly title: string
  readonly onClose: () => void;
}

const Modal: FlowComponent<ModalProps> = (props) => {
  return (
    <>
      <div onClick={props.onClose} class="fixed top-0 left-0 right-0 bottom-0 z-40 bg-gray-700/75"/>
      <div
        class="dark:bg-slate-900 dark:text-slate-300 bg-slate-50 text-slate-800 z-50 flex flex-col fixed max-h-[50%] w-2/4 left-1/4 top-1/4 my-auto flex justify-center items-center">
        <h1 class="border-b-2 w-full flex justify-center">{props.title}</h1>
        {props.children}
      </div>
    </>
  );
};

export default Modal;
