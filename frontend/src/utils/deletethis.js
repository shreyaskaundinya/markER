// cytoscape({
//     container: document.getElementById('cy'),

//     elements: {
//         nodes: [
//             { data: { id: 'table_User', table: true, label: 'User' } },
//             { data: { id: 'attr_User_id', attr: true, label: 'id' } },
//             {
//                 data: {
//                     id: 'attr_User_username',
//                     attr: true,
//                     label: 'username',
//                 },
//             },
//             {
//                 data: {
//                     id: 'attr_User_password',
//                     attr: true,
//                     label: 'password',
//                 },
//             },
//             {
//                 data: {
//                     id: 'attr_User_created_at',
//                     attr: true,
//                     label: 'created_at',
//                 },
//             },
//             { data: { id: 'comp_User_name', compAttr: true, label: 'name' } },
//             {
//                 data: {
//                     id: 'compsub_User_name_fname',
//                     compSubAttr: true,
//                     label: 'fname',
//                 },
//             },
//             {
//                 data: {
//                     id: 'compsub_User_name_lname',
//                     compSubAttr: true,
//                     label: 'lname',
//                 },
//             },
//             { data: { id: 'table_Project', table: true, label: 'Project' } },
//             { data: { id: 'attr_Project_id', attr: true, label: 'id' } },
//             { data: { id: 'attr_Project_name', attr: true, label: 'name' } },
//             {
//                 data: {
//                     id: 'attr_Project_created_at',
//                     attr: true,
//                     label: 'created_at',
//                 },
//             },
//             { data: { id: 'relation_Has', reln: true, label: 'Has' } },
//         ],
//         edges: [
//             {
//                 data: {
//                     id: 'e_1',
//                     source: 'table_User',
//                     target: 'attr_User_id',
//                 },
//             },
//             {
//                 data: {
//                     id: 'e_2',
//                     source: 'table_User',
//                     target: 'attr_User_username',
//                 },
//             },
//             {
//                 data: {
//                     id: 'e_3',
//                     source: 'table_User',
//                     target: 'attr_User_password',
//                 },
//             },
//             {
//                 data: {
//                     id: 'e_4',
//                     source: 'table_User',
//                     target: 'attr_User_created_at',
//                 },
//             },
//             {
//                 data: {
//                     id: 'e_5',
//                     source: 'table_User',
//                     target: 'comp_User_name',
//                 },
//             },
//             {
//                 data: {
//                     id: 'e_6',
//                     source: 'comp_User_name',
//                     target: 'compsub_User_name_fname',
//                 },
//             },
//             {
//                 data: {
//                     id: 'e_7',
//                     source: 'comp_User_name',
//                     target: 'compsub_User_name_lname',
//                 },
//             },
//             {
//                 data: {
//                     id: 'e_8',
//                     source: 'table_Project',
//                     target: 'attr_Project_id',
//                 },
//             },
//             {
//                 data: {
//                     id: 'e_9',
//                     source: 'table_Project',
//                     target: 'attr_Project_name',
//                 },
//             },
//             {
//                 data: {
//                     id: 'e_10',
//                     source: 'table_Project',
//                     target: 'attr_Project_created_at',
//                 },
//             },
//             {
//                 data: {
//                     id: 'e_11',
//                     source: 'table_User',
//                     target: 'relation_Has',
//                 },
//             },
//             {
//                 data: {
//                     id: 'e_12',
//                     source: 'relation_Has',
//                     target: 'table_Project',
//                 },
//             },
//         ],
// },
// style: [
//     {
//         selector: 'node[table]',
//         style: {
//             label: 'data(label)',
//             shape: 'rectangle',
//             width: 'label',
//             height: 'label',
//             padding: '18',
//             'text-halign': 'center',
//             'text-valign': 'center',
//         },
//     },
//     {
//         selector: 'node[attr, compSubAttr, compAttr]',
//         style: {
//             label: 'data(label)',
//             shape: 'ellipse',
//             width: 'label',
//             height: 'label',
//             padding: '18',
//             'text-halign': 'center',
//             'text-valign': 'center',
//         },
//     },
//     {
//         selector: 'node[reln]',
//         style: {
//             label: 'data(id)',
//             shape: 'diamond',
//             width: 'label',
//             height: 'label',
//             padding: '25',
//             'text-halign': 'center',
//             'text-valign': 'center',
//         },
//     },
//     {
//         selector: 'edge',
//         style: {
//             'curve-style': 'bezier',
//             width: 1.5,
//         },
//     },
// ],

//     layout: {
//         name: 'breadthfirst',
//     },
// });
