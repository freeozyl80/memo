import logo from './logo.svg';
import Demo from './demo'
import {useRef, useState} from 'react'

import './App.css';

function App() {
  const elementRef = useRef()
  const [info, setInfo] = useState(0);

  const test = () => {
   setInfo(info+1)
  }
  return (
    <div className="App">
      <Demo className="demo" ref={elementRef} info={info}/>
      <header className="App-header" onClick={test}>
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
