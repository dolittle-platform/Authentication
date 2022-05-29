// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Suspense } from 'react';

import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import CircularProgress from '@mui/material/CircularProgress';
import Link from '@mui/material/Link';
import Typography from '@mui/material/Typography';
import ChevronLeft from '@mui/icons-material/ChevronLeft';

import { configuration } from '../Configuration';
import { SelectTenantForm } from './SelectTenantForm';

export const SelectTenant = (): JSX.Element => {
    return (
        <>
            <Box css={{ paddingTop: '212px', textAlign: 'center' }}>
                <Typography variant="h2" css={{ marginBottom: '30px' }}>Select your tenant</Typography>
                <Suspense fallback={<CircularProgress />}>
                    <SelectTenantForm/>
                </Suspense>
                <Box css={{ marginTop: '3px' }}>
                    { configuration.supportEmail && <Typography>Don't have a tenant? <Link underline="always" color="inherit" href={'mailto:'+configuration.supportEmail}>Email us here.</Link></Typography>}
                </Box>
            </Box>
            <Button startIcon={<ChevronLeft/>} href={configuration.logoutPath} css={{ position: 'absolute', left: '20px', bottom: '20px' }}>Log out</Button>
        </>
    );
};
