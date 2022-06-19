import { editor, Uri } from 'monaco-editor';
import { setDiagnosticsOptions } from 'monaco-yaml';

// The uri is used for the schema file match.
const modelUri = Uri.parse('file://masking.yml');

setDiagnosticsOptions({
  enableSchemaRequest: true,
  hover: true,
  completion: true,
  validate: true,
  format: true,
  schemas: [
    {
      // Id of the first schema
      uri: 'https://raw.githubusercontent.com/CGI-FR/PIMO/main/schema/v1/pimo.schema.json',
      // Associate with our model
      fileMatch: [String(modelUri)],
    },
  ],
});

// editor.defineTheme("PIMO", {
//   base: "vs",
//   inherit: true,
//   rules: [{ background: 'EDF9FA' }],
//   colors: {
//     'editor.background': '#EDF9FA',
//   }
// });
// editor.setTheme("PIMO");

const value = 'version: "1"\nmasking:\n  - selector:\n      jsonpath: "name"\n    mask:\n      randomChoiceInUri: "pimo://nameFR"\n';

var editorYaml = editor.create(document.getElementById('editor-yaml'), {
  automaticLayout: true,
  tabSize: 2,
  scrollBeyondLastLine: false,
  minimap: {enabled: false},
  model: editor.createModel(value, 'yaml', modelUri),
});

var editorJson = editor.create(document.getElementById('editor-json'), {
  automaticLayout: true,
  scrollBeyondLastLine: false,
  minimap: {enabled: false},
  model: editor.createModel('{"name": "Bill"}', 'json', Uri.parse('file://data.jsonl')),
});

var resultJson = editor.create(document.getElementById('result-json'), {
  automaticLayout: true,
  scrollBeyondLastLine: false,
  minimap: {enabled: false},
  readOnly: true,
  model: editor.createModel('', 'json', Uri.parse('file://result.jsonl')),
});

///////////////////////////////////////////////////////////

async function postData() {
  const postData = {
      data: editorJson.getValue(),
      masking: editorYaml.getValue()
  }
  console.log(postData)

  // if (postData.data.length === 0 || postData.masking.length === 0) {
  //     postData.data = example.json
  //     postData.masking = example.yaml
  // }

  try {
      const res = await fetch(`/play`, {
          method: "POST",
          headers: {
              "Content-Type": "application/json"
          },
          body: JSON.stringify(postData)
      })

      if (!res.ok) {
        if (res.status == 500) {
          const data = await res.text()
          throw new Error(data)
        }
        const message = `An error has occurred: ${res.status} - ${res.statusText}`
        throw new Error(message)
      }

      const data = await res.json()

      resultJson.setValue(JSON.stringify(data))
      document.getElementById('result-error').innerText = ""
  } catch (err) {
      console.log(err)
      document.getElementById('result-error').innerText = err
  } finally {
    document.getElementById('label-output').innerText = "Output"
  }
}

function debounce(func, timeout = 300){
    let timer;
    return (...args) => {
        document.getElementById('label-output').innerText = "Output (refreshing...)"
        clearTimeout(timer);
        timer = setTimeout(() => { func.apply(this, args); }, timeout);
    };
}

let autoPostData = debounce(postData, 500);
document.getElementById('editor-yaml').onkeyup = autoPostData;
document.getElementById('editor-yaml').oninput = autoPostData;
document.getElementById('editor-yaml').onpaste = autoPostData;
document.getElementById('editor-yaml').oncut = autoPostData;
document.getElementById('editor-json').onkeyup = autoPostData;
document.getElementById('editor-json').oninput = autoPostData;
document.getElementById('editor-json').onpaste = autoPostData;
document.getElementById('editor-json').oncut = autoPostData;
autoPostData();
