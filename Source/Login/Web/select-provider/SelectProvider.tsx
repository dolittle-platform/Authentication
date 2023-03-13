// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import { Box, CircularProgress, Typography } from '@mui/material';

import { configuration } from '../Configuration';
import { SelectProviderForm } from './SelectProviderForm';

export const SelectProvider = () => {
    const { applicationName, showDolittleHeadline } = configuration;

    return (
        <>
            <Box sx={{ maxWidth: 370, mb: 12.25, mx: 'auto' }}>
                <Typography variant='h1' sx={{ mb: 4 }}>
                    {
                        applicationName
                            ? `Welcome to ${applicationName}`
                            : 'Welcome'
                    }
                </Typography>
                <Typography variant='h5'>
                    {
                        showDolittleHeadline
                            ? 'Transform your business by leveraging real time events.'
                            : 'Sign in to continue.'
                    }
                </Typography>
            </Box>

            <Suspense fallback={<CircularProgress />}>
                <SelectProviderForm />
            </Suspense>
        </>
    );
};
