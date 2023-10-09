// Copyright (c) Aigonix. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Button, Typography } from '@mui/material';
import { ArrowBack } from '@mui/icons-material';

import { configuration } from '../Configuration';

export const LoggedOut = () =>
    <>
        <Typography variant='h2'>You have been logged out.</Typography>

        <Button color='inherit' startIcon={<ArrowBack />} href={configuration.startPath} sx={{ mt: 13.5 }}>
            Return to login page
        </Button>
    </>;
