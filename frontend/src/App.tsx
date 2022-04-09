import React, {useEffect, useState} from 'react';

const App = (): JSX.Element => {
  const [r, setR] = useState<{ ID: string; Name: string }[]>([]);
  useEffect(() => {
    fetch('/all')
      .then(r => {
        console.log(r);
        r.json().then(j => {
          console.log(j);
          setR(j);
        });
      });
  }, []);
  return (
    <>
      <h1>Test</h1>
      {r.map(a => <p key={a.ID}>{a.Name}</p>)}
    </>
  );
};

export default App;
