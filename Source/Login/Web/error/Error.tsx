// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Link from '@mui/material/Link';
import Typography from '@mui/material/Typography';

import { configuration } from '../Configuration';

export const Error = (): JSX.Element => {
    const { logoutPath, supportEmail } = configuration;

    const tryAgainText = supportEmail
        ? <>Please log out and try again, or <Link href={'mailto:' + supportEmail}>contact us</Link> if the issue persists.</>
        : <>Please log out and try again.</>

    return (
        <>
            <Typography variant="h2" sx={{ mb: '1.875rem', letterSpacing: '-0.03125px' }}>
                Oops, something went wrong.
            </Typography>

            <Typography variant='subtitle2'>
                { tryAgainText }
            </Typography>

            <Box sx={{ mt: 8 }}>
                <Button
                    size='large'
                    color='inherit'
                    sx={{ fontSize: '0.875rem' }}
                    href={logoutPath}>
                    Log out
                </Button>
            </Box>
        </>
    );
};
