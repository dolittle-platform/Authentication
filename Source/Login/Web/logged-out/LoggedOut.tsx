// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import ArrowBack from '@mui/icons-material/ArrowBack';

import { configuration } from '../Configuration';

export const LoggedOut = (): JSX.Element => (
    <>
        <Typography variant='h2' sx={{ letterSpacing: '-0.03125rem' }}>You have been logged out.</Typography>
        <Box sx={{ mt: 13.5 }}>
            <Button
                size='large'
                color='inherit'
                startIcon={<ArrowBack />}
                sx={{ letterSpacing: '0.06em' }}
                href={configuration.startPath}>
                Return to login page
            </Button>
        </Box>
    </>
);
