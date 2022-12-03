function ParseER(er) {
    let nodes = [];
    let edges = [];

    let tables = er.tables;
    let relations = er.relations;

    let edge_count = 0;

    const wrap = (d) => {
        return { data: d };
    };

    let start = Date.now();

    // entities
    tables.forEach((t) => {
        let table_name = t.name;
        let table_id = `table_${t.name}`;

        nodes.push(
            wrap({
                id: table_id,
                tableType: t.type,
                table: true,
                label: table_name,
            })
        );

        let attr = t.attr;
        let compAttr = t.compositeAttr;

        // attr
        attr.forEach((a) => {
            let attr_id = `attr_${table_name}_${a.name}`;
            nodes.push(
                wrap({
                    id: attr_id,
                    attr: true,
                    label: a.name,
                    primaryKey: a.properties.primaryKey ? 'true' : 'false',
                    foreignKey: a.properties.foreignKey ? 'true' : 'false',
                    unique: a.properties.unique ? 'true' : 'false',
                    notNull: a.properties.notNull ? 'true' : 'false',
                    derived: a.properties.derived ? 'true' : 'false',
                    multivalue: a.properties.multivalue ? 'true' : 'false',
                })
            );

            edge_count += 1;
            edges.push(
                wrap({
                    id: `e_${edge_count}`,
                    source: table_id,
                    target: attr_id,
                })
            );
        });

        // composite attr
        compAttr.forEach((c) => {
            let compAttr_id = `comp_${table_name}_${c.name}`;
            nodes.push(
                wrap({
                    id: compAttr_id,
                    compAttr: true,
                    label: c.name,
                })
            );

            edge_count += 1;
            edges.push(
                wrap({
                    id: `e_${edge_count}`,
                    source: table_id,
                    target: compAttr_id,
                })
            );

            let compSubAttr = c.attr;
            compSubAttr.forEach((cs) => {
                let compSubAttr_id = `compsub_${table_name}_${c.name}_${cs.name}`;
                nodes.push(
                    wrap({
                        id: compSubAttr_id,
                        compSubAttr: true,
                        label: cs.name,
                    })
                );
                edge_count += 1;
                edges.push(
                    wrap({
                        id: `e_${edge_count}`,
                        source: compAttr_id,
                        target: compSubAttr_id,
                    })
                );
            });
        });
    });

    // relations
    relations.forEach((r) => {
        let reln_id = `relation_${r.name}`;

        nodes.push(
            wrap({
                id: reln_id,
                reln: true,
                isIdentifying: r.identifying ? 'true' : 'false',
                label: r.name,
            })
        );

        let source_table_id = `table_${r.table1}`;
        let dest_table_id = `table_${r.table2}`;

        edge_count += 1;
        edges.push(
            wrap({
                id: `e_${edge_count}`,
                source: source_table_id,
                target: reln_id,
                cardinality: `(${r.cardinality1}, ${r.participation1})`,
            })
        );

        edge_count += 1;
        edges.push(
            wrap({
                id: `e_${edge_count}`,
                source: reln_id,
                target: dest_table_id,
                cardinality: `(${r.cardinality1}, ${r.participation1})`,
            })
        );
    });

    let end = Date.now();
    console.log(end - start);

    console.log(nodes, edges);

    return [nodes, edges];
}

export default ParseER;

/*

Types of nodes : 
  - Entity [rect]
  - Relation [diamond]
  - Attribute | Composite Attribute | Composite Sub Attribute [ecclipse]
  - Derived [underline]
  - Multivalue [ecclipse with double]

Types of edges : 
  - Relation => Entity to Relation & Relation to Entity
  - Attribute => Entity to Attribute
  - Composite Attribute => Entity to Comp Attr
  - Composite Attr => Composite Sub Attr

*/
