// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import Box from '@mui/material/Box';
import CircularProgress from '@mui/material/CircularProgress';
import Typography from '@mui/material/Typography';

import { configuration } from '../Configuration';
import { SelectProviderForm } from './SelectProviderForm';

export const SelectProvider = (): JSX.Element => {
    const { applicationName, showDolittleHeadline } = configuration;

    return (
        <Box>
            <Box sx={{ maxWidth: '23.0625rem', mb: 12.25, ml: 'auto', mr: 'auto' }}>
                <Typography variant='h1' sx={{ mb: '2rem' }}>
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
        </Box>
    );
};
