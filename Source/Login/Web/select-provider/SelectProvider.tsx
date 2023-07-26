// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import { Box, CircularProgress, Link, Typography } from '@mui/material';

import { configuration } from '../Configuration';
import { SelectProviderForm } from './SelectProviderForm';
import { WelcomeHeader } from '../components/WelcomeHeader';

export const SelectProvider = () => {
    const { showDolittleHeadline, supportEmail } = configuration;

    return (
        <>
            <Box sx={{ maxWidth: 370, mb: 12.25, mx: 'auto' }}>
                <WelcomeHeader />

                <Typography variant='h5' sx={{ mt: 4 }}>
                    {showDolittleHeadline ? 'Transform your business by leveraging real time events.' : 'Sign in to continue.'}
                </Typography>
            </Box>

            <Suspense fallback={<CircularProgress />}>
                <SelectProviderForm />
            </Suspense>

            <Typography sx={{ mt: 4 }}>
                Don't have an account? <Link href={'mailto:' + supportEmail}>Contact us</Link> to get started.
            </Typography>
        </>
    );
};
