import { useEffect, useRef, useState } from 'react';
import CytoscapeComponent from 'react-cytoscapejs';

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
                    stylesheet={[
                        {
                            selector: 'node[table]',
                            style: {
                                label: 'data(label)',
                                shape: 'rectangle',
                                width: 'label',
                                height: 'label',
                                padding: '18',
                                'text-halign': 'center',
                                'text-valign': 'center',
                            },
                        },
                        {
                            selector: 'node[attr]',
                            style: {
                                label: 'data(label)',
                                shape: 'ellipse',
                                width: 'label',
                                height: 'label',
                                padding: '18',
                                'text-halign': 'center',
                                'text-valign': 'center',
                            },
                        },
                        {
                            selector: 'node[compSubAttr]',
                            style: {
                                label: 'data(label)',
                                shape: 'ellipse',
                                width: 'label',
                                height: 'label',
                                padding: '18',
                                'text-halign': 'center',
                                'text-valign': 'center',
                            },
                        },
                        {
                            selector: 'node[compAttr]',
                            style: {
                                label: 'data(label)',
                                shape: 'ellipse',
                                width: 'label',
                                height: 'label',
                                padding: '18',
                                'text-halign': 'center',
                                'text-valign': 'center',
                            },
                        },
                        {
                            selector: 'node[reln]',
                            style: {
                                label: 'data(id)',
                                shape: 'diamond',
                                width: 'label',
                                height: 'label',
                                padding: '25',
                                'text-halign': 'center',
                                'text-valign': 'center',
                            },
                        },
                        {
                            selector: 'edge',
                            style: {
                                'curve-style': 'bezier',
                                width: 1.5,
                            },
                        },
                    ]}
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
