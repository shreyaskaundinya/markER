import * as monaco from 'monaco-editor';

// Since packaging is done by you, you need
// to instruct the editor how you named the
// bundles that contain the web workers.


(function () {
	// create div to avoid needing a HtmlWebpackPlugin template
	const div = document.createElement('div');
	div.id = 'root';
	div.style = 'width:800px; height:600px; border:1px solid #ccc;';

	document.body.appendChild(div);
})();



function App(){
  monaco.editor.create(document.getElementById('root'), {
    value: `const foo = () => 0;`,
    language: 'javascript',
    theme: 'vs-dark'
  });
}
export default App;