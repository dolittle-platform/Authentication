// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';

import Logo from '../styles/images/logo.svg';

const useStyles = makeStyles({
    root: {
        position: 'absolute',
        top: '246px',
        left: '144px',
    },
    tagline: {
        marginTop: '9px',
        marginLeft: '102px',
        maxWidth: '480px',
    },
});

export const Headline = (): JSX.Element => {
    const classes = useStyles();
    return (
        <Box className={classes.root}>
            <Logo width={347} />
            <Typography variant="h1" className={classes.tagline}>Transforming your business with real time events.</Typography>
        </Box>
    );
};
