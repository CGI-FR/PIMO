import { merge } from 'webpack-merge';

import common from "./webpack.config.common.js";
import path from 'path';

import AddAssetHtmlPlugin from 'add-asset-html-webpack-plugin';
import CopyWebpackPlugin from 'copy-webpack-plugin';

export default merge(common(true), {
    mode: "production",
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
