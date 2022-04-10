// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Routes, Route } from 'react-router-dom';

import Paper from '@mui/material/Paper';

import { configuration } from '../Configuration';
import { ErrorBoundary } from '../error/ErrorBoundary';
import { Error } from '../error/Error'
import { SelectProvider } from '../select-provider/SelectProvider'
import { SelectTenant } from '../select-tenant/SelectTenant'
import { Background } from './Background';
import { Headline } from './Headline';

export const Layout = (): JSX.Element => {
    return (
        <>
            <Background />
            { configuration.showDolittleHeadline && <Headline /> }
            <Paper
                elevation={24}
                square={true}
                css={{ position: 'absolute', right: '0', height: '100%', width: '45%', maxWidth: '650px', backgroundColor: '#191A21' }}>
                <Routes>
                    <Route path="/.auth/select-provider" element={
                        <ErrorBoundary>
                            <SelectProvider />
                        </ErrorBoundary>
                    }/>
                    <Route path="/.auth/select-tenant" element={
                        <ErrorBoundary>
                            <SelectTenant />
                        </ErrorBoundary>
                    }/>
                    <Route path="/.auth/error" element={
                        <Error />
                    }/>
                </Routes>
            </Paper>
        </>
    );
};

export default Layout;
