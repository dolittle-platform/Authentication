// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import 'reflect-metadata';

import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router } from 'react-router-dom';

import { Layout } from './layouts/Layout';

import './styles/theme';
import './index.scss';

export default function App(this: any) {
    return (
        <Router>
            <Layout />
        </Router>
    );
}

ReactDOM.render(<App />, document.getElementById('root'));
