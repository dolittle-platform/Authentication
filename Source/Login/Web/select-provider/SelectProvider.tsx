// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import { Box, CircularProgress, Typography } from '@mui/material';

import { SelectProviderForm } from './SelectProviderForm';

export const SelectProvider = (): JSX.Element => (
    <>
        <Box sx={{ maxWidth: '23.0625rem', mb: 12.25, ml: 'auto', mr: 'auto' }}>
            <Typography variant='h1' sx={{ mb: '2rem' }}>
                Welcome to Dolittle Studio
            </Typography>
            <Typography variant='h5'>
                Transform your business by leveraging real time events.
            </Typography>
        </Box>

        <Suspense fallback={<CircularProgress />}>
            <SelectProviderForm />
        </Suspense>
    </>
);
