// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.
import { Typography } from '@mui/material';

import { configuration } from '../Configuration';

export const WelcomeHeader = () => {
    const { applicationName } = configuration;

    return <Typography variant='h1' sx={{ mb: '2rem' }}>
        {
            applicationName
                ? `Welcome to ${applicationName}`
                : 'Welcome'
        }
    </Typography>
}