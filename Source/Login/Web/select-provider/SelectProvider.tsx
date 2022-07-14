// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import { Box, CircularProgress, Link, Typography } from '@mui/material';

import { SelectProviderForm } from './SelectProviderForm';
import { LoginWrapper } from '../layouts/LoginWrapper'
import { configuration } from '../Configuration';

const unicodeSpaceChar = '\u0020';

export const SelectProvider = (): JSX.Element => {
    return (
        <LoginWrapper>
            <Box mb={12.25} ml='auto' mr='auto' sx={{ maxInlineSize: '369px' }}>
                <Typography variant='h1' sx={{ marginBlockEnd: '32px' }}>
                    Welcome to Dolittle Studio
                </Typography>
                <Typography variant='h5'>
                    Transform your business by leveraging real time events.
                </Typography>
            </Box>

            <Suspense fallback={<CircularProgress />}>
                <SelectProviderForm />
            </Suspense>

            <Typography variant='subtitle2' sx={{ marginBlockStart: '40px' }}>
                Don't have an account?{unicodeSpaceChar}
                <Link href={'mailto:' + configuration.supportEmail}>Contact us</Link>
                {unicodeSpaceChar}to get started.
            </Typography>
        </LoginWrapper>
    );
};
