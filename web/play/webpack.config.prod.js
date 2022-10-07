import { merge } from 'webpack-merge';

import common from "./webpack.config.common.js";

export default merge(common(true), {
  mode: "production",

});
