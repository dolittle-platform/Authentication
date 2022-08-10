// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Routes as RouterRoutes, Route } from 'react-router-dom';

import { ErrorBoundary } from '../error/ErrorBoundary';
import { Error } from '../error/Error'
import { SelectProvider } from '../select-provider/SelectProvider'
import { SelectTenant } from '../select-tenant/SelectTenant'
import { LoggedOut } from '../logged-out/LoggedOut'

export const Routes = (): JSX.Element => (
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