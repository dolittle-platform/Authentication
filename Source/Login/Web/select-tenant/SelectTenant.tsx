// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import CircularProgress from '@mui/material/CircularProgress';
import Link from '@mui/material/Link';
import Typography from '@mui/material/Typography';

import { configuration } from '../Configuration';
import { Theme } from '../styles/Theme';
import { SelectTenantForm } from './SelectTenantForm';

export const SelectTenant = (): JSX.Element => {
    return (
        <>
            <Typography variant='h2' css={{ marginBottom: '66px' }}>Select your tenant</Typography>
            <Suspense fallback={<CircularProgress />}>
                <SelectTenantForm/>
            </Suspense>
            <Box css={{ margin: '18px 16vw 0 16vw', [Theme.breakpoints.up('sm')]: { margin: '75px 0 0 0' } }}>
                { configuration.supportEmail && <Typography>Don't have access to a tenant? <Link href={'mailto:'+configuration.supportEmail}>Contact us</Link> to get started.</Typography>}
            </Box>
            <Button size='large' color='inherit' href={configuration.logoutPath} css={{ marginTop: '128px', [Theme.breakpoints.up('sm')]: { marginTop: '25px' } }}>Log out</Button>
        </>
    );
};
