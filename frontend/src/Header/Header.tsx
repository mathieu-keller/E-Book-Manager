import React, {useEffect, useState} from 'react';
import Upload from "../Upload/Upload";
import Button from "../UI/Button";

const Header = () => {
  const [isDarkMode, setDarkMode] = useState<boolean>(window.matchMedia('(prefers-color-scheme: dark)').matches);
  const [uploadFile, setUploadFile] = useState<boolean>(false);
  const setDark = () => {
    setDarkMode(!isDarkMode);
  };

  useEffect(() => {
    if (isDarkMode) {
      document.documentElement.classList.add('dark');
    } else {
      document.documentElement.classList.remove('dark');
    }
  }, [isDarkMode]);

  return (
    <>
      {uploadFile ? <Upload onClose={() => setUploadFile(false)}/> : null}
      <div className="flex flex-row justify-between border-b-2">
        <p>Icon</p>
        <h1 className="text-3xl m-2">E-Book-Manager!</h1>
        <div>
          <Button onClick={setDark}>dark</Button>
          <Button onClick={() => setUploadFile(true)}>Upload!</Button>
        </div>
      </div>
    </>
  );
};

export default Header;
