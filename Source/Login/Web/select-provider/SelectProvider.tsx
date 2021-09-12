// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React, { Suspense } from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import CircularProgress from '@material-ui/core/CircularProgress';
import Typography from '@material-ui/core/Typography';

import { SelectProviderForm } from './SelectProviderForm';

const useStyles = makeStyles({
    root: {
        padding: '158px 64px 0 64px',
    },
    title: {
        marginBottom: '20px',
    },
    subtitle: {
        marginBottom: '30px',
    },
    form: {
        textAlign: 'center',
    }
});

export const SelectProvider = (): JSX.Element => {
    const classes = useStyles();
    return (
        <Box className={classes.root}>
            <Typography variant="h1" className={classes.title}>Welcome to Dolittle Studio!</Typography>
            <Typography variant="h2" className={classes.subtitle}>Sign in to continue</Typography>
            <Box className={classes.form}>
                <Suspense fallback={<CircularProgress />}>
                    <SelectProviderForm/>
                </Suspense>
            </Box>
        </Box>
    );
};
