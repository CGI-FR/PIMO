"use strict";
import { editor, Uri } from "monaco-editor";
import { setDiagnosticsOptions } from "monaco-yaml";
import { Elm } from "./Main";
import LZString from "lz-string";
import mermaid from "mermaid";
import * as d3 from "d3";

var app = Elm.Main.init({ flags: "{{ version }}" });

let lastSandbox;
let debounceTimeOut = 500;

if (typeof Go !== "undefined") {
  initWasm();
  debounceTimeOut = 10; // wasm is faster !
} else {
  initBackend();
}

function initWasm() {
  if (!WebAssembly.instantiateStreaming) {
    // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
      const source = await (await resp).arrayBuffer();
      return await WebAssembly.instantiate(source, importObject);
    };
  }

  const go = new Go();

  let mod, inst;

  WebAssembly.instantiateStreaming(fetch("pimo.wasm"), go.importObject).then(async (result) => {
    mod = result.module;
    inst = result.instance;
    console.log("start go");

    go.run(inst);

    /**
     * Pimo mask
     */

    pimo
      .play(lastSandbox.masking, lastSandbox.input)
      .then(app.ports.outputUpdater.send)
      .catch((err) => {
        app.ports.errorUpdater.send(String(err));
      });
    pimo
      .flow(lastSandbox.masking)
      .then(app.ports.flowUpdater.send)
      .catch((err) => {
        app.ports.errorUpdater.send(String(err));
      });
    app.ports.pimoMask.subscribe((sandbox) => {
      pimo
        .flow(sandbox.masking)
        .then(app.ports.flowUpdater.send)
        .catch((err) => {
          app.ports.errorUpdater.send(String(err));
        });
      pimo
        .play(sandbox.masking, sandbox.input)
        .then(app.ports.outputUpdater.send)
        .catch((err) => {
          app.ports.errorUpdater.send(String(err));
        });
    });
    inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
  });

  // Save the sandbox
  app.ports.pimoMask.subscribe((sandbox) => (lastSandbox = sandbox));
}

function initBackend() {
  /**
   * Pimo mask
   */

  app.ports.pimoMask.subscribe((sandbox) => {
    sandbox["data"] = sandbox["input"];
    delete sandbox["input"];

    // masking data
    fetch("/play", {
      method: "POST",
      body: JSON.stringify(sandbox),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => response.text())
      .then(app.ports.outputUpdater.send)
      .catch((err) => {
        console.log(err);
        app.ports.errorUpdater.send(String(err));
      });

    // flow chart
    fetch("/flow", {
      method: "POST",
      body: JSON.stringify(sandbox),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => response.text())
      .then(app.ports.flowUpdater.send)
      .catch((err) => {
        console.log(err);
        app.ports.errorUpdater.send(String(err));
      });
  });
}

// The uri is used for the schema file match.
const modelUri = Uri.parse("file://masking.yml");

setDiagnosticsOptions({
  enableSchemaRequest: true,
  hover: true,
  completion: true,
  validate: true,
  format: true,
  schemas: [
    {
      // Id of the first schema
      uri: "https://raw.githubusercontent.com/CGI-FR/PIMO/{{ version }}/schema/v1/pimo.schema.json",
      // Associate with our model
      fileMatch: [String(modelUri)],
    },
  ],
});

function debounce(func, timeout = 300) {
  let timer;
  return (...args) => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      func.apply(this, args);
    }, timeout);
  };
}

function updateUrl() {
  // update URL for sharing
  var c = LZString.compressToEncodedURIComponent(editorYaml.getValue());
  var i = LZString.compressToEncodedURIComponent(editorJson.getValue());
  window.history.replaceState(null, null, `${location.protocol}//${location.host}${location.pathname}#c=${c}&i=${i}`);
}

/**
 * Masking.yml Editor
 */

let editorYaml;
let editorJson;
let urlParams;
if (window.location.hash !== "") {
  urlParams = new URLSearchParams(window.location.hash.replace("#", ""));
} else {
  urlParams = new URLSearchParams(window.location.search); // Legacy insecure url read only compatibilites to avoid broken urls
}

let masking = 'version: "1"\nmasking:\n  - selector:\n      jsonpath: "name"\n    mask:\n      randomChoiceInUri: "pimo://nameFR"';
let input = '{\n    "name": "Bill"\n  }';

if (urlParams.has("c")) {
  masking = LZString.decompressFromEncodedURIComponent(urlParams.get("c"));
}
if (urlParams.has("i")) {
  input = LZString.decompressFromEncodedURIComponent(urlParams.get("i"));
}

app.ports.maskingAndinputUpdater.send({ masking: masking, input: input });

editorYaml = editor.create(document.getElementById("editor-yaml"), {
  automaticLayout: true,
  tabSize: 2,
  scrollBeyondLastLine: false,
  minimap: { enabled: false },
  model: editor.createModel(masking, "yaml", modelUri),
});

let updateMasking = debounce(() => {
  app.ports.maskingUpdater.send(editorYaml.getValue());
  updateUrl();
}, debounceTimeOut);
document.getElementById("editor-yaml").onkeyup = updateMasking;
document.getElementById("editor-yaml").oninput = updateMasking;
document.getElementById("editor-yaml").onpaste = updateMasking;
document.getElementById("editor-yaml").oncut = updateMasking;

app.ports.updateMaskingEditor.subscribe((masking) => {
  if (editorYaml == undefined) {
    return;
  }
  editorYaml.setValue(masking);
  updateUrl();
});

/**
 * Input Json Editor
 */

editorJson = editor.create(document.getElementById("editor-json"), {
  automaticLayout: true,
  tabSize: 2,
  scrollBeyondLastLine: false,
  minimap: { enabled: false },
  model: editor.createModel(input, "json", Uri.parse("file://input.jsonl")),
});

let updateInput = debounce(() => {
  app.ports.inputUpdater.send(editorJson.getValue());
  updateUrl();
}, debounceTimeOut);

document.getElementById("editor-json").onkeyup = updateInput;
document.getElementById("editor-json").oninput = updateInput;
document.getElementById("editor-json").onpaste = updateInput;
document.getElementById("editor-json").oncut = updateInput;

app.ports.updateInputEditor.subscribe((input) => {
  if (editorJson == undefined) {
    return;
  }
  editorJson.setValue(JSON.stringify(JSON.parse(input), null, 2));
  updateUrl();
});

/**
 * Result JSON editor
 */

let resultJson;

resultJson = editor.create(document.getElementById("result-json"), {
  automaticLayout: true,
  tabSize: 2,
  scrollBeyondLastLine: false,
  minimap: { enabled: false },
  model: editor.createModel("", "json", Uri.parse("file://result.jsonl")),
});

app.ports.updateOutputEditor.subscribe((output) => {
  if (resultJson == undefined) {
    return;
  }

  if (output === "") {
    resultJson.setValue("");
  } else {
    try {
      resultJson.setValue(JSON.stringify(JSON.parse(output), null, 2));
    } catch (error) {
      console.log(error);
      resultJson.setValue(output);
    }
  }
});

var resultFlowchart = document.getElementById("flowchart");
mermaid.initialize({ startOnLoad: true });

app.ports.updateFlow.subscribe((data) => {
  try {
    mermaid.parse(data);
    const cb = function (svgGraph) {
      if (document.getElementById("div")) {
        resultFlowchart.removeChild(document.getElementById("dflowchartGraph"));
      }
      var graph = document.createElement("dflowchartGraph");
      graph.id = "dflowchartGraph";
      graph.innerHTML = svgGraph;
      resultFlowchart.appendChild(graph);
    };
    mermaid.render("flowchartGraph", data, cb, resultFlowchart);
  } catch (error) {
    let error_msg = document.createElement("span");
    error_msg.appendChild(document.createTextNode("flow chart is not in mermaid format : " + error.str));
    let src = document.createElement("span");
    src.appendChild(document.createTextNode(data));

    resultFlowchart.replaceChildren(error_msg, src);
  }
});

// CTRL + S download masking.yaml file
document.addEventListener(
  "keydown",
  function (e) {
    if ((e.key === "s" || e.key === "S") && (navigator.userAgentData.platform.match("Mac") ? e.metaKey : e.ctrlKey)) {
      e.preventDefault();

      var encodedMasking = encodeURIComponent(editorYaml.getValue());
      var aDownloadMasking = document.createElement("a");

      aDownloadMasking.setAttribute("href", `data:text/yaml,${encodedMasking}`);
      aDownloadMasking.setAttribute("download", "masking.yml");
      aDownloadMasking.click();
      aDownloadMasking.remove();
    }
  },
  false
);

window.addEventListener("load", function () {
  var svgs = d3.selectAll("dflowchartGraph > svg");
  svgs.each(function () {
    var svg = d3.select(this);
    svg.html("<g>" + svg.html() + "</g>");
    var inner = svg.select("g");
    var zoom = d3.zoom().on("zoom", function (event) {
      inner.attr("transform", event.transform);
    });
    svg.call(zoom);
  });
});

window.addEventListener("wheel", function () {
  var svgs = d3.selectAll("dflowchartGraph > svg");
  svgs.each(function () {
    var svg = d3.select(this);
    svg.html("<g>" + svg.html() + "</g>");
    var inner = svg.select("g");
    var zoom = d3.zoom().on("zoom", function (event) {
      inner.attr("transform", event.transform);
    });
    svg.call(zoom);
  });
});

// Examples ///////////////////////////////////////////////

function loadExample(params) {
  app.ports.maskingAndinputUpdater.send({
    masking: LZString.decompressFromEncodedURIComponent(params[0]),
    input: LZString.decompressFromEncodedURIComponent(params[1]),
  });
}

let examples_generation = new Map([
  [
    "Generate first name, last name and email from referentials",
    [
      "MQAgKgFglgziCmAPAhgWwA4Bt4hhA9gO4gHEAu+u82AxmSAG7KYCu8cAZgE76ohfwO8AQDsyUZjABQoAIL9Bw+GImYQNZCJAAjHMhCZ8GtRyjYANCH2F42hTHwsuNHPi5WQABQCSAWQDyIFBiwiLMCkKi4swy4BBKQXD6QsQCkcrRmHDoyDQA1vAAJkEiMFCFOGTxIABE6FCo+DU6wchcAJ6WsRxuCCgY2CD1jQBcAPRjYajwAKIAcur4YsjBSQaw9PgcCCIA5piwECBT7AB0UlIMwmVLI7UAjDVSqMgwecG7IxcgIKCmXDB6NgyABybJQfIgFjoDxcTSFXjHNA4biI-QHQEgLZI6bSH6gKo4E4gQhmNS6dQCZBkIolMoVDymajFE7FGr-QEAfROT3xIEJIA59Beb1qyEKhWasGO8CKtIoVgl-OqTMwxWCypwACkAMr+BaOMjoFhkKR83QaFgwHA2ElkkC7KBXKxaeAYMjtEC6-UlY30BXDfBm36amWIYWvPJ2zBqa20ej6OEiBF8JisHAagU1E4AMQASs00koVDEfgBaKjxtxfH61kAAKwcIhyVTu7KgALI3ORvNrIryMBrdZAFfFhTbvbrFaTKYAwgQIfBvCIAKpcKBtwPjSbI-NPYOgGDIxYiFzoMicXoCzCvegnSzaE12qogQpQDjpMQRYuZF3FVUssieIjpW8B0NWwZ1o2SwthAbY3lyPKQSA-aDsh5aKuOtSTrW07wrw874Iuy5rhutRbhMMBOLmBbfCGApuisJhQMy0bkjgNC8H6tKonwAroAIDBQI4cBpmwwEEpQCLKrAli2laei4B8gz9li2weugOA1DSAzUvAvYEtUOlYHp-LIAUawVDANDruewlaNipD8pQj5mMUAoCFRmD0GJKI8Hw+CEu4AGMMw4kHoKvSNAIJQ9FwLziEsIB3BAZBGoOEzaIYuynLs+DoPEcIaBUqDtKcnGoGM4pXGIZYAEwAAz3AA7GMVofGWeVlsZN40jAYwRYYNjuPCIBzPgsg0C4X5tCiLCnolpSWIGfQ0smcACnl-JuiZfXPkcSBkHCgrzXQ9kSZF7jRRmIhxQl9nJSQaXoBlYwijS8XBIUMC5VAVQsNopzCWMMACVAuwDRhcZgRQXBDlBTawW2jFmDh-bw7WY5tgA3tjICnEK3bTCAAA+BhEMIpNjRNU0ZCAAC+9OnLj+MIV2xJk0NlNk+Nk3TfQjMAALtPlLxmATXBPEAA",
      "N4XyA",
    ],
  ],
  [
    "Generate first name, last name and email from sample data",
    [
      "MQAgKgFglgziCmAPAhgWwA4Bt4hhA9gO4gHEAu+IAxgE7zJk4B28xAbspgK7xxkQMQmfPgDWcTFFE46nEB268AUKADqOLjBzIQAWWQ1R+NtQFQmICiAAmlfrAA0l6HFghkFzgHN8NKP1RnQSoPEAAjHDIaZHN4axB8C1lMGwZkJSU2eBoYKESALhAAIgBGIqVUZBhRcy98jJAQUH4cSsNjEErqkEIoTBSveBZoxnd5Th4gshA3XNQ+g0tKBR5XCxaQLygspPgAM2yhsihOJUbQCEOeyOjzZ1aDIxNUfGt4FN7+ECY0XhmmKjcN7xO5FdBQF75AD0UJgXBoP1Q8AAYgAlIogM4gAC0uHe8CoFBo9UapJAACsYIl0AwIIUinsoDkyAB9RHwcpkrriElknHuazWemcvm4tpPXl80kwNBYeD08GQmFwhG-NHlLGgGVI6iJKjwdBkOB7Xz3IRVabspxhLjTT4QGxQPYHOhMaZ0F1HE4pDzxRnveLsmBY3FabCE3yS0mU6m0+mYC1s34ixrcmBRxq45CC4VYslix7GDN87VyhUQ-DQqHs9UNJpm+CVPogf2YeK9frhHBUfAYW1xFs0Xtm9B0Nh5TTjRTB85LGx2FxOQgaLRjXJMLzYTpVUQJPaWACe6BwRUYGATjBFzUulkbWAYkWQ0jgOjeMFoUENeQs+H3pDnNp9PEGx0HCmDTCsOB7EOgT4C0NAtlAAZTqsmotqaLx0P8Jo0JUxyJCAhQQGQZDoOmMJhMIXgAHQ+OglzRCEbyoAe1E9qgULZjsZDYgATAADCUADsUKaLU2I+NiZ73owMBQmhwjLghvogAAcvgACCVD6m67hYXsXAAvhTAwE4iqUEgjBMNYfA3j4t7ng+cD2ggiBRDoBlGd+M71jhnS+Dg5g4Xh36ESQJFkVWXSMLh5g2bR-gQFwYTUXksKjlAXjyZmeLhkSxYxkwNL8PSjYxJgKbbtUxbZkKxQAN71SA1GMsySY6gAPkIRDZCAXXqVpOnTF1cJhDAUQgPxIAlCAAC+s3UY1zUJuN7U4F1im9f1mnaUcc2zQAAge+DoE2mAtTQ5RAA",
      "N4XyA",
    ],
  ],
  [
    "Generate a fake phone number",
    [
      "MQAgKgFglgziCmAPAhgWwA4Bt4hhA9gO4gHEAu+IA5vAHbwBOyZOAZsgNY4xkNS1UQAN2SYArvDhkIzEA0np4AYzIhkIVvgapmAKFAB1HISiZM1Oo2Y50BeiFpjUAI0ZwAFADEmtJTmS0ACYgAKoAdADKYRpaOmQwAJQgYjD8gtI4AETyNIiZIDowHLq6Qm5Q+LQAXCCZAIyZuoUcaVUlICAAtLjw2CpabR1DIABWMJXozBA1mbaV8AD6rAyNw80wg8NdaoGBM6tb3TlIMwAMANp1nQBsALruIOennQCctwDeAEwAvgnvACzfRq6DrdGC9ZQUBibYZjCZTGZzegLFIHDrrGFDbrIXb7EGHOTwXIzc7uW5PV4fADM33OCVuj2eb3eNM6FOZgMaQA",
      "N4XyA",
    ],
  ],
  [
    "Generate a valid NIR (french citizen identification number)",
    [
      "MQAgKgFglgziCmAPAhgWwA4Bt4hhA9gO4gHEAu+IA5vAHbwBOyZOAZsgNY4BEAcgJIAlbiAAUrBnQDGEEFKhkoALzogoAEzqLWUKcyj5aIWgFdUAI0YBKAFChIOAYLVwp+DNkQhLekzBwKCLRumiBkEDjmUAzhIOrM8AA0uEggyLTqJMhwAEwAtOpQVIFutGQM+JjGZpYMNjYAbowwBrQAXCDcOdw2qNkcULRUbfUgIKD+2FJkaSBMGe7UdJp1Y3kpUxQMI2O7IABWMIbozBAd3DQZjD17fTAcMDt7IOvI6urnN8-r8+ruAMIEXTwJ7PF6dACyXzB624ADEejYxhN4JtZr9FlEYrJ4iwkeDJvBpvhtvi9odjqdzljwgB9XHwaEgO4PUG7V7vT5k9lzdLqAAiCTZYIZEMG5wAjABOACsAAY8nKJYqJWA5XK2urNXKAFpMkUJCHIRDnHJys0qlVqjVa9V67lrOIJAAKyAY-lJYL2+BMZHQvrhJL6ZHO5oA9Eqw2a5QA2RHIjZEmbIXkLVBxeAnGKoLTVCyMfHrQnEz3PCm0E7hc6aLNkHNlWkhRnclmPB3gt4fTr68EY1D8MrC56ocUgCXt27GjpSqWjcaJ6bovmLeRkACeedqhYXWyH5crZ06q7XjfwmiZraHHK73B7P2X-cHE92I-aY+fYz6JpAM9n+JRaIpn2IAkis27Fru7b7lSnSgdcLb9G2MJpJy3btveaYDiGH7MqO45ei+U4-jOc6gKUegzPATQMGu4SDFQYSUFIkgJGEEQgDo7ozLWIGsGxjhCOBqJJiSe5HBWMHcLQ0S0jAZBumQF79EOnZ5OU6QtFoHQAOQAN66VAfHwAAjiAAB0lwrJCIgAL42RK+mov4dk5I5GR2fpZk0hA9KsQAPrgJjmHJDAgAAHGOcoebpXnRHSDIgAFMBBSFIAAMwgDK0VmTWCn1mQp6hAF6AMIMZB8dwACk5rqNw2XHoVODFaVZQVdVaW1dlcGhc1ZVtXKHV1TZ2lIv+cjuP6LBpBk43BMw-EgPkhTFDMpTlJUm4Fo6EGiVB4kHuc0kMKeZQVJgSn3Cp7xqfMmmDiAem6YF5g-gA7GIqBnmIZUxgALOZR2yfJMRWGMUqvVYdkjWNOi0MgmCYBu4Q4A08MaBxkjBLITggAAhEJmy7WC0FVp0R0XRwV1do9ZmA3JCnZYDa1nVDNhAA",
      "N4XyA",
    ],
  ],
  [
    "Generate a valid SIRET (french business identification number)",
    [
      "MQAgKgFglgziCmAPAhgWwA4Bt4hhA9gO4gHEAu+IA5vAHbwBOyZOAZsgNY4BEAygJIAlAKJhuIABSsGdAMYQQAIwCuMKPRgx4cKABM6ZKKyizmUfLRC1lqRYwCUAKFCQcAkWBCwQs-BnxauiD4rCBkhJToyAxkcBLuwgByIMi0QYn8AML2ADQIyPIgUTEkyHCpXrQsVAxQZACePhDwshwgulBUdY6OAG6MahYAXCDcAEzcjqhlHOpUQz0gIKDGDDBkRdEbIWHNIAme3sggTGl+IAAcALQdXRvWtow+yJjYQQmJjktXuPDYshQGAsliCQAArGAWKJkCAjPhCJKTUHTGAcGDA0EgH7IXS6OFIzE-GQ0RBwgDaAAYrgBOAC6AG8LgBfSZfZa-XxpTYlHYwtwIw7lE6pXTnAAsN06dSsNjsDGer3g6SybJ+Wn+gIxoIhUOYsNGGUyBKWKLRWpB2Nx+LZoKJ8BJ5KpdPpYpZi3ZfP2Aq8cF8-kCwVCno+OTZK3wryISqUjWOaloVGwTRabVu0okmGUEEsLyo+FqMNQ9jClE5ZAYEcq1QLjV5exDYZArAjmCjQUUjU9htDS3DkcI0Y7KRAmDKG3jiZw8hT7SlGwzWZzmDzBYgRZLPgs5cr6mrdVrQb2hAISYOPaxvw1+fNSx1tGh+vhHmNIFN6JtFpSVtG9PpADoPiZN1MQvTNsxGekmQ-b4wngDBRxYOFfwAgUgOQw0gJfGCwNoCCoKAA",
      "N4XyA",
    ],
  ],
]);

let examples_anonymization = new Map([
  [
    "Anonymize a dataset by suppression",
    [
      "MQAgggdg9hCeC2BLAXgQwC6JiRBnE6AFgKYgAOATlAMbG75QBmBFqEujUFSEA5iKhAATDKlzF0AGgBQoRBAEgA7qlgFCGEBWKIhxCJkax5-QfKGIAbroCuqADbliFXNjw54ZKPUQAje8QAdLLgCsRiaipq6FAC1ISIxJakiOg4+DFaxPBQycKi6prUbCC+pDbiQgSxuvqGaq7wxDBB0tLJLlgQAFwgAEQAjH3S8GIA1ibdbSAgoAAqCfiMiC5po7hjWTnJ+I2ky8T2VUSaqNrCK8TUabUGiAcuIEyKZM6uENIzALQg4gHXXFwUxmIJAPwAVu8yBhCL0+stVgB9CCoJrDUHfECQmDQohw+xidDI1HEdEg9ZjYEY7TbYi9dAUGzEaazEALfTKUgQYjEY6xMogYraRg2RxKVKEHAQCzaa44PR3B64YIzUCQKUyq43BWGRIUSSCqDwXzyXnKCUCaBEZwgAAUZBs-kQ1CeFHIFCsGGIAEoZKrBSVxBQ8pltLd6i83jBgp9WQAxLggYgAD1RZACIHpJCl1CNKXwZmllzl4fueoNsdAZWKFVI1pACV4JFwNwguaaICQTbSAvrAGEAKIAeQA9AmbNKbc8iO529C4CrWQB1DRpJSkYoKIQ1fDadOoWjqUiWBxM0pqQSsPjEP2sthVUZjUi4Gznevxbwcq+8OjmoREQVtC9EBeCoGwyHwcUAM0AJCRAABpJwXBgXBb1AGd8Hcd8YFoMg0meeCvjYGAEFSWAY39dkQG5ZM1nGc17EcPcCUPet5HbUh7xAAkW0RXwYAqEAT3sJkCy0Ngf0XH4-i1QEqVBCEoRhOF2LzMkFKxJS8X6HiiX4iAKnUztxnkkFvzpEAAFYAAZbOsllQDjeQHHsWADXXaieT5cgzkwFy1D0BkoGibMTQoIhEREdBj1PVp-QASWYa1zncGBXMUXATAzV5kIUfiKAUeQQAAZgGEdrMskcBgATgAdgANilDLRkYw1PDYNzK3UWcSgFfQoBsJtqiyUsjEUWUuCqIr62IuAkGQM0orECRF3mRZPNo4yNgYxw0rUJ9iDII8QFgcI3WwMKIqizjpXyaLESYREEjfWJKCSOpuuVWNpMOWSXFMzFsQgXFYX6S7CEir0jMBrTQb6a6HsYJ7LiMikAbu4gAAUziDdGZnkB10ATbgMDhayACZyrK8nbPq6GQQG9BCeJ0Z0DhGnrLp6QgA&",
      "N4KABGBEBmCWBOBnALgfQHYEMC2BTSAXFAFID2AFupADThQA2mKGO+RkAIqfrRJAEYJk5VABNMyNlADMARgD0ABgCs82QE4A7ADYadSLmwAHeqQCeueISgAZTAHcL8AOSIwAYVLHM6M3r7ikqik0KjkCFKQigqK0vIATIqJ-lCw6ADGXlKyyop5irwMTGj8pOgArojWiXkgAL5AA",
    ],
  ],
  [
    "Anonymize a dataset by randomisation",
    [
      "MQAgggdg9hCeC2BLAXgQwC6JiRBnE6AFgKYgAOATlAMbG75QBmBFqEujUFSEA5iKhAATDKlzF0AGgBQoRBAEgA7qlgFCGEBWKIhxCJkax5-QfKGIAbroCuqADbliFXNjw54ZKPUQAje8QAdLLgCsRiaipq6FAC1ISIxJakiOg4+DFaxGT2qLTCouqa1GwgvqQ24kIEsbr6hmqu8MQwQSEA8skUShSpJgXogkqphCDUUJ4B6MT2aqwQQhMglg42dCAaGSQguKjNO8TUNr3oamI7NmSUdLj9IoMyoL42aakgANbE2fiCCbyjlVQvkQ9lSak4FHUpHuYgkwWkXVuMAAXCAAEQARjR0ngYneJmR0mkIBAoAAKgl8IxEC40rjcO8sjk8qQmqRqTNqkRNKhtMIaYdXnoDIgOS4QExFGRnK4IMSQABaA4BagxFyEkmaxUgABWsrIGEIqLR1NpAH0IHtiNitSSlXqYAaiMbcrh0BarTbNfT3hrbfNFvAAMKEKCIWgASQgAFVesayIh4FBkQB6FOW5oAMQAStj5eSSAolKQIF8ubFymNecRGDZHMMiDgFgLVThhYZEi5giTQJAmxZtK26iKxZIxhNgaXqg3RmwoERnCAABRkGz+cMSyGUKwYYgAShkParCnEFGSNSyw4aUplMHhR8zXBAxAAHnscqRUQum+N9u4zM2g5CvUoqdmO+ZlIcqCVKQ35-CQbo-hMpBIP8aSVt+QYAKLtCmj42Asi6SkQ7i-gacD3qSIAUiWr50niyggo42jMvk37yL+pBsNUrrur4MCVMsqx0N22riCqaq4H6Wr2vqhrGhxyFejJupyc66K8Wa-EQJUykgD60magGExRughlajwqIAKwAAx2fKtr6agL6ohidnuUSD7yA4sxjsWIClsQ5ZMrkbHbMCFBEGa9xcQsAzEGaTBmgkfIrPYay4KJoARswC58u4MCzIotx8AETguNg-EUAo8ggAAzBiKY2VZKYYgAnAA7AAbE2xW4vYjhkWwsCHlRJH4CUCiVvoUA2P8F7aFeopnFk4wUNUtXfnOcBIMgQUDLC6BZdRlIBXRTkMoxA0hSyULiEJ6XrL4czcUsMWZfKSriYKXBSQ5mqyY68nohFUUxXpdqqUD6lojFiWMMlAp6T6f2OUqAYACK7uZtoxQAsvIxrtR1NkKjZGJkxiZJ2ciNN2QAWhDuO7njznGgATDZnOU5T1M2bT-MM0zSoxQACryp441qs3oKu6CPtwGAc3ZXU81z2JAA&",
      "N4KABGBEBmCWBOBnALgfQHYEMC2BTSAXFAFID2AFupADThQA2mKGO+RkAIqfrRJAEYJk5VABNMyNlACMATgDsANgC0ABgCsygMzSadSLmwAHeqQCeueISgAZTAHcL8AOSIwAYVLHM6M3r7ikqik0KjkCFKQAEyqMWpaarq8ULDoAMZeUtLqqrmqyZCMzPyk6ACuiNYxuSAAvkA",
    ],
  ],
  [
    "Anonymize a technical ID (like a plate number)",
    [
      "MQAgggdg9hCeC2BLAXgQwC6JiKAzE6ApgMYAWEixqANiIgCaESa6KEBOAziABQAO1DIRAQArvABGHTgBoQnPOgDuqdsM4dENEAGtCsWSD6kYwsZOlyAdDYCUAKFAAVUsPTtUETsSiMQ8VE4dECVEalopEEJAtnYCKBBRDQJSTxThNQBzQgAPf0Dg9ATPGAQUNxJySm0GJhZYq3t7ADdpLAgALhAAIgBGbvsAoMQITI6mkBBnVwIPLx8-IeDQ8JA1AVRiYRpaMlVNoi46CHpEZoZRHdh7SdAAIVgQRlxUUWp0EHHbkABaEABVAAKgIAogAlADCYAAyiCQNRCOhDtwVhEMoQNlt6CAJI85vQoPAAcDwVDYfDEciblNfvCoEoOFRkgikdIQmE0WsMYIsTi8Z4CUTqPTGYFhCyqd8-qdMoh0CiOTj0ZjCNjcWsBYSnohZfLqaA-jscOhXHE9h5iMj2atIqJmFBRGRVdS-hoEZaoFwvpMfX8AFYKCB8DCkLrdDZEAD65ik7AGPt9IADMGDJrDGnYWmo8Z9S29Cfcnm8vkIXQA3gBfCY0lzCc0HDggYiCTgabhUCBKpswViZURqbFFJupUZuGbPV7vJWpc6el3yQjuorsfM+5NBkNh1w5SPNGiiQg5yZ56kFubFxirhPNwJtq+JkZ8UToMMABl6ACYAMwAFgArAAbAA7AAHAAnGAdwQgAIiCABiR4Jj6DroE+L49O+37-sB4GQTB8EDPqIAuIg3BLE2aQ0AoXZJG4CScKIfBqK2TwYKg85uiQy73uuqahj0VDsPQ0biLGiEnkhsxFgspankhN6toQnD3pMfyPs+b6fr+gGgWBiGSShaFhgAVAMQA&",
      "N4KABGBEAOA2CGAXApgfQHYFcC2AjZATpAFxQCCAQgLQCMATAMxUDCAIpADThQDOhAlvFgko9BgBYArADZajcXIlVKbTt0gALZAA9UANyGZkIyGRoAOAKKopFMqgBiFyajYB2N2oiQAxvAIAJhg4+ESkkGJSspGKCirsIAC+QA",
    ],
  ],
]);

let examples_pseudonymization = new Map([
  [
    "Add noise to data",
    [
      "MQAgCgzgpgrgJgewHYE8C2BLAXgQwC4bIgYQh4AWUIADgE4IDGUEpCAZmbTkhGwrZiQBzEDhBx8OaHmJJRINNyRRaIAFCgK+MpXGSQtKABt8GYWQTyzcDADcMcGDiM0VEIg24gklo8iEqIABGUBqiLBhCynAgAO4YFAgwMmwwtBSBZnwCpsgAdGEAqkhGGADWVNzI6Ni5SAA0sjTQ8NWYuAREYhJ4UlAyJCAQBEYuhiZ4UDF4lhkg-JFmznq90gWgACq6hlJEMyAw0M2wiKiYEHUgnnIhIABcYQC0OlRoUNxmIuwvK2KDGGhqPxekgZPsKlBqCAABRQPJCPKiWwqHABUi2ZwwZiNQQAekUAA9GkgYGgQqpvjZhmYGDIMUYsRA8syAJRPH4MBCUQxIJjzDhzHp-UgAoFcUEwuEIgyQkx8i5vED0rFxBLkH4KqjixBoJWYqCNHC2BAOK4IUYkQg8RrMvJs0DPDKGYikMTKKYWGUOKCgjBsFAHJCKCAVGLK5iS+GItjWeR0BBBIxQRQEBgyzm0GJmH4LIRLFxC6RssIASTkFEGUAJOEBSca4dI7umlluDD8RzY9F1c1z+b1DOY6xAADF+DpBrEqNcDkc5sGyqRtQg0AARNKXMeTWv4OFhLaDeNMFguq7OJMxABEibStE+F-mqgvODgNnMPhIUAvBTUyNoECtdwgBeACMF5qPOnwPGoIAgJs5AThgoxDPBbBgroqJUN8czUG4RAhHgk4+iAjwAEwPiAADUZFLmgRgoNBxFDMYUC0vwDwwRxIAAFbuEg1D4OQgFPgEYGcfOEDsZxjFbtQExQEJADeCnPjEeQYTC2pljIpEgCRLIAL76aJUnPJ2y48cgQkYfeoCGJySC-jIQpkCgOGemYkwBLQagMXBCFIRAKFoVQPSYRwQhcI4lz4YRciPAAzAADIl5EUUlKUSCgi7cDqdEMc80BJqxtCSZxFl8QJQkRc+TiTHAAD6yDGTB4mlRxzyhWAOB-iobVSbI1DJKOOR4EJJHJQAbI8iXAdNJHNe1Bg5cua5cJ0SB9VJghCY8YDpSuC39YSQkUXtyUHQxnEdTuXU9SVl39fMySDXgw0pmNk3TbNiXzWoQA&",
      "N4KABGBECGDmCmkBcYCMBOANOKsBO0AJgK7QAu8hA+gPYB2yUG66AtAAwBsrAzO5CAC+QA",
    ],
  ],
  [
    "Preserve coherence with non-reversible functions",
    [
      "MQAgCgzgpgrgJgewHYE8C2BLAXgQwC4bIgYQh4AWUIADgE4IDGUEpCAZmbTkhGwrZiQBzEDhBx8OaHmJJRINNyRRaIAFCgK+MpXGSQtKABt8GYWQTyzcDADcMcGDiM0VEIg24gklo8iEqIABGUBqiLBhCynAgAO4YFAgwMmwwtBSBZnwCpsgAdGEAqkhGGADWVNzI6Ni5SAA0sjTQ8NWYuAREYhJ4UlAyJCAQBEYuhiZ4UDF4lhkg-JFmznq90gWgACq6hlJEMyAw0M2wiKiYEHUgnnIhIABcYQC0OlRoUNxmIuwvK2KDGGhqPxekgZPsKlBqCAABRQPJCPKiWwqHABUi2ZwwZiNQQAekUAA9GkgYGgQqpvjZhmYGDIMUYsRA8syAJRPH4MBCUQxIJjzDhzHp-UgAoFcUEwuEIgyQkx8i5vED0rFxBLkH4KqjixBoJWYqCNHC2BAOK4IUYkQg8RrMvJs0DPDKGYikMTKKYWGUOKCgjBsFAHJCKCAVGLK5iS+GItjWeR0BBBIxQRQEBgyzm0GJmH4LIRLFxC6RssIASTkFEGUAJOEBSca4ddzsM1DlHs53J9eCMAeh2e6-RUghIqYUShULPqYVi5AwDHVb24rpAeeRchgQakod+qtGoiMsRwKFIt2DW-iFA1Nao4cnoCCyXmJQDPiQj0MyNoEAwiaoqV5nR4HcXFuQ4plvOIZznBR3mtVVEgfKtJloJBliFRpDEeb1fRjTwAJ3YIqFFBAIh-Ao1A-L9kDuEAACIAEYaLUU9PgeNQQBAUAAEEQDYKBYlHENG0IpAO1BbtxAHAQzGHWdGlKCoQHIKR1X4RTlLLQpaAwNiQGeaAk1pfgHnYkyQAAK3cJBqHwchqJomNPzwAB9FC3kY0zTwgYzTN0tSIHIDStLs6gAQQO5cVxVyoAAMQAJXc3z9KgQzaG80yLOQayKDskxhhcq8EvYzy0pM54lP8wKMGC0LwtxCA0iiuLGJ00BotUhAnQEsoIEaDIA2uAilORCSkKHakGHE891TmHA4BsPDvjmAADaApiWmgcC4N4kJ0vTjGSmZUp09LLKy2zaKCDB0nIJyeigQqhigKYStMmNjDgOzLuu278HujiQDa1QoA-AMAG9QbyL6KB+yYAF9Yb1BkDR+ZjzHiXcVyoOZNRleqjDwY6uq8wn2OebUABFfpeny7oAWTMOy6IATgAdgABkeNm6M5uiNjZtm7n5wW2YALQenyVigWmcAJRmmaZjmuZ5vmBaF-mxZJ3y7rATboCOiXTKSPBqGSQGUzsgAmfmADYec5i3GKAA&",
      "N4KABGBEBmCWBOBnALgfQHYEMC2BTSAXFAFID2AFupADThQA2mKGO+RkAIqfrRJAEYJk5VABNMyNlACMATgDsANgC0ABgCsygMzTIIAL5A",
    ],
  ],
  [
    "Preserve coherence and enable reversibility with a cache",
    [
      "MQAgCgzgpgrgJgewHYE8C2BLAXgQwC4bIgYQh4AWUIADgE4IDGUEpCAZmbTkhGwrZiQBzEDhBx8OaHmJJRINNyRRaIAFCgK+MpXGSQtKABt8GYWQTyzcDADcMcGDiM0VEIg24gklo8iEqIABGUBqiLBhCynAgAO4YFAgwMmwwtBSBZnwCpsgAdGEAqkhGGADWVNzI6Ni5SAA0sjTQ8NWYuAREYhJ4UlAyJCAQBEYuhiZ4UDF4lhkg-JFmznq90gWgACq6hlJEMyAw0M2wiKiYEHUgnnIhIABcYQC0OlRoUNxmIuwvK2KDGGhqPxekgZPsKlBqCAABRQPJCPKiWwqHABUi2ZwwZiNQQAekUAA9GkgYGgQqpvjZhmYGDIMUYsRA8syAJRPH4MBCUQxIJjzDhzHp-UgAoFcUEwuEIgyQkx8i5vED0rFxBLkH4KqjixBoJWYqCNHC2BAOK4IUYkQg8RrMvJs0DPDKGYikMTKKYWGUOKCgjBsFAHJCKCAVGLK5iS+GItjWeR0BBBIxQRQEBgyzm0GJmH4LIRLFxC6RssIASTkFEGUAJOEBSca4ddzsM1DlHs53J9eCMAeh2bmmpAVagDGSVpZ9TCsXIGAY6re3FdIDzyLkMCDUlDv1Vo1ERliOBQpFuwc38QoGprVHD6xAAHVKkZ3HFuGDLIoKkM0lQnVR4jvbj4ZoWhAVpNHMSR4NQyS-NIYSwlGIAAIKiEgMQAELBAg54nh6Z7qv2l56gyUAsjeGyWDgs4YFAyLEHgjSeLOEZ-i4tyHFMjQhJ47HeMgjyGMitAgYmVCpLynQ8FcXjxo4fKcsBVpMmEWwqFA2QGi8zoscEVADiSZKBJSJAEOJRGMmBug4TEhb9AUaiCSByB3CAABEACMLlqGEADClGUKQ7pwKQ+y3HAw4mIYcCNLEVAMDskwgAATFJTEQI02QgDGQl4AA+kghHcDEJjDHll4+X5EbXDp4iktQHp7LoMZJhAh6TLqeE-AABo8jyOICjyMbog1QAAcpeAC8TU+peeQAFbuCUnXzNQEnlalUk3FQfg4GFMQhOpCgbp88g2fRqrnnM3WPNtcADRVKWUGNbyTRgSb5W8c0LUYS0ICtVprf5G3VQJbgejgbCTKoN3HdCub5mZVAADyPAAfIdIYeg2ojOhA+7UHVcDjmEFakMuEZzECERBK9CQBvsINCd+uh4OKvD8CmoHQgEMhiHD+UuOGmX0LqcxWQjpFecNEAPCAmUYNlpVvDlw0y7LgYYAAjlizksyqoBrprWLjbrrzvJJYhi4LVXIN2oj4+8FK8hpT6AfJpSOXI2nkDg9tIGosvFbl71QMrFWq7LBta1AOu0Fi-tq6ADPQDHKpeSenwPPHzzQEmtL8OHavzcg1D4OQzkuVlJXB55asPdHrmV0Hl6h0xNdqye0vx7XzzaggaDeeQJpMGWhS0Bg5fUACCB3LiuLBwAYgASp5WdDMYw4zLQBey0XSAlxQ5eB4rUBt7LKuuUfwct5Qp-o2Une17LPeFX3A9D1AI9jxPU8z7iEBpAvZeaggA&",
      "N4KABGBEBmCWBOBnALgfQHYEMC2BTSAXFAFID2AFupADThQA2mKGO+RkAIqfiAL5A",
    ],
  ],
  [
    "Preserve coherence and enable reversibility with encryption",
    [
      "MQAgCgzgpgrgJgewHYE8C2BLAXgQwC4bIgYQh4AWUIADgE4IDGUEpCAZmbTkhGwrZiQBzEDhBx8OaHmJJRINNyRRaIAFCgK+MpXGSQtKABt8GYWQTyzcDADcMcGDiM0VEIg24gklo8iEqIABGUBqiLBhCynAgAO4YFAgwMmwwtBSBZnwCpsgAdGEAqkhGGADWVNzI6Ni5SAA0sjTQ8NWYuAREYhJ4UlAyJCAQBEYuhiZ4UDF4lhkg-JFmznq90gWgACq6hlJEMyAw0M2wiKiYEHUgnnIhIABcYQC0OlRoUNxmIuwvK2KDGGhqPxekgZPsKlBqCAABRQPJCPKiWwqHABUi2ZwwZiNQQAekUAA9GkgYGgQqpvjZhmYGDIMUYsRA8syAJRPH4MBCUQxIJjzDhzHp-UgAoFcUEwuEIgyQkx8i5vED0rFxBLkH4KqjixBoJWYqCNHC2BAOK4IUYkQg8RrMvJs0DPDKGYikMTKKYWGUOKCgjBsFAHJCKCAVGLK5iS+GItjWeR0BBBIxQRQEBgyzm0GJmH4LIRLFxC6RssIASTkFEGUAJOEBSca4ddzsM1DlHs53J9eCMAez3X6KkEJFTCiUKnqYVi5AwDHVb24rpAeeRchgQakod+qtGoiMsRwKFIt2DG-iFA1Nao4fHoCCAZ9DFoKGoBHMDdVZ7E0Af-RAFRQ6xALZBjeRIYiNE0Yk5C0ICtV0kBiJcIzmIEIiCDBSjwAN9kMZFaCOOY8HFXh+BTK0YUMR5vV9f0WQKNRcJg5A7hAAAiABGFi1GPT4HjUEAQGeaAk1pfgHn48SQAAK3cJBqHwchmJYqiCBjFROIk48xIkkA2DYNitO0gwcBsAlmLYgAGPjDN-KAUAAMXoNAAFEkFsRS7LstiAH0nIAOQAYQAJQATTADYSwAeV8ryAGknOCziwg8tiDmgBcYwJD0dRwbM4CgGMkASMjoVKYZ+R3FxnD8WI23IHAuFpNwmi8JI8GoZIQHvR9nw9YZaE+WiwnMtiACYAGYABYAFYADYAHYAA4AE4cCCBg8rYIQp0ksojDQHxqAARzwvAYFsWICRQLAAEEACF-IAEScuyAHEAAkSwAKRigAZABZXyIrAABFQKAGUNkKAA1AB1AANYKAC0wi2LVjIwAkaHqi9JlUPLccEKhJ2ndU5PScrBQQRRcvyswiqIeJtxwWknFGANbkOKYALs-hOurWsDXkLgTPKiytxceNHD5N9OVBHLCvMZBu3ERY8FIMJoTmMWYxOq46oa3HWAFXRUm3bKaYK+mkEGzYO2yQW5nDcrhYxhRDhkJMWHmVQoEO1mdC8GaRsRa6jHceQne+FLefMl1ZHpBw6P4zZLDyh8nxkMQnc58QoHT59mMIrErME4w85mWgDPE6TkDkihFOUv0MBULy0+6yY4HU8TjwgKv+OeYy4EUgBvYe8kb1TaAAXynruJOeXT9Ks6yjJMszLJX-i24zwvaGLzebPsxyXLc1jkp8gKQrCyLorihK1CAA&",
      "N4KABGBECWAmCmA7ALtAZteAnSAuKAHAEwDMBBAnACykkCsVlFkIAvkA",
    ],
  ],
]);

let examples_other = new Map([
  [
    "Preserve null and/or empty values",
    [
      "G4UwTgzglg9gdgLgAQCICMKBQBbAhhAayjgHMFNMkkBiJABzBAnFGRQCMAbXOAlJAO5ROnJHBgAXJLjp1OATyR5CxEknhiAriPVgkIbHQmKAFChQBKJMFydNTSkgC0SZpxABjCTDDkq-pAArCHg6XAkACzZOKAgJADobOxAsAOUCPwCkD3g4ngk2dJAAE1T-BiYWEDYuHj4KKloK5jBWVDhtTn4hHXEpGTlFdNV1OC0dJPsIRybGFraUAyN5buFRPulZBSV8IlJR-UNjJDNLa1spzCA&",
      "N4KABGBEA2CWDOAXSAuMBtcEzEgNwENoBXAU1TAAYBfAGiwl0JPLUksjoZ3yLIsid62Hs35oAdsWjRqWALohqQA",
    ],
  ],
  ["Apply a template on each value of arrays", ["G4UwTgzglg9gdgLgAQCICMKBQBbAhhAayjgHMFNMkkBaJCEAGxAGMAXGMcq7pAKwngAHXKwAWyFMFwMAriAhYeeQlx5JWIbIIYiQ1ELmbjKaqlA3YJU2SEWmkxACYgAHhKiOXd0xe26JAN4BAHTWckgAPkgygoLgAL7xSCJIgjDQrLBwSEHBHi6JWEA&", "N4KABGBEBuCGA2BXApgZ0gLjAbUrSANFAEaFQDGZkAJpALogC+QA"]],
  [
    "How to handle nested/complex structure",
    [
      "G4UwTgzglg9gdgLgAQCICMKBQEQgCbIAsATJgLYCGEA1lHAOYKZJIC0SOANiAMYAuMMExYikAYiR8AFlAhIADhWkKYdPnIFIKcLWDAUAnkhgAzBeAjwIzUUgBWluIunIUg+tqgAvJbDgQAOnkLKyxRShphWwkoPiQAdyhOTiQAIxAFKGC8SRhJKQyI2gZM4M46DLwQEzpYvzSQThh4m1F5LJAo23F8jMUwEDg4mFS7XjiACgokPBhKOgBKBKSU9K0eHhAIaFTuZeVpDJQAfRQkYAowKApdjLgKMhBW2zox-gAFS8G+V1Pn8KoxUY-1s7C440EXW6tgc8GcUlc90eYWhAMiINREhqjTwcjo0CqvXMkHgxlG4yQPG0DXWm22+BmUAG-E4Bgx0L4IDI8k4Sk6qAA3gLJLE9gEkRkAL6SlGosGNCFCdndWFOJQI1AQACuYAlstRSCKUINki5PL5riFIr4Yu1uoeUplyrYHAV-EhztEqvhri5FCS+tRRs9Igkh3OFE4WoyEAEAxy+KghPD-W+ZLecVktK2Oz20jAMC19CkqFO+W01A0eRTXyGSFe4z8IZYnO5vM5luFTXi4CQ4odSGlASt3d7ATtEsHkoAAlaAscArN5jppVggA&",
      "N4KABBYEQPYE4HMCGA7AlgLyQFzTFAzlAFxgDa4kVoVtkUAJjALZJoonQDGLADqgE8AdD2ZQANJTpUovAKZwC+IqQrT1YGhu1QUSZnM5QANnPxIJU7dKgEArnD0GjzNMdNxL1jVDms3RlBW3gC+kt7UwREQuvqGpFAAVjCG4dG0tg5O8dAMSABuaERp6fR+bMaBUdoh1bQAunVhdVo6TP4cCaL8KMIAZp4lPvKKypxq6a2l0NlGSMZsHEOlmY5xLgpcaApe0zHlAQlBpc2lUyuzCX3G8Giou3url9CmCHCoDA-Tvv6VR3XqWreRqAqyNEJAA",
    ],
  ],
  ["Temporary fields", ["G4UwTgzglg9gdgLgAQCICMKBQEQgCbIAsATJgLYCGEA1lHAOYKaZJIC0SOANiAMYAuMME1aikAKwjwADhX4ALZCn5gKcaCDj8A+gDMoILnixjKNEWKQU8eNirUatSgN7PVcPAEEu0+RQByAK5kSADMAL7hWCzsnIZ8gsIxYpIycoqo+obGyUhm1BZi1gSoACryIEjyUDaaSMAUXIGVUBBIyK4AdPbqBlp6BkaRWEA&", "N4XyA"]],
  //["Multiple mask and/or multiple selector syntax", ["", ""]],
  //["Seed parameter", ["", ""]],
  //["Caches", ["", ""]],
  //["Change date formats", ["", ""]],
  //["Parse raw JSON", ["", ""]],
  //["Generate sequences", ["", ""]],
]);

let examples_generation_a = document.getElementById("example-generation");
examples_generation.forEach((params, name) => {
  let link = examples_generation_a.cloneNode(true);
  link.innerText = name;
  link.onclick = () => {
    loadExample(params);
  };
  examples_generation_a.parentElement.appendChild(link);
});
examples_generation_a.remove();

let examples_anonymization_a = document.getElementById("example-anonymization");
examples_anonymization.forEach((params, name) => {
  let link = examples_anonymization_a.cloneNode(true);
  link.innerText = name;
  link.onclick = () => {
    loadExample(params);
  };
  examples_anonymization_a.parentElement.appendChild(link);
});
examples_anonymization_a.remove();

let examples_pseudonymization_a = document.getElementById("example-pseudonymization");
examples_pseudonymization.forEach((params, name) => {
  let link = examples_pseudonymization_a.cloneNode(true);
  link.innerText = name;
  link.onclick = () => {
    loadExample(params);
  };
  examples_pseudonymization_a.parentElement.appendChild(link);
});
examples_pseudonymization_a.remove();

let examples_other_a = document.getElementById("example-other");
examples_other.forEach((params, name) => {
  let link = examples_other_a.cloneNode(true);
  link.innerText = name;
  link.onclick = () => {
    loadExample(params);
  };
  examples_other_a.parentElement.appendChild(link);
});
examples_other_a.remove();
