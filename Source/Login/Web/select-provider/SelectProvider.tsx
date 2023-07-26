// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';
import { Link } from 'react-router-dom';

import { Box, CircularProgress, Divider, Link as MuiLink, Typography } from '@mui/material';

import { configuration } from '../Configuration';
import { SelectProviderForm } from './SelectProviderForm';
import { WelcomeHeader } from '../components/WelcomeHeader';

export const SelectProvider = () => {
    const { showDolittleHeadline } = configuration;

    return (
        <>
            <Box sx={{ maxWidth: 370, mb: 12.25, mx: 'auto' }}>
                <WelcomeHeader />

                <Typography variant='h5'>
                    {showDolittleHeadline ? 'Transform your business by leveraging real time events.' : 'Sign in to continue.'}
                </Typography>
            </Box>

            <Suspense fallback={<CircularProgress />}>
                <SelectProviderForm />
            </Suspense>

            <Divider variant='middle' sx={{ my: 4, backgroundColor: '#3B3D48' }} />

            <Typography variant='body2' color='gray'>
                Don't have an account? <MuiLink component={Link} to='/.auth/no-tenant'>Get access</MuiLink>
            </Typography>
        </>
    );
};
