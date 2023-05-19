import React, { Suspense } from 'react';

import { Box, Button, CircularProgress, Link, Typography } from '@mui/material';

import { configuration } from '../Configuration';
import { WelcomeHeader } from '../components/WelcomeHeader';

export const NoTenant = (): JSX.Element => {
    const { logoutPath, supportEmail } = configuration;

    return (
        <Box>
            <WelcomeHeader />

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