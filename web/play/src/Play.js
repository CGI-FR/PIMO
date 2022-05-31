import React, { useRef, useState } from 'react'
import './styles/play.css'

const example = {
    json: "{\n \"age\": 35,\n \"name\": \"Benjamin\", \n \"surname\": \"Toto\",\n \"address\":\n   {\n    \"town\":\"Nantes\"\n   },\n \"mail\": \"Benjamin.Toto@hotmail.fr\"}\n",
    yaml: "version: \"1\"\n" +
        "seed: 42\n" +
        "masking:\n" +
        "  - selector:\n" +
        "      jsonpath: \"customer.phone\"\n" +
        "    mask:\n" +
        "      regex: \"0[1-7]( ([0-9]){2}){4}\"\n" +
        "  - selector:\n" +
        "      jsonpath: \"name\"\n" +
        "    mask:\n" +
        "      constant: \"Toto\"\n" +
        "  - selector:\n" +
        "      jsonpath: \"name2\"\n" +
        "    mask:\n" +
        "      randomChoice:\n" +
        "       - \"Mickael\"\n" +
        "       - \"Mathieu\"\n" +
        "       - \"Marcelle\"\n" +
        "  - selector:\n" +
        "      jsonpath: \"age\"\n" +
        "    mask:\n" +
        "      randomInt:\n" +
        "        min: 25\n" +
        "        max: 32\n" +
        "  - selector:\n" +
        "      jsonpath: \"name3\"\n" +
        "    mask:\n" +
        "      command: \"echo Dorothy\"\n" +
        "  - selector:\n" +
        "      jsonpath: \"surname\"\n" +
        "    mask:\n" +
        "      weightedChoice:\n" +
        "        - choice: \"Dupont\"\n" +
        "          weight: 9\n" +
        "        - choice: \"Dupond\"\n" +
        "          weight: 1\n" +
        "  - selector:\n" +
        "      jsonpath: \"address.town\"\n" +
        "    mask:\n" +
        "      hash:\n" +
        "        - \"Emerald City\"\n" +
        "        - \"Ruby City\"\n" +
        "        - \"Sapphire City\"\n" +
        "  - selector:\n" +
        "      jsonpath: \"date\"\n" +
        "    mask:\n" +
        "      randDate:\n" +
        "        dateMin: \"1970-01-01T00:00:00Z\"\n" +
        "        dateMax: \"2020-01-01T00:00:00Z\"\n" +
        "  - selector:\n" +
        "      jsonpath: \"name4\"\n" +
        "    mask:\n" +
        "      replacement: \"name\"\n" +
        "  - selector:\n" +
        "      jsonpath: \"mail\"\n" +
        "    mask:\n" +
        "      template: \"{{.surname}}.{{.name}}@gmail.com\"\n" +
        "  - selector:\n" +
        "      jsonpath: \"last_contact\"\n" +
        "    mask:\n" +
        "      duration: \"-P60D\""
}


export default function Play() {
    const baseURL = "http://localhost:3010/play"

    const jsonInput = useRef(null)
    const yamlInput = useRef(null)

    const [getResult, setGetResult] = useState([])
    const [postResult, setPostResult] = useState()


    const clearPostOutput = () => {
        setPostResult(null)
        jsonInput.current.value = null
        yamlInput.current.value = null
    }

    async function postData() {
        const postData = {
            data: jsonInput.current.value,
            masking: yamlInput.current.value
        }

        if (postData.data.length === 0 || postData.masking.length === 0) {
            postData.data = example.json
            postData.masking = example.yaml
        }

        try {
            const res = await fetch(`${baseURL}`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(postData)
            })

            if (!res.ok) {
                const message = `An error has occurred: ${res.status} - ${res.statusText}`
                throw new Error(message)
            }

            const data = await res.json()

            const result = {
                status: res.status + "-" + res.statusText,
                headers: {
                    "Content-Type": res.headers.get("Content-Type"),
                    "Content-Length": res.headers.get("Content-Length")
                },
                data: Object.assign([], data)
            }

            setPostResult(result)
            setGetResult(result.data)
        } catch (err) {
            setGetResult(err.message)
        }
    }

    const result = JSON.stringify(Object.assign({}, getResult))

    return (
        <div className="page-wrap">
            <h2 className="text-center">Pimo Play!</h2>
            <div className="row">
                <div className="col-sm12 col-md-6">
                    <div className="form-group text-center">
                        <p>JSON</p>
                        <textarea name="data" rows="12" className="form-control" ref={jsonInput} placeholder={example.json} />
                    </div><br />
                </div>
                <div className="col-sm12 col-md-6">
                    <div className="form-group text-center">
                        <p>YAML</p>
                        <textarea name="masking" className="form-control" ref={yamlInput} rows="12" placeholder={example.yaml} />
                    </div>
                </div>
            </div>
            <div className="row d-flex justify-content-center">
                <button className="btn btn-sm btn-primary col-md-2" onClick={postData}>Post Json</button>
                <button className="btn btn-sm btn-warning col-md-2" onClick={clearPostOutput}>Reset</button>
            </div>
            <div className="form-group"><br />
                <p className="text-center">RESULT</p>
            </div>
            {getResult && <div className="alert alert-secondary mt-2" role="alert">
                <p>{result}</p>
            </div>}
        </div>
    )
}
