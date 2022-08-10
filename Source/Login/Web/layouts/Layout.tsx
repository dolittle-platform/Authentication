// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Routes, Route } from 'react-router-dom';

import { Box } from '@mui/material';

import { BgLogo } from '../styles/logos';
import Logo from '../styles/images/logo.svg';
import { ErrorBoundary } from '../error/ErrorBoundary';
import { Error } from '../error/Error'
import { SelectProvider } from '../select-provider/SelectProvider'
import { SelectTenant } from '../select-tenant/SelectTenant'
import { LoggedOut } from '../logged-out/LoggedOut'

const styles = {
    backgroundLogoContainer: {
        maxInlineSize: '793px',
        minBlockSize: '100vh'
    },
    mainContainer: {
        inlineSize: '100%',
        maxInlineSize: '541px',
        position: 'absolute',
        top: '20%',
        right: '20%',
        transform: 'translate(20%, 0%)',
        textAlign: 'center',
        padding: '20px'
    }
};

export const Layout = (): JSX.Element => (
    <>
        <Box sx={styles.backgroundLogoContainer}>
            <BgLogo />
        </Box>
        <Box sx={styles.mainContainer}>
            <Routes>
                <Route path="/.auth/select-provider" element={
                    <ErrorBoundary>
                        <SelectProvider />
                    </ErrorBoundary>
                } />
                <Route path="/.auth/select-tenant" element={
                    <ErrorBoundary>
                        <SelectTenant />
                    </ErrorBoundary>
                } />
                <Route path="/.auth/logged-out" element={
                    <ErrorBoundary>
                        <LoggedOut />
                    </ErrorBoundary>
                } />
                <Route path="/.auth/error" element={
                    <Error />
                } />
            </Routes>
            <Logo sx={{ width: 166, height: 39, mt: 18.5, mb: 18.5 }}/>
        </Box>
    </>
);

export default Layout;
