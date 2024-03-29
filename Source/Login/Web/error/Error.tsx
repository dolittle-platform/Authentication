// Copyright (c) Aigonix. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Box, Button, Link, Typography } from '@mui/material';

import { configuration } from '../Configuration';

export const Error = () => {
    const { logoutPath, supportEmail } = configuration;

    const tryAgainText = supportEmail
        ? <>Please log out and try again, or <Link href={'mailto:' + supportEmail}>contact us</Link> if the issue persists.</>
        : <>Please log out and try again.</>

    return (
        <>
            <Typography variant='h2' sx={{ mb: 4 }}>
                Oops, something went wrong.
            </Typography>

            <Typography variant='subtitle2'>
                {tryAgainText}
            </Typography>

            <Box sx={{ mt: 8 }}>
                <Button color='inherit' sx={{ fontSize: 13 }} href={logoutPath}>
                    Log out
                </Button>
            </Box>
        </>
    );
};
