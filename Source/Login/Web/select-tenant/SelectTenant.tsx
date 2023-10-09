// Copyright (c) Aigonix. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import { Button, CircularProgress, Link, Typography } from '@mui/material';

import { configuration } from '../Configuration';
import { SelectTenantForm } from './SelectTenantForm';

export const SelectTenant = () =>
    <>
        <Typography variant='h2' sx={{ mb: 4 }}>
            Select your customer
        </Typography>

        <Suspense fallback={<CircularProgress />}>
            <SelectTenantForm />
        </Suspense>

        {configuration.supportEmail &&
            <Typography variant='subtitle2' sx={{ mt: 8 }}>
                Don't have access to a customer? <Link href={'mailto:' + configuration.supportEmail}>Contact us</Link> to get started.
            </Typography>
        }

        <Button color='inherit' href={configuration.logoutPath} sx={{ mt: 4 }}>
            Log out
        </Button>
    </>;
