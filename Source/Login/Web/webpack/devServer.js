// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

module.exports = (basePath) => {
    return {
        historyApiFallback: {
            index: `${basePath}index.html`
        },
        host: '0.0.0.0',
        port: 8091,
        publicPath: basePath,
        contentBase: process.cwd(),
        disableHostCheck: true
    };
};