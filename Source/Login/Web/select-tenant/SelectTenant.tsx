// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import { Box, Button, CircularProgress, Link, Typography } from '@mui/material';

import { configuration } from '../Configuration';
import { SelectTenantForm } from './SelectTenantForm';
import { LoginWrapper } from '../layouts/LoginWrapper';

const unicodeSpaceChar = '\u0020';

export const SelectTenant = (): JSX.Element => {
    const { logoutPath, supportEmail } = configuration;

    return (
        <LoginWrapper>
            <Typography
                variant='h2'
                sx={{ marginBlockEnd: '32px' }}>
                Select your customer
            </Typography>

            <Suspense fallback={<CircularProgress />}>
                <SelectTenantForm />
            </Suspense>

            <Box mt={12.5} mb={5}>
                {supportEmail &&
                    <Typography variant='subtitle2'>Don't have access to a tenant?{unicodeSpaceChar}
                        <Link href={'mailto:' + supportEmail}>Contact us</Link>
                        {unicodeSpaceChar}to get started.
                    </Typography>}
            </Box>

            <Box>
                <Button
                    size='large'
                    color='inherit'
                    href={logoutPath}>
                    Log out
                </Button>
            </Box>
        </LoginWrapper>
    );
};
