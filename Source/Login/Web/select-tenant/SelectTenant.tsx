// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import { Box, Button, CircularProgress, Link, Typography } from '@mui/material';

import { configuration } from '../Configuration';
import { SelectTenantForm } from './SelectTenantForm';

export const SelectTenant = (): JSX.Element => {
    const { logoutPath, supportEmail } = configuration;

    return (
        <Box>
            <Typography
                variant='h2'
                sx={{ mb: '2rem' }}>
                Select your customer
            </Typography>

            <Suspense fallback={<CircularProgress />}>
                <SelectTenantForm />
            </Suspense>

            <Box sx={{ mt: 12.5, mb: 5 }}>
                {
                    supportEmail &&
                        <Typography variant='subtitle2'>
                            Don't have access to a customer? <Link href={'mailto:' + supportEmail}>Contact us</Link> to get started.
                        </Typography>
                }
            </Box>

            <Box>
                <Button
                    size='large'
                    color='inherit'
                    href={logoutPath}>
                    Log out
                </Button>
            </Box>
        </Box>
    );
};
