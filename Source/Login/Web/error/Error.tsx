// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Box, Button, Link, Typography } from '@mui/material';
import { Theme } from '../styles/Theme';

import { configuration } from '../Configuration';
import { LoginWrapper } from '../layouts/LoginWrapper';

const unicodeSpaceChar = '\u0020';

const styles = {
    transition: 'all .3s',
    '&:hover': { color: Theme.palette.primary.light }
}

export const Error = (): JSX.Element => {
    const { logoutPath, supportEmail } = configuration;

    return (
        <LoginWrapper>
            <Typography variant="h2" sx={{ marginBottom: '30px', letterSpacing: '-0.5px' }}>
                Oops, something went wrong.
            </Typography>

            <Typography variant='subtitle2'>
                Please log out and try again, or
                {unicodeSpaceChar}
                <Link href={'mailto:' + supportEmail} sx={{ ...styles, textDecoration: 'underline' }}>contact us</Link>
                {unicodeSpaceChar}
                if the issue persists.
            </Typography>

            <Box mt={8}>
                <Button
                    size='large'
                    color='inherit'
                    sx={{ ...styles, fontSize: '14px' }}
                    href={logoutPath}>
                    Log out
                </Button>
            </Box>
        </LoginWrapper>
    );
};
