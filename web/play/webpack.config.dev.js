import { merge } from 'webpack-merge';

import common from "./webpack.config.common.js";

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
  }
});
