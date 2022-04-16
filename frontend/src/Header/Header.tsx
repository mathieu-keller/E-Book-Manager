import React, {useEffect, useState} from 'react';
import Upload from "../Upload/Upload";
import Button, {PrimaryButton} from "../UI/Button";
import {useDispatch, useSelector} from "react-redux";
import {AppStore} from "../Store/Store.types";
import {useLocation, useNavigate, useSearchParams} from "react-router-dom";
import {ApplicationReducer} from "../Reducers/HeaderReducer";

const Header = (): JSX.Element => {
  const [isDarkMode, setDarkMode] = useState<boolean>(window.matchMedia('(prefers-color-scheme: dark)').matches);
  const [uploadFile, setUploadFile] = useState<boolean>(false);
  const setDark = (): void => {
    setDarkMode(!isDarkMode);
  };

  useEffect((): void => {
    if (isDarkMode) {
      document.documentElement.classList.add('dark');
    } else {
      document.documentElement.classList.remove('dark');
    }
  }, [isDarkMode]);
  const headerText = useSelector<AppStore, string>((store): string => store.application.headerText);

  useEffect((): void => {
    document.title = `E-Book - ${headerText}`;
  }, [headerText]);
  const [search, setSearch] = useState<string | null>(null);
  const [, setSearchParams] = useSearchParams();
  const navigate = useNavigate();
  const loc = useLocation();
  const dispatch = useDispatch();
  const [timeout, setTimeout] = useState<number | null>(null);

  useEffect((): void => {
    if (loc.state !== null && loc.state !== undefined) {
      setSearchParams(`q=${loc.state as string}`);
      setSearch(loc.state as string);
    }
  }, [loc.state]);

  useEffect((): void => {
    if (search !== null && search.trim() !== "") {
      if (loc.pathname !== '/search') {
        navigate(`/search?q=${search.trim()}`);
      } else {
        if (timeout !== null) {
          window.clearTimeout(timeout);
        }
        setTimeout(window.setTimeout((): void => {
          setSearchParams(`q=${search.trim()}`);
        }, 500));
      }
      dispatch(ApplicationReducer.actions.setHeaderText(`Search: ${search}`));
    } else if (search === null || search.trim() === "") {
      if (timeout !== null) {
        window.clearTimeout(timeout);
      }
      setSearch(null);
      if (loc.pathname === '/search') {
        navigate('/books');
      }
    }
  }, [search]);


  return (
    <>
      {uploadFile ? <Upload onClose={(): void => setUploadFile(false)}/> : null}
      <div className="flex flex-row justify-between border-b-2">
        <div>
          <Button onClick={(): void => navigate('/')}>Home</Button>
          <Button onClick={(): void => setDark()}>{isDarkMode ? 'Light mode' : 'Dark mode'}</Button>
        </div>
        <h1 className="text-5xl m-2 font-bold">{headerText}</h1>
        <PrimaryButton onClick={(): void => setUploadFile(true)}>Upload!</PrimaryButton>
      </div>
      <input
        className="w-[100%] text-5xl bg-slate-300 dark:bg-slate-700"
        placeholder="Search Books, Authors and Subjects"
        value={search || ''}
        onChange={(e): void => setSearch(e.currentTarget.value)}
      />
    </>
  );
};

export default Header;
