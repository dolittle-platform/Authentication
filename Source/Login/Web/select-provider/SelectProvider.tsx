// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import { Link } from 'react-router-dom';

import { Box, Button, CircularProgress, Typography } from '@mui/material';
import { ArrowForward } from '@mui/icons-material';

import { configuration } from '../Configuration';
import { SelectProviderForm } from './SelectProviderForm';
import { WelcomeHeader } from '../components/WelcomeHeader';

export const SelectProvider = () =>
    <>
        <Box sx={{ maxWidth: 370, mb: 8, mx: 'auto' }}>
            <WelcomeHeader />

            <Typography variant='h5' sx={{ mt: 4 }}>
                {configuration.showDolittleHeadline ? 'Transform your business by leveraging real time events.' : 'Sign in to continue.'}
            </Typography>
        </Box>

        <Suspense fallback={<CircularProgress />}>
            <SelectProviderForm />
        </Suspense>

        <Typography sx={{ mt: 8 }}>Don't have an account?</Typography>
        <Button component={Link} to='/.auth/no-tenant' endIcon={<ArrowForward />}>Get access</Button>
    </>;
