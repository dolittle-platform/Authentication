// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Route, Routes as RouterRoutes } from 'react-router-dom';

import { ErrorBoundary } from '../error/ErrorBoundary';
import { Error } from '../error/Error';
import { LoggedOut } from '../logged-out/LoggedOut';
import { SelectProvider } from '../select-provider/SelectProvider';
import { SelectTenant } from '../select-tenant/SelectTenant';
import { NoTenant } from '../no-tenant/NoTenant';

export const Routes = () => (
    <RouterRoutes>
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

        <Route path="/.auth/no-tenant" element={

            <ErrorBoundary>
                <NoTenant />
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
    </RouterRoutes>
);

export default Routes;
