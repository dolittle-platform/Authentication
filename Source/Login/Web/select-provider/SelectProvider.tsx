// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import CircularProgress from '@mui/material/CircularProgress';
import Typography from '@mui/material/Typography';

import { configuration } from '../Configuration';
import { Theme } from '../styles/Theme';
import { SelectProviderForm } from './SelectProviderForm';

export const SelectProvider = (): JSX.Element => {
    const title = configuration.applicationName ? `Welcome to ${configuration.applicationName}` : 'Welcome';
    const subtitle = configuration.showDolittleHeadline ? 'Transforming your business with real time events.' : 'Sign in to continue.';
    return (
        <>
            <Typography variant='h1' css={{ marginBottom: '32px' }}>{title}</Typography>
            <Typography variant='h5' css={{ margin: '0 16vw 80px 16vw', [Theme.breakpoints.up('sm')]: { margin: '0 0 60px 0' } }}>{subtitle}</Typography>
            <Suspense fallback={<CircularProgress />}>
                <SelectProviderForm/>
            </Suspense>
        </>
    );
};
