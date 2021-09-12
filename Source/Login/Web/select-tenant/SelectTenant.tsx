// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React, { Suspense } from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import Button from '@material-ui/core/Button';
import CircularProgress from '@material-ui/core/CircularProgress';
import Link from '@material-ui/core/Link';
import Typography from '@material-ui/core/Typography';
import ChevronLeft from '@material-ui/icons/ChevronLeft';

import { SelectTenantForm } from './SelectTenantForm';

const useStyles = makeStyles({
    root: {
        paddingTop: '212px',
        textAlign: 'center',
    },
    subtitle: {
        marginBottom: '30px',
    },
    noTenant: {
        marginTop: '3px',
    },
    backButton: {
        position: 'absolute',
        left: '20px',
        bottom: '20px',
    }
});

export const SelectTenant = (): JSX.Element => {
    const classes = useStyles();
    return (
        <>
            <Box className={classes.root}>
                <Typography variant="h2" className={classes.subtitle}>Select your tenant</Typography>
                <Suspense fallback={<CircularProgress />}>
                    <SelectTenantForm/>
                </Suspense>
                <Box className={classes.noTenant}>
                    <Typography>Don't have a tenant? <Link underline="always" color="inherit" href="mailto:support@dolittle.com">Email us here.</Link></Typography>
                </Box>
            </Box>
            <Button
                startIcon={<ChevronLeft/>}
                className={classes.backButton}
                onClick={() => window.location.href = '/.auth/logout'}
            >Go back</Button>
        </>
    );
};
