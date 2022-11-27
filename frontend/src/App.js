import React from 'react';
import { createRoot } from "react-dom/client";
import Editor from './Editor'
import Graph from './Graph'
import Console from './console'
import './App.css';

export default class App extends React.Component {
  
  render() {
    
    return (
      <div>
        <title>markER</title>
        <meta name="description" content="Generated by create next app" />
      <header class="header font-sans font-bold">
        markER
      </header>
      {/* <h1 className="text-3xl font-bold text-blue-400">
      Hello world!
      </h1> */}
      <div class="sep">
        <div class="code">
          <Editor/>
        </div>
        <div class="right">
          <div class="graph">
              <Graph/>
          </div>
          <div class="console">
          <Console/>
          </div>
        </div>
      </div>
    </div>  
    );
  }
}

const container = document.getElementById('root');
const root = createRoot(container);
root.render(<App />);

