import {Component, createSignal, onCleanup, onMount, Show, Accessor, Setter} from "solid-js";
import Upload from "../Upload/Upload";
import {Button, PrimaryButton} from "../UI/Button";
import upload_icon from "../assets/upload.svg";
import {useNavigate} from "solid-app-router";


const Header: Component = () => {
  const navigate = useNavigate();
  const [isDarkMode, setDarkMode] = createSignal<boolean>(window.matchMedia('(prefers-color-scheme: dark)').matches);
  const [showUploadModal, setShowUploadModal] = createSignal<boolean>(false);

  const setDarkClass = () => {
    if (isDarkMode()) {
      document.documentElement.classList.add('dark');
    } else {
      document.documentElement.classList.remove('dark');
    }
  };

  const setDark = (): void => {
    setDarkMode(old => !old);
    setDarkClass();
  };

  onMount(() => {
    setDarkClass();
  });


  return (
    <>
      <Show when={showUploadModal()}>
        <Upload onClose={() => setShowUploadModal(false)}/>
      </Show>
      <div class="flex flex-row justify-between border-b-2">
        <div>
          <Button onClick={() => navigate("/")}>
            Home
          </Button>
          <Button onClick={setDark}>
            {isDarkMode() ? 'Light mode' : 'Dark mode'}
          </Button>
        </div>
        <h1 class="text-5xl m-2 font-bold break-all">E-Book-Manager</h1>
        <PrimaryButton onClick={() => setShowUploadModal(true)}>
          <img
            class="dark:invert invert-0 h-8 mr-1"
            src={upload_icon}
            alt="upload"
          /> Upload!
        </PrimaryButton>
      </div>
    </>
  );
};

export default Header;
