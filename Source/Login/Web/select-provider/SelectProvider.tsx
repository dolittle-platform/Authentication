// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import { Box, CircularProgress, Typography } from '@mui/material';

import { configuration } from '../Configuration';
import { SelectProviderForm } from './SelectProviderForm';
import { WelcomeHeader } from '../components/WelcomeHeader';

export const SelectProvider = (): JSX.Element => {
    const { showDolittleHeadline } = configuration;

    return (
        <Box>
            <Box sx={{ maxWidth: '23.0625rem', mb: 12.25, ml: 'auto', mr: 'auto' }}>
                <WelcomeHeader />
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
