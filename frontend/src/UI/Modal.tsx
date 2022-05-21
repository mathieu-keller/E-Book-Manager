import { FlowComponent } from 'solid-js';

type ModalProps = {
  readonly title: string
  readonly onClose: () => void;
}

const Modal: FlowComponent<ModalProps> = (props) => {
  return (
    <>
      <div class="fixed top-0 left-0 right-0 bottom-0 z-40 flex justify-center items-center">
      <div onClick={props.onClose} class="fixed top-0 left-0 right-0 bottom-0 z-0 bg-gray-700/75 flex justify-center items-center"/>
        <div
          onClick={e => e.preventDefault()}
          class="dark:bg-slate-900 p-5 dark:text-slate-300 bg-slate-50 text-slate-800 z-1 flex flex-col fixed max-h-[50%] max-w-[80%] my-auto flex">
          <h1 class="border-b-2 w-full flex justify-center text-5xl font-bold pb-5">{props.title}</h1>
          {props.children}
        </div>
      </div>

    </>
  );
};

export default Modal;
