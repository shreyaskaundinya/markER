export const graphStyleSheet = [
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
        selector: 'node[tableType="weak"]',
        style: {
            'border-width': '4px',
        },
    },
    {
        selector: 'node[isIdentifying="true"]',
        style: {
            'border-width': '4px',
        },
    },
    {
        selector: 'node[primaryKey="true"]',
        style: {
            'background-color': '#ff0000',
        },
    },
    {
        selector: 'node[derived="true"]',
        style: {
            'background-color': '#0000ff',
        },
    },
    {
        selector: 'node[multivalue="true"]',
        style: {
            'border-width': '4px',
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
            label: 'data(cardinality)',
        },
    },
];
