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
  const headerText = useSelector<AppStore, string>((store) => store.application.headerText);

  useEffect(() => {
    document.title = `E-Book - ${headerText}`;
  }, [headerText]);
  const [search, setSearch] = useState<string | null>(null);
  const [, setSearchParams] = useSearchParams();
  const navigate = useNavigate();
  const loc = useLocation();
  const dispatch = useDispatch();
  const [timeout, setTimeout] = useState<number | null>(null);

  useEffect(() => {
    if (loc.state !== null && loc.state !== undefined) {
      setSearchParams(`q=${loc.state as string}`);
      setSearch(loc.state as string);
    }
  }, [loc.state]);

  useEffect(() => {
    if (search !== null && search.trim() !== "") {
      if (loc.pathname !== '/search') {
        navigate(`/search?q=${search.trim()}`);
      } else {
        if (timeout !== null) {
          window.clearTimeout(timeout);
        }
        setTimeout(window.setTimeout(() => {
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
        <Button onClick={() => navigate('/')}>Home</Button>
        <h1 className="text-5xl m-2 font-bold">{headerText}</h1>
        <PrimaryButton onClick={(): void => setUploadFile(true)}>Upload!</PrimaryButton>
      </div>
      <input className="w-[100%]" value={search || ''} onChange={(e) => setSearch(e.currentTarget.value)}/>
    </>
  );
};

export default Header;
