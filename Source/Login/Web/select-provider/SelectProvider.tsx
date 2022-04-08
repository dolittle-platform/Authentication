// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import Box from '@mui/material/Box';
import CircularProgress from '@mui/material/CircularProgress';
import Typography from '@mui/material/Typography';

import { configuration } from '../Configuration';
import { SelectProviderForm } from './SelectProviderForm';

export const SelectProvider = (): JSX.Element => {
    const title = configuration.applicationName ? `Welcome to ${configuration.applicationName}!` : 'Welcome!';
    return (
        <Box css={{ padding: '158px 64px 0 64px' }}>
            <Typography variant="h1" css={{ marginBottom: '20px' }}>{title}</Typography>
            <Typography variant="h2" css={{ marginBottom: '30px' }}>Sign in to continue</Typography>
            <Box css={{ textAlign: 'center' }}>
                <Suspense fallback={<CircularProgress />}>
                    <SelectProviderForm/>
                </Suspense>
            </Box>
        </Box>
    );
};
