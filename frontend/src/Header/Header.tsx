import React, {useEffect, useState} from 'react';
import Upload from "../Upload/Upload";
import Button, {PrimaryButton} from "../UI/Button";
import {useSelector} from "react-redux";
import {AppStore} from "../Store/Store.types";

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

  return (
    <>
      {uploadFile ? <Upload onClose={(): void => setUploadFile(false)}/> : null}
      <div className="flex flex-row justify-between border-b-2">
        <Button onClick={setDark}>dark</Button>
        <h1 className="text-5xl m-2 font-bold">{headerText}</h1>
        <PrimaryButton onClick={(): void => setUploadFile(true)}>Upload!</PrimaryButton>
      </div>
    </>
  );
};

export default Header;
