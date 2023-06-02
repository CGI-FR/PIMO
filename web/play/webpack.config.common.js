import HtmlWebPackPlugin from 'html-webpack-plugin';
import MonacoWebpackPlugin from 'monaco-editor-webpack-plugin';


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
        devServer: {
            proxy: {
                '/play': 'http://localhost:3010',
                '/flow': 'http://localhost:3010',
            },
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
            })
        ],
    }
};
