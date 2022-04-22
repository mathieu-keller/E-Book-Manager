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
  const [searchParams, setSearchParams] = useSearchParams();
  const navigate = useNavigate();
  const loc = useLocation();
  const dispatch = useDispatch();
  const querySearch = searchParams.get("q");
  useEffect((): void => {
    if (loc.state !== null && loc.state !== undefined) {
      setSearchParams(`q=${loc.state as string}`);
    }
  }, [loc.state]);
  useEffect((): void => {
    if (querySearch !== null && querySearch.trim() !== "") {
      if (loc.pathname !== '/search') {
        navigate(`/search?q=${querySearch}`);
      }
      dispatch(ApplicationReducer.actions.setHeaderText(`Search: ${querySearch}`));
    } else if (querySearch === null || querySearch.trim() === "") {
      setSearchParams("");
      if (loc.pathname === '/search') {
        navigate('/books');
      }
    }
  }, [querySearch]);

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
        value={querySearch || ''}
        onChange={(e): void => setSearchParams("q=" + e.currentTarget.value)}
      />
    </>
  );
};

export default Header;
