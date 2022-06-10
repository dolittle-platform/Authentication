// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Link from '@mui/material/Link';
import Typography from '@mui/material/Typography';

import { configuration } from '../Configuration';
import { Theme } from '../styles/Theme';

export const Error = (): JSX.Element => {
    return (
        <>
            <Typography variant="h2" css={{ marginBottom: '30px' }}>We're sorry, but something went wrong.</Typography>
            <Box css={{ margin: '0 16vw', [Theme.breakpoints.up('sm')]: { margin: '0 0' } }}>{
                configuration.supportEmail
                    ? <Typography>You can log out and try again, or <Link href={'mailto:'+configuration.supportEmail}>contact us</Link> if it still doesn't work.</Typography>
                    : <Typography>You can log out and try again by clicking below.</Typography>
            }</Box>
            <Button size='large' color='inherit' href={configuration.logoutPath} css={{ marginTop: '128px', [Theme.breakpoints.up('sm')]: { marginTop: '25px' } }}>Log out and try again</Button>
        </>
    );
};
