import React from 'react';

const Header = () => {
  return (
    <div className="flex flex-row justify-between border-b-2">
      <p>Icon</p>
      <h1 className="text-3xl m-2">E-Book-Manager!</h1>
      <button>Upload!</button>
    </div>
  );
};

export default Header;
