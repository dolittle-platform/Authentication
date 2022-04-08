// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Link from '@mui/material/Link';
import Typography from '@mui/material/Typography';

import { configuration } from '../Configuration';

export const Error = (): JSX.Element => {
    return (
        <Box css={{ padding: '158px 64px 0 64px' }}>
            <Typography variant="h2" css={{ marginBottom: '30px' }}>We're sorry, but something went wrong.</Typography>
            {
                configuration.supportEmail
                    ? <Typography>You can log out and try again, or <Link underline="always" color="inherit" href={'mailto:'+configuration.supportEmail}>email us here.</Link> if it still doesn't work.</Typography>
                    : <Typography>You can log out and try again by clicking below.</Typography>
            }
            <Button
                variant="contained"
                onClick={() => window.location.href = '/.auth/logout'}
                css={{ marginTop: '28px' }}
            >Log out and try again</Button>
        </Box>
    );
};
