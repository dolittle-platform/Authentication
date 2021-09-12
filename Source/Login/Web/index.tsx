// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React, { lazy, Suspense } from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router } from 'react-router-dom';

import { CacheProvider } from 'rest-hooks';
import { ThemeProvider } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline'

import { Theme } from './styles/Theme';
const Layout = lazy(() => import('./layouts/Layout'));

export default function App(this: any) {
    return (
        <CacheProvider>
            <ThemeProvider theme={Theme}>
                <CssBaseline/>
                <Router>
                    <Suspense fallback={<></>}>
                        <Layout/>
                    </Suspense>
                </Router>
            </ThemeProvider>
        </CacheProvider>
    );
}

ReactDOM.render(<App />, document.getElementById('root'));
