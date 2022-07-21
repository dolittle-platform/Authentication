// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Box, Button, Link, Typography } from '@mui/material';

import { configuration } from '../Configuration';
import { LoginWrapper } from '../layouts/LoginWrapper';

export const Error = (): JSX.Element => {
    const { logoutPath, supportEmail } = configuration;

    return (
        <LoginWrapper>
            <Typography variant="h2" sx={{ marginBottom: '1.875rem', letterSpacing: '-0.5px' }}>
                Oops, something went wrong.
            </Typography>

            <Typography variant='subtitle2'>
                {'Please log out and try again, or '}
                <Link href={'mailto:' + supportEmail} sx={{ textDecoration: 'underline' }}>contact us</Link>
                {' if the issue persists.'}
            </Typography>

            <Box mt={8}>
                <Button
                    size='large'
                    color='inherit'
                    sx={{ fontSize: '0.875rem' }}
                    href={logoutPath}>
                    Log out
                </Button>
            </Box>
        </LoginWrapper>
    );
};
