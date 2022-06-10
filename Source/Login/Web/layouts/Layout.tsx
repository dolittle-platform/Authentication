// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Routes, Route } from 'react-router-dom';

import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';

import { configuration } from '../Configuration';
import { ErrorBoundary } from '../error/ErrorBoundary';
import { Error } from '../error/Error'
import { Theme } from '../styles/Theme';
import { SelectProvider } from '../select-provider/SelectProvider'
import { SelectTenant } from '../select-tenant/SelectTenant'
import { LoggedOut } from '../logged-out/LoggedOut'
import { Headline } from './Headline';

export const Layout = (): JSX.Element => {
    return (
        <>
            {/* { configuration.showDolittleHeadline && <Headline /> } */}
            <Grid container>
                <Grid item xs={0} sm={6} />
                <Grid item xs={12} sm={6}>
                    <Box css={{ 
                        paddingTop: '16vh',
                        textAlign: 'center',
                        [Theme.breakpoints.up('sm')]: {
                            paddingTop: '17vh',
                            maxWidth: '30em'
                        },
                    }}>
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
                            <Route path="/.auth/logged-out" element={
                                <ErrorBoundary>
                                    <LoggedOut />
                                </ErrorBoundary>
                            }/>
                            <Route path="/.auth/error" element={
                                <Error />
                            }/>
                        </Routes>
                    </Box>
                </Grid>
            </Grid>
        </>
    );
};

export default Layout;
