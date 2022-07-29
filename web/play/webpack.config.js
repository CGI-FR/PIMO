import HtmlWebPackPlugin from 'html-webpack-plugin';
import MonacoWebpackPlugin from 'monaco-editor-webpack-plugin';

export default {

    module: {
        rules: [
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
        poll: 1000, // Check for changes every second
    },
    devServer: {
        proxy: {
            '/play': 'http://localhost:3010',
        },
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
    ],
};
