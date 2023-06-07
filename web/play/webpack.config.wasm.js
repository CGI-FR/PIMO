import { merge } from 'webpack-merge';

import common from "./webpack.config.common.js";

import AddAssetHtmlPlugin from 'add-asset-html-webpack-plugin';
import CopyWebpackPlugin from 'copy-webpack-plugin';

import path from 'path';

export default merge(common(false), {
    mode: "development",
    module: {
        rules: [
            {
                test: /index\.es$/,
                loader: 'string-replace-loader',
                options: {
                    search: '{{ version }}',
                    replace: 'master',
                    flags: 'g'
                }
            }
        ]
    },
    plugins: [
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
    ]
});
