// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Box, Button, Typography } from '@mui/material';
import { ArrowBack } from '@mui/icons-material';

import { configuration } from '../Configuration';
import { LoginWrapper } from '../layouts/LoginWrapper';

export const LoggedOut = (): JSX.Element => {
    return (
        <LoginWrapper>
            <Typography variant='h2' sx={{ letterSpacing: '-0.5px' }}>You have been logged out.</Typography>
            <Box mt={13.5}>
                <Button
                    size='large'
                    color='inherit'
                    startIcon={<ArrowBack />}
                    sx={{ letterSpacing: '0.06em' }}
                    href={configuration.startPath}>
                    Return to login page
                </Button>
            </Box>
        </LoginWrapper>
    );
}
