import React,{useRef} from "react";
// import raw from './code1.txt';

import Editor from "@monaco-editor/react";


export function Edit2() {
    const editorRef = useRef(null);
    // function onFileChange (e)  {
     
    //   // Update the state
    //   console.log(e.target.files[0] );
    //  reader = new FileReader()
    // };
    function onFileChange(event) {
      var file = event.target.files[0];
      var reader = new FileReader();
      reader.onload = function(event) {
        // The file's text will be printed here
        console.log(event.target.result)
        var text=event.target.result;
        data(text)
      };
    
      reader.readAsText(file);
    }
    var code="";

    function data(data){
        code=data;
        console.log("code=",code)
    }
    function exportData() {
        const fileData = editorRef.current.getValue();
        const blob = new Blob([fileData], { type: "text/plain" });
        const url = URL.createObjectURL(blob);
        const link = document.createElement("a");
        link.download = "code.txt";
        link.href = url;
        link.click();
      }
      
    function handleEditorDidMount(editor, monaco) {
      editorRef.current = editor; 
    }
    // function handleEditorWillMount(editor, monaco) {
    //   fetch(raw)
    //   .then((r) => r.text())
    //   .then(text  => {
    //     console.log(text);
    //     data(text)
    //   })
    //   }
    var v="";
    function clear(){
      editorRef.current.setValue(v);
    }
      function importData(){
        // editorRef.current.setValue(null);
        editorRef.current.setValue(code);
        console.log(editorRef)
      }
    
    
    function showValue() {
      fetch("http://localhost:5050/parse",editorRef.current.getValue());
    }
  return (  

    <>
    <br/>
    <div class="edit space-x-3 pl-5">
    <button class="btn btn-success" onClick={clear}> Clear Editor</button>
    <button class="btn btn-success" onClick={showValue}>Compile Code</button>
    <button class="btn btn-success" onClick={exportData}> Download Code</button>
    
    </div>
    <br/>
    <div class="rounded-1xl border-solid border-2 border-green-500 ">
   <Editor
    height="50vh"
     defaultLanguage="sql"
    //  beforeMount={handleEditorWillMount}
     defaultValue="-- Write ur code below "
     theme = "vs-dark"
     onMount={handleEditorDidMount}
   />
   </div>
   <br/>
   <div class="edit space-x-3 pl-3">
   <input type="file" class="file-input file-input-bordered file-input-success w-full max-w-xs" onChange={onFileChange}/>
    <button class="btn btn-success" onClick={importData}> Upload Code</button>
    </div>
   </>
  );
}