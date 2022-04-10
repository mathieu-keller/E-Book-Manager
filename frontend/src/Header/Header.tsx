import React, {useEffect, useState} from 'react';

const Header = () => {
  const [isDarkMode, setDarkMode] = useState<boolean>(window.matchMedia('(prefers-color-scheme: dark)').matches);
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
    <div className="flex flex-row justify-between border-b-2">
      <p>Icon</p>
      <h1 className="text-3xl m-2">E-Book-Manager!</h1>
      <div>
        <button onClick={setDark}>dark</button>
        <button>Upload!</button>
      </div>
    </div>
  );
};

export default Header;
