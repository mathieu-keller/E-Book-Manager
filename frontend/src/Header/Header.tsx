import { Component, createSignal, onMount, Show } from 'solid-js';
import Upload from '../Upload/Upload';
import { Button, PrimaryButton } from '../UI/Button';
import uploadIcon from '../assets/upload.svg';
import { useNavigate } from 'solid-app-router';
import { headerStore, setHeaderStore } from '../Store/HeaderStore';
import { setSearchStore } from '../Store/SearchStore';

const Header: Component = () => {
  const navigate = useNavigate();
  const [isDarkMode, setDarkMode] = createSignal<boolean>(window.matchMedia('(prefers-color-scheme: dark)').matches);
  const [showUploadModal, setShowUploadModal] = createSignal<boolean>(false);
  const [search, setSearch] = createSignal<string>('');

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

  const [timer, setTimer] = createSignal<number | null>(null);
  const setSearchValue = (inputValue: string) => {
    setSearch(inputValue);
    if (timer() == null) {
      setTimer(setTimeout(() => {
        setSearchStore({ search: search() });
        setTimer(null);
      }, 1000));
    }
  };

  return (
    <>
      <Show when={showUploadModal()}>
        <Upload onClose={() => setShowUploadModal(false)}/>
      </Show>
      <div class="flex flex-row justify-between border-b-2">
        <div>
          <Button onClick={() => navigate('/')}>
            Home
          </Button>
          <Button onClick={setDark}>
            {isDarkMode() ? 'Light mode' : 'Dark mode'}
          </Button>
        </div>
        <h1 class="text-5xl m-2 font-bold break-all">{headerStore.title}</h1>
        <PrimaryButton onClick={() => setShowUploadModal(true)}>
          <img
            class="dark:invert invert-0 h-8 mr-1"
            src={uploadIcon}
            alt="upload"
          /> Upload!
        </PrimaryButton>
      </div>
      <input
        class="w-[100%] text-5xl bg-slate-300 dark:bg-slate-700"
        placeholder="Search Books, Authors and Subjects"
        value={search()}
        onInput={e => setSearchValue(e.currentTarget.value)}
      />
    </>
  );
};

export default Header;
