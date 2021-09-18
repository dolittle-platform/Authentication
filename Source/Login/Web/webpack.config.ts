// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { CleanWebpackPlugin } from 'clean-webpack-plugin';
import HtmlWebpackPlugin from 'html-webpack-plugin'
import path from 'path'
import TerserWebpackPlugin from 'terser-webpack-plugin';
import { Configuration } from 'webpack';

type WebpackArguments = {
    mode: 'none' | 'production' | 'development',
}

const globalFrontendConfiguration = (args: WebpackArguments): string => {
    if (args.mode === 'production') {
        return `
            window.configuration = {
                showDolittleHeadline: {{ .ShowDolittleHeadline }},
                applicationName: "{{ .ApplicationName }}",
                supportEmail: "{{ .SupportEmail }}",
            };
        `.split('\n').map(_ => _.trim()).join('');
    } else {
        return `
            window.configuration = {
                showDolittleHeadline: true,
                applicationName: "Dolittle Studio",
                supportEmail: "support@dolittle.com",
            };
        `.trim();
    }
}

export default (_env: any, args: WebpackArguments): Configuration => {

    return {
        entry: './index.tsx',
        output: {
            path: path.join(__dirname, 'wwwroot'),
            filename: '[name].[chunkhash].bundle.js',
            chunkFilename: '[name].[chunkhash].chunk.js',
        },
        optimization: {
            runtimeChunk: 'single',
            minimize: true,
            minimizer: [
                new TerserWebpackPlugin({
                    extractComments: false,
                }),
            ],
        },
        resolve: {
            extensions: [ '.tsx', '.ts', '.js' ],
        },
        module: {
            rules: [
                {
                    test: /\.tsx?$/,
                    exclude: /node_modules/,
                    loader: 'ts-loader',
                },
                {
                    test: /\.woff2?$/,
                    exclude: /node_modules/,
                    loader: 'file-loader',
                },
                {
                    test: /\.(png|jpg)$/,
                    exclude: /node_modules/,
                    loader: 'file-loader',
                },
                {
                    test: /\.svg$/,
                    exclude: /node_modules/,
                    loader: '@svgr/webpack',
                },
            ],
        },
        plugins: [
            new HtmlWebpackPlugin({
                template: 'index.html',
                templateParameters: {
                    configuration: globalFrontendConfiguration(args),
                },
            }),
            new CleanWebpackPlugin(),
        ],
        devServer: {
            port: 8091,
            historyApiFallback: {
                index: '/.auth/assets/index.html',
            },
            hot: true,
            devMiddleware: {
                publicPath: '/.auth/assets/',
            },
        },
    };
};
