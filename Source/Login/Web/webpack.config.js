// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

const output = require('./webpack/output');
const optimization = require('./webpack/optimization');
const resolve = require('./webpack/resolve');
const rules = require('./webpack/rules');
const plugins = require('./webpack/plugins');
const devServer = require('./webpack/devServer');

module.exports = (env, argv) => {
    const production = argv.mode === 'production';
    const basePath = '/.auth/assets/';
    const title = 'Dolittle Studio'

    const config = {
        entry: './index.tsx',
        target: 'web',
        output: output(env, argv, basePath),
        optimization: optimization,
        resolve: resolve,
        module: {
            rules: rules
        },
        plugins: plugins(basePath, title),
        devtool: production ? '' : 'inline-source-map',
        devServer: devServer(basePath)
    };

    return config;
};
