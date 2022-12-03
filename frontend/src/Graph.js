import { useEffect, useRef, useState } from 'react';
import CytoscapeComponent from 'react-cytoscapejs';
import { graphStyleSheet } from './utils/graphStyleSheet';

export default function Graph({ er }) {
    const cyRef = useRef(null);
    const [layout, setLayout] = useState({
        name: 'cola',
        animate: true,
    });
    const [elements, setElements] = useState([]);

    useEffect(() => {
        setElements(() =>
            CytoscapeComponent.normalizeElements({
                nodes: er.nodes,
                edges: er.edges,
            })
        );

        setLayout((prev) => ({
            ...prev,
            name: prev.name === 'breadthfirst' ? 'cola' : 'breadthfirst',
            animate: true,
        }));
    }, [er]);

    return (
        <div className='font-serif text-xl font-bold'>
            Graph
            <div>
                <CytoscapeComponent
                    wheelSensitivity={0.1}
                    elements={elements}
                    layout={layout}
                    cy={(ref) => {
                        cyRef.current = ref;
                    }}
                    stylesheet={graphStyleSheet}
                    style={{
                        width: '100%',
                        height: '600px',
                        border: '1px solid white',
                        backgroundColor: 'white',
                    }}
                />
            </div>
        </div>
    );
}
