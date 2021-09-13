// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import Button from '@material-ui/core/Button';
import Link from '@material-ui/core/Link';
import Typography from '@material-ui/core/Typography';

import { configuration } from '../Configuration';

const useStyles = makeStyles({
    root: {
        padding: '158px 64px 0 64px',
    },
    subtitle: {
        marginBottom: '30px',
    },
    button: {
        marginTop: '28px',
    },
});

export const Error = (): JSX.Element => {
    const classes = useStyles();
    return (
        <Box className={classes.root}>
            <Typography variant="h2" className={classes.subtitle}>We're sorry, but something went wrong.</Typography>
            {
                configuration.supportEmail
                    ? <Typography>You can log out and try again, or <Link underline="always" color="inherit" href={'mailto:'+configuration.supportEmail}>email us here.</Link> if it still doesn't work.</Typography>
                    : <Typography>You can log out and try again by clicking below.</Typography>
            }
            <Button
                variant="contained"
                className={classes.button}
                onClick={() => window.location.href = '/.auth/logout'}
            >Log out and try again</Button>
        </Box>
    );
};
