// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import { Box, CircularProgress, Link, Typography } from '@mui/material';

import { SelectProviderForm } from './SelectProviderForm';
import { configuration } from '../Configuration';

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

        <Typography variant='subtitle2' sx={{ mt: '2.5rem' }}>
            Don't have an account? <Link href={'mailto:' + configuration.supportEmail} sx={{ textDecoration: 'underline' }}>Contact us</Link> to get started.
        </Typography>
    </>
);
