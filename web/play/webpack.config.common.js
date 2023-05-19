import HtmlWebPackPlugin from 'html-webpack-plugin';
import MonacoWebpackPlugin from 'monaco-editor-webpack-plugin';
import AddAssetHtmlPlugin from 'add-asset-html-webpack-plugin';
import CopyWebpackPlugin from 'copy-webpack-plugin';

import path from 'path';

export default (isProd) => {
    return {
        resolve: {
            modules: ["./src", "node_modules"],
            extensions: [".js", ".es", ".elm", ".scss", ".png"]
        },
        module: {
            rules: [
                {
                    test: /\.elm$/,
                    exclude: [/elm-stuff/, /node_modules/],
                    use: {
                        loader: "elm-webpack-loader",
                        options: {
                            debug: !isProd
                        }
                    }
                },
                {
                    test: /\.css$/,
                    use: ['style-loader', 'css-loader', 'postcss-loader'],
                },
                {
                    test: /\.ttf$/,
                    type: 'asset',
                },
            ],
        },
        watchOptions: {
            ignored: /node_modules/,
            aggregateTimeout: 200,
            poll: 1000
        },
        plugins: [
            new HtmlWebPackPlugin(),
            new MonacoWebpackPlugin({
                languages: ['yaml', 'json'],
                customLanguages: [
                    {
                        label: 'yaml',
                        entry: 'monaco-yaml',
                        worker: {
                            id: 'monaco-yaml/yamlWorker',
                            entry: 'monaco-yaml/yaml.worker',
                        },
                    },
                ],
            }),
            new CopyWebpackPlugin({
                patterns: [
                    {
                        from: '../../bin/pimo.wasm',
                        to: path.join('pimo.wasm')
                    }
                ]
            }),
            new AddAssetHtmlPlugin({
                publicPath: ``,
                outputPath: ``,
                filepath: path.join('src', `wasm_exec.js`),
            })
        ],
    }
};
