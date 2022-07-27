import './style.css';
import LZString from 'lz-string';
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
      uri: 'https://raw.githubusercontent.com/CGI-FR/PIMO/{{ version }}/schema/v1/pimo.schema.json',
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

var masking = 'version: "1"\nmasking:\n  - selector:\n      jsonpath: "name"\n    mask:\n      randomChoiceInUri: "pimo://nameFR"\n';
var input = '{\n  "name": "Bill"\n}';

const urlParams = new URLSearchParams(window.location.search);
if (urlParams.has('c')) {
    masking = LZString.decompressFromEncodedURIComponent(urlParams.get('c'));
}
if (urlParams.has('i')) {
    input = LZString.decompressFromEncodedURIComponent(urlParams.get('i'));
}

var editorYaml = editor.create(document.getElementById('editor-yaml'), {
  automaticLayout: true,
  tabSize: 2,
  scrollBeyondLastLine: false,
  minimap: {enabled: false},
  model: editor.createModel(masking, 'yaml', modelUri),
});

var editorJson = editor.create(document.getElementById('editor-json'), {
  automaticLayout: true,
  scrollBeyondLastLine: false,
  minimap: {enabled: false},
  model: editor.createModel(input, 'json', Uri.parse('file://data.jsonl')),
});

var resultJson = editor.create(document.getElementById('result-json'), {
  automaticLayout: true,
  scrollBeyondLastLine: false,
  minimap: {enabled: false},
  readOnly: true,
  model: editor.createModel('', 'json', Uri.parse('file://result.jsonl')),
});

document.getElementById('loading').remove();

// Examples ///////////////////////////////////////////////

function loadExample(params) {
    editorYaml.setValue(LZString.decompressFromEncodedURIComponent(params[0]));
    editorJson.setValue(LZString.decompressFromEncodedURIComponent(params[1]));
    resultJson.setValue("");
    autoPostData();
}

let examples_generation = new Map([
    ["Generate first name, last name and email from referentials", ["MQAgKgFglgziCmAPAhgWwA4Bt4hhA9gO4gHEAu+u82AxmSAG7KYCu8cAZgE76ohfwO8AQDsyUZjABQoAIL9Bw+GImYQNZCJAAjHMhCZ8GtRyjYANCH2F42hTHwsuNHPi5WQABQCSAWQDyIFBiwiLMCkKi4swy4BBKQXD6QsQCkcrRmHDoyDQA1vAAJkEiMFCFOGTxIABE6FCo+DU6wchcAJ6WsRxuCCgY2CD1jQBcAPRjYajwAKIAcur4YsjBSQaw9PgcCCIA5piwECBT7AB0UlIMwmVLI7UAjDVSqMgwecG7IxcgIKCmXDB6NgyABybJQfIgFjoDxcTSFXjHNA4biI-QHQEgLZI6bSH6gKo4E4gQhmNS6dQCZBkIolMoVDymajFE7FGr-QEAfROT3xIEJIA59Beb1qyEKhWasGO8CKtIoVgl-OqTMwxWCypwACkAMr+BaOMjoFhkKR83QaFgwHA2ElkkC7KBXKxaeAYMjtEC6-UlY30BXDfBm36amWIYWvPJ2zBqa20ej6OEiBF8JisHAagU1E4AMQASs00koVDEfgBaKjxtxfH61kAAKwcIhyVTu7KgALI3ORvNrIryMBrdZAFfFhTbvbrFaTKYAwgQIfBvCIAKpcKBtwPjSbI-NPYOgGDIxYiFzoMicXoCzCvegnSzaE12qogQpQDjpMQRYuZF3FVUssieIjpW8B0NWwZ1o2SwthAbY3lyPKQSA-aDsh5aKuOtSTrW07wrw874Iuy5rhutRbhMMBOLmBbfCGApuisJhQMy0bkjgNC8H6tKonwAroAIDBQI4cBpmwwEEpQCLKrAli2laei4B8gz9li2weugOA1DSAzUvAvYEtUOlYHp-LIAUawVDANDruewlaNipD8pQj5mMUAoCFRmD0GJKI8Hw+CEu4AGMMw4kHoKvSNAIJQ9FwLziEsIB3BAZBGoOEzaIYuynLs+DoPEcIaBUqDtKcnGoGM4pXGIZYAEwAAz3AA7GMVofGWeVlsZN40jAYwRYYNjuPCIBzPgsg0C4X5tCiLCnolpSWIGfQ0smcACnl-JuiZfXPkcSBkHCgrzXQ9kSZF7jRRmIhxQl9nJSQaXoBlYwijS8XBIUMC5VAVQsNopzCWMMACVAuwDRhcZgRQXBDlBTawW2jFmDh-bw7WY5tgA3tjICnEK3bTCAAA+BhEMIpNjRNU0ZCAAC+9OnLj+MIV2xJk0NlNk+Nk3TfQjMAALtPlLxmATXBPEAA", "N4XyA"]],
    ["Generate first name, last name and email from sample data", ["MQAgKgFglgziCmAPAhgWwA4Bt4hhA9gO4gHEAu+IAxgE7zJk4B28xAbspgK7xxkQMQmfPgDWcTFFE46nEB268AUKADqOLjBzIQAWWQ1R+NtQFQmICiAAmlfrAA0l6HFghkFzgHN8NKP1RnQSoPEAAjHDIaZHN4axB8C1lMGwZkJSU2eBoYKESALhAAIgBGIqVUZBhRcy98jJAQUH4cSsNjEErqkEIoTBSveBZoxnd5Th4gshA3XNQ+g0tKBR5XCxaQLygspPgAM2yhsihOJUbQCEOeyOjzZ1aDIxNUfGt4FN7+ECY0XhmmKjcN7xO5FdBQF75AD0UJgXBoP1Q8AAYgAlIogM4gAC0uHe8CoFBo9UapJAACsYIl0AwIIUinsoDkyAB9RHwcpkrriElknHuazWemcvm4tpPXl80kwNBYeD08GQmFwhG-NHlLGgGVI6iJKjwdBkOB7Xz3IRVabspxhLjTT4QGxQPYHOhMaZ0F1HE4pDzxRnveLsmBY3FabCE3yS0mU6m0+mYC1s34ixrcmBRxq45CC4VYslix7GDN87VyhUQ-DQqHs9UNJpm+CVPogf2YeK9frhHBUfAYW1xFs0Xtm9B0Nh5TTjRTB85LGx2FxOQgaLRjXJMLzYTpVUQJPaWACe6BwRUYGATjBFzUulkbWAYkWQ0jgOjeMFoUENeQs+H3pDnNp9PEGx0HCmDTCsOB7EOgT4C0NAtlAAZTqsmotqaLx0P8Jo0JUxyJCAhQQGQZDoOmMJhMIXgAHQ+OglzRCEbyoAe1E9qgULZjsZDYgATAADCUADsUKaLU2I+NiZ73owMBQmhwjLghvogAAcvgACCVD6m67hYXsXAAvhTAwE4iqUEgjBMNYfA3j4t7ng+cD2ggiBRDoBlGd+M71jhnS+Dg5g4Xh36ESQJFkVWXSMLh5g2bR-gQFwYTUXksKjlAXjyZmeLhkSxYxkwNL8PSjYxJgKbbtUxbZkKxQAN71SA1GMsySY6gAPkIRDZCAXXqVpOnTF1cJhDAUQgPxIAlCAAC+s3UY1zUJuN7U4F1im9f1mnaUcc2zQAAge+DoE2mAtTQ5RAA", "N4XyA"]],
    ["Generate a fake phone number", ["MQAgKgFglgziCmAPAhgWwA4Bt4hhA9gO4gHEAu+IA5vAHbwBOyZOAZsgNY4xkNS1UQAN2SYArvDhkIzEA0np4AYzIhkIVvgapmAKFAB1HISiZM1Oo2Y50BeiFpjUAI0ZwAFADEmtJTmS0ACYgAKoAdADKYRpaOmQwAJQgYjD8gtI4AETyNIiZIDowHLq6Qm5Q+LQAXCCZAIyZuoUcaVUlICAAtLjw2CpabR1DIABWMJXozBA1mbaV8AD6rAyNw80wg8NdaoGBM6tb3TlIMwAMANp1nQBsALruIOennQCctwDeAEwAvgnvACzfRq6DrdGC9ZQUBibYZjCZTGZzegLFIHDrrGFDbrIXb7EGHOTwXIzc7uW5PV4fADM33OCVuj2eb3eNM6FOZgMaQA", "N4XyA"]],
    ["Generate a valid NIR (french citizen identification number)", ["MQAgKgFglgziCmAPAhgWwA4Bt4hhA9gO4gHEAu+IA5vAHbwBOyZOAZsgNY4BEAcgJIAlbiAAUrBnQDGEEFKhkoALzogoAEzqLWUKcyj5aIWgFdUAI0YBKAFChIOAYLVwp+DNkQhLekzBwKCLRumiBkEDjmUAzhIOrM8AA0uEggyLTqJMhwAEwAtOpQVIFutGQM+JjGZpYMNjYAbowwBrQAXCDcOdw2qNkcULRUbfUgIKD+2FJkaSBMGe7UdJp1Y3kpUxQMI2O7IABWMIbozBAd3DQZjD17fTAcMDt7IOvI6urnN8-r8+ruAMIEXTwJ7PF6dACyXzB624ADEejYxhN4JtZr9FlEYrJ4iwkeDJvBpvhtvi9odjqdzljwgB9XHwaEgO4PUG7V7vT5k9lzdLqAAiCTZYIZEMG5wAjABOACsAAY8nKJYqJWA5XK2urNXKAFpMkUJCHIRDnHJys0qlVqjVa9V67lrOIJAAKyAY-lJYL2+BMZHQvrhJL6ZHO5oA9Eqw2a5QA2RHIjZEmbIXkLVBxeAnGKoLTVCyMfHrQnEz3PCm0E7hc6aLNkHNlWkhRnclmPB3gt4fTr68EY1D8MrC56ocUgCXt27GjpSqWjcaJ6bovmLeRkACeedqhYXWyH5crZ06q7XjfwmiZraHHK73B7P2X-cHE92I-aY+fYz6JpAM9n+JRaIpn2IAkis27Fru7b7lSnSgdcLb9G2MJpJy3btveaYDiGH7MqO45ei+U4-jOc6gKUegzPATQMGu4SDFQYSUFIkgJGEEQgDo7ozLWIGsGxjhCOBqJJiSe5HBWMHcLQ0S0jAZBumQF79EOnZ5OU6QtFoHQAOQAN66VAfHwAAjiAAB0lwrJCIgAL42RK+mov4dk5I5GR2fpZk0hA9KsQAPrgJjmHJDAgAAHGOcoebpXnRHSDIgAFMBBSFIAAMwgDK0VmTWCn1mQp6hAF6AMIMZB8dwACk5rqNw2XHoVODFaVZQVdVaW1dlcGhc1ZVtXKHV1TZ2lIv+cjuP6LBpBk43BMw-EgPkhTFDMpTlJUm4Fo6EGiVB4kHuc0kMKeZQVJgSn3Cp7xqfMmmDiAem6YF5g-gA7GIqBnmIZUxgALOZR2yfJMRWGMUqvVYdkjWNOi0MgmCYBu4Q4A08MaBxkjBLITggAAhEJmy7WC0FVp0R0XRwV1do9ZmA3JCnZYDa1nVDNhAA", "N4XyA"]],
    ["Generate a valid SIRET (french business identification number)", ["MQAgKgFglgziCmAPAhgWwA4Bt4hhA9gO4gHEAu+IA5vAHbwBOyZOAZsgNY4BEAygJIAlAKJhuIABSsGdAMYQQAIwCuMKPRgx4cKABM6ZKKyizmUfLRC1lqRYwCUAKFCQcAkWBCwQs-BnxauiD4rCBkhJToyAxkcBLuwgByIMi0QYn8AML2ADQIyPIgUTEkyHCpXrQsVAxQZACePhDwshwgulBUdY6OAG6MahYAXCDcAEzcjqhlHOpUQz0gIKDGDDBkRdEbIWHNIAme3sggTGl+IAAcALQdXRvWtow+yJjYQQmJjktXuPDYshQGAsliCQAArGAWKJkCAjPhCJKTUHTGAcGDA0EgH7IXS6OFIzE-GQ0RBwgDaAAYrgBOAC6AG8LgBfSZfZa-XxpTYlHYwtwIw7lE6pXTnAAsN06dSsNjsDGer3g6SybJ+Wn+gIxoIhUOYsNGGUyBKWKLRWpB2Nx+LZoKJ8BJ5KpdPpYpZi3ZfP2Aq8cF8-kCwVCno+OTZK3wryISqUjWOaloVGwTRabVu0okmGUEEsLyo+FqMNQ9jClE5ZAYEcq1QLjV5exDYZArAjmCjQUUjU9htDS3DkcI0Y7KRAmDKG3jiZw8hT7SlGwzWZzmDzBYgRZLPgs5cr6mrdVrQb2hAISYOPaxvw1+fNSx1tGh+vhHmNIFN6JtFpSVtG9PpADoPiZN1MQvTNsxGekmQ-b4wngDBRxYOFfwAgUgOQw0gJfGCwNoCCoKAA", "N4XyA"]]
]);

let examples_anonymization = new Map([
    ["Anonymize a dataset by suppression", ["MQAgggdg9hCeC2BLAXgQwC6JiRBnE6AFgKYgAOATlAMbG75QBmBFqEujUFSEA5iKhAATDKlzF0AGgBQoRBAEgA7qlgFCGEBWKIhxCJkax5-QfKGIAbroCuqADbliFXNjw54ZKPUQAje8QAdLLgCsRiaipq6FAC1ISIxJakiOg4+DFaxPBQycKi6prUbCC+pDbiQgSxuvqGaq7wxDBB0tLJLlgQAFwgAEQAjH3S8GIA1ibdbSAgoAAqCfiMiC5po7hjWTnJ+I2ky8T2VUSaqNrCK8TUabUGiAcuIEyKZM6uENIzALQg4gHXXFwUxmIJAPwAVu8yBhCL0+stVgB9CCoJrDUHfECQmDQohw+xidDI1HEdEg9ZjYEY7TbYi9dAUGzEaazEALfTKUgQYjEY6xMogYraRg2RxKVKEHAQCzaa44PR3B64YIzUCQKUyq43BWGRIUSSCqDwXzyXnKCUCaBEZwgAAUZBs-kQ1CeFHIFCsGGIAEoZKrBSVxBQ8pltLd6i83jBgp9WQAxLggYgAD1RZACIHpJCl1CNKXwZmllzl4fueoNsdAZWKFVI1pACV4JFwNwguaaICQTbSAvrAGEAKIAeQA9AmbNKbc8iO529C4CrWQB1DRpJSkYoKIQ1fDadOoWjqUiWBxM0pqQSsPjEP2sthVUZjUi4Gznevxbwcq+8OjmoREQVtC9EBeCoGwyHwcUAM0AJCRAABpJwXBgXBb1AGd8Hcd8YFoMg0meeCvjYGAEFSWAY39dkQG5ZM1nGc17EcPcCUPet5HbUh7xAAkW0RXwYAqEAT3sJkCy0Ngf0XH4-i1QEqVBCEoRhOF2LzMkFKxJS8X6HiiX4iAKnUztxnkkFvzpEAAFYAAZbOsllQDjeQHHsWADXXaieT5cgzkwFy1D0BkoGibMTQoIhEREdBj1PVp-QASWYa1zncGBXMUXATAzV5kIUfiKAUeQQAAZgGEdrMskcBgATgAdgANilDLRkYw1PDYNzK3UWcSgFfQoBsJtqiyUsjEUWUuCqIr62IuAkGQM0orECRF3mRZPNo4yNgYxw0rUJ9iDII8QFgcI3WwMKIqizjpXyaLESYREEjfWJKCSOpuuVWNpMOWSXFMzFsQgXFYX6S7CEir0jMBrTQb6a6HsYJ7LiMikAbu4gAAUziDdGZnkB10ATbgMDhayACZyrK8nbPq6GQQG9BCeJ0Z0DhGnrLp6QgA&", "N4KABGBEBmCWBOBnALgfQHYEMC2BTSAXFAFID2AFupADThQA2mKGO+RkAIqfrRJAEYJk5VABNMyNlADMARgD0ABgCs82QE4A7ADYadSLmwAHeqQCeueISgAZTAHcL8AOSIwAYVLHM6M3r7ikqik0KjkCFKQigqK0vIATIqJ-lCw6ADGXlKyyop5irwMTGj8pOgArojWiXkgAL5AA"]],
    ["Anonymize a dataset by randomisation", ["MQAgggdg9hCeC2BLAXgQwC6JiRBnE6AFgKYgAOATlAMbG75QBmBFqEujUFSEA5iKhAATDKlzF0AGgBQoRBAEgA7qlgFCGEBWKIhxCJkax5-QfKGIAbroCuqADbliFXNjw54ZKPUQAje8QAdLLgCsRiaipq6FAC1ISIxJakiOg4+DFaxGT2qLTCouqa1GwgvqQ24kIEsbr6hmqu8MQwQSEA8skUShSpJgXogkqphCDUUJ4B6MT2aqwQQhMglg42dCAaGSQguKjNO8TUNr3oamI7NmSUdLj9IoMyoL42aakgANbE2fiCCbyjlVQvkQ9lSak4FHUpHuYgkwWkXVuMAAXCAAEQARjR0ngYneJmR0mkIBAoAAKgl8IxEC40rjcO8sjk8qQmqRqTNqkRNKhtMIaYdXnoDIgOS4QExFGRnK4IMSQABaA4BagxFyEkmaxUgABWsrIGEIqLR1NpAH0IHtiNitSSlXqYAaiMbcrh0BarTbNfT3hrbfNFvAAMKEKCIWgASQgAFVesayIh4FBkQB6FOW5oAMQAStj5eSSAolKQIF8ubFymNecRGDZHMMiDgFgLVThhYZEi5giTQJAmxZtK26iKxZIxhNgaXqg3RmwoERnCAABRkGz+cMSyGUKwYYgAShkParCnEFGSNSyw4aUplMHhR8zXBAxAAHnscqRUQum+N9u4zM2g5CvUoqdmO+ZlIcqCVKQ35-CQbo-hMpBIP8aSVt+QYAKLtCmj42Asi6SkQ7i-gacD3qSIAUiWr50niyggo42jMvk37yL+pBsNUrrur4MCVMsqx0N22riCqaq4H6Wr2vqhrGhxyFejJupyc66K8Wa-EQJUykgD60magGExRughlajwqIAKwAAx2fKtr6agL6ohidnuUSD7yA4sxjsWIClsQ5ZMrkbHbMCFBEGa9xcQsAzEGaTBmgkfIrPYay4KJoARswC58u4MCzIotx8AETguNg-EUAo8ggAAzBiKY2VZKYYgAnAA7AAbE2xW4vYjhkWwsCHlRJH4CUCiVvoUA2P8F7aFeopnFk4wUNUtXfnOcBIMgQUDLC6BZdRlIBXRTkMoxA0hSyULiEJ6XrL4czcUsMWZfKSriYKXBSQ5mqyY68nohFUUxXpdqqUD6lojFiWMMlAp6T6f2OUqAYACK7uZtoxQAsvIxrtR1NkKjZGJkxiZJ2ciNN2QAWhDuO7njznGgATDZnOU5T1M2bT-MM0zSoxQACryp441qs3oKu6CPtwGAc3ZXU81z2JAA&", "N4KABGBEBmCWBOBnALgfQHYEMC2BTSAXFAFID2AFupADThQA2mKGO+RkAIqfrRJAEYJk5VABNMyNlACMATgDsANgC0ABgCsygMzSadSLmwAHeqQCeueISgAZTAHcL8AOSIwAYVLHM6M3r7ikqik0KjkCFKQAEyqMWpaarq8ULDoAMZeUtLqqrmqyZCMzPyk6ACuiNYxuSAAvkA"]],
    ["Anonymize a technical ID (like a plate number)", ["MQAgggdg9hCeC2BLAXgQwC6JiKAzE6ApgMYAWEixqANiIgCaESa6KEBOAziABQAO1DIRAQArvABGHTgBoQnPOgDuqdsM4dENEAGtCsWSD6kYwsZOlyAdDYCUAKFAAVUsPTtUETsSiMQ8VE4dECVEalopEEJAtnYCKBBRDQJSTxThNQBzQgAPf0Dg9ATPGAQUNxJySm0GJhZYq3t7ADdpLAgALhAAIgBGbvsAoMQITI6mkBBnVwIPLx8-IeDQ8JA1AVRiYRpaMlVNoi46CHpEZoZRHdh7SdAAIVgQRlxUUWp0EHHbkABaEABVAAKgIAogAlADCYAAyiCQNRCOhDtwVhEMoQNlt6CAJI85vQoPAAcDwVDYfDEciblNfvCoEoOFRkgikdIQmE0WsMYIsTi8Z4CUTqPTGYFhCyqd8-qdMoh0CiOTj0ZjCNjcWsBYSnohZfLqaA-jscOhXHE9h5iMj2atIqJmFBRGRVdS-hoEZaoFwvpMfX8AFYKCB8DCkLrdDZEAD65ik7AGPt9IADMGDJrDGnYWmo8Z9S29Cfcnm8vkIXQA3gBfCY0lzCc0HDggYiCTgabhUCBKpswViZURqbFFJupUZuGbPV7vJWpc6el3yQjuorsfM+5NBkNh1w5SPNGiiQg5yZ56kFubFxirhPNwJtq+JkZ8UToMMABl6ACYAMwAFgArAAbAA7AAHAAnGAdwQgAIiCABiR4Jj6DroE+L49O+37-sB4GQTB8EDPqIAuIg3BLE2aQ0AoXZJG4CScKIfBqK2TwYKg85uiQy73uuqahj0VDsPQ0biLGiEnkhsxFgspankhN6toQnD3pMfyPs+b6fr+gGgWBiGSShaFhgAVAMQA&", "N4KABGBEAOA2CGAXApgfQHYFcC2AjZATpAFxQCCAQgLQCMATAMxUDCAIpADThQDOhAlvFgko9BgBYArADZajcXIlVKbTt0gALZAA9UANyGZkIyGRoAOAKKopFMqgBiFyajYB2N2oiQAxvAIAJhg4+ESkkGJSspGKCirsIAC+QA"]],
]);

let examples_pseudonymization = new Map([
    ["Add noise to data", ["MQAgCgzgpgrgJgewHYE8C2BLAXgQwC4bIgYQh4AWUIADgE4IDGUEpCAZmbTkhGwrZiQBzEDhBx8OaHmJJRINNyRRaIAFCgK+MpXGSQtKABt8GYWQTyzcDADcMcGDiM0VEIg24gklo8iEqIABGUBqiLBhCynAgAO4YFAgwMmwwtBSBZnwCpsgAdGEAqkhGGADWVNzI6Ni5SAA0sjTQ8NWYuAREYhJ4UlAyJCAQBEYuhiZ4UDF4lhkg-JFmznq90gWgACq6hlJEMyAw0M2wiKiYEHUgnnIhIABcYQC0OlRoUNxmIuwvK2KDGGhqPxekgZPsKlBqCAABRQPJCPKiWwqHABUi2ZwwZiNQQAekUAA9GkgYGgQqpvjZhmYGDIMUYsRA8syAJRPH4MBCUQxIJjzDhzHp-UgAoFcUEwuEIgyQkx8i5vED0rFxBLkH4KqjixBoJWYqCNHC2BAOK4IUYkQg8RrMvJs0DPDKGYikMTKKYWGUOKCgjBsFAHJCKCAVGLK5iS+GItjWeR0BBBIxQRQEBgyzm0GJmH4LIRLFxC6RssIASTkFEGUAJOEBSca4dI7umlluDD8RzY9F1c1z+b1DOY6xAADF+DpBrEqNcDkc5sGyqRtQg0AARNKXMeTWv4OFhLaDeNMFguq7OJMxABEibStE+F-mqgvODgNnMPhIUAvBTUyNoECtdwgBeACMF5qPOnwPGoIAgJs5AThgoxDPBbBgroqJUN8czUG4RAhHgk4+iAjwAEwPiAADUZFLmgRgoNBxFDMYUC0vwDwwRxIAAFbuEg1D4OQgFPgEYGcfOEDsZxjFbtQExQEJADeCnPjEeQYTC2pljIpEgCRLIAL76aJUnPJ2y48cgQkYfeoCGJySC-jIQpkCgOGemYkwBLQagMXBCFIRAKFoVQPSYRwQhcI4lz4YRciPAAzAADIl5EUUlKUSCgi7cDqdEMc80BJqxtCSZxFl8QJQkRc+TiTHAAD6yDGTB4mlRxzyhWAOB-iobVSbI1DJKOOR4EJJHJQAbI8iXAdNJHNe1Bg5cua5cJ0SB9VJghCY8YDpSuC39YSQkUXtyUHQxnEdTuXU9SVl39fMySDXgw0pmNk3TbNiXzWoQA&", "N4KABGBECGDmCmkBcYCMBOANOKsBO0AJgK7QAu8hA+gPYB2yUG66AtAAwBsrAzO5CAC+QA"]],
    ["Preserve coherence with non-reversible functions", ["MQAgCgzgpgrgJgewHYE8C2BLAXgQwC4bIgYQh4AWUIADgE4IDGUEpCAZmbTkhGwrZiQBzEDhBx8OaHmJJRINNyRRaIAFCgK+MpXGSQtKABt8GYWQTyzcDADcMcGDiM0VEIg24gklo8iEqIABGUBqiLBhCynAgAO4YFAgwMmwwtBSBZnwCpsgAdGEAqkhGGADWVNzI6Ni5SAA0sjTQ8NWYuAREYhJ4UlAyJCAQBEYuhiZ4UDF4lhkg-JFmznq90gWgACq6hlJEMyAw0M2wiKiYEHUgnnIhIABcYQC0OlRoUNxmIuwvK2KDGGhqPxekgZPsKlBqCAABRQPJCPKiWwqHABUi2ZwwZiNQQAekUAA9GkgYGgQqpvjZhmYGDIMUYsRA8syAJRPH4MBCUQxIJjzDhzHp-UgAoFcUEwuEIgyQkx8i5vED0rFxBLkH4KqjixBoJWYqCNHC2BAOK4IUYkQg8RrMvJs0DPDKGYikMTKKYWGUOKCgjBsFAHJCKCAVGLK5iS+GItjWeR0BBBIxQRQEBgyzm0GJmH4LIRLFxC6RssIASTkFEGUAJOEBSca4ddzsM1DlHs53J9eCMAeh2e6-RUghIqYUShULPqYVi5AwDHVb24rpAeeRchgQakod+qtGoiMsRwKFIt2DW-iFA1Nao4cnoCCyXmJQDPiQj0MyNoEAwiaoqV5nR4HcXFuQ4plvOIZznBR3mtVVEgfKtJloJBliFRpDEeb1fRjTwAJ3YIqFFBAIh-Ao1A-L9kDuEAACIAEYaLUU9PgeNQQBAUAAEEQDYKBYlHENG0IpAO1BbtxAHAQzGHWdGlKCoQHIKR1X4RTlLLQpaAwNiQGeaAk1pfgHnYkyQAAK3cJBqHwchqJomNPzwAB9FC3kY0zTwgYzTN0tSIHIDStLs6gAQQO5cVxVyoAAMQAJXc3z9KgQzaG80yLOQayKDskxhhcq8EvYzy0pM54lP8wKMGC0LwtxCA0iiuLGJ00BotUhAnQEsoIEaDIA2uAilORCSkKHakGHE891TmHA4BsPDvjmAADaApiWmgcC4N4kJ0vTjGSmZUp09LLKy2zaKCDB0nIJyeigQqhigKYStMmNjDgOzLuu278HujiQDa1QoA-AMAG9QbyL6KB+yYAF9Yb1BkDR+ZjzHiXcVyoOZNRleqjDwY6uq8wn2OebUABFfpeny7oAWTMOy6IATgAdgABkeNm6M5uiNjZtm7n5wW2YALQenyVigWmcAJRmmaZjmuZ5vmBaF-mxZJ3y7rATboCOiXTKSPBqGSQGUzsgAmfmADYec5i3GKAA&", "N4KABGBEBmCWBOBnALgfQHYEMC2BTSAXFAFID2AFupADThQA2mKGO+RkAIqfrRJAEYJk5VABNMyNlACMATgDsANgC0ABgCsygMzTIIAL5A"]],
    ["Preserve coherence and enable reversibility with a cache", ["MQAgCgzgpgrgJgewHYE8C2BLAXgQwC4bIgYQh4AWUIADgE4IDGUEpCAZmbTkhGwrZiQBzEDhBx8OaHmJJRINNyRRaIAFCgK+MpXGSQtKABt8GYWQTyzcDADcMcGDiM0VEIg24gklo8iEqIABGUBqiLBhCynAgAO4YFAgwMmwwtBSBZnwCpsgAdGEAqkhGGADWVNzI6Ni5SAA0sjTQ8NWYuAREYhJ4UlAyJCAQBEYuhiZ4UDF4lhkg-JFmznq90gWgACq6hlJEMyAw0M2wiKiYEHUgnnIhIABcYQC0OlRoUNxmIuwvK2KDGGhqPxekgZPsKlBqCAABRQPJCPKiWwqHABUi2ZwwZiNQQAekUAA9GkgYGgQqpvjZhmYGDIMUYsRA8syAJRPH4MBCUQxIJjzDhzHp-UgAoFcUEwuEIgyQkx8i5vED0rFxBLkH4KqjixBoJWYqCNHC2BAOK4IUYkQg8RrMvJs0DPDKGYikMTKKYWGUOKCgjBsFAHJCKCAVGLK5iS+GItjWeR0BBBIxQRQEBgyzm0GJmH4LIRLFxC6RssIASTkFEGUAJOEBSca4ddzsM1DlHs53J9eCMAeh2bmmpAVagDGSVpZ9TCsXIGAY6re3FdIDzyLkMCDUlDv1Vo1ERliOBQpFuwc38QoGprVHD6xAAHVKkZ3HFuGDLIoKkM0lQnVR4jvbj4ZoWhAVpNHMSR4NQyS-NIYSwlGIAAIKiEgMQAELBAg54nh6Z7qv2l56gyUAsjeGyWDgs4YFAyLEHgjSeLOEZ-i4tyHFMjQhJ47HeMgjyGMitAgYmVCpLynQ8FcXjxo4fKcsBVpMmEWwqFA2QGi8zoscEVADiSZKBJSJAEOJRGMmBug4TEhb9AUaiCSByB3CAABEACMLlqGEADClGUKQ7pwKQ+y3HAw4mIYcCNLEVAMDskwgAATFJTEQI02QgDGQl4AA+kghHcDEJjDHll4+X5EbXDp4iktQHp7LoMZJhAh6TLqeE-AABo8jyOICjyMbog1QAAcpeAC8TU+peeQAFbuCUnXzNQEnlalUk3FQfg4GFMQhOpCgbp88g2fRqrnnM3WPNtcADRVKWUGNbyTRgSb5W8c0LUYS0ICtVprf5G3VQJbgejgbCTKoN3HdCub5mZVAADyPAAfIdIYeg2ojOhA+7UHVcDjmEFakMuEZzECERBK9CQBvsINCd+uh4OKvD8CmoHQgEMhiHD+UuOGmX0LqcxWQjpFecNEAPCAmUYNlpVvDlw0y7LgYYAAjlizksyqoBrprWLjbrrzvJJYhi4LVXIN2oj4+8FK8hpT6AfJpSOXI2nkDg9tIGosvFbl71QMrFWq7LBta1AOu0Fi-tq6ADPQDHKpeSenwPPHzzQEmtL8OHavzcg1D4OQzkuVlJXB55asPdHrmV0Hl6h0xNdqye0vx7XzzaggaDeeQJpMGWhS0Bg5fUACCB3LiuLBwAYgASp5WdDMYw4zLQBey0XSAlxQ5eB4rUBt7LKuuUfwct5Qp-o2Une17LPeFX3A9D1AI9jxPU8z7iEBpAvZeaggA&", "N4KABGBEBmCWBOBnALgfQHYEMC2BTSAXFAFID2AFupADThQA2mKGO+RkAIqfiAL5A"]],
    ["Preserve coherence and enable reversibility with encryption", ["MQAgCgzgpgrgJgewHYE8C2BLAXgQwC4bIgYQh4AWUIADgE4IDGUEpCAZmbTkhGwrZiQBzEDhBx8OaHmJJRINNyRRaIAFCgK+MpXGSQtKABt8GYWQTyzcDADcMcGDiM0VEIg24gklo8iEqIABGUBqiLBhCynAgAO4YFAgwMmwwtBSBZnwCpsgAdGEAqkhGGADWVNzI6Ni5SAA0sjTQ8NWYuAREYhJ4UlAyJCAQBEYuhiZ4UDF4lhkg-JFmznq90gWgACq6hlJEMyAw0M2wiKiYEHUgnnIhIABcYQC0OlRoUNxmIuwvK2KDGGhqPxekgZPsKlBqCAABRQPJCPKiWwqHABUi2ZwwZiNQQAekUAA9GkgYGgQqpvjZhmYGDIMUYsRA8syAJRPH4MBCUQxIJjzDhzHp-UgAoFcUEwuEIgyQkx8i5vED0rFxBLkH4KqjixBoJWYqCNHC2BAOK4IUYkQg8RrMvJs0DPDKGYikMTKKYWGUOKCgjBsFAHJCKCAVGLK5iS+GItjWeR0BBBIxQRQEBgyzm0GJmH4LIRLFxC6RssIASTkFEGUAJOEBSca4ddzsM1DlHs53J9eCMAez3X6KkEJFTCiUKnqYVi5AwDHVb24rpAeeRchgQakod+qtGoiMsRwKFIt2DG-iFA1Nao4fHoCCAZ9DFoKGoBHMDdVZ7E0Af-RAFRQ6xALZBjeRIYiNE0Yk5C0ICtV0kBiJcIzmIEIiCDBSjwAN9kMZFaCOOY8HFXh+BTK0YUMR5vV9f0WQKNRcJg5A7hAAAiABGFi1GPT4HjUEAQGeaAk1pfgHn48SQAAK3cJBqHwchmJYqiCBjFROIk48xIkkA2DYNitO0gwcBsAlmLYgAGPjDN-KAUAAMXoNAAFEkFsRS7LstiAH0nIAOQAYQAJQATTADYSwAeV8ryAGknOCziwg8tiDmgBcYwJD0dRwbM4CgGMkASMjoVKYZ+R3FxnD8WI23IHAuFpNwmi8JI8GoZIQHvR9nw9YZaE+WiwnMtiACYAGYABYAFYADYAHYAA4AE4cCCBg8rYIQp0ksojDQHxqAARzwvAYFsWICRQLAAEEACF-IAEScuyAHEAAkSwAKRigAZABZXyIrAABFQKAGUNkKAA1AB1AANYKAC0wi2LVjIwAkaHqi9JlUPLccEKhJ2ndU5PScrBQQRRcvyswiqIeJtxwWknFGANbkOKYALs-hOurWsDXkLgTPKiytxceNHD5N9OVBHLCvMZBu3ERY8FIMJoTmMWYxOq46oa3HWAFXRUm3bKaYK+mkEGzYO2yQW5nDcrhYxhRDhkJMWHmVQoEO1mdC8GaRsRa6jHceQne+FLefMl1ZHpBw6P4zZLDyh8nxkMQnc58QoHT59mMIrErME4w85mWgDPE6TkDkihFOUv0MBULy0+6yY4HU8TjwgKv+OeYy4EUgBvYe8kb1TaAAXynruJOeXT9Ks6yjJMszLJX-i24zwvaGLzebPsxyXLc1jkp8gKQrCyLorihK1CAA&", "N4KABGBECWAmCmA7ALtAZteAnSAuKAHAEwDMBBAnACykkCsVlFkIAvkA"]],
]);

let examples_other = new Map([
    ["Preserve null and/or empty values", ["G4UwTgzglg9gdgLgAQCICMKBQBbAhhAayjgHMFNMkkBiJABzBAnFGRQCMAbXOAlJAO5ROnJHBgAXJLjp1OATyR5CxEknhiAriPVgkIbHQmKAFChQBKJMFydNTSkgC0SZpxABjCTDDkq-pAArCHg6XAkACzZOKAgJADobOxAsAOUCPwCkD3g4ngk2dJAAE1T-BiYWEDYuHj4KKloK5jBWVDhtTn4hHXEpGTlFdNV1OC0dJPsIRybGFraUAyN5buFRPulZBSV8IlJR-UNjJDNLa1spzCA&", "N4KABGBEA2CWDOAXSAuMBtcEzEgNwENoBXAU1TAAYBfAGiwl0JPLUksjoZ3yLIsid62Hs35oAdsWjRqWALohqQA"]],
    ["Apply a template on each value of arrays", ["G4UwTgzglg9gdgLgAQCICMKBQBbAhhAayjgHMFNMkkBaJCEAGxAGMAXGMcq7pAKwngAHXKwAWyFMFwMAriAhYeeQlx5JWIbIIYiQ1ELmbjKaqlA3YJU2SEWmkxACYgAHhKiOXd0xe26JAN4BAHTWckgAPkgygoLgAL7xSCJIgjDQrLBwSEHBHi6JWEA&", "N4KABGBEBuCGA2BXApgZ0gLjAbUrSANFAEaFQDGZkAJpALogC+QA"]],
    ["How to handle nested/complex structure", ["G4UwTgzglg9gdgLgAQCICMKBQEQgCbIAsATJgLYCGEA1lHAOYKZJIC0SOANiAMYAuMMExYikAYiR8AFlAhIADhWkKYdPnIFIKcLWDAUAnkhgAzBeAjwIzUUgBWluIunIUg+tqgAvJbDgQAOnkLKyxRShphWwkoPiQAdyhOTiQAIxAFKGC8SRhJKQyI2gZM4M46DLwQEzpYvzSQThh4m1F5LJAo23F8jMUwEDg4mFS7XjiACgokPBhKOgBKBKSU9K0eHhAIaFTuZeVpDJQAfRQkYAowKApdjLgKMhBW2zox-gAFS8G+V1Pn8KoxUY-1s7C440EXW6tgc8GcUlc90eYWhAMiINREhqjTwcjo0CqvXMkHgxlG4yQPG0DXWm22+BmUAG-E4Bgx0L4IDI8k4Sk6qAA3gLJLE9gEkRkAL6SlGosGNCFCdndWFOJQI1AQACuYAlstRSCKUINki5PL5riFIr4Yu1uoeUplyrYHAV-EhztEqvhri5FCS+tRRs9Igkh3OFE4WoyEAEAxy+KghPD-W+ZLecVktK2Oz20jAMC19CkqFO+W01A0eRTXyGSFe4z8IZYnO5vM5luFTXi4CQ4odSGlASt3d7ATtEsHkoAAlaAscArN5jppVggA&", "N4KABBYEQPYE4HMCGA7AlgLyQFzTFAzlAFxgDa4kVoVtkUAJjALZJoonQDGLADqgE8AdD2ZQANJTpUovAKZwC+IqQrT1YGhu1QUSZnM5QANnPxIJU7dKgEArnD0GjzNMdNxL1jVDms3RlBW3gC+kt7UwREQuvqGpFAAVjCG4dG0tg5O8dAMSABuaERp6fR+bMaBUdoh1bQAunVhdVo6TP4cCaL8KMIAZp4lPvKKypxq6a2l0NlGSMZsHEOlmY5xLgpcaApe0zHlAQlBpc2lUyuzCX3G8Giou3url9CmCHCoDA-Tvv6VR3XqWreRqAqyNEJAA"]],
    ["Temporary fields", ["G4UwTgzglg9gdgLgAQCICMKBQEQgCbIAsATJgLYCGEA1lHAOYKaZJIC0SOANiAMYAuMME1aikAKwjwADhX4ALZCn5gKcaCDj8A+gDMoILnixjKNEWKQU8eNirUatSgN7PVcPAEEu0+RQByAK5kSADMAL7hWCzsnIZ8gsIxYpIycoqo+obGyUhm1BZi1gSoACryIEjyUDaaSMAUXIGVUBBIyK4AdPbqBlp6BkaRWEA&", "N4XyA"]],
    //["Multiple mask and/or multiple selector syntax", ["", ""]],
    //["Seed parameter", ["", ""]],
    //["Caches", ["", ""]],
    //["Change date formats", ["", ""]],
    //["Parse raw JSON", ["", ""]],
    //["Generate sequences", ["", ""]],
]);

let examples_generation_a = document.getElementById("example-generation")
examples_generation.forEach((params, name) => {
    let link = examples_generation_a.cloneNode(true);
    link.innerText = name;
    link.onclick = () => { loadExample(params) }
    examples_generation_a.parentElement.appendChild(link)
});
examples_generation_a.remove()

let examples_anonymization_a = document.getElementById("example-anonymization")
examples_anonymization.forEach((params, name) => {
    let link = examples_anonymization_a.cloneNode(true);
    link.innerText = name;
    link.onclick = () => { loadExample(params) }
    examples_anonymization_a.parentElement.appendChild(link)
});
examples_anonymization_a.remove()

let examples_pseudonymization_a = document.getElementById("example-pseudonymization")
examples_pseudonymization.forEach((params, name) => {
    let link = examples_pseudonymization_a.cloneNode(true);
    link.innerText = name;
    link.onclick = () => { loadExample(params) }
    examples_pseudonymization_a.parentElement.appendChild(link)
});
examples_pseudonymization_a.remove()

let examples_other_a = document.getElementById("example-other")
examples_other.forEach((params, name) => {
    let link = examples_other_a.cloneNode(true);
    link.innerText = name;
    link.onclick = () => { loadExample(params) }
    examples_other_a.parentElement.appendChild(link)
});
examples_other_a.remove()

document.getElementById("reset-link").onclick = () => {
    editorYaml.setValue(masking);
    editorJson.setValue(input);
    resultJson.setValue("");
    autoPostData();
}

/* When the user clicks on the button, toggle between hiding and showing the dropdown content */
document.getElementById("dropdown-button").onclick = () => {
    document.getElementById("myDropdown").classList.toggle("show");
}

// Close the dropdown if the user clicks outside of it
window.onclick = function(e) {
    if (!e.target.matches('.dropbtn')) {
    var myDropdown = document.getElementById("myDropdown");
      if (myDropdown.classList.contains('show')) {
        myDropdown.classList.remove('show');
      }
    }
}

// Share an url link
document.getElementById("btnShareTest").onclick = () => {
    var dummy = document.createElement('input'),
    text = window.location.href;
    document.body.appendChild(dummy);
    dummy.value = text;
    dummy.select();
    document.execCommand('copy');
    document.body.removeChild(dummy);
    alert("URL Copied.");
}


// Download venom test yaml file
/**
 * @param {string} str  The string to be indented.
 * @param {number} numOfSpaces  The amount of spaces to place at the
 *     beginning of each line of the string.
 * @return {string}  The new string with each line beginning with the desired
 *     amount of spaces.
*/
function indent(str, numOfSpaces) {
    str = str.replace(/^(?=.)/gm, new Array(numOfSpaces + 1).join(' '));
    return str
}

document.getElementById("btnDownloadTest").onclick = () => {

    var masking = editorYaml.getValue();
    var output = resultJson.getValue();
    var input = editorJson.getValue();
    var url = window.location.href;

    var template = `name: "test generated  from pimoplay ${url}"
testcases:
- name: pimo play generated test (to change)
  steps:
    - script: rm -f masking.yml
    - script: |-
        cat > masking.yml <<EOF
${indent(masking, 8)}
        EOF
    - script: |-
        jq -c . > input.jsonl <<EOF
${indent(input, 8)}
        EOF
    - script: |-
        jq -c . > expected.jsonl <<EOF
${indent(output, 8)}
        EOF
    - script: |-
        < input.jsonl pimo > result.jsonl
      assertions:
        - result.code ShouldEqual 0
    - script: |-
        diff expected.jsonl result.jsonl
      assertions:
        - result.code ShouldEqual 0
        - result.systemout ShouldBeEmpty`;

    var encodedString = btoa(unescape(encodeURIComponent(template)));
    document.getElementById("aDownloadTest").href = "data:text/yaml;base64," + encodedString
}

///////////////////////////////////////////////////////////

async function postData() {
  const postData = {
      data: editorJson.getValue(),
      masking: editorYaml.getValue()
  }
  console.log(postData)

  // update URL for sharing
  var c = LZString.compressToEncodedURIComponent(postData.masking);
  var i = LZString.compressToEncodedURIComponent(postData.data);
  window.history.replaceState(null, null, `${location.protocol}//${location.host}${location.pathname}?c=${c}&i=${i}`);

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

      resultJson.setValue(JSON.stringify(data, null, 2))
      document.getElementById('result-error').innerText = ""
  } catch (err) {
      console.log(err)
      document.getElementById('result-error').innerText = err
  } finally {
    document.getElementById('refresh-spinner').style.display = 'none';
    document.getElementById('refresh-button').style.display = 'inline';
  }
}

function debounce(func, timeout = 300){
    let timer;
    return (...args) => {
        document.getElementById('refresh-spinner').style.display = 'inline';
        document.getElementById('refresh-button').style.display = 'none';
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
document.getElementById('refresh-button').onclick = autoPostData;
autoPostData();


