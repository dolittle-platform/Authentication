import React, { Suspense } from 'react';

import { Box, Button, CircularProgress, Link, Typography } from '@mui/material';

import { configuration } from '../Configuration';

export const NoTenant = (): JSX.Element => {
    const { logoutPath, supportEmail } = configuration;

    return (
        <Box>
            <Typography
                variant='h2'
                sx={{ mb: '2rem' }}>
                No Access
            </Typography>


            <Box sx={{ mt: 12.5, mb: 5 }}>
                {
                    supportEmail &&
                    <Typography variant='subtitle2'>
                        Don't have access to a customer? <Link href={'mailto:' + supportEmail}>Contact us</Link> or <Link href={'mailto:' + supportEmail}>Register</Link> to be notified when we open for external sign-ups.
                    </Typography>
                }
            </Box>

            <Box>
                <Button
                        size='large'
                    color='inherit'
                    href={logoutPath}>
                    Log out
                </Button>
            </Box>
        </Box>
    );
};