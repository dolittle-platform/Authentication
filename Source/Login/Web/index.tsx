// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { lazy, Suspense } from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter as Router } from 'react-router-dom';

import { CacheProvider } from 'rest-hooks';
import { CssBaseline } from '@mui/material'
import { ThemeProvider } from '@mui/material/styles';

import { Theme } from './styles/Theme';
const Layout = lazy(() => import('./layouts/Layout'));

export default function App(this: any) {
    return (
        <CacheProvider>
            <ThemeProvider theme={Theme}>
                <CssBaseline />
                <Router>
                    <Suspense fallback={<></>}>
                        <Layout />
                    </Suspense>
                </Router>
            </ThemeProvider>
        </CacheProvider>
    );
};

ReactDOM
    .createRoot(document.getElementById('root')!)
    .render(<App />);
