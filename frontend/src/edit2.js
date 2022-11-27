import React,{useRef} from "react";


import Editor from "@monaco-editor/react";


export function Edit2() {
    const editorRef = useRef(null);

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
    
    function showValue() {
      fetch("http://localhost:5050/parse",editorRef.current.getValue());
    }
  return (  

    <>
    <button class="btn btn-success" onClick={showValue}>Show value</button>
    <button class="btn btn-success" onClick={exportData}> download value</button>
   <Editor  
     height="50vh"
     defaultLanguage="sql"
     defaultValue="-- Enter ur code below"
     theme = "vs-dark"
     onMount={handleEditorDidMount}
   />
   </>
  );
}