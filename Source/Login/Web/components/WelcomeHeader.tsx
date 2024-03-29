// Copyright (c) Aigonix. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Typography } from '@mui/material';

import { configuration } from '../Configuration';

export const WelcomeHeader = () =>
    <Typography variant='h1' sx={{ mb: '4' }}>
        {configuration.applicationName ? `Welcome to ${configuration.applicationName}` : 'Welcome'}
    </Typography>;
