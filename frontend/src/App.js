import Cytoscape from 'cytoscape';
import Cola from 'cytoscape-cola';
import React, { useState } from 'react';
import { createRoot } from 'react-dom/client';
import './App.css';
import Console from './console';
import Editor from './Editor';
import Graph from './Graph';
import ParseER from './utils/parser';

Cytoscape.use(Cola);

export default function App() {
    const [er, setEr] = useState({
        nodes: [],
        edges: [],
    });

    const handleCompileCode = (editorCode) => {
        localStorage.setItem('marker-code', editorCode);
        if (editorCode.length > 0) {
            fetch('http://localhost:5050/parse', {
                method: 'POST',
                body: JSON.stringify({
                    code: editorCode,
                }),
            })
                .then((res) => res.json())
                .then((data) => {
                    const [nodes, edges] = ParseER(JSON.parse(data.data));
                    setEr(() => {
                        return {
                            nodes: nodes,
                            edges: edges,
                        };
                    });
                    // console.log(data.data);
                })
                .catch((err) => {
                    console.error(err);
                });
        }
    };

    return (
        <div>
            <header className='header font-sans font-bold'>markER</header>
            <div className='sep'>
                <div className='code'>
                    <Editor handleCompileCode={handleCompileCode} />
                </div>
                <div className='right'>
                    <div className='graph'>
                        <Graph er={er} />
                    </div>
                    <div className='console'>
                        <Console />
                    </div>
                </div>
            </div>
        </div>
    );
}

const container = document.getElementById('root');
const root = createRoot(container);
root.render(<App />);
