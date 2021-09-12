// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';
import { Switch, Route } from 'react-router-dom';

import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';

import { Headline } from './Headline';
import { ErrorBoundary } from '../error/ErrorBoundary';
import { Error } from '../error/Error'
import { SelectProvider } from '../select-provider/SelectProvider'
import { SelectTenant } from '../select-tenant/SelectTenant'

const useStyles = makeStyles({
    root: {
        position: 'absolute',
        right: '0',
        height: '100%',
        width: '45%',
        maxWidth: '650px',
        backgroundColor: '#191A21',
    },
});

export const Layout = (): JSX.Element => {
    const classes = useStyles();
    return (
        <>
            <Headline />
            <Paper className={classes.root} elevation={24} square={true}>
                <Switch>
                    <Route path="/.auth/select-provider">
                        <ErrorBoundary>
                            <SelectProvider />
                        </ErrorBoundary>
                    </Route>
                    <Route path="/.auth/select-tenant">
                        <ErrorBoundary>
                            <SelectTenant />
                        </ErrorBoundary>
                    </Route>
                    <Route path="/.auth/error">
                        <Error />
                    </Route>
                </Switch>
            </Paper>
        </>
    );
};

export default Layout;
