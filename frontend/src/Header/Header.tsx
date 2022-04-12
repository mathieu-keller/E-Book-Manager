import React, {useEffect, useState} from 'react';
import Upload from "../Upload/Upload";
import Button, {PrimaryButton} from "../UI/Button";

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

  return (
    <>
      {uploadFile ? <Upload onClose={(): void => setUploadFile(false)}/> : null}
      <div className="flex flex-row justify-between border-b-2">
        <Button onClick={setDark}>dark</Button>
        <h1 className="text-3xl m-2">E-Book-Manager!</h1>
        <PrimaryButton onClick={(): void => setUploadFile(true)}>Upload!</PrimaryButton>
      </div>
    </>
  );
};

export default Header;
